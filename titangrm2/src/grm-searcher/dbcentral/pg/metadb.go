package pg

import (
	"fmt"
	. "grm-searcher/types"
	"grm-service/crypto"
	. "grm-service/dbcentral/pg"
	"grm-service/log"
	"strings"
)

type MetaDB struct {
	MetaCentralDB
}

func (db MetaDB) GetTotalCountWhere(tableName, where string) int64 {
	var total int64 = 0
	sql := fmt.Sprintf("select count(*) from %s where %s;", tableName, where)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return total
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&total)
		return total
	}
	return total
}

func (db *MetaDB) GetPgStorage(dataid string) (string, error) {
	var storage string
	sql := fmt.Sprintf("SELECT storage FROM data_object where uuid = '%s';", dataid)

	rows, err := db.Conn.Query(sql)
	if err != nil {
		return storage, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&storage)
		if err != nil {
			log.Warn(err)
		}
		if len(storage) > 0 {
			//连接字符串解密
			url, err := crypto.AesDecrypt(storage)
			if err != nil {
				return storage, err
			} else {
				return url, nil
			}
		}
	}
	if err := rows.Err(); err != nil {
		return storage, err
	}
	return storage, nil
}

func (db MetaDB) SearchByGeo(dataids []string, r SearchInfo) ([]*MetaInfo, int64, error) {
	infos := make([]*MetaInfo, 0)

	where := fmt.Sprintf(" uuid in (%s) and ST_Intersects(envelope,ST_GeometryFromText('%s', 4326)) ",
		strings.Join(dataids, ","), r.Geometry)
	total := db.GetTotalCountWhere("data_object", where)

	sql := fmt.Sprintf("select name,data_type,sub_type,path,file_size,uuid,ST_AsGeoJson(envelope)"+
		" from data_object where %s", where)
	if len(r.Limit) > 0 && len(r.Offset) > 0 &&
		len(r.Sort) > 0 && len(r.Order) > 0 {
		sql += " order by " + r.Sort + " " + r.Order +
			" limit " + r.Limit + " offset " + r.Offset
	}

	var name, data_type, path, obj_id, envelope, sub_type string
	var file_size float64
	rows, _ := db.Conn.Query(sql)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &data_type, &sub_type, &path, &file_size, &obj_id, &envelope)
		if err != nil {
			continue
		}
		infos = append(infos, &MetaInfo{
			Name:            name,
			DataType:        data_type,
			SubType:         sub_type,
			Path:            path,
			FileSize:        file_size,
			EnvelopeGeoJson: envelope,
			UUID:            obj_id,
		})
	}
	return infos, total, nil
}

func (db MetaDB) DataFilter(_dataids []string, geo string) ([]*MetaInfo, int64, error) {
	infos := make([]*MetaInfo, 0)
	where := fmt.Sprintf(" uuid in (%s) and ST_Intersects(envelope,ST_GeometryFromText('%s', 4326)) ",
		strings.Join(_dataids, ","), geo)
	total := db.GetTotalCountWhere("data_object", where)

	sql := fmt.Sprintf("select name,data_type,sub_type,path,file_size,uuid,ST_AsText(envelope)"+
		" from data_object where %s", where)

	var name, data_type, path, obj_id, envelope, sub_type string
	var file_size float64
	rowsM, _ := db.Conn.Query(sql)
	defer rowsM.Close()
	for rowsM.Next() {
		err := rowsM.Scan(&name, &data_type, &sub_type, &path, &file_size, &obj_id, &envelope)
		if err != nil {
			continue
		}
		infos = append(infos, &MetaInfo{
			Name:            name,
			DataType:        data_type,
			SubType:         sub_type,
			Path:            path,
			FileSize:        file_size,
			EnvelopeGeoJson: envelope,
			UUID:            obj_id,
		})
	}
	return infos, total, nil
}

package pg

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"applications/data-collection/types"
	. "grm-service/dbcentral/pg"
	. "grm-service/time"

	"grm-service/crypto"
	"grm-service/log"
	"grm-service/util"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

type SystemDB struct {
	SysCentralDB
}

func (db SystemDB) AddDataPub(data *types.DataPub) (string, error) {
	value := "false"
	if data.IsPub {
		value = "true"
	}

	sql := fmt.Sprintf(`insert into data_publish(data_id,publish_url,is_pub,data_name,
							data_path,data_type,create_time,wms,wmts,wfs,srs,
							is_cached,wms_pub,wmts_pub,user_id,style) 
							values ('%s','%s', '%s', '%s', '%s',
							'%s',%s,'%s','%s','%s','%s','%t','%s','%s','%s', '%s')`,
		data.DataId, data.PubUrl, value, data.DataName,
		data.DataPath, data.DataType, util.GetTimeNowDB(),
		data.WMS, data.Wmts, data.Wfs, data.Srs,
		data.IsCached, data.WmsUrl, data.WmtsUrl, data.User, data.Style)
	fmt.Println(sql)
	_, err := db.Conn.Exec(sql)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == "23505" {
				return "", errors.New("data is already published")
			}
		}
		return "", err
	}
	return data.PubUrl, nil
}

// 获取数据存放设备信息
func (db SystemDB) GetDataDevice() (*types.Device, error) {
	sql := fmt.Sprintf(`select id,ip_address,db_port,db_user,db_pwd,geo_store from device
							where storage_type = 'DB' and storage_org = 'PostgreSQL' and geo_store is not NULL`)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ip_address, db_port, db_user, db_pwd, store string
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &ip_address, &db_port, &db_user, &db_pwd, &store)
		if err != nil {
			log.Printf("GetDataDevice rows.Scan error: %s\n", err.Error())
			continue
		} else {
			if len(db_pwd) > 0 {
				if pwdEnc, err := crypto.AesDecrypt(db_pwd); err == nil {
					db_pwd = pwdEnc
				}
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &types.Device{
		Id:         id,
		IpAddress:  ip_address,
		DBPort:     db_port,
		DBUser:     db_user,
		DBPwd:      db_pwd,
		DBGeoStore: store,
	}, nil
}

// 添加调查
func (db SystemDB) AddCollection(data *types.Collection) (string, error) {
	if len(data.Id) < 1 {
		data.Id = util.NewUUID()
	}

	fields, err := json.Marshal(data.Fields)
	if err != nil {
		return "", err
	}

	sql := fmt.Sprintf(`insert into data_collector(id,name,user_id,type,description,start_time,end_time,fields) 
							values ('%s','%s','%s','%s','%s','%s','%s','%s')`,
		data.Id, data.Name, data.User, data.Type, data.Description, data.StartTime, data.EndTime, string(fields))
	_, err = db.Conn.Exec(sql)
	if err != nil {
		return "", err
	}
	return data.Id, nil
}

func (db SystemDB) GetCollectionType(colId string) (string, error) {
	sql := fmt.Sprintf(`select type from data_collector where id = '%s'`, colId)
	var typeStr string
	err := db.Conn.QueryRow(sql).Scan(&typeStr)
	if err != nil {
		return "", err
	}
	return typeStr, nil
}

// 获取用户调查列表
func (db SystemDB) GetCollections(user string) ([]*types.Collection, error) {
	var layers []*types.Collection
	sql := fmt.Sprintf(`select id,name,type,description,create_time,start_time,end_time,fields from data_collector where user_id = '%s'`, user)
	//fmt.Println(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return layers, err
	}
	defer rows.Close()

	for rows.Next() {
		var layer types.Collection
		layer.User = user

		var create_time, start_time, end_time time.Time
		var fields string
		err = rows.Scan(&layer.Id, &layer.Name, &layer.Type, &layer.Description, &create_time, &start_time,
			&end_time, &fields)
		if err != nil {
			log.Printf("GetCollections rows.Scan error: %s\n", err.Error())
			return layers, err
		}
		layer.CreateTime = GetTimeStd(create_time)
		layer.StartTime = GetDateStd(start_time)
		layer.EndTime = GetDateStd(end_time)

		if err := json.Unmarshal([]byte(fields), &layer.Fields); err != nil {
			return layers, err
		}
		layers = append(layers, &layer)
	}
	if err := rows.Err(); err != nil {
		return layers, err
	}
	return layers, nil
}

// 移除调查
func (db SystemDB) DelCollection(user, colId string) error {
	sql := fmt.Sprintf(`delete from data_collector where user_id = '%s' and id = '%s'`, user, colId)
	_, err := db.Conn.Exec(sql)
	if err != nil {
		return err
	}
	sql = fmt.Sprintf(`delete from ref_user_data where data_id = '%s'`, colId)
	_, err = db.Conn.Exec(sql)
	return err
}

// 更新调查信息
func (db SystemDB) UpdateCollection(args *types.Collection) error {
	var comma, filter string
	if len(args.Name) > 0 {
		filter = fmt.Sprintf("%s %s name = '%s'", filter, comma, args.Name)
		comma = ","
	}

	if len(args.Description) > 0 {
		filter = fmt.Sprintf("%s %s description = '%s'", filter, comma, args.Description)
		comma = ","
	}

	sql := fmt.Sprintf(`update data_collector set %s  where id = '%s' and user_id = '%s'`,
		filter, args.Id, args.User)
	//fmt.Println(sql)
	_, err := db.Conn.Exec(sql)
	return err
}

package pg

import (
	"database/sql"
	"fmt"
	. "grm-searcher/types"
	. "grm-service/dbcentral/pg"
	"strings"
)

type DataDB struct {
	DataCentralDB
}

func (db DataDB) GetTableData(r SearchInfo) (*TableData, error) {
	var tableData TableData
	var tableName, sqlStr string
	tableName = fmt.Sprintf(`"%s"`, r.DataId)

	var where string = " 1 = 1"
	if r.Geometry != nil {
		where = fmt.Sprintf("ST_Intersects(ST_Transform(geom,4326),ST_GeometryFromText('%s', 4326)) ", r.Geometry)
	}

	//按照概要字段进行查询   TODO

	total := db.GetTotalCountWhere(tableName, where)
	tableData.Total = int(total)

	sqlStr = fmt.Sprintf(`select *, ST_AsGeoJSON(geom) as geom_json from %s  where %s`, tableName, where)
	if len(r.Sort) > 0 && len(r.Order) > 0 {
		sqlStr = fmt.Sprintf(`%s order by %s %s`, sqlStr, r.Sort, r.Order)
	} else {
		sqlStr = fmt.Sprintf(`%s order by gid`, sqlStr)
	}
	if len(r.Offset) > 0 && len(r.Limit) > 0 {
		sqlStr = fmt.Sprintf(`%s limit %s offset %s`, sqlStr, r.Limit, r.Offset)
	}
	rows, err := db.Conn.Query(sqlStr)
	if err != nil {
		return &tableData, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return &tableData, err
	}

	var colName Row
	var _geoIndex int = -1
	for geoIndex, val := range cols {
		if val == "geom" {
			_geoIndex = geoIndex
			continue
		}
		colName.Rows = append(colName.Rows, val)
	}
	tableData.Datas = append(tableData.Datas, colName)

	values := make([]sql.NullString, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(scanArgs...)
		var row Row
		for i, col := range values {

			if _geoIndex >= 0 && i == _geoIndex {
				continue
			}

			var value string
			if col.Valid {
				value = col.String
			}

			row.Rows = append(row.Rows, value)
		}
		tableData.Datas = append(tableData.Datas, row)
	}
	if err := rows.Err(); err != nil {
		return &tableData, err
	}
	return &tableData, nil
}

func (db DataDB) GetDataIds(userid string, ds, data_ids []string) ([]string, error) {
	sql := fmt.Sprintf("select data_id from ref_data_user where data_set in (%s) and user_id = '%s'",
		strings.Join(ds, ","), userid)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return data_ids, err
	}
	var data_id string
	for rows.Next() {
		err := rows.Scan(&data_id)
		if err != nil {
			continue
		}
		data_ids = append(data_ids, data_id)
	}
	rows.Close()
	return data_ids, nil
}

func (db DataDB) GetTotalCountWhere(tableName, where string) int64 {
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

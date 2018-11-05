package pg

import (
	"database/sql"
	"fmt"
	"time"

	"applications/data-collection/types"

	. "data-manager/types"
	. "grm-service/dbcentral/pg"
	"grm-service/util"
)

type DataDB struct {
	DataCentralDB
}

// 创建数据表格
func (db DataDB) CreateDataTable(args *types.Collection) error {
	cols := `gid serial, 
			名称 varchar(256) NOT NULL,
			geom geometry,
			ts_vector tsvector default ''::tsvector`
	comma := ","
	for _, field := range args.Fields {
		if field.Name == "gid" || field.Name == "名称" || field.Name == "geom" {
			continue
		}
		cols = fmt.Sprintf(`%s %s %s varchar(512) NOT NULL DEFAULT ''`, cols, comma, field.Name)
	}

	sql := fmt.Sprintf(`CREATE TABLE "ftr-%s" (%s);
	ALTER TABLE "ftr-%s" ADD PRIMARY KEY (gid);`, args.Id, cols, args.Id)
	fmt.Println(sql)
	_, err := db.Conn.Exec(sql)
	return err
}

// 添加数据记录
func (db DataDB) AddTableData(dataId string, args *types.DataFields) error {
	var cols, comma, values, tsVector string
	for _, field := range *args {
		if field.Name == "gid" || field.Name == "create_time" {
			continue
		}
		cols = fmt.Sprintf(`%s %s %s`, cols, comma, field.Name)
		if field.Name == "geom" {
			values = fmt.Sprintf(`%s %s ST_GeomFromGeoJSON('%s')`, values, comma, field.Value)
		} else {
			values = fmt.Sprintf(`%s %s '%s'`, values, comma, field.Value)
			if field.Name == "名称" {
				tsVector = field.Value
			}
		}
		comma = ","
	}

	sql := fmt.Sprintf(`insert into "ftr-%s"(%s,ts_vector) values (%s,to_tsvector('zh_CN', '%s')) returning gid`,
		dataId, cols, values, tsVector)
	fmt.Println(sql)

	var gid string
	//var create_time time.Time
	if err := db.Conn.QueryRow(sql).Scan(&gid); err != nil {
		return err
	}
	*args = append(*args, types.DataField{Name: "gid", Value: gid})
	//*args = append(*args, types.DataField{Name: "create_time", Value: util.GetTimeStd(create_time)})
	return nil
}

// 移除数据记录
func (db DataDB) DelTableData(colId, dataId string) error {
	sql := fmt.Sprintf(`delete from "ftr-%s" where gid = %s`, colId, dataId)
	_, err := db.Conn.Exec(sql)
	return err
}

// 编辑数据记录
func (db DataDB) UpdateTableData(colId, dataId string, args *types.DataFields) (*types.DataFields, error) {
	var comma, values string
	for _, field := range *args {
		if field.Name == "geom" {
			values = fmt.Sprintf("%s %s %s = ST_GeomFromGeoJSON('%s')", values, comma, field.Name, field.Value)
		} else {
			values = fmt.Sprintf("%s %s %s = '%s'", values, comma, field.Name, field.Value)
			if field.Name == "名称" {
				values = fmt.Sprintf(`%s,ts_vector = to_tsvector('zh_CN', '%s')`, values, field.Value)
			}
		}
		comma = ","
	}

	sqlStr := fmt.Sprintf(`update "ftr-%s" set %s where gid = %s returning *, ST_AsGeoJSON(geom) as geom_json`, colId, values, dataId)
	fmt.Println(sqlStr)

	rows, err := db.Conn.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	fields := types.DataFields{}
	_geoIndex := -1
	time_index := -1
	for index, val := range cols {
		if val == "geom" {
			_geoIndex = index
			//continue
		} else if val == "create_time" {
			time_index = index
		} else if val == "geom_json" {
			val = "geom"
		}
		fields = append(fields, types.DataField{Name: val})
	}

	retValues := make([]sql.NullString, len(cols))
	scanArgs := make([]interface{}, len(retValues))
	for i := range retValues {
		scanArgs[i] = &retValues[i]
	}

	if rows.Next() {
		rows.Scan(scanArgs...)
		for i, col := range retValues {
			if _geoIndex >= 0 && i == _geoIndex {
				continue
			}

			var value string
			if col.Valid {
				value = col.String
			}
			//fmt.Println(value)
			if i == time_index {
				create_time, err := time.Parse("2006-01-02T15:04:05Z", value)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fields[i].Value = util.GetTimeStd(create_time)
			} else {
				fields[i].Value = value
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fields = append(fields[:_geoIndex], fields[_geoIndex+1:]...)
	return &fields, nil
}

///////////////////////////////////////////////////////////////
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

func (db DataDB) GetTableData(r *DataSearch) (*TableData, error) {
	var tableData TableData
	var tableName, sqlStr string
	tableName = fmt.Sprintf(`"ftr-%s"`, r.DataId)

	where := " 1 = 1"
	if len(r.Key) > 0 {
		val := `'%` + r.Key + `%'`
		where = fmt.Sprintf(`%s and "名称" ilike %s`, where, val)
	}

	total := db.GetTotalCountWhere(tableName, where)
	tableData.Total = int(total)

	//if len(r.Limit) == 0 && len(r.Offset) == 0 {
	//	r.Offset = "0"
	//	r.Limit = "20"
	//}

	sqlStr = fmt.Sprintf(`select *, ST_AsGeoJSON(geom) as geom_json from %s  where %s`, tableName, where)
	if len(r.Sort) > 0 && len(r.Order) > 0 {
		sqlStr = fmt.Sprintf(`%s order by %s %s`, sqlStr, r.Sort, r.Order)
	} else {
		sqlStr = fmt.Sprintf(`%s order by gid desc`, sqlStr)
	}
	if len(r.Offset) > 0 && len(r.Limit) > 0 {
		sqlStr = fmt.Sprintf(`%s limit %s offset %s`, sqlStr, r.Limit, r.Offset)
	}
	fmt.Println(sqlStr)
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
	_geoIndex := -1
	time_index := -1

	for geoIndex, val := range cols {
		if val == "geom" {
			_geoIndex = geoIndex
			continue
		}
		if val == "create_time" {
			time_index = geoIndex
		}
		if val == "geom_json" {
			val = "geom"
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

			if i == time_index {
				create_time, err := time.Parse("2006-01-02T15:04:05Z", value)
				if err != nil {
					fmt.Println(err)
					continue
				}
				row.Rows = append(row.Rows, util.GetTimeStd(create_time))
			} else {
				row.Rows = append(row.Rows, value)
			}
		}
		tableData.Datas = append(tableData.Datas, row)
	}
	if err := rows.Err(); err != nil {
		return &tableData, err
	}
	return &tableData, nil
}

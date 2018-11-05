package pg

import (
	//	"database/sql"
	"fmt"
	. "grm-service/dbcentral/pg"
	"grm-service/log"
	"strings"
	. "titan-statistics/types"
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

func (db MetaDB) UserDataTypeStat(dataids []string) (*TypeStatList, error) {
	sql := fmt.Sprintf("select count(*),data_type from data_object where uuid in (%s) group by data_type;",
		strings.Join(dataids, ","))
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list TypeStatList
	for rows.Next() {
		var data_type string
		var count int
		err = rows.Scan(&count, &data_type)
		if err != nil {
			log.Warn("DataTypeStat.rows.Scan error:", err)
		} else {
			list = append(list, TypeStat{data_type, count})
		}
	}
	return &list, nil
}

func (db MetaDB) DataTypeStat() (*TypeStatList, error) {
	sql := "select count(*),data_type from data_object group by data_type;"
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list TypeStatList
	for rows.Next() {
		var data_type string
		var count int
		err = rows.Scan(&count, &data_type)
		if err != nil {
			log.Warn("DataTypeStat.rows.Scan error:", err)
		} else {
			list = append(list, TypeStat{data_type, count})
		}
	}
	return &list, nil
}

func (db MetaDB) UserSubTypeStat(data_type string, dataids []string) (*SubtypeList, error) {
	sql := fmt.Sprintf("select count(*),sub_type from data_object where data_type = '%s' and uuid in (%s) group by sub_type;",
		data_type, strings.Join(dataids, ","))
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list SubtypeList = SubtypeList{DataType: data_type}
	for rows.Next() {
		var data_type string
		var count int
		err = rows.Scan(&count, &data_type)
		if err != nil {
			log.Warn("DataTypeStat.rows.Scan error:", err)
		} else {
			list.TypeStatList = append(list.TypeStatList, SubtypeStat{data_type, count})
		}
	}
	return &list, nil
}

func (db MetaDB) SubTypeStat(data_type string) (*SubtypeList, error) {
	sql := fmt.Sprintf("select count(*),sub_type from data_object where data_type = '%s' group by sub_type;", data_type)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list SubtypeList = SubtypeList{DataType: data_type}
	for rows.Next() {
		var data_type string
		var count int
		err = rows.Scan(&count, &data_type)
		if err != nil {
			log.Warn("DataTypeStat.rows.Scan error:", err)
		} else {
			list.TypeStatList = append(list.TypeStatList, SubtypeStat{data_type, count})
		}
	}
	return &list, nil
}

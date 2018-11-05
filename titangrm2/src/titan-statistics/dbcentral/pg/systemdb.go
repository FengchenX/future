package pg

import (
	"fmt"
	//	. "grm-searcher/types"
	//"grm-service/common"
	//"grm-service/crypto"
	. "grm-service/dbcentral/pg"
	//	"strings"
)

type SystemDB struct {
	SysCentralDB
}

func (db SystemDB) GetMarketsDataIds() ([]string, error) {
	var data_ids []string = make([]string, 0)
	sql := fmt.Sprintf("select data_id from ref_data_user where is_market = true;")
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	var data_id string
	for rows.Next() {
		err := rows.Scan(&data_id)
		if err != nil {
			continue
		}
		data_ids = append(data_ids, fmt.Sprintf("'%s'", data_id))
	}
	rows.Close()
	return data_ids, nil
}

func (db SystemDB) GetDataIds(userid string) ([]string, error) {
	var data_ids []string = make([]string, 0)
	sql := fmt.Sprintf("select data_id from ref_data_user where user_id = '%s';", userid)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return data_ids, err
	}

	for rows.Next() {
		var data_id string
		err := rows.Scan(&data_id)
		if err != nil {
			continue
		}
		data_ids = append(data_ids, fmt.Sprintf("'%s'", data_id))
	}
	rows.Close()
	return data_ids, nil
}

func (db SystemDB) GetDataIdsByOne(ds string) ([]string, error) {
	var data_ids []string = make([]string, 0)
	sql := fmt.Sprintf("select data_id from ref_data_user where data_set = '%s';", ds)
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

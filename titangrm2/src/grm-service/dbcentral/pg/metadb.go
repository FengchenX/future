package pg

import (
	"fmt"
	"strings"

	"grm-service/crypto"
)

type MetaCentralDB struct {
	Central
}

func ConnectMetaDB(host, user, password string) (MetaCentralDB, error) {
	central, err := ConnectDB(host, MetaDBName, user, password)
	if err != nil {
		return MetaCentralDB{}, err
	}
	return MetaCentralDB{central}, err
}

// 获取设备geo_storage
func (db MetaCentralDB) GetDeviceGeoStorage(data string) (string, error) {
	sql := fmt.Sprintf(`select storage from data_object where uuid = '%s'`, data)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var storage string
	if rows.Next() {
		err = rows.Scan(&storage)
		if err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	// 获取ip
	url, err := crypto.AesDecrypt(storage)
	if err != nil {
		return "", err
	}
	pre_index := strings.Index(url, "@")
	if pre_index == -1 {
		return "", err
	}
	ip := url[pre_index+1:]

	comma_index := strings.Index(ip, ":")
	if comma_index == -1 {
		return "", err
	}
	return "postgis_" + ip[:comma_index], nil
}

// 获取数据信息
func (db MetaCentralDB) GetDataPath(dataId string) (string, string, error) {
	sql := fmt.Sprintf(`select name,path from data_object where uuid = '%s'`, dataId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	var name, path string
	if rows.Next() {
		err = rows.Scan(&name, &path)
		if err != nil {
			return "", "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", "", err
	}
	return name, path, nil
}

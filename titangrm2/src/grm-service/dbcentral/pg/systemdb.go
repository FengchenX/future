package pg

import (
	"fmt"
	"time"

	"data-manager/types"
	"grm-service/common"
	"grm-service/crypto"
	"grm-service/log"
	. "grm-service/time"
	"grm-service/util"
)

type SysCentralDB struct {
	Central
}

func ConnectSystemDB(host, user, password string) (SysCentralDB, error) {
	central, err := ConnectDB(host, SysDBName, user, password)
	if err != nil {
		return SysCentralDB{}, err
	}
	return SysCentralDB{central}, err
}

// 获取类型信息
func (db SysCentralDB) GetTypesInfo(dataType string) (map[string]*types.DataType, error) {
	datas := make(map[string]*types.DataType, 40)

	var sql string
	if len(dataType) == 0 {
		sql = fmt.Sprintf(`select name,label,parent,is_obsoleted,extension,create_time,description from data_type`)
	} else {
		sql = fmt.Sprintf(`select name,label,parent,is_obsoleted,extension,create_time,description 
							from data_type where name = '%s'`, dataType)
	}
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data types.DataType
		var create_time time.Time
		err = rows.Scan(&data.Name, &data.Label, &data.Parent, &data.IsObsoleted, &data.Extensions,
			&create_time, &data.Description)
		if err != nil {
			log.Error("GetTypesInfo rows.Scan error: %s\n", err.Error())
			continue
		} else {
			data.CreateTime = GetTimeStd(create_time)
			datas[data.Name] = &data
		}
	}
	if err := rows.Err(); err != nil {
		return datas, err
	}
	return datas, nil
}

// 获取设备数据库连接信息 TODO: 移除
func (db SysCentralDB) GetDataDBInfo(device string) (*ConnConfig, error) {
	var config ConnConfig
	sql := fmt.Sprintf(`select ip_address,db_port,db_user,db_pwd from device where id = %s and storage_type = '%s'`,
		device, common.DBType)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ip, port, user, pwd string
	if rows.Next() {
		err = rows.Scan(&ip, &port, &user, &pwd)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(pwd) > 0 {
		if pwdEnc, err := crypto.AesDecrypt(pwd); err == nil {
			pwd = pwdEnc
		}
	}
	config = ConnConfig{
		Host:     ip,
		Port:     port,
		Database: DataDBName,
		User:     user,
		Password: pwd,
	}
	return &config, nil
}

// 获取设备连接信息
func (db SysCentralDB) GetDeviceStr(device int) (string, error) {
	sql := fmt.Sprintf(`select storage_org,storage_type,file_sys,mount_path,
						ip_address,db_port,db_user,db_pwd 
						from device where id = %d`, device)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var ip, port, user, pwd string
	var storage_org, storage_type, file_sys, mount_path string
	if rows.Next() {
		err = rows.Scan(&storage_org, &storage_type, &file_sys, &mount_path, &ip, &port, &user, &pwd)
		if err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	var devStr string
	// 判定storage_type
	switch storage_type {
	case common.DBType:
		{
			if len(user) == 0 || len(pwd) == 0 || len(ip) == 0 || len(port) == 0 {
				return devStr, fmt.Errorf(util.TR("Invalid database connection info:%s", device))
			}
			if pwdEnc, err := crypto.AesDecrypt(pwd); err == nil {
				pwd = pwdEnc
			}
			if storage_org == common.POSTGRESQL {
				devStr = fmt.Sprintf("postgres://%s:%s@%s:%s/TitanCloud.Data",
					user, pwd, ip, port)
			}
		}
	case common.DFSType:
		devStr = "dfs"
	case common.NFSType:
		if len(mount_path) == 0 || len(file_sys) == 0 {
			return devStr, fmt.Errorf(util.TR("Invalid nfs file system info :%s", device))
		}
		devStr = fmt.Sprintf("nfs:%s", mount_path)
	}
	return devStr, nil
}

// 获取设备geo_storage
func (db SysCentralDB) GetDeviceGeoStorage(device int) (string, error) {
	sql := fmt.Sprintf(`select geo_storage from device where id = %d`, device)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var geo_storage string
	if rows.Next() {
		err = rows.Scan(&geo_storage)
		if err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	return geo_storage, nil
}

// 注册数据图层
func (db SysCentralDB) AddDataLayer(layer *common.DataLayer) (*common.DataLayer, error) {
	isDefault := "false"
	if layer.IsDefault {
		isDefault = "true"
	}
	if len(layer.Layer) == 0 {
		layer.Layer = util.NewUUID()
	}
	sql := fmt.Sprintf(`insert into ref_data_layer(id,data_id,user_id,name,style,description,is_default,
							srs, wms, wmts, wfs, wms_pub, wmts_pub) 
							values ('%s','%s','%s','%s', '%s','%s',%s,'%s','%s','%s','%s', '%s','%s') returning create_time`,
		layer.Layer, layer.Data, layer.User, layer.Name, layer.Style, layer.Description, isDefault,
		layer.Srs, common.GisPre+layer.WMS, common.GisPre+layer.Wmts,
		common.GisPre+layer.Wfs, common.GisPre+layer.WmsUrl, common.GisPre+layer.WmtsUrl)
	fmt.Println(sql)
	stmt, err := db.Conn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var createTime time.Time
	err = stmt.QueryRow().Scan(&createTime)
	if err != nil {
		return nil, err
	}
	layer.CreateTime = GetTimeStd(createTime)
	return layer, nil
}

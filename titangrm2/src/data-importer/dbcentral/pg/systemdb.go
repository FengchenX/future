package pg

import (
	"fmt"
	"grm-service/common"
	"grm-service/crypto"
	"grm-service/util"
	"math/rand"
	"strconv"

	. "grm-service/dbcentral/pg"

	"data-manager/types"
	"grm-service/log"
	. "storage-manager/types"
	"strings"
	"time"
)

type SystemDB struct {
	SysCentralDB
}

func (db SystemDB) GetDataSetType(dataset string) (string, error) {
	sql := fmt.Sprintf(`select type from data_set where id = '%s'`, dataset)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var dataSetType string
	if rows.Next() {
		err := rows.Scan(&dataSetType)
		if err != nil {
			log.Errorf("GetDataSetType rows.Scan error: %s\n", err.Error())
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	return dataSetType, nil
}

func (db SystemDB) GetDeviceStrByType(dataTypes string) (string, string, error) {
	sql := fmt.Sprintf(`select id,label,storage_type,data_type,
							ip_address,server_name,db_port,db_user,db_pwd,file_sys,
							mount_path,create_time,description,volume,storage_org 
							from device order by create_time`)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	var sysTypes map[string]*types.DataType
	if len(dataTypes) > 0 {
		sysTypes, err = db.GetTypesInfo("")
		if err != nil {
			log.Error("GetTypesInfo: %s", err.Error())
			return "", "", err
		}
	}

	var devices DeviceList
	var create_time time.Time
	for rows.Next() {
		var dev Device
		err = rows.Scan(&dev.Id, &dev.Label, &dev.StorageType, &dev.DataType, &dev.IpAddress, &dev.ServiceName,
			&dev.DBPort, &dev.DBUser, &dev.DBPwd, &dev.FileSys, &dev.MountPath, &create_time,
			&dev.Description, &dev.Volume, &dev.StorageOrg)
		if err != nil {
			log.Error("GetDevices rows.Scan error: %s", err.Error())
			continue
		} else {
			if len(dataTypes) > 0 {
				//fmt.Println(dev.DataType)
				has := false
				if !strings.Contains(dev.DataType, dataTypes) {
					devTypes := strings.Split(dev.DataType, ",")
					for _, val := range devTypes {
						//fmt.Println(val)
						if sysTypes[val].Parent == dataTypes || sysTypes[dataTypes].Parent == val {
							has = true
							break
						}
					}
				} else {
					has = true
				}
				if !has {
					continue
				}
			}
			devices = append(devices, dev)
		}
	}

	if err := rows.Err(); err != nil {
		return "", "", err
	}
	if len(devices) == 0 {
		return "", "", err
	}

	// 判定storage_type
	var devStr string
	dev := devices[rand.Intn(len(devices))]
	switch dev.StorageType {
	case common.DBType:
		{
			if len(dev.DBUser) == 0 || len(dev.DBPwd) == 0 || len(dev.IpAddress) == 0 || len(dev.DBPort) == 0 {
				return "", devStr, fmt.Errorf(util.TR("Invalid database connection info:%s", dev))
			}
			if pwdEnc, err := crypto.AesDecrypt(dev.DBPwd); err == nil {
				dev.DBPwd = pwdEnc
			}
			if dev.StorageOrg == common.POSTGRESQL {
				devStr = fmt.Sprintf("postgres://%s:%s@%s:%s/TitanCloud.Data",
					dev.DBUser, dev.DBPwd, dev.IpAddress, dev.DBPort)
			}
		}
	case common.DFSType:
		devStr = "dfs"
	case common.NFSType:
		if len(dev.MountPath) == 0 || len(dev.FileSys) == 0 {
			return "", devStr, fmt.Errorf(util.TR("Invalid nfs file system info :%s", dev))
		}
		devStr = fmt.Sprintf("nfs:%s", dev.MountPath)
	}
	return strconv.Itoa(dev.Id), devStr, nil
}

package pg

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"

	"data-manager/types"
	. "storage-manager/types"

	"grm-service/crypto"
	. "grm-service/dbcentral/pg"
	"grm-service/log"
	. "grm-service/time"
	//. "grm-service/time"
	//"grm-service/log"
)

type SystemDB struct {
	SysCentralDB
}

// 注册存储设备
func (db SystemDB) DeviceRegistry(device *Device) (*Device, error) {
	if len(device.DBPwd) > 0 {
		if pwdEnc, err := crypto.AesEncrypt(device.DBPwd); err == nil {
			device.DBPwd = pwdEnc
		}
	}

	sql := fmt.Sprintf(`insert into device(label,storage_type,data_type,
							ip_address,server_name,db_port,db_user,db_pwd,file_sys,
							mount_path,description,volume,storage_org,geo_storage) values ('%s','%s','%s',
							'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s') returning id`,
		device.Label, device.StorageType, device.DataType, device.IpAddress, device.ServiceName,
		device.DBPort, device.DBUser, device.DBPwd, device.FileSys, device.MountPath,
		device.Description, device.Volume, device.StorageOrg, device.GeoStorage)
	fmt.Println("RegistryDevice:", sql)
	stmt, err := db.Conn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == "23505" {
				return nil, ErrDeviceNameExists
			}
		}
		return nil, err
	}
	device.Id = id
	return device, nil
}

// 获取当前系统存储设备
func (db SystemDB) GetDevices(dataTypes string) (DeviceList, error) {
	var devices DeviceList
	sql := fmt.Sprintf(`select id,label,storage_type,data_type,
							ip_address,server_name,db_port,db_user,db_pwd,file_sys,
							mount_path,create_time,description,volume,storage_org 
							from device order by create_time`)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sysTypes map[string]*types.DataType
	if len(dataTypes) > 0 {
		sysTypes, err = db.GetTypesInfo("")
		if err != nil {
			log.Error("GetTypesInfo: %s", err.Error())
			return nil, err
		}
	}

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
				fmt.Println(dev.DataType)
				has := false
				if !strings.Contains(dev.DataType, dataTypes) {
					devTypes := strings.Split(dev.DataType, ",")
					for _, val := range devTypes {
						fmt.Println(val)
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
			dev.CreateTime = GetTimeStd(create_time)
			devices = append(devices, dev)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		return nil, nil
	}
	return devices, nil
}

// 更新存储设备信息
func (db SystemDB) UpdateDevice(id string, device *UpdateDeviceRequest) error {
	if len(device.DBPwd) > 0 {
		if pwdEnc, err := crypto.AesEncrypt(device.DBPwd); err == nil {
			device.DBPwd = pwdEnc
		}
	}

	var volume string
	if len(device.Volume) > 0 {
		volume = strings.Replace(device.Volume, " ", "", -1)
		volume = strings.ToUpper(volume)
		if strings.HasSuffix(volume, "K") ||
			strings.HasSuffix(volume, "M") ||
			strings.HasSuffix(volume, "G") ||
			strings.HasSuffix(volume, "T") {
			volume = volume + "B"
		} else if strings.HasSuffix(volume, "KB") ||
			strings.HasSuffix(volume, "MB") ||
			strings.HasSuffix(volume, "GB") ||
			strings.HasSuffix(volume, "TB") {
		} else {
			return ErrDeviceVolume
		}
		_, err := strconv.Atoi(volume[:len(volume)-2])
		if err != nil {
			return ErrDeviceVolume
		}
	}
	sql := fmt.Sprintf(`update device set `)

	var comma string
	if device.Label != "" {
		sql = fmt.Sprintf(`%s %s label = '%s'`, sql, comma, device.Label)
		comma = ","
	}
	if device.DateType != "" {
		sql = fmt.Sprintf(`%s %s data_type = '%s'`, sql, comma, device.DateType)
		comma = ","
	}
	if device.IpAddress != "" {
		sql = fmt.Sprintf(`%s %s ip_address = '%s'`, sql, comma, device.IpAddress)
		comma = ","
	}
	if device.DBPort != "" {
		sql = fmt.Sprintf(`%s %s db_port = '%s'`, sql, comma, device.DBPort)
		comma = ","
	}
	if device.DBUser != "" {
		sql = fmt.Sprintf(`%s %s db_user = '%s'`, sql, comma, device.DBUser)
		comma = ","
	}
	if device.DBPwd != "" {
		sql = fmt.Sprintf(`%s %s db_pwd = '%s'`, sql, comma, device.DBPwd)
		comma = ","
	}
	if device.FileSys != "" {
		sql = fmt.Sprintf(`%s %s file_sys = '%s'`, sql, comma, device.FileSys)
		comma = ","
	}
	if device.MountPath != "" {
		sql = fmt.Sprintf(`%s %s mount_path = '%s'`, sql, comma, device.MountPath)
		comma = ","
	}
	if device.Description != "" {
		sql = fmt.Sprintf(`%s %s description = '%s'`, sql, comma, device.Description)
		comma = ","
	}
	if device.GeoServer != "" {
		sql = fmt.Sprintf(`%s %s geo_storage = '%s'`, sql, comma, device.GeoServer)
		comma = ","
	}
	if device.Volume != "" {
		sql = fmt.Sprintf(`%s %s volume = '%s'`, sql, comma, volume)
		comma = ","
	}
	sql = fmt.Sprintf("%s where id = %s", sql, id)
	fmt.Println("UpdateDevice:", sql)

	if len(comma) > 0 {
		_, err := db.Conn.Exec(sql)
		if err != nil {
			if pgErr, ok := err.(pgx.PgError); ok {
				if pgErr.Code == "23505" {
					return ErrDeviceNameExists
				}
			}
			return err
		}
	}
	return nil
}

func (db SystemDB) DeleteDevice(id string) error {
	sql := fmt.Sprintf(`delete from device where id = %s`, id)
	_, err := db.Conn.Exec(sql)
	return err
}

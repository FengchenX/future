package types

import (
	"errors"

	. "grm-service/util"
)

var (
	ErrInvalidDBInfo    = errors.New(TR("Invalid database connection info"))
	ErrInvalidNFSInfo   = errors.New(TR("Invalid nfs file system info"))
	ErrDeviceNameExists = errors.New(TR("device name already exists"))
	ErrDeviceVolume     = errors.New(TR("invalid volume size"))
	ErrGetDeviceInfo    = errors.New("Failed to get storage volume")
)

// 存储设备
type Device struct {
	Id          int    `json:"id"`
	Label       string `json:"label"`
	StorageType string `json:"storage_type"`
	StorageOrg  string `json:"storage_org"`
	DataType    string `json:"data_types"`
	IpAddress   string `json:"ip_address"`
	ServiceName string `json:"server_name"`
	DBPort      string `json:"db_port"`
	DBUser      string `json:"db_user"`
	DBPwd       string `json:"db_pwd"`
	GeoStorage  string `json:"geo_storage"`
	FileSys     string `json:"file_sys,omitempty"`
	MountPath   string `json:"mount_path,omitempty"`
	Volume      string `json:"total_volume"`
	Used        string `json:"used"`
	UsedPercent string `json:"used_percent"`
	CreateTime  string `json:"create_time"`
	Description string `json:"description"`
}

type DeviceList []Device

// 设备更新
type UpdateDeviceRequest struct {
	Label       string `json:"label"`
	DateType    string `json:"data_type"`
	IpAddress   string `json:"ip_address"`
	DBPort      string `json:"db_port"`
	DBUser      string `json:"db_user"`
	DBPwd       string `json:"db_pwd"`
	FileSys     string `json:"file_sys"`
	MountPath   string `json:"mount_path"`
	Description string `json:"description"`
	Volume      string `json:"volume"`
	GeoServer   string `json:"geo_server"`
}

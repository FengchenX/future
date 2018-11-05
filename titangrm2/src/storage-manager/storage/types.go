package storage

import (
	"time"
)

const (
	DefaultTimerInterval = time.Second * 60
)

// 设备注册
type registryDeviceRequest struct {
	Label       string `json:"label"`
	StorageType string `json:"storage_type",description:"DB/NFS/DFS"`
	StorageOrg  string `json:"storage_org"`
	DateType    string `json:"data_type"`
	IpAddress   string `json:"ip_address"`
	ServiceName string `json:"server_name"`
	DBPort      string `json:"db_port"`
	DBUser      string `json:"db_user"`
	DBPwd       string `json:"db_pwd"`
	FileSys     string `json:"file_sys"`
	MountPath   string `json:"mount_path"`
	Volume      string `json:"volume"`
	Description string `json:"description"`
}

package storage

import (
	"strconv"
	"strings"
	"time"

	"grm-service/dbcentral/pg"
	"grm-service/log"

	"github.com/emicklei/go-restful"

	"grm-service/common"
	"grm-service/geoserver"
	. "grm-service/util"
	. "storage-manager/types"
)

var (
	volUnit = map[string]string{"K": "KB", "M": "MB", "G": "GB", "T": "TB"}
)

// POST http://localhost:8080/devices
func (s StorageSvc) deviceRegistry(req *restful.Request, res *restful.Response) {
	reqInfo := registryDeviceRequest{}
	err := req.ReadEntity(&reqInfo)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	if len(reqInfo.Volume) > 0 {
		volume := strings.Replace(reqInfo.Volume, " ", "", -1)
		volume = strings.ToUpper(volume)
		if _, ok := volUnit[volume[len(volume)-1:]]; ok {
			reqInfo.Volume = volume + "B"
		}

		_, err := strconv.Atoi(volume[:len(volume)-2])
		if err != nil {
			ResWriteError(res, ErrDeviceVolume)
			return
		}
	}

	// 检查信息完整性
	var geoStorage string
	switch reqInfo.StorageType {
	case common.DBType:
		{
			if len(reqInfo.IpAddress) == 0 ||
				len(reqInfo.DBUser) == 0 ||
				len(reqInfo.DBPort) == 0 ||
				len(reqInfo.DBPwd) == 0 {
				ResWriteError(res, ErrInvalidDBInfo)
				return
			}

			// 注册geoserver中的storage
			if reqInfo.StorageOrg == common.POSTGRESQL {
				geoStorage = "postgis_" + reqInfo.IpAddress
				err := s.GeoServer.AddPgStore(geoserver.GeoWorkSpace, geoStorage,
					reqInfo.IpAddress, reqInfo.DBPort, pg.DataDBName, reqInfo.DBUser, reqInfo.DBPwd)
				if err != nil {
					log.Error("Failed to add pg store:", err.Error())
					return
				}
			}
		}
	case common.NFSType:
		{
			if len(reqInfo.FileSys) == 0 ||
				len(reqInfo.MountPath) == 0 {
				ResWriteError(res, ErrInvalidNFSInfo)
				return
			}
		}
	case common.DFSType:
		{

		}
	}

	dev := Device{
		Label:       reqInfo.Label,
		StorageType: reqInfo.StorageType,
		StorageOrg:  reqInfo.StorageOrg,
		DataType:    reqInfo.DateType,
		IpAddress:   reqInfo.IpAddress,
		ServiceName: reqInfo.ServiceName,
		DBPort:      reqInfo.DBPort,
		DBUser:      reqInfo.DBUser,
		DBPwd:       reqInfo.DBPwd,
		GeoStorage:  geoStorage,
		FileSys:     reqInfo.FileSys,
		MountPath:   reqInfo.MountPath,
		Description: reqInfo.Description,
		Volume:      reqInfo.Volume,
	}
	ret, err := s.SysDB.DeviceRegistry(&dev)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, ret)
}

// GET http://localhost:8080/devices
func (s StorageSvc) deviceList(req *restful.Request, res *restful.Response) {
	ret, err := s.SysDB.GetDevices(req.QueryParameter("data-type"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	for index, _ := range ret {
		if err := s.DynamicDB.GetStorage(&ret[index]); err != nil {
			log.Error(err)
		}
	}
	ResWriteEntity(res, &ret)
}

// 更新存储设备信息
func (s StorageSvc) updateDeviceInfo(req *restful.Request, res *restful.Response) {
	devId := req.PathParameter("device-id")

	var args UpdateDeviceRequest
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}

	if err := s.SysDB.UpdateDevice(devId, &args); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

// 移除存储设备
func (s StorageSvc) delDevice(req *restful.Request, res *restful.Response) {
	devId := req.PathParameter("device-id")
	if err := s.SysDB.DeleteDevice(devId); err != nil {
		ResWriteError(res, err)
		return
	}

	// TODO: 移除etcd中设备信息
	ResWriteEntity(res, nil)
}

func (s StorageSvc) deviceLoop() error {
	devices, err := s.SysDB.GetDevices("")
	if err != nil {
		return err
	}

	for i, _ := range devices {
		dev := &devices[i]
		var info *DeviceInfo
		var err error
		if dev.StorageType == common.NFSType {
			info, err = GetDeviceInfo(dev.MountPath)
		} else if dev.StorageType == common.DBType && dev.StorageOrg == common.POSTGRESQL {
			info, err = GetDeviceDBInfo(dev)
		} else if dev.StorageType == common.DBType && dev.StorageOrg == common.MONGODB {
			info, err = GetDeviceMongoInfo(dev)
		} else {
			dev.Used = "-1"
			dev.UsedPercent = "0%"
		}
		if err != nil {
			log.Println("Failed to get device space info")
			//dev.Volume = "-1"
			//dev.Free = "-1"
			dev.Used = "-1"
			dev.UsedPercent = "0%"
		} else {
			dev.FileSys = info.FileSystem
			dev.Volume = info.Volume
			dev.Used = info.Used
			dev.UsedPercent = info.UsedPercent
		}

		if err := s.DynamicDB.UpdateStorage(dev.Id, dev.Used, dev.UsedPercent); err != nil {
			return err
		}
	}
	return nil
}

func (s StorageSvc) DeviceLoop() {
	t := time.NewTicker(DefaultTimerInterval)
	for {
		select {
		case <-t.C:
			s.deviceLoop()
		}
	}
}

package storage

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"storage-manager/dbcentral/etcd"
	"storage-manager/dbcentral/pg"
	. "storage-manager/types"

	"grm-service/geoserver"
	. "grm-service/util"
)

type StorageSvc struct {
	SysDB     *pg.SystemDB
	DynamicDB *etcd.DynamicDB
	GeoServer *geoserver.GeoserverUtil

	DataDir   string
	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s StorageSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/devices").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("storage manager")}

	// 注册存储设备
	ws.Route(ws.POST("/").To(s.deviceRegistry).
		Doc(TR("device registry")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(registryDeviceRequest{}).
		Writes(Device{}))

	// 获取存储设备列表
	ws.Route(ws.GET("/").To(s.deviceList).
		Doc(TR("get device list")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("data-type", "device for data type").DataType("string")).
		Writes(DeviceList{}))

	// 修改存储设备信息
	ws.Route(ws.PUT("/{device-id}").To(s.updateDeviceInfo).
		Doc(TR("update device info")).
		Param(ws.PathParameter("device-id", "device id").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags).Reads(UpdateDeviceRequest{}))

	// 移除存储设备
	ws.Route(ws.DELETE("/{device-id}").To(s.delDevice).
		Doc(TR("update device info")).
		Param(ws.PathParameter("device-id", "device id").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags))

	return ws
}

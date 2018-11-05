package layers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"grm-service/common"
	"grm-service/geoserver"
	"grm-service/log"
	. "grm-service/util"

	"data-manager/types"

	"github.com/emicklei/go-restful"
	//"grm-service/log"
)

func (s DataLayerSvc) getLayers(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	dataId := req.PathParameter("data-id")
	// 获取数据信息
	name, _, err := s.MetaDB.GetDataPath(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	ret, err := s.SysDB.GetDataLayers(dataId, userId, name)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	if len(ret) > 0 {
		ResWriteHeaderEntity(res, &ret)
	} else {
		ResWriteHeaderEntity(res, nil)
	}
}

func (s DataLayerSvc) getLayer(req *restful.Request, res *restful.Response) {
	dataId := req.PathParameter("data-id")
	layerId := req.PathParameter("layer-id")

	// 获取数据信息
	name, _, err := s.MetaDB.GetDataPath(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	ret, err := s.SysDB.GetDataLayer(dataId, layerId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ret.DataName = name
	ResWriteHeaderEntity(res, ret)
}

func (s DataLayerSvc) addLayer(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	var args addLayerReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(args.Name) == 0 || len(userId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid dataset name or user id")))
		return
	}
	dataId := req.PathParameter("data-id")
	layer := common.DataLayer{
		Layer:       NewUUID(),
		Name:        args.Name,
		Data:        dataId,
		User:        userId,
		Style:       args.Style,
		Description: args.Description,
		IsDefault:   args.IsDefault,
	}

	// 发布geoserver 图层
	geoStorage, err := s.MetaDB.GetDeviceGeoStorage(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	srs, wmsUrl, wmtsUrl, wms, wfs, wmts, err := s.GeoServer.AddShpLayer(geoserver.GeoWorkSpace,
		geoStorage, dataId, layer.Layer)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	layer.Srs = srs
	layer.WmsUrl = wmsUrl
	layer.WmtsUrl = wmtsUrl
	layer.WMS = wms
	layer.Wfs = wfs
	layer.Wmts = wmts

	// 设置图层样式
	layerId := fmt.Sprintf("lyr-%s_%s", dataId, layer.Layer)
	var style string
	switch layer.Style {
	case "1":
		style = "point"
	case "2":
		style = "line"
	case "3":
		style = "polygon"
	default:
		style = userId + "_" + layer.Style
	}
	if len(layer.Style) > 0 {
		fmt.Println("layer:", layerId, ",style:", style)
		if err := s.GeoServer.SetLayerStyle(layerId, style); err != nil {
			ResWriteError(res, err)
			return
		}
	}

	ret, err := s.SysDB.AddDataLayer(&layer)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, ret)
}

func (s DataLayerSvc) delLayer(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	if err := s.SysDB.DelDataLayer(req.PathParameter("layer-id"), userId); err != nil {
		ResWriteError(res, err)
		return
	}
	name := fmt.Sprintf("lyr-%s_%s", req.PathParameter("data-id"), req.PathParameter("layer-id"))
	if err := s.GeoServer.DeleteShpLayer(name); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

func (s DataLayerSvc) updateLayer(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	dataId := req.PathParameter("data-id")
	layerId := req.PathParameter("layer-id")
	layer := fmt.Sprintf("lyr-%s_%s", dataId, layerId)

	var args types.UpdateLayerReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}

	// 编辑style
	if len(args.Style) > 0 {
		if err := s.GeoServer.SetLayerStyle(layer, args.Style); err != nil {
			ResWriteError(res, err)
			return
		}
	}

	// 更新pg
	if err := s.SysDB.UpdateDataLayer(dataId, layerId, userId, &args); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// 更新数据快试图
func (s DataLayerSvc) updateDataSnapShot(req *restful.Request, res *restful.Response) {
	// 数据路径
	dataId := req.PathParameter("data-id")
	layerId := req.PathParameter("layer-id")
	dataDir := filepath.Join(s.ConfigDir, "data", dataId)
	if err := CheckDir(dataDir); err != nil {
		log.Error("Failed to create data dir :", dataDir)
		ResWriteError(res, err)
		return
	}

	var thumbFile string
	var pic []byte
	var err error

	uploadType := req.QueryParameter("type")
	fmt.Println(uploadType)
	if uploadType == common.FileUpload {
		file, fh, err := req.Request.FormFile("snapshot")
		if err != nil {
			ResWriteError(res, err)
			return
		}
		defer file.Close()

		pic, err = ioutil.ReadAll(file)
		if err != nil {
			ResWriteError(res, err)
			return
		}
		thumbFile = filepath.Join(dataDir, fh.Filename)
	} else if uploadType == common.Base64Upload {
		var args types.SnapshotReq
		if err := req.ReadEntity(&args); err != nil {
			ResWriteError(res, err)
			return
		}
		if len(args.Image) == 0 {
			ResWriteError(res, fmt.Errorf(TR("Invalid snapshot image")))
			return
		}
		picbase64 := args.Image[strings.Index(args.Image, ",")+1:]
		pic, err = base64.StdEncoding.DecodeString(picbase64)
		if err != nil {
			ResWriteError(res, err)
			return
		}

		if len(args.FileName) > 0 {
			thumbFile = fmt.Sprintf("%s/%s", dataDir, args.FileName)
		} else {
			ext := args.Image[strings.Index(args.Image, "/")+1 : strings.Index(args.Image, ";")]
			thumbFile = fmt.Sprintf("%s/%s.%s", dataDir, layerId, ext)
		}
	}

	fmt.Println(thumbFile)
	if err := ioutil.WriteFile(thumbFile, pic, os.ModePerm); err != nil {
		ResWriteError(res, err)
		return
	}
	url := strings.Replace(thumbFile, s.ConfigDir, common.FilePre, -1)
	if err := s.SysDB.UpdateLayerSnap(layerId, url); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

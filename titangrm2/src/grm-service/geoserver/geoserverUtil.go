package geoserver

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"path/filepath"
	//	"strconv"
	"grm-service/util"
	"sort"
	"strings"
)

type GeoserverUtil struct {
	Lan []string
	//	Wan      string     wan地址在webui中进行设置，服务只提供相对路径
	UserName string
	Pwd      string
}

func (svc GeoserverUtil) AddWorkspace(name string) error {
	url := util.GetRandomStr(svc.Lan) + "/rest/workspaces?default=true"
	ws := Workspaces{&NameObject{name}}
	data, _ := json.Marshal(ws)
	resp, err := util.HttpAuthPost(url, data, svc.UserName, svc.Pwd)
	fmt.Println("AddWorkspace:", resp, err)
	return err
}

func (svc GeoserverUtil) AddPgStore(ws, name, host, port, database, user, passwd string) error {
	url := util.GetRandomStr(svc.Lan) + "/rest/workspaces/" + ws + "/datastores"

	ents := make([]*Entry, 0)
	ents = append(ents, &Entry{"host", host})
	ents = append(ents, &Entry{"port", port})
	ents = append(ents, &Entry{"database", database})
	ents = append(ents, &Entry{"user", user})
	ents = append(ents, &Entry{"passwd", passwd})
	ents = append(ents, &Entry{"dbtype", "postgis"})

	store := DataStores{&DataStore{name, &ConnectionParameters{ents}}}

	data, _ := json.Marshal(store)
	resp, err := util.HttpAuthPost(url, data, svc.UserName, svc.Pwd)
	fmt.Println("AddPgStore:", resp, err)
	return err
}

func (svc GeoserverUtil) AddShpLayer(ws, store, dataid, index string) (string, string, string, string, string, string, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/workspaces/" + ws + "/datastores/" + store + "/featuretypes"
	//fmt.Println(url)
	name := fmt.Sprintf("lyr-%s_%s", dataid, index)
	ftAdd := FeatureTypeJson{&FeatureType{Name: name, NativeName: dataid}}
	data, _ := json.Marshal(ftAdd)
	resp, err := util.HttpAuthPost(url, data, svc.UserName, svc.Pwd)
	fmt.Println("AddShpLayer:", resp, err)
	if err != nil {
		fmt.Println("AddShpLayerStore error:", err)
	}

	//http://192.168.1.189:8181/geoserver/TitanGRM/wms?service=WMS&version=1.1.0&request=GetMap&layers=TitanGRM:cityboundary&styles=&bbox=73.1794815063477,17.9404830932617,135.405303955078,53.7387275695801&width=768&height=441&srs=EPSG:404000&format=application/openlayers
	//获取这个图层的信息
	ft, err := svc.GetShapeLayerInfo(name)
	if err != nil {
		fmt.Println("AddShpLayer.GetShapeLayerInfo error:", err)
		return "", "", "", "", "", "", err
	}

	openlayerUrl := fmt.Sprintf("/%s/wms?service=WMS&version=1.1.0&request=GetMap&layers=%s&bbox=%f,%f,%f,%f&width=%d&height=%d&srs=%s&format=application/openlayers",
		ws, ws+":"+name, ft.FeatureType.NativeBoundingBox.Minx, ft.FeatureType.NativeBoundingBox.Miny,
		ft.FeatureType.NativeBoundingBox.Maxx, ft.FeatureType.NativeBoundingBox.Maxy, 754, 388, ft.FeatureType.Srs)

	//http://192.168.1.189:8181/grm/gwc/demo/TitanGRM:gf2?gridSet=EPSG:4326&format=image/png
	wmtsUrl := fmt.Sprintf("/gwc/demo/%s?gridSet=EPSG:4326&format=image/png", ws+":"+name)

	wms := fmt.Sprintf("/%s/wms?service=WMS&version=1.1.0&request=GetCapabilities&layers=%s:%s", ws, ws, name)
	wfs := fmt.Sprintf("/%s/wfs?service=WFS&version=1.0.0&request=GetCapabilities&layers=%s:%s", ws, ws, name)
	wmts := fmt.Sprintf("/gwc/service/wmts?REQUEST=getcapabilities")
	return ft.FeatureType.Srs, openlayerUrl, wmtsUrl, wms, wfs, wmts, nil
}

func (svc GeoserverUtil) AddShpLayerWithBBox(ws, store, dataid, index string, minx, maxx, miny, maxy float64, srs string) (string, string, string, string, string, string, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/workspaces/" + ws + "/datastores/" + store + "/featuretypes"
	//fmt.Println(url)
	var name string
	if len(index) == 0 {
		name = fmt.Sprintf("%s", dataid)
	} else {
		name = fmt.Sprintf("%s_%s", dataid, index)
	}

	ftAdd := FeatureTypeJson{&FeatureType{
		Name:       name,
		NativeName: dataid,
		Srs:        srs,
		NativeBoundingBox: &NativeBBox{
			Minx: minx,
			Maxx: maxx,
			Miny: miny,
			Maxy: maxy,
		},
		LatLonBoundingBox: &WGS84BBox{
			Minx: minx,
			Maxx: maxx,
			Miny: miny,
			Maxy: maxy,
		},
	}}
	data, _ := json.Marshal(ftAdd)
	resp, err := util.HttpAuthPost(url, data, svc.UserName, svc.Pwd)
	fmt.Println("AddShpLayer:", resp, err)
	if err != nil {
		fmt.Println("AddShpLayerStore error:", err)
	}

	//http://192.168.1.189:8181/geoserver/TitanGRM/wms?service=WMS&version=1.1.0&request=GetMap&layers=TitanGRM:cityboundary&styles=&bbox=73.1794815063477,17.9404830932617,135.405303955078,53.7387275695801&width=768&height=441&srs=EPSG:404000&format=application/openlayers
	//获取这个图层的信息
	ft, err := svc.GetShapeLayerInfo(name)
	if err != nil {
		fmt.Println("AddShpLayer.GetShapeLayerInfo error:", err)
		return "", "", "", "", "", "", err
	}

	openlayerUrl := fmt.Sprintf("/%s/wms?service=WMS&version=1.1.0&request=GetMap&layers=%s&bbox=%f,%f,%f,%f&width=%d&height=%d&srs=%s&format=application/openlayers",
		ws, ws+":"+name, ft.FeatureType.NativeBoundingBox.Minx, ft.FeatureType.NativeBoundingBox.Miny,
		ft.FeatureType.NativeBoundingBox.Maxx, ft.FeatureType.NativeBoundingBox.Maxy, 754, 388, ft.FeatureType.Srs)

	//http://192.168.1.189:8181/grm/gwc/demo/TitanGRM:gf2?gridSet=EPSG:4326&format=image/png
	wmtsUrl := fmt.Sprintf("/gwc/demo/%s?gridSet=EPSG:4326&format=image/png", ws+":"+name)

	wms := fmt.Sprintf("/%s/wms?service=WMS&version=1.1.0&request=GetCapabilities&layers=%s:%s", ws, ws, name)
	wfs := fmt.Sprintf("/%s/wfs?service=WFS&version=1.0.0&request=GetCapabilities&layers=%s:%s", ws, ws, name)
	wmts := fmt.Sprintf("/gwc/service/wmts?REQUEST=getcapabilities")
	return ft.FeatureType.Srs, openlayerUrl, wmtsUrl, wms, wfs, wmts, nil
}

func (svc GeoserverUtil) GetShapeLayerInfo(layer string) (*FeatureTypeJson, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/layers/" + layer
	//fmt.Println("URL:", url)
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	var layerinfo LayerJson
	err = json.Unmarshal([]byte(resp), &layerinfo)
	if err != nil {
		fmt.Println("GetShapeLayerInfo.resp:", resp)
		return nil, err
	}
	//	fmt.Println("shape layer url :", layerinfo.Layer.Resource.Href)
	respInfo, err := util.HttpAuthGet(layerinfo.Layer.Resource.Href, svc.UserName, svc.Pwd)
	var ft FeatureTypeJson
	err = json.Unmarshal([]byte(respInfo), &ft)
	if err != nil {
		fmt.Println("GetShapeLayerInfo.FeatureType.resp:", resp)
		return nil, err
	}
	return &ft, nil
}

func (svc GeoserverUtil) GetShapeLayerDefaultStyle(layer string) (string, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/layers/" + layer
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	var layerinfo LayerJson
	err = json.Unmarshal([]byte(resp), &layerinfo)
	if err != nil {
		fmt.Println("GetShapeLayerInfo.resp:", resp)
		return "", err
	}
	return layerinfo.Layer.DefaultStyle.Name, nil
}

func (svc GeoserverUtil) GetRasterLayerInfo(layer string) (*CoverageJson, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/layers/" + layer

	//	fmt.Println("URL:", url)
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	var layerinfo LayerJson
	err = json.Unmarshal([]byte(resp), &layerinfo)
	if err != nil {
		fmt.Println("GetRasterLayerInfo.resp:", resp)
		return nil, err
	}
	//	fmt.Println("raster layer url :", layerinfo.Layer.Resource.Href)
	respInfo, err := util.HttpAuthGet(layerinfo.Layer.Resource.Href, svc.UserName, svc.Pwd)
	var ft CoverageJson
	err = json.Unmarshal([]byte(respInfo), &ft)
	if err != nil {
		fmt.Println("GetRasterLayerInfo.Coverage.resp:", resp)
		return nil, err
	}
	return &ft, nil
}

func (svc GeoserverUtil) DeleteShpLayer(layer string) error {
	url := util.GetRandomStr(svc.Lan) + "/rest/layers/" + layer + "?recurse=true"
	resp, err := util.HttpAuthDelete(url, svc.UserName, svc.Pwd)
	fmt.Println("DeleteShpLayer:", resp, err)
	return nil
}

func (svc GeoserverUtil) GetAllStyles() ([]*Style, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/styles"
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	if err != nil {
		fmt.Println("GetAllStyle.HttpAuthGet error:", err)
		return nil, err
	}
	var styles StylesJson
	err = json.Unmarshal([]byte(resp), &styles)
	if err != nil {
		fmt.Println("GetAllStyle.json.Unmarshal error:", err)
		return nil, err
	}

	for _, s := range styles.Styles.Styles {
		//http://192.168.1.189:8181/grm/wms?REQUEST=GetLegendGraphic&VERSION=1.0.0&FORMAT=image/png&WIDTH=40&HEIGHT=40&STRICT=false&style=green
		s.Pic = fmt.Sprintf("/wms?REQUEST=GetLegendGraphic&VERSION=1.0.0&FORMAT=image/png&STRICT=false&style=%s",
			s.Name)
	}
	return styles.Styles.Styles, err
}

func (svc GeoserverUtil) GetLayerStyles(layer string) ([]*UserStyle, error) {
	styles := make([]*UserStyle, 0)
	url := fmt.Sprintf("%s/wms?SERVICE=WMS&REQUEST=GetStyles&VERSION=1.1.1&FORMAT=application/vnd.ogc.sld+xml&Layers=%s",
		util.GetRandomStr(svc.Lan), layer)
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	if err != nil {
		fmt.Println("GetLayerStyles.HttpAuthGet error:", err, resp)
		return styles, err
	}

	//解析sld
	//	fmt.Printf("GetLayerStyles.Resp:%+v\n", resp)

	var sldStruct StyledLayerDescriptor
	err = xml.Unmarshal([]byte(resp), &sldStruct)
	if err != nil {
		fmt.Println("GetLayerStyles.xml.Unmarshal error:", err, resp)
		return styles, nil
	}

	//	fmt.Printf("sldStruct:%+v\n", sldStruct)
	//	for _, s := range sldStruct.NamedLayer.UserStyle {
	//		if fmt.Sprintf("%v", s.FeatureTypeStyle.Rule.Filter.PropertyIsBetween) == "{}" {
	//			s.FeatureTypeStyle.Rule.Filter.PropertyIsBetween = nil
	//		}
	//	}

	sort.Sort(sldStruct.NamedLayer.UserStyle)

	//添加pic地址
	for _, s := range sldStruct.NamedLayer.UserStyle {
		//http://192.168.1.189:8181/grm/wms?REQUEST=GetLegendGraphic&VERSION=1.0.0&FORMAT=image/png&WIDTH=40&HEIGHT=40&STRICT=false&style=green
		s.Pic = fmt.Sprintf("/wms?REQUEST=GetLegendGraphic&VERSION=1.0.0&FORMAT=image/png&STRICT=false&style=%s",
			s.Name)
	}

	return sldStruct.NamedLayer.UserStyle, nil
}

func (svc GeoserverUtil) AddStyle(sldname, sld string) error {
	for _, geo_url := range svc.Lan {
		url := geo_url + "/rest/styles/"
		st := fmt.Sprintf("<style><name>%s</name><filename>%s.sld</filename></style>", sldname, sldname)
		resp, err := util.HttpAuthPostByContentType(url, []byte(st), svc.UserName, svc.Pwd,
			"application/xml; charset=UTF-8")
		if err != nil {
			fmt.Println("AddStyle error:", err, resp)
			return err
		}

		//	fmt.Printf("AddStyle.Post.Resp:%+v\n", resp)

		url = geo_url + "/rest/styles/" + sldname + "?raw=true"
		resp, err = util.HttpAuthPutByContentType(url, []byte(sld), svc.UserName, svc.Pwd,
			"application/vnd.ogc.sld+xml; charset=UTF-8")
		if err != nil {
			fmt.Println("EditStyle error:", err, resp)
			return err
		}
	}
	//	fmt.Printf("AddStyle.Put.Resp:%+v\n", resp)
	return nil
}

func (svc GeoserverUtil) EditStyle(sldname, sld string) error {
	for _, geo_url := range svc.Lan {
		url := geo_url + "/rest/styles/" + sldname + "?raw=true"
		resp, err := util.HttpAuthPutByContentType(url, []byte(sld), svc.UserName, svc.Pwd,
			"application/vnd.ogc.sld+xml; charset=UTF-8")
		if err != nil {
			fmt.Println("EditStyle error:", err, resp)
			return err
		}
	}
	return nil
}

func (svc GeoserverUtil) DeleteStyle(sldname string) error {
	for _, geo_url := range svc.Lan {
		url := geo_url + "/rest/styles/" + sldname + "?recurse=true"
		resp, err := util.HttpAuthDelete(url, svc.UserName, svc.Pwd)
		if err != nil {
			fmt.Println("DeleteStyle error:", err, resp)
			return err
		}
	}
	return nil
}

//func (svc GeoserverUtil) SetStyle(style, sld string) error {
//	for _, geo_url := range svc.Lan {
//		url := geo_url + "/rest/styles/" + style + ".xml"
//		resp, err := util.HttpAuthPutByContentType(url, []byte(sld), svc.UserName, svc.Pwd,
//			"application/vnd.ogc.sld+xml; charset=UTF-8")
//		if err != nil {
//			fmt.Println("SetStyle error:", err, resp)
//			return err
//		}
//	}
//	return nil
//}

func (svc GeoserverUtil) GetStyle(style string) ([]*UserStyle, error) {
	styles := make([]*UserStyle, 0)
	url := util.GetRandomStr(svc.Lan) + "/rest/styles/" + style
	resp, err := util.HttpAuthGetByContentType(url, svc.UserName, svc.Pwd,
		"application/json", "application/vnd.ogc.sld+xml")
	if err != nil {
		fmt.Println("GetStyle error:", err, resp)
		return styles, err
	}

	var sldStruct StyledLayerDescriptor
	err = xml.Unmarshal([]byte(resp), &sldStruct)
	if err != nil {
		fmt.Println("GetStyle.xml.Unmarshal error:", err, resp)
		return styles, nil
	}

	//	fmt.Printf("sldStruct:%+v\n", sldStruct)
	//	for _, s := range sldStruct.NamedLayer.UserStyle {
	//		if fmt.Sprintf("%v", s.FeatureTypeStyle.Rule.Filter.PropertyIsBetween) == "{}" {
	//			s.FeatureTypeStyle.Rule.Filter.PropertyIsBetween = nil
	//		}
	//	}

	sort.Sort(sldStruct.NamedLayer.UserStyle)

	//添加pic地址
	for _, s := range sldStruct.NamedLayer.UserStyle {
		//http://192.168.1.189:8181/grm/wms?REQUEST=GetLegendGraphic&VERSION=1.0.0&FORMAT=image/png&WIDTH=40&HEIGHT=40&STRICT=false&style=green
		s.Pic = fmt.Sprintf("/wms?REQUEST=GetLegendGraphic&VERSION=1.0.0&FORMAT=image/png&STRICT=false&style=%s",
			s.Name)
	}

	return sldStruct.NamedLayer.UserStyle, nil
}

func (svc GeoserverUtil) SetLayerStyle(layer, style string) error {
	fmt.Println("SetLayerStyle:", layer, style)
	for _, geo_url := range svc.Lan {
		var styles []*Style = make([]*Style, 0)
		ls, err := svc.GetLayerStyles(layer)
		if err != nil {
			fmt.Println("SetLayerStyle.GetLayerStyles error:", err)
		} else {
			var snames []string = make([]string, 0)
			for _, s := range ls {
				snames = append(snames, s.Name)
			}
			snames = util.RemoveDuplicatesAndEmpty(snames)
			for _, sn := range snames {
				styles = append(styles, &Style{Name: sn})
			}
		}

		styles = append(styles, &Style{Name: style})
		url := geo_url + "/rest/layers/" + layer
		lry := LayerJson{&Layer{
			DefaultStyle: &NameObject{style},
			Styles:       &StylesStruct{styles},
		}}
		data, err := json.Marshal(lry)
		if err != nil {
			fmt.Println("SetLayerStyle.json.Marshal error:", err, lry)
			return err
		}

		fmt.Println("data:", string(data))
		resp, err := util.HttpAuthPut(url, data, svc.UserName, svc.Pwd)
		if err != nil {
			fmt.Println("SetLayerStyle error:", err, resp)
			return err
		}
		fmt.Println("resp:", resp)
	}
	return nil
}

func (svc GeoserverUtil) GetLayerFields(layer string) (string, error) {
	///wfs?SERVICE=WFS&REQUEST=DescribeFeatureType&TYPENAME=geonode%3Astreetline
	url := util.GetRandomStr(svc.Lan) + "/wfs?SERVICE=WFS&REQUEST=DescribeFeatureType&TYPENAME=" + layer
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	if err != nil {
		fmt.Println("GetLayerFields error:", err, resp)
		return "", err
	}
	return resp, nil
}

func (svc GeoserverUtil) GetLayerStyle(layer string) (string, error) {
	///wms?SERVICE=WMS&SERVICE=WMS&VERSION=1.1.1&REQUEST=GetStyles&LAYERS=geonode%3Astreetline
	url := util.GetRandomStr(svc.Lan) + "/wms?SERVICE=WMS&VERSION=1.1.1&REQUEST=GetStyles&LAYERS=" + layer
	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
	if err != nil {
		fmt.Println("GetLayerStyle error:", err, resp)
		return "", err
	}
	return resp, nil
}

//func (svc GeoserverUtil) GetLayerCache(layer string) (string, error) {
//	url := svc.Address + "/gwc/rest/layers/" + layer
//	resp, err := util.HttpAuthGet(url, svc.UserName, svc.Pwd)
//	return resp, err
//}
//
//func (svc GeoserverUtil) SetLayerCache(layer, cache string) error {
//	url := svc.Address + "/gwc/rest/layers/" + layer
//	data, _ := xml.Marshal(cache)
//	resp, err := util.HttpAuthXmlPost(url, data, svc.UserName, svc.Pwd)
//	fmt.Println(resp, err)
//	return nil
//}

func (svc GeoserverUtil) AddTifStore(ws, name, tifFile string) error {
	url := util.GetRandomStr(svc.Lan) + "/rest/workspaces/" + ws + "/coveragestores"
	cs := CoverageStores{&CoverageStore{name, "GeoTIFF", "file://" + tifFile, true, &NameObject{ws}}}
	data, err := json.Marshal(cs)
	resp, err := util.HttpAuthPost(url, data, svc.UserName, svc.Pwd)
	fmt.Println("AddTifStore.resp:", resp)
	return err
}

func (svc GeoserverUtil) AddTifLayer(ws, store, name string) (string, string, string, string, string, error) {
	url := util.GetRandomStr(svc.Lan) + "/rest/workspaces/" + ws + "/coveragestores/" + store + "/coverages"

	fullname := name
	fileSuffix := filepath.Ext(name)
	if strings.ToLower(fileSuffix) == ".tif" || strings.ToLower(fileSuffix) == ".tiff" {
		name = strings.TrimSuffix(name, fileSuffix)
	}

	ftAdd := CoverageJson{&CoverageInfo{Name: name, NativeName: name}}
	data, err := json.Marshal(ftAdd)
	resp, err := util.HttpAuthPost(url, data, svc.UserName, svc.Pwd)
	fmt.Println("AddTifLayer.resp:", resp)
	if err != nil {
		fmt.Println("AddTifLayerStore error:", err)
	}

	//获取这个图层的信息
	//	fmt.Println("fullname:", fullname)
	ft, err := svc.GetRasterLayerInfo(fullname)
	if err != nil {
		fmt.Println("AddTifLayerStore.GetRasterLayerInfo error:", err)
		return "", "", "", "", "", err
	}

	//http://192.168.1.189:8181/geoserver/TitanGRM/wms?service=WMS&version=1.1.0&request=GetMap&layers=TitanGRM:1&bbox=114.798176,32.177711,117.473456,34.161067&width=1366&height=700&srs=EPSG:4326&format=application/openlayers

	openlayerUrl := fmt.Sprintf("/%s/wms?service=WMS&version=1.1.0&request=GetMap&layers=%s&bbox=%f,%f,%f,%f&width=%d&height=%d&srs=%s&format=application/openlayers",
		ws, ws+":"+name, ft.Coverage.NativeBoundingBox.Minx, ft.Coverage.NativeBoundingBox.Miny,
		ft.Coverage.NativeBoundingBox.Maxx, ft.Coverage.NativeBoundingBox.Maxy, 754, 388, ft.Coverage.Srs)

	//http://192.168.1.189:8181/grm/gwc/demo/TitanGRM:gf2?gridSet=EPSG:4326&format=image/png
	wmtsUrl := fmt.Sprintf("/gwc/demo/%s?gridSet=EPSG:4326&format=image/png", ws+":"+name)

	wms := fmt.Sprintf("/%s/wms?service=WMS&version=1.1.0&request=GetCapabilities&layers=%s", ws, ws+":"+name)
	wmts := fmt.Sprintf("/gwc/service/wmts?REQUEST=getcapabilities")
	return ft.Coverage.Srs, openlayerUrl, wmtsUrl, wms, wmts, nil
}

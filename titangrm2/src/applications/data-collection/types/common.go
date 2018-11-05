package types

//
const (
	PointType   = "point"
	PolygonType = "polygon"
)

// 数据发布信息
type DataPub struct {
	DataId     string `json:"data_id"`
	User       string `json:"user"`
	DataName   string `json:"data_name"`
	DataType   string `json:"data_type"`
	DataPath   string `json:"data_path"`
	PubUrl     string `json:"pub_url"`
	IsPub      bool   `json:"is_pub"`
	CreateTime string `json:"create_time"`
	WMS        string `json:"wms"`
	Wmts       string `json:"wmts"`
	Wfs        string `json:"wfs"`
	Srs        string `json:"srs"`
	WmsUrl     string `json:"wmsurl,omitempty"`
	WmtsUrl    string `json:"wmtsurl,omitempty"`
	IsCached   bool   `json:"is_cached"`
	Style      string `json:"style"`
}

// 存储设备
type Device struct {
	Id          int    `json:"id"`
	Label       string `json:"label"`
	StorageType string `json:"storage_type"`
	StorageOrg  string `json:"storage_org"`
	DateType    string `json:"data_types"`
	IpAddress   string `json:"ip_address"`
	DBPort      string `json:"db_port"`
	DBUser      string `json:"db_user"`
	DBPwd       string `json:"db_pwd"`
	DBGeoStore  string `json:"geo_store,omitempty"`
	FileSys     string `json:"file_sys,omitempty"`
	MountPath   string `json:"mount_path,omitempty"`
	Volume      string `json:"total_volume"`
	Used        string `json:"used"`
	UsedPercent string `json:"used_percent"`
	CreateTime  string `json:"create_time"`
	Description string `json:"description,omitempty"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type" description:"number,text,time,geom,area"` // 数字 文本 时间 经纬度	 面积
}

type Collection struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	User        string  `json:"user"`
	Description string  `json:"description"`
	CreateTime  string  `json:"create_time"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Type        string  `json:"type" description:"line/polygon"`
	Fields      []Field `json:"fields"`
}

type Colllections []Collection

type DataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DataFields []DataField

type Meta struct {
	Name     string      `json:"name"`
	Title    string      `json:"title"`
	Required interface{} `json:"required"`
	Value    interface{} `json:"value"`
	Type     string      `json:"type"`
	Classify interface{} `json:"classify"`
	Modified interface{} `json:"is_modified"`
	System   interface{} `json:"system"`
}

type Group struct {
	Group string  `json:"group"`
	Label string  `json:"label"`
	Value []*Meta `json:"value"`
}

type DataMeta struct {
	FullValid    interface{} `json:"full_valid"`
	MetaData     []*Group    `json:"metadata"`
	Type         string      `json:"type"`
	Label        string      `json:"label"`
	DisplayField string      `json:"display_field"`
}

// 更新元数据信息请求
type UpdateMetaRequest struct {
	DataId       string  `json:"data_id"`
	DataType     string  `json:"data_type"`
	Group        string  `json:"group"`
	Metas        []*Meta `json:"metas"`
	Tags         string  `json:"tags,omitempty"`
	Style        string  `json:"style,omitempty"`
	DisplayField string  `json:"display_field,omitempty"`
	AuthUser     string  `json:"user"`
}

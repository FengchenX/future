package types

import (
	"github.com/twpayne/go-geom/encoding/geojson"
	//	"time"
	"encoding/json"
)

// 数据类型
type DataType struct {
	Name        string      `json:"name" description:"name of datatype"`
	Label       string      `json:"label" description:"label of datatype"`
	Parent      string      `json:"parent" description:"parent of datatype"`
	IsObsoleted bool        `json:"is_obsoleted" description:"status of datatype"`
	Extensions  string      `json:"extension" description:"extensions of datatype"`
	CreateTime  string      `json:"create_time" description:"create_time of datatype"`
	Metas       interface{} `json:"metas,omitempty" description:"metas of datatype"`
	Description string      `json:"description" description:"dataset description"`
}

// 数据类型集合
type DataTypeList []DataType

// 元元数据
type TypeMeta struct {
	DataId   string           `json:"id" description:"id of data"`
	DataType string           `json:"data_type" description:"name of datatype"`
	SubType  string           `json:"sub_type" description:"name of perent type"`
	Label    string           `json:"label" description:"label of datatype"`
	Status   string           `json:"status" description:"status of datatype"`
	User     string           `json:"owner" description:"user id of the dataset"`
	Geometry *json.RawMessage `json:"geometry" description:"user id of the dataset"`

	SnapPath   string `json:"snapshot"`
	ThumbPath  string `json:"thumb"`
	CreateTime string `json:"create_time"`
	Detail     string `json:"detail"`
	DataUrl    string `json:"data_url"`

	ViewCount int64  `json:"view_cnt"`
	Tags      string `json:"tags"`

	BasicInformation *LabelMeta `json:"Basic Information" description:"metas of Basic Information"`
	Content          *LabelMeta `json:"Content" description:"metas of Content"`
	Cover            *LabelMeta `json:"Cover" description:"metas of Cover"`
	DataQuality      *LabelMeta `json:"Data quality" description:"metas of Data quality"`
	Distribute       *LabelMeta `json:"Distribute" description:"metas of Distribute"`
	Entity           *LabelMeta `json:"Entity" description:"metas of Entity"`
	Identification   *LabelMeta `json:"Identification" description:"metas of Identification"`
	Limit            *LabelMeta `json:"Limit" description:"metas of Limit"`
	Reference        *LabelMeta `json:"Reference" description:"metas of Reference"`
	RRUnit           *LabelMeta `json:"Reference and Responsible unit" description:"metas of Reference and Responsible unit"`
}

type LabelMeta struct {
	Label *string              `json:"label" description:"label"`
	Metas map[string]MetaValue `json:"metas" description:"metas"`
}

type MetaValue struct {
	Value    interface{} `json:"value"`
	Title    string      `json:"title"`
	Type     string      `json:"type"`
	Required bool        `json:"required"`
	ReadOnly bool        `json:"readonly"`
	Query    bool        `json:"queryable"`
	Classify bool        `json:"classify"`
	Modified bool        `json:"is_modified"`
}

type GroupMeta struct {
	Group  string      `json:"group" description:"name of group"`
	Label  string      `json:"label" description:"label of group"`
	Values []MetaValue `json:"value" description:"metas of group"`
}

// 数据集
type DataSet struct {
	ID          string `json:"id,omitempty" description:"identifier of the dataset"`
	Name        string `json:"name" description:"name of the dataset"`
	User        string `json:"user" description:"user id of the dataset"`
	Type        string `json:"type" description:"dataset type"`
	Description string `json:"description" description:"dataset description"`
}

// 数据集集合
type DataSetList []DataSet

/////////////////////////////////////

type FieldInfo struct {
	Name     string      `json:"field,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Operand  string      `json:"operand,omitempty"`
	IsTime   bool        `json:"is_time,omitempty"`
	MinValue string      `json:"min_value,omitempty"`
	MaxValue string      `json:"max_value,omitempty"`
}

type SearchInfo struct {
	Attrs    []FieldInfo       `json:"attrs,omitempty"`
	Geometry *geojson.Geometry `json:"geo_json,omitempty"`
	DataType string            `json:"data_type,omitempty"`
	DataId   string            `json:"data_id,omitempty"`
	//	DatasetShow bool              `json:"dataset_show,omitempty"`
	Only_geom  bool     `json:"only_geom,omitempty"`
	DatasetIds []string `json:"dataset_ids,omitempty"`

	Order  string `json:"order,omitempty"`
	Limit  string `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
	Sort   string `json:"sort,omitempty"`
	Key    string `json:"key,omitempty"`

	Userid string `json:"userid,omitempty"`
}

type History struct {
	Name       string     `json:"name" description:"name of history"`
	Id         string     `json:"id" description:"id of history"`
	SearchInfo SearchInfo `json:"body" description:"body of history"`
	Path       string     `json:"path" description:"path of history"`
	Userid     string     `json:"userid,omitempty"`
	DateTime   string     `json:"date_time,omitempty"`
}

// 数据类型集合
type HistoryList []History

type DataFilterRequest struct {
	GeoRegion string   `json:"geo_wkt,omitempty"`
	DataSets  []string `json:"data_sets,omitempty"`
	Datas     []string `json:"datas,omitempty"`
	UserId    string   `json:"user_id"`
}

type TableData struct {
	Total int   `json:"total"`
	Datas []Row `json:"table"`
}

type Row struct {
	Rows []interface{} `json:"row"`
}

type MetaInfo struct {
	Name      string `json:"name"`
	DataType  string `json:"data_type"`
	SubType   string `json:"sub_type"`
	TypeLabel string `json:"type_label,omitempty"`

	Path     string  `json:"path"`
	PathUrl  string  `json:"path_url,omitempty"`
	FileSize float64 `json:"file_size"`
	UUID     string  `json:"uuid"`

	Resolution      float64     `json:"resolution,omitempty"`
	Size            string      `json:"size,omitempty"`
	SRS             string      `json:"ref_system,omitempty"`
	EnvelopeGeoJson interface{} `json:"envelope_geojson,omitempty"`
	CentryPoint     interface{} `json:"centry_point,omitempty"`

	MetaJson string `json:"meta_json,omitempty"`

	DataTime    string `json:"data_time,omitempty"`
	ModifyTime  string `json:"modify_time,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	Detail      string `json:"detail,omitempty"`
	ShpType     string `json:"shp_type,omitempty"`
	FeatureNums int64  `json:"feature_nums,omitempty"`

	ReceiveTime  string `json:"receive_time,omitempty"`
	SatType      string `json:"sat_type,omitempty"`
	Sensor       string `json:"sensor,omitempty"`
	Band         string `json:"band,omitempty"`
	CloudPercent int    `json:"cloud_percent,omitempty"`

	SnapPath  string `json:"snapshot"`
	ThumbPath string `json:"thumb"`

	PublishUrl string `json:"publish_url"`
	IsPublish  bool   `json:"is_pub,omitempty"`
	Dataset    string `json:"dataset,omitempty"`

	LoadUser  string `json:"owner"`
	ViewCount int64  `json:"view_cnt"`
	Tags      string `json:"tags"`

	Style  string `json:"style,omitempty"`
	Status string `json:"status"`

	ModelPos  string `json:"model_pos,omitempty"`
	OriginPos string `json:"origin_pos,omitempty"`

	DisplayField string `json:"display_field,omitempty"`
	DataUrl      string `json:"data_url"`
}

type MetaInfosTotalReply struct {
	Total     int64       `json:"total"`
	MetaInfos []*MetaInfo `json:"metaInfos,omitempty"`
}

//type DatasetMetaInfosTotalReply struct {
//	Total     int64               `json:"total,omitempty"`
//	MetaInfos []*DatasetMetaInfos `json:"datasetMetaInfos,omitempty"`
//}

//type DatasetMetaInfos struct {
//	Dataset     string      `json:"dataset,omitempty"`
//	DatasetName string      `json:"dataset_name,omitempty"`
//	MetaInfos   []*MetaInfo `json:"metaInfos,omitempty"`
//}

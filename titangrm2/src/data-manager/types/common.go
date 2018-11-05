package types

import (
	"grm-service/common"
)

const (
	DataCreated    = "created"
	DataSubscribed = "subscribed"
	DataShared     = "shared"

	ClassMarket = "market"
	ClassUeser  = "user"
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
type Meta struct {
	Type      string      `json:"type" description:"name of datatype"`
	Label     string      `json:"label" description:"label of datatype"`
	DataMeta  []GroupMeta `json:"metadata" description:"metas of data type"`
	ChildMeta []Meta      `json:"childmeta,omitempty" description:"metas of child type"`
}

type MetaValue struct {
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	Value      interface{} `json:"value"`
	Type       string      `json:"type"`
	Required   bool        `json:"required"`
	ReadOnly   bool        `json:"readonly"`
	Query      bool        `json:"query"`
	Classify   bool        `json:"classify"`
	IsModified bool        `json:"is_modified"`
}

type GroupMeta struct {
	Group  string      `json:"group" description:"name of group"`
	Label  string      `json:"label" description:"label of group"`
	Values []MetaValue `json:"value" description:"metas of group"`
}

// 更新元元数据信息
type MetaFieldReq struct {
	Group    string `json:"group"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	ReadOnly bool   `json:"readonly"`
	//Query    bool   `json:"query"`
	Classify bool `json:"classify"`
}

// 数据集
type DataSet struct {
	ID          string `json:"id,omitempty" description:"identifier of the dataset"`
	Name        string `json:"name" description:"name of the dataset"`
	User        string `json:"user" description:"user id of the dataset"`
	Type        string `json:"type" description:"dataset type"`
	Class       string `json:"class" description:"dataset of market"`
	CreateTime  string `json:"create_time,omitempty"`
	Description string `json:"description" description:"dataset description"`
}

// 数据集集合
type DataSetList []DataSet

// 数据图层
type UpdateLayerReq struct {
	Name        string `json:"layer_name"`
	Style       string `json:"style_sld"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
}

// 数据信息
type DataInfo struct {
	DataId      string          `json:"data_id"`
	DataName    string          `json:"data_name"`
	DataType    string          `json:"data_type"`
	SubType     string          `json:"sub_type"`
	Path        string          `json:"path"`
	Size        string          `json:"size"`
	ViewCnt     string          `json:"view_cnt"`
	DataTime    string          `json:"data_time"`
	CreateTime  string          `json:"create_time"`
	DeleteTime  string          `json:"delete_time"`
	Status      string          `json:"status"`
	DataSource  string          `json:"data_source"`
	Abstract    string          `json:"abstract"`
	Description string          `json:"description"`
	SnapShot    string          `json:"snapshot"`
	Thumb       string          `json:"thumb"`
	Tags        string          `json:"tags"`
	Owner       common.UserInfo `json:"owner"`
	DataUrl     string          `json:"data_url"`
	Storage     string          `json:"storage"`
	Attribute   []interface{}   `json:"attribute"`
}

type SnapshotReq struct {
	Image    string `json:"image"`
	FileName string `json:"file_name,omitempty"`
}

// 更新数据信息
type UpdateDataInfoReq struct {
	Abstract    string `json:"abstract" description:"abstract or \"omit\""`
	Description string `json:"description" description:"description or \"omit\""`
	Tags        string `json:"tags" description:"tags or \"omit\""`
}

type UpdateMeta struct {
	Group string `json:"group"`
	Metas []struct {
		Field string      `json:"field"`
		Value interface{} `json:"value"`
	} `json:"fields"`
}

type UpdateDataMetaReq struct {
	Metas []UpdateMeta `json:"metas"`
}

// 表格内容
type Row struct {
	Rows []interface{} `json:"row"`
}

type TableData struct {
	Total int   `json:"total"`
	Datas []Row `json:"table"`
}

type DataSearch struct {
	DataId   string `json:"data_id,omitempty"`
	DataType string `json:"data_type"`

	Order  string `json:"order,omitempty"`
	Limit  string `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
	Sort   string `json:"sort,omitempty"`
	Key    string `json:"key,omitempty"`
}

// 子对象
type SubGroup struct {
	Count int64             `json:"count"`
	Datas map[string]string `json:"datas"`
}

type DataSubObj struct {
	Group map[string]*SubGroup `json:"groups"`
}

// 数据评论
type Comment struct {
	Id         string           `json:"id"`
	DataId     string           `json:"data_id"`
	CreateTime string           `json:"create_time"`
	ToUser     *common.UserInfo `json:"reply_user,omitempty"`
	FromUser   *common.UserInfo `json:"user"`
	Content    string           `json:"content"`
}

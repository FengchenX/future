package datas

import (
	"data-manager/types"
)

const (
	FeatureType   = "Feature"
	ShapeType     = "Shape"
	DocType       = "Document"
	MediaType     = "Media"
	OrthoType     = "Ortho"
	DEMType       = "DEM"
	AerialType    = "Aerial"
	CalTablelType = "CalTable"

	BasicGroup = "Basic Information"
)

type dataContent struct {
	Id        string            `json:"data_id"`
	Name      string            `json:"data_name"`
	DataType  string            `json:"data_type"`
	SubType   string            `json:"sub_type"`
	DataUrl   string            `json:"data_url,omitempty"`
	SubObj    *types.DataSubObj `json:"sub_object,omitempty"`
	TableData *types.TableData  `json:"table_data,omitempty"`
}

type addCommentReq struct {
	ToUser  string `json:"reply_user"`
	Content string `json:"content"`
}

type commentListResp struct {
	Cmts  []*types.Comment `json:"commnets"`
	Total int              `json:"total"`
}

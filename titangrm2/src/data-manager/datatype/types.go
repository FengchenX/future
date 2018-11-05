package datatype

// 更新数据类型信息
type updateTypeInfoReq struct {
	Label       string `json:"label" description:"label of datatype"`
	IsObsoleted bool   `json:"is_obsoleted" description:"status of datatype"`
	Extensions  string `json:"extension" description:"extensions of datatype"`
	Description string `json:"description" description:"dataset description"`
}

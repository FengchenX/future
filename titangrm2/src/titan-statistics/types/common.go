package types

import ()

type TypeStat struct {
	DataType string `json:"data_type"`
	Count    int    `json:"count"`
}

type TypeStatList []TypeStat

type SubtypeStat struct {
	SubType string `json:"sub_type"`
	Count   int    `json:"count"`
}

type SubtypeList struct {
	DataType     string        `json:"data_type"`
	TypeStatList []SubtypeStat `json:"sub_types,omitempty"`
}

type TypeAggr struct {
	DataType string     `json:"data_type,omitempty"`
	Aggr     []AggrAttr `json:"aggr"`
}

type AggrAttr struct {
	AttrTitle string `json:"attr_title,omitempty"`
	AttrName  string `json:"attr_name,omitempty"`
	AttrType  string `json:"attr_type,omitempty"`
	AttrSize  string `json:"attr_size,omitempty"`
}

type DistinctValues []string

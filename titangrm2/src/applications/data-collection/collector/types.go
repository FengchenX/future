package collector

import (
	"applications/data-collection/types"
)

type addCollection struct {
	Name        string        `json:"name"`
	StartTime   string        `json:"start_time"`
	EndTime     string        `json:"end_time"`
	Description string        `json:"description"`
	Type        string        `json:"type" description:"point/polygon"`
	Fields      []types.Field `json:"fields"`
}

type updateCollection struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type addColDataReq struct {
	Datas types.DataFields `json:"datas"`
}

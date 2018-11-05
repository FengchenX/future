package common

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// 分页筛选参数
type PageFilter struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
}

// 数据图层
type DataLayer struct {
	Layer       string `json:"layer_id"`
	Name        string `json:"layer_name"`
	Data        string `json:"data_id"`
	DataName    string `json:"data_name"`
	User        string `json:"user_id"`
	Style       string `json:"style"`
	Description string `json:"description"`
	SnapShot    string `json:"snapshot"`
	CreateTime  string `json:"create_time"`
	IsDefault   bool   `json:"is_default"`
	Srs         string `json:"srs"`
	WMS         string `json:"wms"`
	Wmts        string `json:"wmts"`
	Wfs         string `json:"wfs"`
	WmsUrl      string `json:"wmsurl,omitempty"`
	WmtsUrl     string `json:"wmtsurl,omitempty"`
}

// 图层样式
type LayerStyle struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	User        string      `json:"user"`
	Sld         interface{} `json:"sld"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	CreateTime  string      `json:"create_time"`
}

// 用户概要信息
type UserInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

package theme

import "tile-manager/types"

type addThemeReq struct {
	//Id             string  `json:"id"`
	Name           string  `json:"name"`
	ImageType      string  `json:"image_type"`
	TileSize       uint32  `json:"tile_size"`
	TileFormat     string  `json:"tile_format"`
	Srs            string  `json:"srs"`
	NoData         float32 `json:"no_data"`
	TimeResolution string  `json:"time_resolution"`
	Device         int     `json:"device"`
	Description    string  `json:"description"`
	Transparency   string  `json:"transparency"`
	//AuthUser       string  `json:"user"`
	Session    string `json:"session"`
	Projection string `json:"projection"`
}
type addThemeResp struct {
	Id int `json:"id"`
}
type updateThemeReq struct {
	Name           string  `json:"name"`
	ImageType      string  `json:"image_type"`
	TileSize       uint32  `json:"tile_size"`
	TileFormat     string  `json:"tile_format"`
	Srs            string  `json:"srs"`
	NoData         float32 `json:"no_data"`
	TimeResolution string  `json:"time_resolution"`
	StorageType    string  `josn:"storage_type"`
	DataUrl        string  `json:"data_url"`
	DBUser         string  `json:"db_user"`
	DBPwd          string  `json:"db_pwd"`
	GmsdLan        string  `json:"gms_lan"`
	GmsdWan        string  `json:"gms_wan"`
	Description    string  `json:"description"`
	//AuthUser       string  `json:"user"`
	Transparency string `json:"transparency"`
	Projection   string `json:"projection"`

	Session string
}

type updateThemePicReq struct {
	Picture string `json:"picture"`
	//AuthUser string `json:"user"`
}

type getThemesResp struct {
	Total  int
	Themes []types.Theme
}

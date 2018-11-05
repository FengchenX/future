package types

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
	//Free        string `json:"free_volume"`
	Used        string `json:"used"`
	UsedPercent string `json:"used_percent"`
	CreateTime  string `json:"create_time"`
	Description string `json:"description,omitempty"`
}

type Theme struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	ImageType      string  `json:"image_type"`
	TileSize       uint32  `json:"tile_size"`
	TileFormat     string  `json:"tile_format"`
	Srs            string  `json:"srs"`
	NoData         float32 `json:"no_data"`
	TimeResolution string  `json:"time_resolution"`
	AuthUser       string  `json:"user"`
	CreateTime     string  `json:"create_time"`
	Description    string  `json:"description,omitempty"`
	Transparency   string  `json:"transparency"`
	Thumb          string  `json:"picture"`
	DeviceId       int     `json:"device_id"`
	DeviceInfo     Device  `json:"device"`
	Projection     string  `json:"projection"`

	Center string `json:"Center"`

	Session string
}

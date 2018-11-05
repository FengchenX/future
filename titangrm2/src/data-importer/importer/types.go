package importer

const (
	ProcessMsg  = "Process"
	ProgressMsg = "Progress"
	StatusMsg   = "Status"
	FeatureMsg  = "Feature"
	OfficeMsg   = "Office"
)

// FileSysDomain
type domain struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Domain string `json:"domain"`
	IsDir  bool   `json:"isParent"`
}

type sysDomains []domain

type domainReq struct {
	Path string `json:"path,omitempty"  description:"domain path"`
}

// 数据扫描
type dataScanRequest struct {
	DataType string `json:"data_type"`
	PreType  string `json:"pre_type"`
	DataDir  string `json:"data_dir"`
	Device   int    `json:"storage_device"`
	DataSet  string `json:"data_set"`
	//Config   int    `json:"config,omitempty"`
	//IsCopy   bool   `json:"is_copy"`
	TaskName string `json:"task_name"`
	//GeoType  string `json:"geo_type"`
}

type dataScanReply struct {
	JobId string `json:"job_id"`
}

// 数据入库
type dataLoadRequest struct {
	ScanTask    string   `json:"scan_task"`
	AllFileLoad bool     `json:"load_all"`
	FileIds     []string `json:"file_ids"`
}

type dataLoadReply struct {
	JobId string `json:"job_id"`
}

// 数据上传
type dataUploadRequest struct {
	DataType string `json:"data_type"`
	PreType  string `json:"pre_type"`
	DataDir  string `json:"data_dir"`
	Device   string `json:"storage_device"`
	DataSet  string `json:"data_set"`
	//GeoType  string `json:"geo_type"`
}

package types

import (
	"errors"

	. "grm-service/util"
)

var (
	ErrDataSetIdNULL      = errors.New(TR("id of new dataset is null"))
	ErrScanDataDirInvalid = errors.New(TR("scan data dir is empty or not exists"))
	ErrInvalidDataType    = errors.New(TR("invalid data type"))
	ErrInvalidDataSet     = errors.New(TR("invalid dataset"))
	ErrTaskIdNULL         = errors.New(TR("id of task is null"))
	ErrScanTaskIdInvalid  = errors.New(TR("invalid scan task id"))
	ErrNoDataLoad         = errors.New(TR("no selected data for load"))
	ErrInvalidDataInfo    = errors.New(TR("invalid data info"))
)

var (
	ScanTask     = "scan"
	LoadTask     = "load"
	UploadTask   = "upload"
	DispatchTask = "dispatch"

	CmdExec = "DataWorker"
	CmdPath = "DataWorker"

	TaskIdle       = "Idle"
	TaskRunning    = "Running"
	TaskFinished   = "Finished"
	TaskTerminated = "Terminated"
)

// task信息
type TaskInfo struct {
	TaskType     string `json:"task_type"`
	TaskId       string `json:"task_id"`
	TaskName     string `json:"task_name"`
	DataType     string `json:"data_type"`
	PreType      string `json:"pre_type"`
	Status       string `json:"status"`
	Progress     string `json:"progress"`
	Process      string `json:"process,omitempty"`
	StartTime    string `json:"start_time"`
	FinishedTime string `json:"finished_time"`
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`

	ScanDir   string `json:"scan_dir,omitempty"`
	DataSet   string `json:"data_set,omitempty"`
	Device    int    `json:"device,omitempty"`
	IsCopy    bool   `json:"is_copy,omitempty"`
	AllLoaded bool   `json:"all_loaded"`
}

type TaskList struct {
	Total int64       `json:"total"`
	Tasks []*TaskInfo `json:"rows"`
}

// 扫描任务结果筛选
type ResultFilter struct {
	FileName      string `json:"file"`
	FileSizeMin   string `json:"size_min"`
	FileSizeMax   string `json:"size_max"`
	CreateTimeMin string `json:"create_time_min"`
	CreateTimeMax string `json:"create_time_max"`
	//ResolutionMin float64 `json:"resolution_min"`
	//ResolutionMax float64 `json:"resolution_max"`
	//RefSystem     string  `json:"ref_system"`
	//Pyramid       string  `json:"has_pyramid"`
}

type ScanResults struct {
	Total int          `json:"total"`
	Datas []ScanResult `json:"rows,omitempty"`
}

// 扫描结果
type ScanResult struct {
	FileId     string  `json:"uuid"`
	FileName   string  `json:"name"`
	FilePath   string  `json:"path"`
	FileSize   string  `json:"file_size"`
	FileType   string  `json:"data_type"`
	TypeLabel  string  `json:"type_label"`
	SubType    string  `json:"sub_type"`
	SubLabel   string  `json:"sub_label"`
	Meta       string  `json:"meta,omitempty"`
	Tags       string  `json:"tags"`
	Resolution float64 `json:"resolution,omitempty"`
	RefSystem  string  `json:"ref_system,omitempty"`
	Pyramid    string  `json:"has_pyramid,omitempty"`
	CreateTime string  `json:"create_time"`
}

package tasks

/*
type TaskInfo struct {
	TaskId       string `json:"task_id"`
	TaskType     string `json:"task_type"`
	TaskName     string `json:"task_name"`
	DataType     string `json:"data_type"`
	Status       string `json:"status"`
	Progress     string `json:"progress"`
	Process      string `json:"process"`
	StartTime    string `json:"start_time"`
	FinishedTime string `json:"finished_time"`
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	AllLoaded    bool   `json:"all_loaded"`
}

type TaskList struct {
	Total int         `json:"total"`
	Tasks []*TaskInfo `json:"rows"`
}

*/

type TaskLogReply struct {
	Log string `json:"log"`
}

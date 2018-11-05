package tasks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"grm-service/log"

	"github.com/emicklei/go-restful"

	"data-importer/executor"
	"data-importer/types"
	. "grm-service/util"
)

// GET http://localhost:8080/tasks/{task-type}
func (s TasksSvc) getTask(req *restful.Request, res *restful.Response) {
	taskType := req.PathParameter("task-type")
	if len(taskType) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid task type")))
		return
	}

	page := ParserPageArgs(req)
	ret, err := s.DynamicDB.GetTasks(taskType)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 分页
	offset, _ := strconv.Atoi(page.Offset)
	limit, _ := strconv.Atoi(page.Limit)
	if offset >= 0 && limit > 0 {
		if offset+limit >= len(ret.Tasks)-1 {
			ret.Tasks = ret.Tasks[offset:]
		} else {
			ret.Tasks = ret.Tasks[offset : offset+limit]
		}
	}

	// 排序
	taskSort(ret.Tasks, page.Sort, page.Order)

	// 判断当前任务是否有未导入数据
	if taskType == types.ScanTask {
		for index, task := range ret.Tasks {
			result, err := s.MetaDB.GetScanDatas(task.TaskId, "", nil, nil, true)
			if err != nil {
				ResWriteError(res, err)
				return
			}
			if result.Total == 0 {
				ret.Tasks[index].AllLoaded = true
			}
		}
	}
	ResWriteHeaderEntity(res, ret)
}

// PUT http://localhost:8080/tasks/{task-id}/status
func (s TasksSvc) terminateTask(req *restful.Request, res *restful.Response) {
	taskId := req.PathParameter("task-id")
	if len(taskId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid task id")))
		return
	}
	taskType := taskId[:strings.Index(taskId, "_")]

	pid, err := s.DynamicDB.GetTaskPid(taskType, taskId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	// 结束进程
	if err := executor.KillProcess(pid); err != nil {
		log.Error("Failed to terminated task:", taskId)
		ResWriteError(res, err)
		return
	}
	// 修改状态信息
	if err := s.DynamicDB.UpdateTaskStatus(taskType, taskId, types.TaskTerminated); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// DELETE http://localhost:8080/tasks/{task-id}
func (s TasksSvc) deleteTask(req *restful.Request, res *restful.Response) {
	taskType := req.PathParameter("task-type")
	if len(taskType) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid task type")))
		return
	}
	taskId := req.QueryParameter("task-id")
	start_time := req.QueryParameter("start-time")
	end_time := req.QueryParameter("end-time")
	if err := s.DynamicDB.DelTask(taskType, taskId, start_time, end_time); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// GET http://localhost:8080/tasks/{task-id}/logs/{log-type}
func (s TasksSvc) getTaskLog(req *restful.Request, res *restful.Response) {
	taskId := req.PathParameter("task-id")
	if len(taskId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid task id")))
		return
	}
	taskType := taskId[:strings.Index(taskId, "_")]
	logType := req.PathParameter("log-type")
	if len(logType) == 0 {
		logType = "stdout"
	}

	logFile := filepath.Join(s.ConfigDir, "dataworker", taskType, taskId, "log", logType+".log")
	data, err := ioutil.ReadFile(logFile)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, &TaskLogReply{string(data)})
}

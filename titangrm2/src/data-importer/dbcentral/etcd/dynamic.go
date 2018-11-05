package etcd

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	//"github.com/coreos/etcd/clientv3"
	//github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	"grm-service/util"

	"data-importer/types"

	"github.com/coreos/etcd/clientv3"
)

type DynamicDB struct {
	DynamicEtcd
}

// 初始化任务信息
func (e DynamicDB) InitTaskInfo(taskType, taskId string, args map[string]string) (string, error) {
	if len(taskId) == 0 {
		taskId = util.NewUUID()
	}
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_DATAWORKER, taskType, taskId)
	for k, v := range args {
		if _, err := e.Cli.Put(context.Background(), key+"/"+k, v); err != nil {
			log.Error("Failed to init task info(%s) : %s", key+"/"+k, err)
			return "", err
		}
	}
	return taskId, nil
}

// 更新task进度
func (e DynamicDB) UpdateTaskProgress(taskType, taskId, progress string) error {
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_DATAWORKER, taskType, taskId)

	// 判定是否出现进度变小的情况
	resp, err := e.Cli.Get(context.Background(), key+"/progress")
	if err == nil && len(resp.Kvs) > 0 {
		var lastProgress int
		val := string(resp.Kvs[0].Value)
		if len(val) > 0 {
			pro, err := strconv.Atoi(val[:len(val)-1])
			if err == nil {
				lastProgress = pro
			}
			if curProgress, err := strconv.Atoi(progress[:len(progress)-1]); err == nil {
				if curProgress <= lastProgress {
					return nil
				}
			}
		}
	}

	// 判定状态
	resp, err = e.Cli.Get(context.Background(), key+"/status")
	if err == nil && len(resp.Kvs) > 0 {
		status := string(resp.Kvs[0].Value)
		if status == types.TaskFinished || status == types.TaskTerminated {
			return nil
		}
	}

	// 更新进度
	if _, err := e.Cli.Put(context.Background(), key+"/progress", progress); err != nil {
		return err
	}
	return nil
}

// 更新task状态
func (e DynamicDB) UpdateTaskStatus(taskType, taskId, status string) error {
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_DATAWORKER, taskType, taskId)

	// 更新pid
	if index := strings.Index(status, "("); index != -1 {
		pid := status[strings.Index(status, "(")+1 : strings.Index(status, ")")]
		if _, err := e.Cli.Put(context.Background(), key+"/pid", pid); err != nil {
			return err
		}
		status = status[:strings.Index(status, "(")]
	}

	// 更新状态
	fmt.Println(key + "/status")
	if _, err := e.Cli.Put(context.Background(), key+"/status", status); err != nil {
		return err
	}

	// 更新进度和完成时间
	if status == types.TaskFinished {
		if _, err := e.Cli.Put(context.Background(), key+"/progress", "100%"); err != nil {
			return err
		}
		if _, err := e.Cli.Put(context.Background(), key+"/finished_time", time.Now().Format("2006-01-02 15:04:05")); err != nil {
			return err
		}
	}
	return nil
}

// 更新task当前处理文件
func (e DynamicDB) UpdateTaskProcess(taskType, taskId, msg string) error {
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_DATAWORKER, taskType, taskId)

	// 判定状态
	resp, err := e.Cli.Get(context.Background(), key+"/status")
	if err == nil && len(resp.Kvs) > 0 {
		status := string(resp.Kvs[0].Value)
		if status == types.TaskFinished || status == types.TaskTerminated {
			return nil
		}
	}

	if _, err := e.Cli.Put(context.Background(), key+"/process", msg); err != nil {
		return err
	}
	return nil
}

// 获取task信息
func (e DynamicDB) GetTaskInfo(taskType, taskId string) (*types.TaskInfo, error) {
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_DATAWORKER, taskType, taskId)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	task := types.TaskInfo{TaskType: taskType, TaskId: taskId}
	for _, ev := range resp.Kvs {
		key := string(ev.Key)
		key = key[strings.LastIndex(key, "/")+1:]
		value := string(ev.Value)

		switch key {
		case "data_type":
			task.DataType = value
		case "pre_type":
			task.PreType = value
		case "finished_time":
			task.FinishedTime = value
		case "progress":
			task.Progress = value
		case "process":
			task.Process = value
		case "start_time":
			task.StartTime = value
		case "status":
			task.Status = value
		case "task_name":
			task.TaskName = value
		case "user_id":
			task.UserId = value
		case "scan_dir":
			task.ScanDir = value
		case "data_set":
			task.DataSet = value
		case "device":
			{
				if dev, err := strconv.Atoi(string(value)); err == nil {
					task.Device = dev
				}
			}
		case "is_copy":
			task.IsCopy = (string(value) == "true")
		}
	}
	return &task, nil
}

// 获取task pid
func (e DynamicDB) GetTaskPid(taskType, taskId string) (int, error) {
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_DATAWORKER, taskType, taskId)
	resp, err := e.Cli.Get(context.Background(), key+"/pid")
	if err != nil {
		return -1, err
	}

	if len(resp.Kvs) > 0 {
		return strconv.Atoi(string(resp.Kvs[0].Value))
	}
	return -1, nil
}

// 获取指定类型task任务列表
func (e DynamicDB) GetTasks(taskType string) (*types.TaskList, error) {
	var ret types.TaskList
	key := fmt.Sprintf("%s/%s", KEY_GRM_DATAWORKER, taskType)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	nodes := make(map[string]*types.TaskInfo)
	for _, kv := range resp.Kvs {
		//fmt.Printf("%s : %s\n", kv.Key, kv.Value)
		value := string(kv.Value)

		keys := strings.Split(string(kv.Key), KEY_SPLIT)
		taskId := keys[len(keys)-2]
		if task, ok := nodes[taskId]; !ok || task == nil {
			nodes[taskId] = &types.TaskInfo{TaskType: taskType, TaskId: taskId}
			ret.Total++
		}

		switch keys[len(keys)-1] {
		case "data_type":
			nodes[taskId].DataType = value
		case "pre_type":
			nodes[taskId].PreType = value
		case "finished_time":
			nodes[taskId].FinishedTime = value
		case "progress":
			nodes[taskId].Progress = value
		//case "process":
		//nodes[taskId].Process = value
		case "start_time":
			nodes[taskId].StartTime = value
		case "status":
			nodes[taskId].Status = value
		case "task_name":
			nodes[taskId].TaskName = value
		case "user_id":
			nodes[taskId].UserId = value
		case "scan_dir":
			nodes[taskId].ScanDir = value
		case "data_set":
			nodes[taskId].DataSet = value
		case "device":
			{
				if dev, err := strconv.Atoi(string(value)); err == nil {
					nodes[taskId].Device = dev
				}
			}
		case "is_copy":
			nodes[taskId].IsCopy = (string(value) == "true")
		}
	}

	for _, v := range nodes {
		userName, err := e.GetUserName(v.UserId)
		if err == nil {
			v.UserName = userName
		}
		ret.Tasks = append(ret.Tasks, v)
	}
	return &ret, nil
}

// 移除task记录
func (e DynamicDB) DelTask(taskType, taskId, start, end string) error {
	key := fmt.Sprintf("%s/%s", KEY_GRM_DATAWORKER, taskType)
	if len(taskId) > 0 {
		_, err := e.Cli.Delete(context.Background(), key+"/"+taskId, clientv3.WithPrefix())
		if err != nil {
			return err
		}
		return nil
	}

	format := "2006-01-02 15:04:05"
	var startTime, endTime time.Time
	var err error

	if len(start) == 0 {
		start = "1976-01-01 00:00:00"
	} else {
		start = fmt.Sprintf("%s 00:00:00", start)
	}

	if startTime, err = time.Parse(format, start); err != nil {
		return err
	}
	end = fmt.Sprintf("%s 23:59:59", end)
	if endTime, err = time.Parse(format, end); err != nil {
		return err
	}
	fmt.Println("start:", startTime)
	fmt.Println("end:", endTime)

	tasks, err := e.GetTasks(taskType)
	if err != nil {
		return err
	}
	for _, task := range tasks.Tasks {
		t, err := time.Parse(format, task.StartTime)
		if err != nil {
			fmt.Println("Failed to get task start time: ", task.TaskId)
			e.DelTask(taskType, task.TaskId, "", "")
			continue
		}
		fmt.Println("task time:", t)
		if (t.After(startTime) && t.Before(endTime)) ||
			t.Equal(endTime) || t.Equal(startTime) {
			if err := e.DelTask(taskType, task.TaskId, "", ""); err != nil {
				return err
			}
		}
	}
	return nil
}

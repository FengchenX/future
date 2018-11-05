package tasks

import (
	"sort"

	"data-importer/types"
)

// 任务排序
type TasklList []*types.TaskInfo

func (s TasklList) Len() int { return len(s) }

func (s TasklList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func taskSort(tasks []*types.TaskInfo, sortArgs, order string) {
	switch sortArgs {
	case "task_name":
		{
			if order == "desc" {
				sort.Sort(ByTaskNameDesc{tasks})
			} else {
				sort.Sort(sort.Reverse(ByTaskNameDesc{tasks}))
			}
		}
	case "status":
		{
			if order == "desc" {
				sort.Sort(ByTaskStatusDesc{tasks})
			} else {
				sort.Sort(sort.Reverse(ByTaskStatusDesc{tasks}))
			}
		}
	case "progress":
		{
			if order == "desc" {
				sort.Sort(ByTaskProgressDesc{tasks})
			} else {
				sort.Sort(sort.Reverse(ByTaskProgressDesc{tasks}))
			}
		}
	case "start_time":
		{
			if order == "desc" {
				sort.Sort(ByStartTimeDesc{tasks})
			} else {
				sort.Sort(sort.Reverse(ByStartTimeDesc{tasks}))
			}
		}
	case "finished_time":
		{
			if order == "desc" {
				sort.Sort(ByFinishedTimeDesc{tasks})
			} else {
				sort.Sort(sort.Reverse(ByFinishedTimeDesc{tasks}))
			}
		}
	case "user_name":
		{
			if order == "desc" {
				sort.Sort(ByUserNameDesc{tasks})
			} else {
				sort.Sort(sort.Reverse(ByUserNameDesc{tasks}))
			}
		}
	}
}

// 任务名
type ByTaskNameDesc struct{ TasklList }

func (s ByTaskNameDesc) Less(i, j int) bool {
	return s.TasklList[i].TaskName > s.TasklList[j].TaskName
}

// 状态
type ByTaskStatusDesc struct{ TasklList }

func (s ByTaskStatusDesc) Less(i, j int) bool {
	return s.TasklList[i].Status > s.TasklList[j].Status
}

// 进度
type ByTaskProgressDesc struct{ TasklList }

func (s ByTaskProgressDesc) Less(i, j int) bool {
	return s.TasklList[i].Progress > s.TasklList[j].Progress
}

// 用户名
type ByUserNameDesc struct{ TasklList }

func (s ByUserNameDesc) Less(i, j int) bool {
	return s.TasklList[i].UserName > s.TasklList[j].UserName
}

// 开始时间
type ByStartTimeDesc struct{ TasklList }

func (s ByStartTimeDesc) Less(i, j int) bool {
	return s.TasklList[i].StartTime > s.TasklList[j].StartTime
}

// 结束时间
type ByFinishedTimeDesc struct{ TasklList }

func (s ByFinishedTimeDesc) Less(i, j int) bool {
	return s.TasklList[i].FinishedTime > s.TasklList[j].FinishedTime
}

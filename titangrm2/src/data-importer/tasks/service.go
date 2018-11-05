package tasks

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"data-importer/dbcentral/etcd"
	"data-importer/dbcentral/pg"
	"data-importer/types"

	"grm-service/mq"
	. "grm-service/util"
)

type TasksSvc struct {
	SysDB     *pg.SystemDB
	MetaDB    *pg.MetaDB
	DynamicDB *etcd.DynamicDB
	MsgQueue  *mq.RabbitMQ

	DataDir   string
	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s TasksSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/tasks")
	//ws.Consumes(restful.MIME_JSON, restful.MIME_XML)
	ws.Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data importer tasks")}

	// 获取任务列表
	ws.Route(ws.GET("/{task-type}").To(s.getTask).
		Doc(TR("get task(s) info")).
		Param(ws.PathParameter("task-type", TR("task type"))).
		Param(ws.QueryParameter("limit", "limit").DataType("string").Required(false)).
		Param(ws.QueryParameter("offset", "offset").DataType("string").Required(false)).
		Param(ws.QueryParameter("sort", "sort").DataType("string").Required(false)).
		Param(ws.QueryParameter("order", "order").DataType("string").Required(false)).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(types.TaskList{}))

	// 中止任务
	ws.Route(ws.PUT("/{task-id}/status").To(s.terminateTask).
		Doc(TR("terminate task by id")).
		Param(ws.PathParameter("task-id", TR("task id"))).
		Metadata(api.KeyOpenAPITags, tags))

	// 移除任务
	ws.Route(ws.DELETE("/{task-type}").To(s.deleteTask).
		Doc(TR("delete task by task id or a range of time")).
		Param(ws.PathParameter("task-type", TR("task type")).Required(true)).
		Param(ws.QueryParameter("task-id", TR("task id")).Required(false)).
		Param(ws.QueryParameter("start-time", TR("start time")).Required(false)).
		Param(ws.QueryParameter("end-time", TR("end time")).Required(false)).
		Metadata(api.KeyOpenAPITags, tags))

	// 获取任务输出
	ws.Route(ws.GET("/{task-id}/logs/{log-type}").To(s.getTaskLog).
		Doc(TR("get task logs")).
		Param(ws.PathParameter("task-id", TR("task id"))).
		Param(ws.PathParameter("log-type", TR("log type: stdout/stderr"))).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(TaskLogReply{}))

	return ws
}

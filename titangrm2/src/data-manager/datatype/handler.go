package datatype

import (
	"fmt"

	"github.com/emicklei/go-restful"

	"data-manager/types"
	"grm-service/log"
	. "grm-service/util"
)

// GET http://localhost:8080/types
func (s DataTypeSvc) getDatatypeList(req *restful.Request, res *restful.Response) {
	types, err := s.SysDB.GetDataTypes(false)
	if err != nil {
		log.Error("Failed to get data types:", err)
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, &types)
}

// PUT http://localhost:8080/types/{type-name}
func (s DataTypeSvc) updateTypeInfo(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type-name")
	var args updateTypeInfoReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}

	typeInfo := types.DataType{
		Name:        typeName,
		Label:       args.Label,
		IsObsoleted: args.IsObsoleted,
		Description: args.Description,
	}

	err := s.SysDB.UpdateTypeInfo(&typeInfo)
	if err != nil {
		log.Error("Failed to update type info:", err)
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// GET http://localhost:8080/types/meta/{type-name}
func (s DataTypeSvc) getTypeMeta(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type-name")
	metas, err := s.SysDB.GetTypeMetas(typeName, true)
	if err != nil {
		log.Error("Failed to get type meta:", err)
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, &metas)
}

// PUT http://localhost:8080/types/meta/{type-name}
func (s DataTypeSvc) updateMetaField(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type-name")

	var args types.MetaFieldReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(typeName) == 0 || len(args.Group) == 0 ||
		len(args.Title) == 0 || len(args.Name) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid type meta info")))
		return
	}

	meta, err := s.SysDB.UpdateMetaField(typeName, &args)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, meta)
}

// POST http://localhost:8080/types/meta/{type-name}
func (s DataTypeSvc) addMetaField(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type-name")

	var args types.MetaFieldReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(typeName) == 0 || len(args.Title) == 0 || len(args.Name) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid type meta info")))
		return
	}
	if len(args.Group) == 0 {
		args.Group = "Basic Information"
	}

	meta, err := s.SysDB.AddMetaField(typeName, &args)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, meta)
}

// Delete http://localhost:8080/types/meta/{type-name}
func (s DataTypeSvc) delMetaField(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type-name")
	group := req.PathParameter("group")
	field := req.PathParameter("field")

	if len(group) == 0 || len(field) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid type meta info")))
		return
	}

	meta, err := s.SysDB.DelMetaField(typeName, group, field)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, meta)
}

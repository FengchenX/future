package datas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/go-restful"

	. "data-manager/dbcentral/pg"
	"data-manager/types"

	"grm-service/common"
	"grm-service/crypto"
	"grm-service/dbcentral/pg"
	"grm-service/log"
	"grm-service/time"
	. "grm-service/util"
)

var DataAttribute = map[string](map[string][]string){
	"Feature-Shape": {
		BasicGroup: {"shp_type", "feature_class_count", "feature_nums", "ref_system"},
	},
}

type attrField struct {
	Name  string      `json:"name"`
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

// 获取数据信息
func (s DataObjectSvc) getDataInfo(req *restful.Request, res *restful.Response) {
	dataId := req.PathParameter("data-id")
	info, err := s.MetaDB.GetDataInfo(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 修改浏览次数
	go func() {
		cnt, err := s.MetaDB.UpdateDataViewCnt(dataId)
		if err != nil {
			log.Error("Failed to update view count: ", err.Error())
		}

		// 修改ES
		if err := s.EsCli.UpdateDataViewCnt(dataId, cnt); err != nil {
			log.Error("Failed to update view count: ", err.Error())
		}
	}()

	// 用户信息
	userName, err := s.DynamicDB.GetUserName(info.Owner.Id)
	if err == nil {
		info.Owner.Name = userName
	}

	// 数据特有属性信息
	metaValue, err := s.EsCli.GetDataMeta(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	key := fmt.Sprintf("%s-%s", info.DataType, info.SubType)
	fieldMap, ok := DataAttribute[key]
	//fmt.Println(key, fieldMap)
	if ok {
		for group, fields := range fieldMap {
			ret, err := json.Marshal(metaValue[group])
			if err != nil {
				ResWriteError(res, err)
				return
			}

			var metas groupMetas
			if err := json.Unmarshal(ret, &metas); err != nil {
				ResWriteError(res, err)
				return
			}

			for _, field := range fields {
				val := metas.Metas[field]
				attr := attrField{
					Name:  field,
					Label: val.Title,
					Value: val.Value,
				}
				info.Attribute = append(info.Attribute, attr)
			}
		}
	}
	ResWriteHeaderEntity(res, info)
}

// 更新数据信息
func (s DataObjectSvc) updateDataInfo(req *restful.Request, res *restful.Response) {
	var args types.UpdateDataInfoReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	dataId := req.PathParameter("data-id")
	if err := s.MetaDB.UpdateDataInfo(dataId, &args); err != nil {
		ResWriteError(res, err)
		return
	}

	// 修改ES中的tags
	if args.Tags != common.OmitArg {
		if err := s.EsCli.UpdateDataTags(dataId, args.Tags); err != nil {
			ResWriteError(res, err)
			return
		}
	}
	ResWriteEntity(res, nil)
}

// 更新数据快试图
func (s DataObjectSvc) updateDataSnapShot(req *restful.Request, res *restful.Response) {
	// 数据路径
	dataId := req.PathParameter("data-id")
	dataDir := filepath.Join(s.ConfigDir, "data", dataId)
	if err := CheckDir(dataDir); err != nil {
		log.Error("Failed to create data dir :", dataDir)
		ResWriteError(res, err)
		return
	}

	var thumbFile string
	var pic []byte
	var err error

	uploadType := req.QueryParameter("type")
	fmt.Println(uploadType)
	if uploadType == common.FileUpload {
		file, fh, err := req.Request.FormFile("snapshot")
		if err != nil {
			ResWriteError(res, err)
			return
		}
		defer file.Close()

		pic, err = ioutil.ReadAll(file)
		if err != nil {
			ResWriteError(res, err)
			return
		}
		thumbFile = filepath.Join(dataDir, fh.Filename)
	} else if uploadType == common.Base64Upload {
		var args types.SnapshotReq
		if err := req.ReadEntity(&args); err != nil {
			ResWriteError(res, err)
			return
		}
		if len(args.Image) == 0 {
			ResWriteError(res, fmt.Errorf(TR("Invalid snapshot image")))
			return
		}
		picbase64 := args.Image[strings.Index(args.Image, ",")+1:]
		pic, err = base64.StdEncoding.DecodeString(picbase64)
		if err != nil {
			ResWriteError(res, err)
			return
		}

		if len(args.FileName) > 0 {
			thumbFile = fmt.Sprintf("%s/%s", dataDir, args.FileName)
		} else {
			ext := args.Image[strings.Index(args.Image, "/")+1 : strings.Index(args.Image, ";")]
			thumbFile = fmt.Sprintf("%s/%s.%s", dataDir, dataId, ext)
		}
	}

	fmt.Println(thumbFile)
	if err := ioutil.WriteFile(thumbFile, pic, os.ModePerm); err != nil {
		ResWriteError(res, err)
		return
	}
	url := strings.Replace(thumbFile, s.ConfigDir, common.FilePre, -1)
	if err := s.MetaDB.UpdateDataSnap(dataId, url); err != nil {
		ResWriteError(res, err)
		return
	}

	if err := s.EsCli.UpdateDataSnap(dataId, url); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

type groupMetas struct {
	Label string                     `json:"label" description:"label of group"`
	Metas map[string]types.MetaValue `json:"metas" description:"metas of group"`
}

// 获取数据元数据信息
func (s DataObjectSvc) getDataMeta(req *restful.Request, res *restful.Response) {
	dataId := req.PathParameter("data-id")
	dataInfo, err := s.MetaDB.GetDataInfo(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	dataType := dataInfo.DataType
	if len(dataInfo.SubType) > 0 {
		dataType = dataInfo.SubType
	}
	meta, err := s.SysDB.GetTypeMetas(dataType, false)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	//fmt.Println(meta)
	metaValue, err := s.EsCli.GetDataMeta(dataId)
	for groupKey, v := range metaValue {
		ret, err := json.Marshal(v)
		if err != nil {
			continue
		}

		var group groupMetas
		if err := json.Unmarshal(ret, &group); err != nil {
			continue
		}

		// 分组元数据赋值
		for i, metas := range meta.DataMeta {
			if metas.Group == groupKey {
				for j, _ := range meta.DataMeta[i].Values {
					meta := &meta.DataMeta[i].Values[j]
					if val, ok := group.Metas[meta.Name]; ok {
						//fmt.Println(val)
						meta.IsModified = val.IsModified
						meta.Value = val.Value
					}
				}
			}
		}
	}
	ResWriteHeaderEntity(res, meta)
}

// 更新数据元数据信息
func (s DataObjectSvc) updateDataMeta(req *restful.Request, res *restful.Response) {
	data := req.PathParameter("data-id")
	var args types.UpdateDataMetaReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}

	if err := s.EsCli.UpdateDataMeta(data, args); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// 移除数据
func (s DataObjectSvc) delData(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	dataId := req.PathParameter("data-id")
	info, err := s.MetaDB.GetDataInfo(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if userId != info.Owner.Id {
		ResWriteError(res, fmt.Errorf(TR("Don't have permission to perform this action")))
		return
	}

	// 移除数据，修改数据状态
	if err := s.MetaDB.UpdateDataStatus(dataId, common.DataObsoleted); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// 获取表格数据内容
func (s *DataObjectSvc) getTableData(dataId, storage string, filter *types.DataSearch) (*types.TableData, error) {
	var dataDB *DataDB
	_dataDb, err := pg.ConnectDataDBUrl(storage)
	if err != nil {
		return nil, err
	}
	defer _dataDb.DisConnect()
	dataDB = &DataDB{_dataDb}

	filter.DataId = dataId
	data, err := dataDB.GetTableData(filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 获取数据内容
func (s DataObjectSvc) getDataContent(req *restful.Request, res *restful.Response) {
	dataId := req.PathParameter("data-id")
	info, err := s.MetaDB.GetDataInfo(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	content := dataContent{
		Id:       dataId,
		Name:     info.DataName,
		DataType: info.DataType,
		SubType:  info.SubType,
	}
	if info.DataType == DocType || info.DataType == MediaType {
		content.DataUrl = info.DataUrl
	} else if info.DataType == OrthoType || info.DataType == DEMType {

	} else if info.DataType == FeatureType || info.DataType == CalTablelType {
		if len(info.Storage) > 0 {
			storage, err := crypto.AesDecrypt(info.Storage)
			if err != nil {
				ResWriteError(res, err)
				return
			}

			args := ParserPageArgs(req)
			filter := types.DataSearch{
				DataType: info.DataType,
				Order:    args.Order,
				Limit:    args.Limit,
				Offset:   args.Offset,
				Sort:     args.Sort,
			}
			data, err := s.getTableData(dataId, storage, &filter)
			if err != nil {
				ResWriteError(res, err)
				return
			}
			content.TableData = data
		}
	} else if info.DataType == AerialType {
		subs, err := s.MetaDB.GetDataSubObjs(dataId)
		if err != nil {
			ResWriteError(res, err)
			return
		}
		content.SubObj = subs
	}
	ResWriteHeaderEntity(res, &content)
}

// 添加数据评论
func (s DataObjectSvc) addComment(req *restful.Request, res *restful.Response) {
	dataId := req.PathParameter("data-id")
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	var request addCommentReq
	if err := req.ReadEntity(&request); err != nil {
		ResWriteError(res, err)
		return
	}
	comment := types.Comment{
		Id:         time.GetUnixNano(),
		DataId:     dataId,
		CreateTime: GetTimeNowStd(),
		FromUser:   &common.UserInfo{Id: userId},
		Content:    request.Content,
	}

	if user, err := s.DynamicDB.GetUserName(userId); err == nil {
		comment.FromUser.Name = user
	}

	if len(request.ToUser) > 0 {
		comment.ToUser = &common.UserInfo{Id: request.ToUser}
		if user, err := s.DynamicDB.GetUserName(request.ToUser); err == nil {
			comment.ToUser.Name = user
		}
	}
	if err := s.DynamicDB.CreateComment(&comment); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, &comment)
}

// 获取数据评论
func (s DataObjectSvc) queryComments(req *restful.Request, res *restful.Response) {
	dataId := req.PathParameter("data-id")
	comments, err := s.DynamicDB.QueryComments(dataId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	var resp commentListResp
	resp.Cmts = comments
	resp.Total = len(comments)
	ResWriteEntity(res, &resp)
}

// 移除评论
func (s DataObjectSvc) delComment(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	dataId := req.PathParameter("data-id")
	commentId := req.PathParameter("comment-id")
	if err := s.DynamicDB.DelComment(dataId, commentId, userId); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

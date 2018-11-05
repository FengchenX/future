package importer

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/emicklei/go-restful"

	"data-importer/executor"
	. "data-importer/mq/dataworker"
	. "data-importer/types"
	"grm-service/common"
	"grm-service/crypto"
	"grm-service/geoserver"
	"grm-service/log"
	"grm-service/mq"
	"grm-service/path"
	grmtime "grm-service/time"
	. "grm-service/util"
)

var (
	domainFile = "FileSysDomain.json"
)

// POST http://localhost:8080/sysdomain
func (s ImporterSvc) getSysDomain(req *restful.Request, res *restful.Response) {
	path, err := req.BodyParameter("path")
	if err != nil {
		ResWriteError(res, err)
		return
	}
	request := domainReq{path}
	domains, err := getFileSysDomain(filepath.Join(s.ConfigDir, domainFile))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	var keys []string
	for k := range domains {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var domainInfo sysDomains
	if len(request.Path) == 0 {
		for _, key := range keys {
			domainInfo = append(domainInfo, domain{Name: key, Path: domains[key], Domain: key, IsDir: true})
		}
	} else {
		for _, key := range keys {
			if strings.HasPrefix(request.Path, domains[key]) {
				var err error
				if domainInfo, err = walkDir(request.Path, key, domains[key]); err != nil {
					ResWriteError(res, err)
					return
				}
			}
		}
	}
	ResWriteHeaderEntity(res, &domainInfo)
}

// 发布默认矢量图层
func (s ImporterSvc) publishLayer(dataId, userId string) error {
	layer := common.DataLayer{
		Layer:     NewUUID(),
		Name:      TR("Layer"),
		Data:      dataId,
		User:      userId,
		IsDefault: true,
	}
	geoStorage, _ := s.MetaDB.GetDeviceGeoStorage(dataId)
	srs, wmsUrl, wmtsUrl, wms, wfs, wmts, err := s.GeoServer.AddShpLayer(geoserver.GeoWorkSpace,
		geoStorage, dataId, layer.Layer)
	if err != nil {
		log.Error("Publish feature Layer error :", err.Error())
		return err
	}
	layer.Srs = srs
	layer.WmsUrl = wmsUrl
	layer.WmtsUrl = wmtsUrl
	layer.WMS = wms
	layer.Wfs = wfs
	layer.Wmts = wmts

	name := fmt.Sprintf("lyr-%s_%s", dataId, layer.Layer)
	style, err := s.GeoServer.GetShapeLayerDefaultStyle(name)
	if err != nil {
		log.Error("Get feature Layer style error :", err.Error())
		return err
	}
	if style == "point" {
		layer.Style = "0"
	} else if style == "line" {
		layer.Style = "1"
	} else if style == "polygon" {
		layer.Style = "2"
	}
	if _, err := s.SysDB.AddDataLayer(&layer); err != nil {
		log.Error("Add Data Layer error :", err.Error())
		return err
	}
	return nil
}

// 发布office
func (s ImporterSvc) publishOffice(dataId, file string) error {
	id, err := s.OfficeServer.Upload(file, dataId)
	if err != nil {
		log.Error("Upload office error :", err.Error())
		return err
	}
	url, err := s.OfficeServer.Share(id)
	if err != nil {
		log.Error("Upload office error :", err.Error())
		return err
	}
	fmt.Println(url)
	url = strings.Replace(url, s.OfficeServer.Endpoints, common.OfficePre, -1)
	if err := s.MetaDB.UpdateDataUrl(dataId, url); err != nil {
		log.Error("Update DataUrl Meta error :", err.Error())
		return err
	}
	if err := s.EsServer.UpdateDataUrl(dataId, url); err != nil {
		log.Error("Update DataUrl Es error :", err.Error())
		return err
	}
	return nil
}

func (s ImporterSvc) launchTask(taskType, taskDir, taskId string, args []string) {
	// 创建task目录
	fmt.Println("Task dir: ", taskDir)
	path.CreateAllPath(taskDir + "/log")
	path.CreateAllPath(taskDir + "/temp")

	// 获取可执行程序
	path, err := exec.LookPath(CmdPath)
	if err == nil {
		CmdPath = path
	}

	// 启动程序
	executor := executor.NewExecutor(CmdPath, args)
	err = executor.LaunchCmd(taskDir, func(msg string) {
		executor.MutexLock.Lock()
		defer executor.MutexLock.Unlock()
		if !strings.HasPrefix(msg, "[") || !strings.HasSuffix(msg, "]\n") {
			return
		}
		msg = msg[1 : len(msg)-2]
		fmt.Printf("%s\n", msg)

		index := strings.Index(msg, ":")
		if index == -1 {
			return
		}
		prefix := msg[:index]
		msg = msg[index+1:]
		switch prefix {
		case ProcessMsg:
			{
				if err := s.DynamicDB.UpdateTaskProcess(taskType, taskId, msg); err != nil {
					log.Error("UpdateTaskProcess error :", err.Error())
					return
				}
				// 发送消息
				if err := s.MsgQueue.PublishMessage(TASKMESSAGE,
					mq.Publish(PublishHandler(s.DynamicDB, taskType, taskId))); err != nil {
					log.Error("Publish TaskProcess Message error :", err.Error())
					return
				}
			}
		case ProgressMsg:
			{
				if err := s.DynamicDB.UpdateTaskProgress(taskType, taskId, msg); err != nil {
					log.Error("UpdateTaskProgress error :", err.Error())
					return
				}
				// 发送消息
				if err := s.MsgQueue.PublishMessage(TASKMESSAGE,
					mq.Publish(PublishHandler(s.DynamicDB, taskType, taskId))); err != nil {
					log.Error("Publish TaskProgress Message error :", err.Error())
					return
				}
			}
		case StatusMsg:
			{
				if err := s.DynamicDB.UpdateTaskStatus(taskType, taskId, msg); err != nil {
					log.Error("UpdateTaskStatus error :", err.Error())
					return
				}
				// 发送消息
				if err := s.MsgQueue.PublishMessage(TASKMESSAGE,
					mq.Publish(PublishHandler(s.DynamicDB, taskType, taskId))); err != nil {
					log.Error("Publish TaskStatus Message error :", err.Error())
					return
				}
			}
		case FeatureMsg:
			{
				// 发布矢量默认图层
				index := strings.Index(msg, ",")
				dataId := msg[:index]
				userId := msg[index+1:]
				if err := s.publishLayer(dataId, userId); err != nil {
					log.Error("Publish shape default layer error :", err.Error())
					return
				}
			}
		case OfficeMsg:
			{
				// 发布office文档
				index := strings.Index(msg, ",")
				path := msg[:index]
				id := msg[index+1:]
				if err := s.publishOffice(id, path); err != nil {
					log.Error("Publish office error :", err.Error())
					return
				}
			}
		}
	})
	if err != nil {
		log.Error("error in launching command:", err)
	}
}

// POST http://localhost:8080/datascan
func (s ImporterSvc) dataScan(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	args := dataScanRequest{}
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	log.Info("Scan args:", args)
	if len(args.DataDir) == 0 || !path.Exists(args.DataDir) {
		ResWriteError(res, ErrScanDataDirInvalid)
		return
	}

	// 验证数据集类型和数据类型是否匹配
	if len(args.DataSet) == 0 {
		ResWriteError(res, ErrInvalidDataSet)
		return
	}
	setType, err := s.SysDB.GetDataSetType(args.DataSet)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if setType != args.DataType {
		sysTypes, err := s.SysDB.GetTypesInfo("")
		if err != nil {
			ResWriteError(res, err)
			return
		}
		if len(setType) > 0 {
			if sysTypes[setType].Parent != args.DataType &&
				sysTypes[args.DataType].Parent != setType {
				ResWriteError(res, fmt.Errorf(TR("Invalid dataset type(%s) for data scan(%s)", setType, args.DataType)))
				return
			}
		}
	}

	// 初始化任务信息
	if len(args.TaskName) == 0 {
		args.TaskName = fmt.Sprintf(`%s %s`, args.DataType, time.Now().Format("2006-01-02 15:04:05"))
	}
	info := map[string]string{
		"task_name":     args.TaskName,
		"task_type":     ScanTask,
		"user_id":       userId,
		"pre_type":      args.PreType,
		"data_type":     args.DataType,
		"data_set":      args.DataSet,
		"device":        strconv.Itoa(args.Device),
		"start_time":    time.Now().Format("2006-01-02 15:04:05"),
		"scan_dir":      args.DataDir,
		"status":        TaskIdle,
		"finished_time": "",
		"pid":           "",
		"process":       "",
		"progress":      "",
	}
	taskId := ScanTask + "_" + NewUUID()
	if _, err := s.DynamicDB.InitTaskInfo(ScanTask, taskId, info); err != nil {
		ResWriteError(res, err)
		return
	}

	taskDir := filepath.Join(s.ConfigDir, "dataworker", ScanTask, taskId)
	// 数据库连接信息
	sysDB := fmt.Sprintf("postgres://%s:%s@%s/%s",
		s.SysDB.Config.User, s.SysDB.Config.Password,
		s.SysDB.Config.Host, s.SysDB.Config.Database)

	metaDB := fmt.Sprintf("postgres://%s:%s@%s/%s",
		s.MetaDB.Config.User, s.MetaDB.Config.Password,
		s.MetaDB.Config.Host, s.MetaDB.Config.Database)

	dataDB, err := s.SysDB.GetDeviceStr(args.Device)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 加密设备连接字符串
	if dataDB, err = crypto.AesEncrypt(dataDB); err != nil {
		ResWriteError(res, err)
		return
	}

	// 扫描程序参数
	arg := []string{
		"--mode", ScanTask,
		"--jobid", taskId,
		"--scandir", args.DataDir,
		"--datatype", args.DataType,
		"--sysdb", sysDB,
		"--metadb", metaDB,
		"--datadb", dataDB,
		"--storage", dataDB,
		"--esdb", s.EsUrl,
		"--data_set", args.DataSet,
		"--workdir", taskDir,
		"--datadir", s.ConfigDir,
		"--user", userId,
		//"--Config", "0",
	}
	go s.launchTask(ScanTask, taskDir, taskId, arg)
	ResWriteHeaderEntity(res, &dataScanReply{taskId})
}

// POST http://localhost:8080/dataload
func (s ImporterSvc) dataLoad(req *restful.Request, res *restful.Response) {
	// 解析参数信息
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	args := dataLoadRequest{}
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	log.Info("Load args:", args)
	if len(args.ScanTask) == 0 {
		ResWriteError(res, ErrScanTaskIdInvalid)
		return
	}
	if len(args.FileIds) == 0 && !args.AllFileLoad {
		ResWriteError(res, ErrNoDataLoad)
		return
	}

	// 获取扫描任务参数信息
	scanInfo, err := s.DynamicDB.GetTaskInfo(ScanTask, args.ScanTask)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 生成入库文件参数列表
	fileIds := make([]string, 0, 100)
	if args.AllFileLoad {
		log.Info("Load all files:", args.ScanTask)
		files, err := s.MetaDB.GetScanDatas(args.ScanTask, "", nil, nil, false)
		if err != nil {
			ResWriteError(res, ErrNoDataLoad)
			return
		}
		for _, val := range files.Datas {
			fmt.Println("file:", val.FileId)
			fileIds = append(fileIds, val.FileId)
		}
	} else {
		fileIds = args.FileIds
	}
	if len(fileIds) == 0 {
		ResWriteError(res, ErrNoDataLoad)
		return
	}

	// 初始化任务信息
	info := map[string]string{
		"task_name":     scanInfo.TaskName,
		"task_type":     LoadTask,
		"data_type":     scanInfo.DataType,
		"pre_type":      scanInfo.PreType,
		"user_id":       userId,
		"start_time":    time.Now().Format("2006-01-02 15:04:05"),
		"status":        TaskIdle,
		"finished_time": "",
		"pid":           "",
		"process":       "",
		"progress":      "",
	}

	// 创建任务目录
	taskId := strings.Replace(args.ScanTask, ScanTask, LoadTask, -1)
	taskDir := filepath.Join(s.ConfigDir, "dataworker", LoadTask, taskId)
	path.CreateAllPath(taskDir)

	// 初始化入库文件列表
	fileList := filepath.Join(taskDir, "file_list.txt")
	if err := writeLoadFiles(fileList, fileIds); err != nil {
		log.Error("Failed to init file list to load error :", err.Error())
		ResWriteError(res, err)
		return
	}

	// 初始化任务信息
	if _, err := s.DynamicDB.InitTaskInfo(LoadTask, taskId, info); err != nil {
		ResWriteError(res, err)
		return
	}

	// 数据库连接信息
	sysDB := fmt.Sprintf("postgres://%s:%s@%s/%s",
		s.SysDB.Config.User, s.SysDB.Config.Password,
		s.SysDB.Config.Host, s.SysDB.Config.Database)

	metaDB := fmt.Sprintf("postgres://%s:%s@%s/%s",
		s.MetaDB.Config.User, s.MetaDB.Config.Password,
		s.MetaDB.Config.Host, s.MetaDB.Config.Database)

	dataDB, err := s.SysDB.GetDeviceStr(scanInfo.Device)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 扫描程序参数
	arg := []string{
		"--mode", LoadTask,
		"--jobid", taskId,
		"--load_files", fileList,
		"--datatype", scanInfo.DataType,
		"--data_set", scanInfo.DataSet,
		"--sysdb", sysDB,
		"--metadb", metaDB,
		"--datadb", dataDB,
		"--esdb", s.EsUrl,
		"--workdir", taskDir,
		"--datadir", s.ConfigDir,
		"--user", userId,
		//"--Config", "0",
	}
	go s.launchTask(LoadTask, taskDir, taskId, arg)
	ResWriteHeaderEntity(res, &dataLoadReply{taskId})
}

// GET http://localhost:8080/datascan/{task-id}
func (s ImporterSvc) dataScanResult(req *restful.Request, res *restful.Response) {
	taskId := req.PathParameter("task-id")
	if len(taskId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid scan task id")))
		return
	}

	filter := func(req *restful.Request) *ResultFilter {
		ret := ResultFilter{
			FileName:      req.QueryParameter("file-name"),
			FileSizeMin:   req.QueryParameter("file-size-min"),
			FileSizeMax:   req.QueryParameter("file-size-max"),
			CreateTimeMin: req.QueryParameter("create-time-min"),
			CreateTimeMax: req.QueryParameter("create-time-max"),
		}
		return &ret
	}

	result, err := s.MetaDB.GetScanDatas(taskId, "", ParserPageArgs(req), filter(req), true)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	//  获取所有类型的信息
	types, err := s.SysDB.GetTypesInfo("")
	if err != nil {
		ResWriteError(res, err)
		return
	}
	for i, _ := range result.Datas {
		result.Datas[i].TypeLabel = types[result.Datas[i].FileType].Label
		if len(result.Datas[i].SubType) > 0 {
			result.Datas[i].SubLabel = types[result.Datas[i].SubType].Label
		}
	}
	ResWriteHeaderEntity(res, &result)
}

// 数据上传
func (s ImporterSvc) dataUpload(req *restful.Request, res *restful.Response) {
	// 验证用户信息
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 获取系统类型
	dataSet := req.Request.FormValue("data-set")
	dataType := req.Request.FormValue("data-type")
	sysTypes, err := s.SysDB.GetTypesInfo("")
	if err != nil {
		ResWriteError(res, err)
		return
	}
	preType := sysTypes[dataType].Parent
	fmt.Println("type:", preType, dataType)

	// 获取存储设备信息
	devId, devStr, err := s.SysDB.GetDeviceStrByType(dataType)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	fmt.Println("device:", devId, devStr)

	// 验证数据集类型和数据类型是否匹配
	if len(dataSet) == 0 || len(dataType) == 0 {
		ResWriteError(res, ErrInvalidDataInfo)
		return
	}
	setType, err := s.SysDB.GetDataSetType(dataSet)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if setType != dataType {
		if len(setType) > 0 {
			if sysTypes[setType].Parent != dataType &&
				sysTypes[dataType].Parent != setType {
				ResWriteError(res, fmt.Errorf(TR("Invalid dataset type(%s) for data load(%s)", setType, dataType)))
				return
			}
		}
	}

	// 数据上传并解压
	req.Request.Body = http.MaxBytesReader(res.ResponseWriter, req.Request.Body, 100<<30)
	defer req.Request.Body.Close()
	file, fh, err := req.Request.FormFile("file")
	if err != nil {
		ResWriteError(res, err)
		return
	}
	defer file.Close()

	dataId := GenerateUUID()
	dataPath := filepath.Join(s.ConfigDir, "data", dataId)
	dataFile := filepath.Join(dataPath, fh.Filename)
	scanDir, err := uploadFormFile(dataFile, file)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	fmt.Println("scan dir: ", scanDir)

	// 初始化任务信息
	taskName := fmt.Sprintf(TR(`UploadTask_%s000`, grmtime.GetUnix()))
	info := map[string]string{
		"task_name":     taskName,
		"task_type":     LoadTask,
		"user_id":       userId,
		"pre_type":      preType,
		"data_type":     dataType,
		"data_set":      dataSet,
		"device":        devId,
		"start_time":    time.Now().Format("2006-01-02 15:04:05"),
		"scan_dir":      scanDir,
		"status":        TaskIdle,
		"finished_time": "",
		"pid":           "",
		"process":       "",
		"progress":      "",
	}
	taskId := LoadTask + "_" + NewUUID()
	if _, err := s.DynamicDB.InitTaskInfo(LoadTask, taskId, info); err != nil {
		ResWriteError(res, err)
		return
	}
	taskDir := filepath.Join(s.ConfigDir, "dataworker", LoadTask, taskId)

	// 数据库连接信息
	sysDB := fmt.Sprintf("postgres://%s:%s@%s/%s",
		s.SysDB.Config.User, s.SysDB.Config.Password,
		s.SysDB.Config.Host, s.SysDB.Config.Database)

	metaDB := fmt.Sprintf("postgres://%s:%s@%s/%s",
		s.MetaDB.Config.User, s.MetaDB.Config.Password,
		s.MetaDB.Config.Host, s.MetaDB.Config.Database)

	// 加密设备连接字符串
	dataDB, err := crypto.AesEncrypt(devStr)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 入库程序参数
	arg := []string{
		"--mode", UploadTask,
		"--jobid", taskId,
		"--scandir", scanDir,
		"--datatype", dataType,
		"--sysdb", sysDB,
		"--metadb", metaDB,
		"--datadb", devStr,
		"--storage", dataDB,
		"--esdb", s.EsUrl,
		"--data_set", dataSet,
		"--workdir", taskDir,
		"--datadir", s.ConfigDir,
		"--user", userId,
		//"--Config", "0",
	}
	go s.launchTask(LoadTask, taskDir, taskId, arg)
	ResWriteHeaderEntity(res, &dataLoadReply{taskId})
}

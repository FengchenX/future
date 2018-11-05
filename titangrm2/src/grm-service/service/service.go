package service

import (
	"net/http"
	"time"

	"grm-service/command"
	"grm-service/util"
)

var (
	// For serving
	DefaultName    = "grm-service"
	DefaultVersion = "latest"
	DefaultId      = util.NewUUID()
	DefaultAddress = ":8080"

	// for registration
	DefaultRegisterTTL      = DefaultRegisterInterval + time.Millisecond
	DefaultRegisterInterval = time.Second * 60
)

const (
	GRMAPIService         = "grm-api"
	DataManagerService    = "data-manager"
	DataImporterService   = "data-importer"
	StorageManagerService = "storage-manager"
	SearcherService       = "searcher"
	StatService           = "statistics"
	LabelMgrService       = "label-manager"
	DataCollection        = "data-collection"

	FileService  = "file"
	EtcdService  = "etcd"
	NginxService = "nginx"
	Geoserver    = "geoserver"

	TitanAuthService   = "titan-auth"
	TileManagerService = "tile-manager"
)

type Service interface {
	Init(c *command.Meta) error
	Run() error
	String() string

	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

func NewService(name, version string) Service {
	return newService(name, version)
}

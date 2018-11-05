package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"github.com/micro/go-api"

	"grm-service/command"
	"grm-service/log"
	"grm-service/registry"
	"grm-service/util"
)

type service struct {
	opts Options

	mux *http.ServeMux
	srv *registry.Service

	sync.Mutex
	running bool
	exit    chan chan error
}

func newService(name, ver string) Service {
	options := newOptions(name, ver)
	s := &service{
		opts: options,
		mux:  http.NewServeMux(),
	}
	s.srv = s.genSrv()
	return s
}

func (s *service) Init(c *command.Meta) error {
	// 加载翻译文件
	log.Printf("Initializing %s ...", s.String())
	//util.LoadTranslation("/usr/local/bin/translation", "zh_CN", s.opts.Name)
	util.LoadTranslation("/usr/local/bin/translation", "zh_CN", "titan-grm")
	//util.LoadTranslation("D:\\WorkSpace\\grm2.0\\src\\Services\\titangrm2\\src\\titan-grm\\translation", "zh_CN", "titan-grm")

	// 初始化日志库
	log.InitLog(s.opts.Name, filepath.Join(c.ConfigDir, "logs", s.opts.Name, s.opts.Version),
		log.LogNameComputer, 3*30*24*time.Hour, 7*24*time.Hour)

	// 初始化服务注册
	if len(c.ServiceAddress) > 0 {
		s.opts.Address = c.ServiceAddress
	}
	if len(c.RegistryAddress) > 0 {
		s.opts.RegistryAddr = c.RegistryAddress
	}
	if len(c.ServiceNamespace) > 0 {
		s.opts.Namespace = c.ServiceNamespace
	}

	srv := s.genSrv()
	srv.Endpoints = s.srv.Endpoints
	s.srv = srv

	return nil
}

func (s *service) Run() error {
	log.Printf("Starting %s ...", s.String())

	if err := s.start(); err != nil {
		return err
	}

	if err := s.register(); err != nil {
		log.Fatal(err)
		return err
	}

	// start reg loop
	ex := make(chan bool)
	go s.run(ex)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Println(<-ch)

	close(ex)

	if err := s.deregister(); err != nil {
		return err
	}

	return s.stop()
}

func (s *service) String() string {
	return fmt.Sprintf("Service %s(%s): %s", s.opts.Name, s.opts.Version, s.opts.Id)
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "TitanGRM API",
			Description: "Resource for TitanGRM",
			Contact: &spec.ContactInfo{
				Name:  "titan",
				Email: "",
				URL:   "",
			},
			License: &spec.License{
				Name: "",
				URL:  "",
			},
			Version: "2.0.0",
		},
	}
	swo.Schemes = []string{"http"}
}

func (s *service) Handle(pattern string, handler http.Handler) {
	wc, ok := handler.(*restful.Container)
	if ok && wc != nil {
		config := restfulspec.Config{
			WebServices:                   wc.RegisteredWebServices(),
			APIPath:                       "/apidocs.json",
			PostBuildSwaggerObjectHandler: enrichSwaggerObject,
		}
		wc.Add(restfulspec.NewOpenAPIService(config))

		cors := restful.CrossOriginResourceSharing{
			ExposeHeaders: []string{"X-My-Header"},
			AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length",
				"Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Allow-Headers", "auth-session"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			CookiesAllowed: true,
			Container:      wc,
		}
		wc.Filter(cors.Filter)

		// http://localhost:8080/apidocs.json
		http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("."))))
	}

	var seen bool
	for _, ep := range s.srv.Endpoints {
		if ep.Name == pattern {
			seen = true
			break
		}
	}
	if !seen {
		s.srv.Endpoints = append(s.srv.Endpoints, &registry.Endpoint{
			Name:     pattern,
			Metadata: api.Encode(&api.Endpoint{Name: pattern, Path: []string{pattern}, Handler: api.Web}),
		})
	}

	s.mux.Handle(pattern, handler)
}

func (s *service) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func (s *service) genSrv() *registry.Service {
	// default host:port
	parts := strings.Split(s.opts.Address, ":")
	host := strings.Join(parts[:len(parts)-1], ":")
	port, _ := strconv.Atoi(parts[len(parts)-1])

	// check the advertise address first
	// if it exists then use it, otherwise
	// use the address
	if len(s.opts.Advertise) > 0 {
		parts = strings.Split(s.opts.Advertise, ":")

		// we have host:port
		if len(parts) > 1 {
			// set the host
			host = strings.Join(parts[:len(parts)-1], ":")

			// get the port
			if aport, _ := strconv.Atoi(parts[len(parts)-1]); aport > 0 {
				port = aport
			}
		} else {
			host = parts[0]
		}
	}

	addr, err := util.Extract(host)
	if err != nil {
		// best effort localhost
		addr = "127.0.0.1"
	}

	return &registry.Service{
		Name:      s.opts.Name,
		Version:   s.opts.Version,
		Namespace: s.opts.Namespace,
		Nodes: []*registry.Node{&registry.Node{
			Id:       s.opts.Id,
			Address:  addr,
			Port:     port,
			Metadata: s.opts.Metadata,
		}},
	}
}

func (s *service) start() error {
	s.Lock()
	defer s.Unlock()

	if s.running {
		return nil
	}

	l, err := net.Listen("tcp", s.opts.Address)
	if err != nil {
		return err
	}

	s.opts.Address = l.Addr().String()
	fmt.Println("address:", s.opts.Address)
	srv := s.genSrv()
	srv.Endpoints = s.srv.Endpoints
	s.srv = srv

	var h http.Handler
	h = s.mux

	for _, fn := range s.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	var httpSrv *http.Server
	httpSrv = &http.Server{}
	httpSrv.Handler = h

	go httpSrv.Serve(l)

	for _, fn := range s.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	s.exit = make(chan chan error, 1)
	s.running = true

	go func() {
		ch := <-s.exit
		ch <- l.Close()
	}()

	log.Printf("Listening on %v\n", l.Addr().String())
	return nil
}

func (s *service) run(exit chan bool) {
	if s.opts.RegisterInterval <= time.Duration(0) {
		return
	}

	t := time.NewTicker(s.opts.RegisterInterval)

	for {
		select {
		case <-t.C:
			s.register()
		case <-exit:
			t.Stop()
			return
		}
	}
}

func (s *service) stop() error {
	s.Lock()
	defer s.Unlock()

	if !s.running {
		return nil
	}

	for _, fn := range s.opts.BeforeStop {
		if err := fn(); err != nil {
			return err
		}
	}

	ch := make(chan error, 1)
	s.exit <- ch
	s.running = false

	log.Println("Stopping")

	for _, fn := range s.opts.AfterStop {
		if err := fn(); err != nil {
			if chErr := <-ch; chErr != nil {
				return chErr
			}
			return err
		}
	}

	return <-ch
}

func RegisterConsulChk() registry.Option {
	return func(o *registry.Options) {
		o.Context = context.WithValue(context.Background(), "consul_tcp_check", 2*time.Second)
	}
}

func (s *service) register() error {
	if s.srv == nil {
		return nil
	}
	if len(s.opts.RegistryAddr) > 0 {
		reg := registry.NewRegistry(registry.Addrs(s.opts.RegistryAddr))
		return reg.Register(s.srv, registry.RegisterTTL(s.opts.RegisterTTL))
	}
	return registry.DefaultRegistry.Register(s.srv, registry.RegisterTTL(s.opts.RegisterTTL))
}

func (s *service) deregister() error {
	if s.srv == nil {
		return nil
	}
	return registry.Deregister(s.srv)
}

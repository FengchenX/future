package svc11

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/feng/future/go-kit/agfun/trace/middleware"
)

type httpService struct {
	service Service
}

func (s *httpService) concatHandler(c *gin.Context) {
	v := c.Request.URL.Query()
	result, err := s.service.Concat(c.Request.Context(), v.Get("a"), v.Get("b"))
	if err != nil {
		fmt.Printf("%+v", err)
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *httpService) sumHandler(c *gin.Context) {
	panic("todo")
}

func RouterInit(hostPort string, tracer opentracing.Tracer, service Service) {
	svc := &httpService{service: service}
	concatHandler := gin.HandlerFunc(svc.concatHandler)
	concatHandler = middleware.FromHTTPRequest(tracer, "Concat")(concatHandler)
	sumHandler := gin.HandlerFunc(svc.sumHandler)
	sumHandler = middleware.FromHTTPRequest(tracer, "Sum")(sumHandler)
	router := gin.New()
	router.GET("/concat/", concatHandler)
	router.POST("/sum/", sumHandler)
	router.Run(hostPort)
}

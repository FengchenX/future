package svc11

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

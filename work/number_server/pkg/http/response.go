package http

import (
	"sub_account_service/number_server/pkg/error"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	resp *gin.Context
}

func (h *HttpResponse) Response(httpCode, errCode int, data interface{}) {
	h.resp.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  error.GetMsg(errCode),
		"data": data,
	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

    router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	
	router.POST("/test", func(c *gin.Context){
		
	})
    router.Run(":8000")
}
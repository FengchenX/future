package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

    router.GET("/", func(c *gin.Context) {
		query := c.Request.URL.Query()
		fmt.Println(query.Get("args1"))
		c.String(http.StatusOK, "Hello World")
	})
	
	router.POST("/test", func(c *gin.Context){
		var req struct{}
		if err := c.BindJSON(&req); err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "test")
	})
    router.Run(":8000")
}
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/swag/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})
	r.Run(":8081") //默认8080端口
}

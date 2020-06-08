package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问方式：1、切到router_static文件下
	//2、go run main.go
	//3、打开浏览器输入http://localhost:8080/assets/1.html
	r.Static("/assets", "./assets")
	//http://localhost:8080/png
	r.Static("/png", ".")

	//http://localhost:8080/static/2.html
	r.StaticFS("/static", http.Dir("static"))

	//没下这个文件不演示
	r.StaticFile("/favicon.ico", "ico")
	r.Run()
}

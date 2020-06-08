package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//浏览器填入格式如： http://localhost:8080/lp罗飘/12
	r.GET("/:name/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": context.Param("name"),
			"id":   context.Param("id"),
		})
	})
	r.Run()
}

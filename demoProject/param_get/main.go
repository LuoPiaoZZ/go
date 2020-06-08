package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//http://localhost:8080/get?first=hello&last=world
	r.GET("/get", func(context *gin.Context) {
		//获取值
		firstName := context.Query("first")
		lastName := context.DefaultQuery("last", "默认值")
		context.String(200, "%s,%s", firstName, lastName)
	})
	r.Run(":8080")
}

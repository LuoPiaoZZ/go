package main

//泛绑定
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//所有user目录下的请求都会执行这个方法，打印出helloworld
	r.GET("/user/*action", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	r.Run()
}

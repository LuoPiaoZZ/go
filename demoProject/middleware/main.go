package main

//自定义中间件
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//推荐：第一种自定义中间件
func Middle1(c *gin.Context) {
	t1 := time.Now()
	fmt.Println("1请求之前", t1)
	c.Set("middle", 1)
	c.Next() //请求
	t2 := time.Now()
	fmt.Println("1请求之后", t2)
	t3 := time.Since(t2)
	fmt.Println("1时间差：", t3)
}

//第二种自定义中间件
func Middle2() gin.HandlerFunc {
	return func(context *gin.Context) {
		t1 := time.Now()
		fmt.Println("2请求之前", t1)
		context.Next() //请求
		t2 := time.Now()
		fmt.Println("2请求之后", t2)
		t3 := time.Since(t2)
		fmt.Println("2时间差：", t3)
	}
}
func main() {
	//创建无中间件路由
	r := gin.New()
	//使用中间件，后面无限续中间件
	r.Use(Middle2())
	r.GET("/get", func(context *gin.Context) {
		context.String(200, "get!!!")
	})
	//路由中间件，后面也可以无限续中间件，这里会拿不到middle那个值
	r.POST("/post", func(context *gin.Context) {
		context.String(200, "post!!!", context.PostForm("name"))
	}, Middle1)
	r.Run()
}

package main
//重定向
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r:=gin.Default()
	r.Static("/static","./static")
	r.GET("/url/1", func(c *gin.Context) {
		//外部重定向：使用301重定向
		c.Redirect(http.StatusMovedPermanently,"/static/login.html")
	})
	//内部重定向，显示地址无变化
	r.GET("/url/2", func(c *gin.Context) {
		c.Request.URL.Path = "/static/login.html"
		r.HandleContext(c)
	})
	r.Run(":8080")
}



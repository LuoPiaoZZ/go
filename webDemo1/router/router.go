package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"webDemo1/middleware"
	"webDemo1/model"
	"webDemo1/service"
)

func InitRouter()  {
	router:=gin.Default()
	//跨域中间件
	//router.Use(middleware.Cors())
	//session中间件
	router.Use(sessions.Sessions("mysession",middleware.Store))
	v1:=router.Group("v1")
	{
		v1.POST("/reg", func(c *gin.Context) {
			var user model.User
			err:=c.ShouldBind(&user)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{"error":err})
			}
			result:=service.InsertUser(user.Uname,user.Upassword)
			if result{
				middleware.SetSession(c,user.Uname)
				c.JSON(http.StatusOK,gin.H{"res":result})
			}else {
				c.JSON(http.StatusBadRequest,gin.H{"res":result})
			}
		})
		v1.POST("/login", func(c *gin.Context) {
			var user model.User
			err:=c.ShouldBind(&user)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{"error":err})
			}
			result:=service.CheckUser(user.Uname,user.Upassword)
			if result>0{
				middleware.SetSession(c,user.Uname)
				c.JSON(http.StatusOK,gin.H{"res":result})
			}else {
				c.JSON(http.StatusBadRequest,gin.H{"res":result})
			}
		})
	}
	router.Static("/templates", "./templates")
	router.Run(":8080")
}

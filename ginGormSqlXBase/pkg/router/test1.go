package router

import (
	"ginGormSqlXBase/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func user(engine *gin.Engine) {
	v1 := engine.Group("v1/test1")
	{
		v1.POST("/regUser", handlers.RegUserHandler)
		v1.POST("/getUser", handlers.GetUserHandler)
	}
}

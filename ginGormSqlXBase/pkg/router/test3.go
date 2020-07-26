package router

import (
	"ginGormSqlXBase/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func cache(engine *gin.Engine) {
	v1 := engine.Group("v1/test3")
	{
		v1.POST("/cahce", handlers.GetCacheHandler)
	}
}

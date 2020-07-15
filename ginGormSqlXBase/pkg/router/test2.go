package router

import (
	"ginGormSqlXBase/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func test(engine *gin.Engine) {
	v1 := engine.Group("v1/test2")
	{
		v1.POST("/getTests", handlers.GetTestsHandler)
	}
}

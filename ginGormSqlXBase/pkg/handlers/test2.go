package handlers

import (
	"ginGormSqlXBase/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"net/http"
	"time"
)

func GetTestsHandler(c *gin.Context) {
	//打印日志
	defer log.Infof("GetTestHandler\n,%v\n,%v\n", c, time.Now())
	//获取
	users := models.GetTests()
	c.JSON(http.StatusOK, gin.H{"res": "success", "data": users})
}

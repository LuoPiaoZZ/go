package handlers

import (
	"ginGormSqlXBase/pkg/define"
	"ginGormSqlXBase/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"net/http"
	"time"
)

func RegUserHandler(c *gin.Context) {
	//打印日志
	defer log.Infof("RegUserHandler,%s,%s", c, time.Now())
	//参数绑定
	req := &define.User{}
	if err := c.Bind(req); err != nil {
		log.Errorf("RegUserHandler,%s", "传入参数错误")
		return
	}
	//参数校验
	if err := req.Validate(); err != nil {
		log.Errorf("RegUserHandler,校验有误，%v", err)
		return
	}
	//注册用户
	num, err := models.InsertUser(*req)
	if err != nil {
		log.Errorf("fail:%v", err)
		return
	}
	if num == 0 {
		log.Errorf("fail")
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": "success", "data": req})
}

func GetUserHandler(c *gin.Context) {
	//打印日志
	defer log.Infof("GetUserHandler\n,%v\n,%v\n", c, time.Now())
	//获取用户
	users := models.GetUser()
	c.JSON(http.StatusOK, gin.H{"res": "success", "data": users})
}

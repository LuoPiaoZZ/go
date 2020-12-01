package handlers

import (
	"baseGinPro/pkg/define"
	"baseGinPro/pkg/models/psql"
	"baseGinPro/pkg/models/redis"
	"baseGinPro/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TestPsql(ctx *gin.Context)  {
	psql.CRUD()
	ctx.JSON(http.StatusOK, result(define.SuccessCode, define.StatusOk, "ok"))
}

func TestRedis(ctx *gin.Context)  {
	//限制请求次数，最好使用token，暂时使用IP
	ip:=ctx.ClientIP()
	limit := ctx.MustGet("limit").(*utils.LimitClick)
	if limit.Exists(ip) {
		ctx.JSON(http.StatusOK, result(define.ErrorCode, define.ErrRequest, nil))
		return
	}
	defer limit.Delete(ip)

	redis.SetGetDate(redis.Pool)
	time.Sleep(2*time.Second)//卡住，查看限制请求效果
	ctx.JSON(http.StatusOK, result(define.SuccessCode, define.StatusOk, "ok"))
}

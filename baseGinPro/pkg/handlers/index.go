package handlers

import (
	"baseGinPro/pkg/define"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespInfo 统一返回结构
type RespInfo struct {
	ErrNo  int      `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// 返回赋值
func result(code int, msg string, data interface{}) *RespInfo {
	return &RespInfo{
		ErrNo:  code,
		ErrMsg: msg,
		Data:   data,
	}
}

// Index 首页
func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, result(define.SuccessCode, define.StatusOk, "ok"))
}

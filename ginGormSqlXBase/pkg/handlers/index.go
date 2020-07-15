package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页
func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webDemo1/model"
	"webDemo1/service"
)

//注册
func Res(c *gin.Context)  (res bool,err error){
	var user model.User
	e:=c.ShouldBindJSON(&user)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return false,e
	}
	result:=service.InsertUser(user.Uname,user.Upassword)
	return result,nil
}
//登陆
func Log(c *gin.Context)  (res int,err error){
	var user model.User
	e:=c.ShouldBindJSON(&user)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return -1,e
	}
	result:=service.CheckUser(user.Uname,user.Upassword)
	return result,nil
}

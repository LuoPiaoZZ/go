package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

}
func authHandler(c *gin.Context)  {
	//接收前台的用户名
	username:=c.Query("username");
	if username=="lp" {
		//生成token
		tokenString,_:=GenToken(username)
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"success",
			"data":gin.H{
				"token":tokenString,
			},
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":2020,
		"msg":"鉴权失败",
	})
	return
}

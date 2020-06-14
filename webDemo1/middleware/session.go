package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


var Store=cookie.NewStore([]byte("user-secret"))
//session中间件
func SetSession(c *gin.Context,uname string)string  {
	session:=sessions.Default(c)
	sessionId:="uname"
	session.Set(sessionId,uname)
	session.Save()
	return sessionId
}
func GetSession(c *gin.Context,sessionId string)(uname string)  {
	session:=sessions.Default(c)
	result:=session.Get(sessionId)
	uname=result.(string)
	return
}
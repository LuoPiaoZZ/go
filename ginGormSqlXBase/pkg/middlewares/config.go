package middlewares

import (
	"ginGormSqlXBase/pkg/config"
	"ginGormSqlXBase/pkg/define"
	"github.com/gin-gonic/gin"
)

func Config(tomlConfig *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(define.StrConfig, tomlConfig)
		c.Next()
	}
}
func SetMiddleware(key string, value interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, value)
		c.Next()
	}
}

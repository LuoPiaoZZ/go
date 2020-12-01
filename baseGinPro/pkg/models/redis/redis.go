package redis

import (
	"baseGinPro/pkg/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	Pool *redis.Pool
)

func NewOpDataRedisConn(cacheName string, tomlConfig *config.Config) gin.HandlerFunc {
	cacheConfig, ok := tomlConfig.RedisServerConf(cacheName)
	if !ok {
		panic(fmt.Sprintf("%v not set.", cacheName))
	}
	// 连接数据库
	Pool = newPool(cacheConfig.Addr, cacheConfig.Password, cacheConfig.DB)

	return func(c *gin.Context) {
		c.Set(cacheName, Pool)
		c.Next()
	}

}

// newPool New redis pool.
func newPool(server, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

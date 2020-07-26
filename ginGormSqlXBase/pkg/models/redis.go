package models

import (
	"fmt"
	"ginGormSqlXBase/pkg/config"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	Pool *redis.Pool
)

// Redis 连接redis数据库
func NewRedisConn(cacheName string, tomlConfig *config.Config) {
	cacheConfig, ok := tomlConfig.RedisServerConf(cacheName)
	if !ok {
		panic(fmt.Sprintf("%v not set.", cacheName))
	}
	// 链接数据库
	Pool = newPool(cacheConfig.Addr, cacheConfig.Password, cacheConfig.DB)
}
// newPool New redis pool.
func newPool(server, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

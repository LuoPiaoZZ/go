package redis

import (
	"baseGinPro/pkg/config"
	"baseGinPro/pkg/define"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	UserCache *redis.Pool
)

//  NewUserCachePool New redis pool.
func NewUserCachePool(tomlConfig *config.Config) *redis.Pool {
	cacheConfig, ok := tomlConfig.RedisServerConf(define.MiddlewareUserCache)
	if !ok {
		panic(fmt.Sprintf("%v not set.", define.MiddlewareUserCache))
	}
	UserCache = &redis.Pool{
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cacheConfig.Addr,
				redis.DialDatabase(cacheConfig.DB),
				redis.DialPassword(cacheConfig.Password),
				redis.DialConnectTimeout(500*time.Millisecond),
				redis.DialReadTimeout(500*time.Millisecond),
				redis.DialWriteTimeout(500*time.Millisecond))
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
	return UserCache
}

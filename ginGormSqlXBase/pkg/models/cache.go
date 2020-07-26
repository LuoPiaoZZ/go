package models

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-xweb/log"
)

func Cache()  {
	redisConn := Pool.Get()
	key:="baseTestKey"
	isGetTestKey, err := redis.Bool(redisConn.Do("exists", key))
	if err != nil {
		log.Errorf("Error:%v", err)
		return
	}
	count:=0
	if isGetTestKey{
		log.Infof("存在")
		count, err = redis.Int(redisConn.Do("get", key))
		if err != nil {
			log.Errorf("Error: %v\n", err)
			return
		}
		log.Infof("读出key:%v",count)

		//再次存入
		_, err = redisConn.Do("set", key, count+1)
		if err != nil {
			log.Errorf("Error: %v\n", err)
			return
		}
		//读出存入
		count, err = redis.Int(redisConn.Do("get", key))
		if err != nil {
			log.Errorf("Error: %v\n", err)
			return
		}
		log.Infof("存入key:%v",count)
	}else {
		log.Infof("不存在")
		//存入
		_, err = redisConn.Do("set", key, count+1)
		if err != nil {
			log.Errorf("Error: %v\n", err)
			return
		}
		//读出存入
		count, err = redis.Int(redisConn.Do("get", key))
		if err != nil {
			log.Errorf("Error: %v\n", err)
			return
		}
		log.Infof("存入key:%v",count)
	}
}

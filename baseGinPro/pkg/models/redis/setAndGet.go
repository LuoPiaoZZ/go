package redis

import (
	"baseGinPro/pkg/define"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func SetGetDate(opDataCache *redis.Pool) (date string, ok bool, err error) {
	conn := opDataCache.Get()
	defer conn.Close()

	key := "lp"
	//set
	if _,err=redis.String(conn.Do("set",key,time.Now().Format(define.YMD)));err!=nil{
		if err != redis.ErrNil {
			return "", false, err
		}
		return "", false, nil
	}
	//get
	if date, err = redis.String(conn.Do("get", key)); err != nil {
		if err != redis.ErrNil {
			return "", false, err
		}
		return "", false, nil
	}
	fmt.Println(date)
	return date, true, nil
}

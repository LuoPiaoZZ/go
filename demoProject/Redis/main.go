package main

//redis操作演示
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main()  {
	//获取连接
	conn,err:=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil {
		fmt.Println("connect redis error:",err)
		return
	}
	defer conn.Close()
	//string的get和set
	_,err=conn.Do("SET","name","lp")
	if err!=nil {
		fmt.Println("redis set error:",err)
	}
	name,err:=redis.String(conn.Do("GET","name"))
	if err!=nil {
		fmt.Println("redis get error:",err)
	}else {
		fmt.Println("Get name:",name)
	}
	//设置key过期时间
	_,err=conn.Do("expire","name",5)//10second
	if err!=nil {
		fmt.Println("set expire error:",err)
		return
	}
	time.Sleep(5*time.Second)
	name1,err:=redis.String(conn.Do("GET","name"))
	if err!=nil {
		fmt.Println("redis get error:",err)
	}else {
		fmt.Println("Get name:",name1)
	}
}

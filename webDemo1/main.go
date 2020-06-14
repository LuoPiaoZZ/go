package main

import (
	"webDemo1/config"
	"webDemo1/router"
)

//注册登陆
func main()  {
defer config.DB.Close()
router.InitRouter()
}

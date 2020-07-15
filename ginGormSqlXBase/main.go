package main

import (
	"ginGormSqlXBase/pkg/config"
	"ginGormSqlXBase/pkg/define"
	"ginGormSqlXBase/pkg/models"
	"ginGormSqlXBase/pkg/router"
	"github.com/go-xweb/log"
	"runtime"
)

// init 初始化配置
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 入口函数
func main() {
	engine, tomlConfig, err := router.InitEngine()
	if err != nil {
		log.Errorf("main: 初始化engine出错, err:%v\n", err)
		return
	}

	// 初始化数据库连接
	initDB(tomlConfig)

	// 启动web服务
	log.Debugf("run ginGormSqlXBase at %v\n", tomlConfig.GetListenAddr())
	_ = engine.Run(tomlConfig.GetListenAddr())
}

// initDB 初始化数据库连接
func initDB(tomlConfig *config.Config) {
	// 设置db连接，gorm连接池

	models.NewMyDBConn(define.MyDB, tomlConfig)
	models.NewMyDBSQLXConn(define.MyDB, tomlConfig)

	models.NewMyDB1Conn(define.MyDB1, tomlConfig)

}

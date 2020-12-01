package main

import (
	"baseGinPro/internal/cron"
	"baseGinPro/pkg/config"
	"baseGinPro/pkg/define"
	"baseGinPro/pkg/models/psql"
	"baseGinPro/pkg/models/redis"
	"baseGinPro/pkg/router"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

// init 初始化配置
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 入口函数
func main() {
	defer func() {
		// 停止服务
		fmt.Println("stop server!")
		if err := recover(); err != nil {
			fmt.Errorf("error: %v \r\nStack: %v", err, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	flag.Parse()

	engine, tomlConfig, err := router.InitEngine()
	if err != nil {
		fmt.Errorf("main: 初始化engine出错, err:%v\n", err)
		return
	}
	config.Conf = tomlConfig

	// 初始化数据库连接
	initDB(tomlConfig)

	// 开启定时任务
	go cron.InitCron()

	// 启动web服务
	fmt.Printf("run baseGinPro at %v\n", tomlConfig.GetListenAddr())
	go engine.Run(tomlConfig.GetListenAddr())

	// 阻塞主进程 等待退出信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGKILL)
	<-signalChan
}

// initDB 初始化数据库连接
func initDB(tomlConfig *config.Config) {

	psql.NewLPDBConn(define.LPDB, tomlConfig)

	// 设置Redis连接
	redis.NewOpDataRedisConn(define.LPCache, tomlConfig)
	//token缓存
	redis.NewUserCachePool(tomlConfig)
}

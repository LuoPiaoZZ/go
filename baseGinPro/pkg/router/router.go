package router

import (
	"baseGinPro/pkg/config"
	"baseGinPro/pkg/handlers"
	"baseGinPro/pkg/middlewares"
	"baseGinPro/pkg/utils"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	tomlFilePath, mode string
)

// InitEngine 初始化engine
func InitEngine() (engine *gin.Engine, tomlConfig *config.Config, err error) {
	flag.StringVar(&tomlFilePath, "config", "docs/lp.toml", "服务配置文件")
	flag.StringVar(&mode, "mode", "release", "模型-debug还是release还是test")
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	gin.SetMode(mode)

	// 解析配置文件
	tomlConfig, err = config.UnmarshalConfig(tomlFilePath)
	if err != nil {
		fmt.Errorf("UnmarshalConfig: err:%v\n", err)
	}

	// 绑定路由，及公共的tomlConfig
	engine = gin.New()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"x-xq5-jwt", "Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"x-xq5-jwt", "Download-Status"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(gin.Recovery())

	//限制接口访问次数
	limit := utils.InitLimitClick()
	engine.Use(middlewares.SetMiddleware("limit", limit))

	// 路由配置
	engine.GET("/", handlers.Index) // 首页

	// 测试接口分组
	lpGroup(engine)

	return engine, tomlConfig, nil
}

func lpGroup(engine *gin.Engine) {
	v:=engine.Group("/baseGinPro")
	{
		//数据库操作
		v.GET("/test/psql",handlers.TestPsql)
		//redis操作
		v.GET("/test/redis",handlers.TestRedis)//http://localhost/baseGinPro/test/redis
	}
}

package router

import (
	"flag"
	"ginGormSqlXBase/pkg/config"
	"ginGormSqlXBase/pkg/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"runtime"
	"time"
)

var (
	tomlFilePath, mode string
)

// InitEngine 初始化engine
func InitEngine() (engine *gin.Engine, tomlConfig *config.Config, err error) {
	flag.StringVar(&tomlFilePath, "config", "docs/test.toml", "服务配置文件")
	flag.StringVar(&mode, "mode", "release", "模型-debug还是release还是test")
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	gin.SetMode(mode)

	// 解析配置文件
	tomlConfig, err = config.UnmarshalConfig(tomlFilePath)
	if err != nil {
		log.Errorf("UnmarshalConfig: err:%v\n", err)
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

	// 路由配置
	engine.GET("/", handlers.Index) // 首页

	// 测试1分组
	test1Group(engine)
	// 测试2分组
	test2Group(engine)

	test3Group(engine)

	return engine, tomlConfig, nil
}
func test3Group(engine *gin.Engine) {
	//获取redis缓存
	cache(engine)

}

func test2Group(engine *gin.Engine) {
	//获取test
	test(engine)

}

func test1Group(engine *gin.Engine) {
	//注册获取用户
	user(engine)
}

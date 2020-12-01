package psql

import (
	"baseGinPro/pkg/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	LPDB *gorm.DB
)

// 创建lpdb连接
func NewLPDBConn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	LPDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大连接数
	LPDB.DB().SetMaxOpenConns(10)
}

package models

import (
	"fmt"
	"ginGormSqlXBase/pkg/config"
	"ginGormSqlXBase/pkg/define"
	"github.com/go-xweb/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewMyDB1Conn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	MyDB1, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大链接数
	MyDB1.DB().SetMaxOpenConns(10)
}

// GetTests
func GetTests() []*define.Test {
	tests := make([]*define.Test, 0)
	var err = MyDB1.Table("test").Find(&tests).Error
	if err != nil {
		log.Errorf("数据库%v", err)
		return nil
	}
	return tests
}

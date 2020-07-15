package models

import (
	"fmt"
	"ginGormSqlXBase/pkg/config"
	"ginGormSqlXBase/pkg/define"
	"github.com/go-xweb/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

// 创建mydb连接

func NewMyDBConn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	MyDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大链接数
	MyDB.DB().SetMaxOpenConns(10)
}

func NewMyDBSQLXConn(dbName string, config *config.Config) {
	var err error
	dbCfg, ok := config.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	MyDBSqlx, err = sqlx.Open("postgres", dbCfg.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("sqlx.Open: err:%v", err))
	}

	// 设置最大空闲连接数
	MyDBSqlx.SetMaxIdleConns(15)
	return
}

// InsertUser 插入用户信息
func InsertUser(user define.User) (rowsAffected int64, err error) {
	psql := `insert into t_user(uname,upassword) values($1,$2)`
	num, err := MyDBSqlx.MustExec(psql, user.Uname, user.Upassword).RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, nil
}

// GetUser 获取用户信息
func GetUser() []*define.User {
	//defer MyDB.Close()
	users := make([]*define.User, 0)
	var err = MyDB.Table("t_user").Find(&users).Error
	if err != nil {
		log.Errorf("数据库%v", err)
		return nil
	}
	return users
}

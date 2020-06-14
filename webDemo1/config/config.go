package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const (
	//postgresql数据库参数
	host="localhost"
	post=5432
	user="postgres"
	password="admin"
	dbname="mydb"
)

//利用init函数来进行数据库连接的初始化
var DB *gorm.DB
func init()  {
	psqlInfo:=fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable ",host,post,user,password,dbname)
	var err error
	DB,err=gorm.Open("postgres",psqlInfo)
	if err!=nil {
		log.Fatal("数据库连接出错1",err)
	}
	if DB.Error!=nil {
		log.Fatal("数据库连接出错2",DB.Error)
	}
	//默认表名禁用复数
	DB.SingularTable(true)
	//修改默认表名的规则
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "t_" + defaultTableName
	}
}

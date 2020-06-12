package main

//orm框架gorm，对Postgres数据库进行操作
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"//1、连接方法和驱动选择一定要匹配
)

type User struct {
	ID int//2、一定要有一个主键，默认ID是主键
	Uname string//3、字段开头一定要大写，否则找不到
	Upassword string
}

const (
	host="localhost"
	post=5432
	user="postgres"
	password="admin"
	dbname="mydb"
)
func main()  {
	psqlInfo:=fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable ",host,post,user,password,dbname)
	db,err:=gorm.Open("postgres",psqlInfo)
	if err!=nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}else {
		fmt.Println("数据库连接成功")
	}
	defer db.Close()
	//这里匹配的表名是t_user
	//默认表名禁用复数
	db.SingularTable(true)
	//修改默认表名的规则
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "t_" + defaultTableName
	}
	//如果存在就自动迁移，不存在就自动创建表
	err = db.AutoMigrate(&User{}).Error
	//添加数据
	var u2 User=User{
		Uname: "1gorm",
		Upassword: "gormpassword",
	}
	db.Create(&u2)
	//读取
	var u1 User
	db.First(&u1)
	fmt.Println(u1.ID,u1.Uname,u1.Upassword)
	var u0 User
	db.Last(&u0,20)
	fmt.Println(u0.ID,u0.Uname,u0.Upassword)
	//更新
	db.Model(&u0).Update("uname","gorm").Update("upassword","111")
	db.Last(&u0,20)
	fmt.Println(u0.ID,u0.Uname,u0.Upassword)
	//删除
	//db.Delete(&User{})
}

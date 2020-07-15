package main

//orm框架gorm，对Postgres数据库进行操作
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"//1、连接方法和驱动选择一定要匹配
	"github.com/jmoiron/sqlx/types"
	"net/http"
	"zonst/logging"
)

type Test struct {
	Test types.JSONText `gorm:"test" json:"test"'`
}

const (
	host="localhost"
	post=5432
	user="postgres"
	password="admin"
	dbname="mydb"
)
func main()  {
	r := gin.Default()
	r.POST("/1", func(c *gin.Context) {
		test:=&Test{}
		if err:=c.Bind(test);err!=nil{
			logging.Errorf("err:%v",err)
			return
		}
		psqlInfo:=fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable ",host,post,user,password,dbname)
		db,err:=gorm.Open("postgres",psqlInfo)
		if err!=nil {
			fmt.Println(err)
			panic("数据库连接失败")
		}else {
			fmt.Println("数据库连接成功")
		}
		defer db.Close()

		err=db.Table("test").Create(test).Error
		if err!=nil {
			logging.Errorf("err:%v",err)
			return
		}
		c.String(http.StatusOK, "ok")
	})
	r.Run(":8082") //默认8080端口
}

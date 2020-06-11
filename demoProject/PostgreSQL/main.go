package main

//PostgreSQL的CRUD
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"//必须添加驱动
	"log"
)

const (
	host="localhost"
	post=5432
	user="postgres"
	password="admin"
	dbname="mydb"
)

type userLogin struct {
	uname string
	upassword string
}

type appContext struct {
	db *sql.DB
}

func main()  {
	c,err:=connectDB()
	defer c.db.Close()
	if err!=nil {
		fmt.Println(err)
	}
	c.insert("insert into t_user values('71','71')")//添加,没搞明白为什么添加成功会返回一个nil错误
	//c.updateDelete("update t_user set uname=$1 where upassword=$2","hello","123")//更新
	//c.updateDelete("delete from t_user where upassword=$1 ","11234")//删除
	//c.read("select * from t_user")//查询
	fmt.Println("action done!")
}
//数据库连接
func connectDB() (c *appContext, errMessage error) {
	psqlInfo:=fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable",host,post,user,password,dbname)
	db,err:=sql.Open("postgres",psqlInfo)
	if err!=nil {
		fmt.Println("连接数据库错误："+err.Error())
		return nil,err
	}else {
		fmt.Println("successfully connected!")
	}
	err=db.Ping()
	if err!=nil {
		fmt.Println("DBPing:"+err.Error())
		return nil,err
	}else {
		fmt.Println("DBPingSuccess")
	}
	//返回连接
	return &appContext{db},nil
}
//添加
func (c *appContext)insert(sqlStr string)  {
	err:=c.db.QueryRow(sqlStr)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("insert成功:"+sqlStr)
}
//查询
func (c *appContext)read(sqlStr string)  {
	rows,err:=c.db.Query(sqlStr)
	if err!=nil {
		fmt.Println("查询失败"+err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		p:=new(userLogin)
		err:=rows.Scan(&p.uname,&p.upassword)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(p.upassword,p.uname)
	}
	return
}
//更新或者删除
func (c *appContext)updateDelete(sqlStr string, stars ...string)  {
	stmt,err:=c.db.Prepare(sqlStr)
	if err!=nil {
		log.Fatal(err)
	}
	var result sql.Result
	if len(stars)==1 {
		result, _ =stmt.Exec(stars[0]) //删除只需要一个参数
	}else {
		result, _ =stmt.Exec(stars[0], stars[1]) //更新需要两个参数
	}
	if err!=nil {
		log.Fatal(err)
	}
	//影响行数
	affectNum,err:=result.RowsAffected()
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("affect rows is ",affectNum)
}

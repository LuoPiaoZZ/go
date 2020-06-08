package main

//带body的post请求
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/body", func(context *gin.Context) {
		//body:=context.Request.Body
		//bodyBytes,err:=ioutil.ReadAll(body)
		//if err!=nil{
		//	context.String(http.StatusBadRequest,err.Error())//404
		//	context.Abort()//结束请求
		//}
		////获取body里的参数,需要回填
		//body=ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		name := context.PostForm("name")
		age := context.PostForm("age")
		context.String(http.StatusOK, "%s,%s", age, name) //200
	})
	r.Run(":80")
	//关于post的测试可使用postman来模拟，如
	//curl --location --request POST '127.0.0.1/body' \
	//--header 'Content-Type: application/json' \
	//--data-raw '{
	//	"name":"罗飘",
	//	"age":20
	//}'
}

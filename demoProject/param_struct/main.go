package main

//表单结构体绑定及校验实现
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"time"
)

type Person struct {
	Name string "form:\"name\"" //双引号需要转义，最好就直接使用键盘左上角这个符号，注意不是英文单引号！！！
	//1、自定义校验方法
	Address  string    `form:"address" binding:"required,myvalidate"`
	Birthday time.Time "form:\"time\""
	Age      int       `form:"age" binding:"required,gt=10"` //定义校验规则,必填且大于10
}

//2、实现校验方法
func myvalidate(fl validator.FieldLevel) bool {
	if add := fl.Field().Len(); add > 10 {
		return true
	}
	return false
}

func main() {
	r := gin.Default()
	//3、注册校验方法到校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("myvalidate", myvalidate)
	}
	r.GET("/get", testing)
	r.POST("/post", testing)
	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	//根据请求的content-type来做不同的binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error,%v", err)
	}

}

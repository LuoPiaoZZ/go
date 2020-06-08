package main

//多语言校验实现
import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	en_translations "github.com/go-playground/validator/translations/en"
	zh_translations "github.com/go-playground/validator/translations/zh"
)

type Person struct {
	Name string "form:\"name\""                        //双引号需要转义，最好就直接使用键盘左上角这个符号，注意不是英文单引号！！！
	Age  int    `form:"age" validate:"required,gt=10"` //定义校验规则,必填且大于10
	//a、多语言翻译校验
	Score int `form:"score" validate:"required"`
}

//b、定义翻译器、校验器
var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func main() {
	r := gin.Default()
	//c、定义不同翻译器
	Validate = validator.New()
	zh := zh2.New()
	en := en2.New()
	Uni = ut.New(zh, en)
	r.POST("/post", func(c *gin.Context) {
		locale := c.DefaultPostForm("locale", "zh")
		trans, _ := Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}
		var person Person
		//绑定失败错误
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}
		//d、校验错误
		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(500, "%v", sliceErrs)
			c.Abort()
			return
		}
		c.String(200, "%v", person)
	})
	r.Run()
}

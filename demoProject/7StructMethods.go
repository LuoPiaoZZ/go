package main

//结构体及函数
import "fmt"

//相当于一个Class
type Person struct {
	//定义字段
	name string
	age int
}
//定义实现方法
func (per *Person)hello()string  {
	return fmt.Sprintf("你好,我是%s,今年%d岁",per.name,per.age)
}

func main(){
	//实例化
	per0:=&Person{
		name: "windy",
		age:18,
	}
	//调用方法
	strHello:=per0.hello()
	fmt.Println(strHello)

}
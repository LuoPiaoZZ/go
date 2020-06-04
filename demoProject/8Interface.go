package main

//接口
import (
	"fmt"
)

//一般而言，接口定义了一组方法，接口不能被实例化，一个类型可以实现多个接口
type Person1 interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	//接口实现
	var p Person1 = &Student{
		name: "windy",
		age:  18,
	}
	fmt.Println(p.getName())
	//空接口：如果定义了一个没有任何方法的接口，那么这个接口可以表示任意类型
	fmt.Println("空接口实现：")
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 12
	fmt.Println(m)
}

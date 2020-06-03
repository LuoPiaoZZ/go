package main

//数组、切片、键值对、指针
import (
	"fmt"
)

func main() {
	var arr0 = [5]int{0, 1, 2, 3, 4} //一维数组，声明时初始化
	arr0[0] = 1                      //修改值
	fmt.Println(arr0[0] + arr0[1])   //1+1=2
	//切片是数组的抽象，容量可以自动扩充
	slice0 := make([]int, 3, 10)          //声明一个长度为3，容量为10的切片
	slice1 := make([]int, 5, 10)          //声明一个长度为5，容量为10的切片
	slice0 = append(slice0, 0, 1)         //长度变更为5,3+2
	slice1 = append(slice0, 2, 3, 4)      //长度变更为8，5+3
	fmt.Println(len(slice0), cap(slice0)) //长度为5，容量为10
	fmt.Println(len(slice1), cap(slice1)) //长度为8，容量为10

	//字典map
	map1 := make(map[string]int)
	map1["windy"] = 0
	map2 := map[string]string{
		"Sunday":   "happy",
		"Saturday": "happy+1", //注意即使是最后一个键值对也要写逗号结束
	}
	fmt.Println(map2["Sunday"], map2["Saturday"], map1["windy"])

	//指针pointer，某个值的地址
	str := "123456"
	var p *string = &str //p表示为str的地址
	*p = "654321"        //*p表示指针p对应的value
	fmt.Println(str, *p) //将会修改str变量的值

}

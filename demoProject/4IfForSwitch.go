package main

//控制流程
import "fmt"

func main() {

	age := 12
	if age > 15 {
		fmt.Println("年龄大于15")
	} else if age == 12 {
		fmt.Println("年龄小于15，且为", age, "岁")
	} else {
		fmt.Println("年龄小于15，且不为12岁")
	}
	//声明一个类型别名，没有初始化
	type Gender int
	//定义几个结构体常量
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)
	//gen:=FEMALE//相当于var gen int = 2
	gen := MALE
	switch gen {
	case FEMALE:
		fmt.Println("女", FEMALE)
		//默认自带了break
	case MALE:
		fmt.Println("男", MALE)
		fallthrough //取消break需要使用这个关键字
	default:
		fmt.Println("其他")
	}
	//遍历数组、切片、键值对
	nums := [4]int{1, 2, 3, 4}
	for i, num := range nums {
		fmt.Println("数组", i, num)
	}
	slien0 := make([]int, 0, 6)
	slien0 = append(slien0, 0, 1, 2)
	for i, num := range slien0 {
		fmt.Println("切片", i, num)
	}
	map0 := map[string]string{
		"0map": "0",
		"1map": "1",
	}
	for k, v := range map0 {
		fmt.Println("字典", k, v)
	}
	//累加
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
		if sum >= 5 {
			break
		}
	}
	fmt.Println("累加结果为", sum) //5
	//死循环
	for true {
		sum++
		if sum >= 10 {
			break
		}
	}
	fmt.Println("累加结果为", sum) //10
}

package main

//简单函数
import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println("加结果为", sum) //3
	quotient, remainder := div(9, 4)
	fmt.Println("商为", quotient, "余为", remainder) //2,1
}

//函数名后面，第一个是参数（参数名，参数类型），第二个是返回值，返回值可以不需要参数名
func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

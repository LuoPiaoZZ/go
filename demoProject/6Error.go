package main

//自定义异常以及捕捉异常
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(because(""))
	get(5) //数组越界
}

//自定义异常
func because(because string) error {
	if len(because) == 0 {
		return errors.New("自定义的异常：输入的原因为空")
	}
	fmt.Println("正确", because)
	return nil //空值
}

//定义一个获取数组相应下标结果的函数
func get(index int) (returnValue int) {
	//捕获异常
	defer func() {
		if er := recover(); er != nil {
			fmt.Println("捕捉到出现的错误", er)
			returnValue = -1 //出现异常返回-1
		}
	}()
	arr := [2]int{1, 2}
	return arr[index]
}

package main

import "testing"

//因为包里面main方法只能有一个的原因，使用test命令行运行测试会报错
//测试指令如：  go test -v

func addT(num1 int, num2 int) int {
	return num1 + num2
}
func testAdd(t *testing.T) {
	if sum := addT(1, 2); sum != 3 {
		t.Error("1+2等于3")
	}
}

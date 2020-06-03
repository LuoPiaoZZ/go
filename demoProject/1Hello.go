package main //一定要有且只有一个main包
//第一个go程序
import "fmt" //引入输入输出标准库

func main() { //一定要有且只有一个main入口方法，特别约定为func main()
	fmt.Print("Hello,世界")
}

//针对多个demo可以选择在终端打开使用go run xxx.go来进行运行，避免go build同一文件夹产生的冲突，每次都要修改删除

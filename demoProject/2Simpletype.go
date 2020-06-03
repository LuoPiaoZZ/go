package main

//数据声明方式
import "fmt"

func main() {
	var a int8
	var b int8
	fmt.Printf("%d,%d\n", a, b) //0,0
	a, b = 1, 2
	fmt.Printf("%d,%d\n", a, b) //1,2
	c := 3
	d := "three三"
	const e bool = true
	f := false
	fmt.Printf("%d,%s,%c,%t,%t\n", c, d, d[0], e, f)
	var g = 'g'
	fmt.Printf("%c\n", g)
	//	注意问题，d[5]乱码，因为string是byte数组形式保存的，类型uint8，占1个byte，打印时需要进行类型转化
	runeArr := []rune(d)           //将字符串转化为rune数组
	fmt.Printf(string(runeArr[5])) //正常打印出中文，三

}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//锁
func main()  {
	//testMap1()
	//testMap2()
	//testMap3()
}
//1、不加锁
func testMap1(){
	//初始化赋值
	var a map[int]int
	a=make(map[int]int,5)
	for i:=0;i<5;i++ {
		a[i]=i
	}
	//竞争
	for i:=0;i<2;i++{
		go func(b map[int]int) {
			//写操作
			b[0]=rand.Intn(100)
		}(a)
	}
	//读操作
	fmt.Println(a)
}
//2、互斥锁
var lock1 sync.Mutex
func testMap2(){
	//初始化赋值
	var a map[int]int
	a=make(map[int]int,5)
	for i:=0;i<5;i++ {
		a[i]=i
	}
	//竞争
	for i:=0;i<2;i++{
		go func(b map[int]int) {
			//写操作
			lock1.Lock()
			b[0]=rand.Intn(100)
			lock1.Unlock()
		}(a)
	}
	//读操作
	lock1.Lock()
	fmt.Println(a)
	lock1.Unlock()

	time.Sleep(time.Second)
}
//3、读写锁
var lock2 sync.RWMutex
func testMap3(){
	//初始化赋值
	var a map[int]int
	a=make(map[int]int,5)
	for i:=0;i<5;i++ {
		a[i]=i
	}
	//a、写操作
	for i:=0;i<5;i++{
		go func(b map[int]int) {
			lock2.Lock()
			b[0]=rand.Intn(100)
			lock2.Unlock()
		}(a)
	}
	//b、读操作
	for i:=0;i<100;i++{
		go func() {
			lock2.RLock()
			fmt.Println(a)
			lock2.RUnlock()
		}()
	}

	time.Sleep(5*time.Second)

}




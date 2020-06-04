package main

//信道，可以阻塞等待并发协程返回消息，在协程之间返回消息
import (
	"fmt"
	"time"
)

var ch =make(chan string,10)//创建大小为10的缓冲通道

func download1(url string){
	fmt.Println("strat to download",url)
	time.Sleep(time.Second)
	ch<-url//将URL传给信道
}
func main(){
	for i:=0;i<3;i++ {
		go download1("a.com/"+string(i+'0'))
	}
	for i:=0;i<3;i++ {
		msg:=<-ch//等待信道返回的消息
		fmt.Println("finsh",msg)
	}
	fmt.Println("Done!")
}

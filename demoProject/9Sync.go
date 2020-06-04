package main

//一组线程执行完以后再执行后面的操作
//应用举例：如在下载中可以将各个下载文件并发下载，但是全部下载完成以后再通知用户
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string)  {
	fmt.Println("start to download",url)
	time.Sleep(1*time.Second)//延时1秒，模拟耗时操作
	wg.Done()
}

func main(){
	for i:=0;i<3;i++ {
		wg.Add(1)//添加到被等待的线程组里面
		go download("a.com"+string(i+'0'))
	}
	wg.Wait()//确保以上操作并发执行完以后再执行下面操作
	fmt.Println("Done!")
}


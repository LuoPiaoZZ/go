package main

//net包
//要查看客户端是否可以接收成功：
//1、可以使用brew安装一个netcat，在netcat里面输入对应地址即可，完整指令：nc localhost 8000
//2、使用telnet简单创建一个tcp连接
import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil { //连接错误打印日志，一直尝试连接
			log.Print(err)
			continue
		}
		go handleConn(conn) //使用并发以后就可以让多个客户端可以收到日期打印消息
	}
}

func handleConn(c net.Conn) {
	defer c.Close() //异常关闭连接
	for {
		//注意：go的格式化字符串是1234567即1月2日下午3点4分5秒06年UTC-0700
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

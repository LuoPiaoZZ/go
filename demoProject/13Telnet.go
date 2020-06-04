package main

//telnet程序
//接收服务端信息打印出来的客户端，可以和nc一起模仿多个客户端并发
import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	//将输出写入到控制台
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

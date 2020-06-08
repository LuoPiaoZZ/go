package main

//优雅关停
import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//优雅关停服务器
func main() {
	r := gin.Default()
	r.GET("/shoutdown", func(context *gin.Context) {
		//模拟超时请求
		time.Sleep(9 * time.Second)
		context.String(200, "超时输出")
	})
	server := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen%s", err)
		}
	}()
	//将当前状态放入信道
	quit := make(chan os.Signal)
	//将输入信号转发到quit，只转发列出来的两种
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shoutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server shutdown", err)
	}
	fmt.Println("server exiting")
}

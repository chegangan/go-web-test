package main

import (
	"context"
	"fmt"
	"github.com/fvbock/endless"
	"go-web-test/pkg/setting"
	"go-web-test/routers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           go-web-test API
// @version         0.5
// @description     This is a simple test of go web server
/**
* @BasePath   routers/api/v1
* @host      localhost:8000
 */
func main() {
	// endless实现优雅的热重启
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

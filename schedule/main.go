package schedule

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"schedule/common"
	"schedule/schedule/bootstrap"
	"syscall"
	"time"
)

var (
	configPath string
)

func InitArgs() {
	flag.StringVar(&configPath, "c", "./config.yaml", "config file path")
	flag.Parse()
}

func main() {

	v, err := common.InitConfig(configPath)
	if err != nil {
		log.Println("config init error:", err)
		return
	}

	app, err := bootstrap.App(&v)
	if err != nil {
		log.Println("app init error:", err)
		return
	}

	go func() {
		if err = app.HttpServer.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("http serve error:", err)
			return
		}
	}()

	// 监听信号
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// 停止服务
	shutdownCtx, shutdownCancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer shutdownCancel()
	if err = app.HttpServer.HttpServer.Shutdown(shutdownCtx); err != nil {
		log.Println("http shutdown error:", err)
		return
	}

}

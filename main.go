package main

import (
	"context"
	"fmt"
	"go-generator/dao/mysql"
	"go-generator/dao/redis"
	"go-generator/logger"
	"go-generator/routers"
	"go-generator/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%#v\n", err)
		return
	}
	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%#v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//连接数据库
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%#v\n", err)
		return
	}
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%#v\n", err)
		return
	}
	//注册路由
	r := routers.Setup()
	//优雅关机

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1) //
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	//超过10s强制关闭
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}

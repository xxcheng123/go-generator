package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
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

	"go.uber.org/zap"
)

var configFile string
var configPath string
var configName string
var configType string

func init() {
	flag.StringVar(&configFile, "f", "", "配置文件全路径，如果配置了当前路径，会忽略其他参数")
	flag.StringVar(&configPath, "p", ".", "配置文件路径")
	flag.StringVar(&configName, "n", "config", "配置文件名字")
	flag.StringVar(&configType, "t", "", "配置文件格式")
}

func main() {
	flag.Parse()
	//configPath, _ = filepath.Abs(configPath)
	//fmt.Printf("Config Name:%s,Path:%s\n", configFile, configPath)
	viperConfig := settings.ViperConfig{
		ConfigFile: configFile,
		ConfigPath: configPath,
		ConfigName: configName,
		ConfigType: configType,
	}
	//加载配置
	if err := settings.Init(&viperConfig); err != nil {
		fmt.Printf("init settings failed, err:%#v\n", err)
		return
	}
	//初始化日志
	if err := logger.Init(settings.Config.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%#v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//连接数据库
	if err := mysql.Init(settings.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%#v\n", err)
		return
	}
	defer mysql.Close()
	if err := redis.Init(settings.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%#v\n", err)
		return
	}
	defer redis.Close()
	//注册路由
	r := routers.Setup(settings.Config)
	//测试版本号查看
	r.GET("/version", func(ctx *gin.Context) {
		ctx.String(200, fmt.Sprintf("current Version:%s", settings.Config.Version))
	})
	//优雅关机

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Config.Port),
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

package logger

import (
	"go-generator/settings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 获取日志格式
func getLogEncoder() zapcore.Encoder {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	return encoder
}

// 获取输入
func getLogWriteSyncer(filename string, maxSize int, maxAge int, maxBackups int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Init(cfg *settings.LogConfig) (err error) {
	encoder := getLogEncoder()
	writeSyncer := getLogWriteSyncer(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxAge,
		cfg.MaxBuckups,
	)
	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}
	zapCore := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(zapCore, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return err
}

func GinLogger() gin.HandlerFunc {
	return ginzap.Ginzap(zap.L(), time.RFC3339, true)
}

func GinRecover() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(zap.L(), true)
}

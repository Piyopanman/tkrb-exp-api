package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//LogMode ロギングのモード
type logMode string

const (
	//Dev 開発用
	Dev logMode = "development"
	//Prod 本番用
	Prod logMode = "production"
)

//ILogger Loggerのインターフェース
type ILogger interface {
	NewLogger() *zap.Logger
}

//DLogger 開発用loggerの構造体
type DLogger struct{}

//PLogger 本番用loggerの構造体
type PLogger struct{}

var Logger *zap.Logger

//InitLogger 開発用か本番用でLoggerを初期化
func InitLogger(mode logMode) {
	if mode == Dev {
		var dLogger ILogger = DLogger{}
		Logger = dLogger.NewLogger()
	} else if mode == Prod {
		var pLogger ILogger = PLogger{}
		Logger = pLogger.NewLogger()
	}
}

//NewLogger 開発用loggerの設定
func (dLogger DLogger) NewLogger() *zap.Logger {
	devLogConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "Msg",
			LevelKey:       "Level",
			TimeKey:        "Time",
			CallerKey:      "Caller",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stderr"},
	}
	devLogger, _ := devLogConfig.Build()
	devLogger.Debug("development mode")
	return devLogger
}

//NewLogger 本番用loggerの設定
func (pLogger PLogger) NewLogger() *zap.Logger {
	proLogConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "Msg",
			LevelKey:       "Level",
			TimeKey:        "Time",
			CallerKey:      "Caller",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"./product.log"},
	}
	proLogger, _ := proLogConfig.Build()
	proLogger.Debug("production mode")
	return proLogger
}

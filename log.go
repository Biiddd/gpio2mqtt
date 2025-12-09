package main

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger(config LogConfig) *zap.SugaredLogger {
	level := parseLogLevel(config.Level)

	// lumberjack 用于日志轮转
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.File,       // 日志文件路径
		MaxSize:    config.MaxSize,    // 单个日志文件最大大小（MB）
		MaxBackups: config.MaxBackups, // 最多保留多少个备份文件
		MaxAge:     config.MaxAge,     // 日志保留天数
		Compress:   config.Compress,   // 是否压缩旧日志
	}

	// zap 编码配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色输出
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 文件名:行号
	}

	fileWriter := zapcore.AddSync(lumberJackLogger)
	consoleWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, level),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return logger.Sugar()
}

func parseLogLevel(lvl string) zapcore.Level {
	switch strings.ToLower(lvl) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

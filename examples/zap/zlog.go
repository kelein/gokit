package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zlog *zap.Logger

// Sugar for default logger
var Sugar *zap.SugaredLogger


var infoFileConfig = lumberjack.Logger{
	Filename:   "info.log",
	MaxSize:    1024 * 2, // 2GB
	MaxBackups: 5,
	MaxAge:     7,
	Compress:   true,
}

var errorFileConfig = lumberjack.Logger{
	Filename:   "error.log",
	MaxSize:    1024 * 2, // 2GB
	MaxBackups: 5,
	MaxAge:     7,
	Compress:   true,
}

// Setting return a optional logger instance
func Setting() *zap.Logger {
	cores := []zapcore.Core{}

	config := zap.NewProductionEncoderConfig()
	// config.EncodeTime = zapcore.ISO8601TimeEncoder
	// config.EncodeTime = zapcore.RFC3339TimeEncoder
	// config.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncodeCaller = zapcore.ShortCallerEncoder
	// config.EncodeCaller = zapcore.FullCallerEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(config)

	high := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})

	low := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.DebugLevel && level < zap.ErrorLevel
	})

	infoFileSyncer := zapcore.AddSync(&infoFileConfig)
	errorFileSyncer := zapcore.AddSync(&errorFileConfig)

	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(
		infoFileSyncer, zapcore.AddSync(os.Stdout)), low)

	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(
		errorFileSyncer, zapcore.AddSync(os.Stdout)), high)

	cores = append(cores, infoFileCore, errorFileCore)

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}

// Sync calls zap Sync method to flush buffered log entries.
// Applications should take care to call Sync before exit.
func Sync() { zlog.Sync() }

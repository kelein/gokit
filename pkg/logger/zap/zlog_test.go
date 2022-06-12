package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestSetting(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zlog := Setting()
			zlog.Debug("Test Zap Logger Debug", zap.String("func", "Test"), zap.Float64("value", 123))
			zlog.Info("Test Zap Logger Info", zap.String("func", "Test"), zap.Float64("value", 123))
			zlog.Warn("Test Zap Logger Warn", zap.String("func", "Test"), zap.Float64("value", 123))
			zlog.Error("Test Zap Logger Error", zap.String("func", "Test"), zap.Float64("value", 123))
			zlog.DPanic("Test Zap Logger Panic", zap.String("func", "Test"), zap.Float64("value", 123))
			zlog.Sync()
		})
	}
}

func TestNewColoredConsoleEncoder(t *testing.T) {
	tests := []struct{ name string }{
		{"A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zlog := Setting()
			zlog.Debug("message with level debug",
				zap.String("func", "debug"), zap.Int8("value", int8(zapcore.DebugLevel)))
			zlog.Info("message with level info",
				zap.String("func", "info"), zap.Int8("value", int8(zapcore.InfoLevel)))
			zlog.Warn("message with level warn",
				zap.String("func", "warn"), zap.Int8("value", int8(zapcore.WarnLevel)))
			zlog.Error("message with level error",
				zap.String("func", "error"), zap.Int8("value", int8(zapcore.ErrorLevel)))
			zlog.DPanic("message with level panic",
				zap.String("func", "panic"), zap.Int8("value", int8(zapcore.DPanicLevel)))
			zlog.Sync()
		})
	}
}

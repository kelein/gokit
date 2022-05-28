package logger

import (
	"testing"

	"go.uber.org/zap"
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
			zlog.Info("Test Zap Logger Info")
			zlog.Warn("Test Zap Logger Warn")
			zlog.Error("Test Zap Logger Error")
			zlog.Debug("Test Zap Logger Debug")
			zlog.Error("Test Zap With Field", zap.String("func", "Test"), zap.Float64("value", 123))
		})
	}
}

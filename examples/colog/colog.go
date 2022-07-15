package main

import (
	"fmt"
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Foreground colors.
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

// ColoredTimeEncoder custom encoder for time field
func ColoredTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// s := t.Format("2006-01-02T15:04:05.000")
	// s := t.Format(time.ISO8601TimeEncoder)
	s := t.Format("2006-01-02T15:04:05.000")
	enc.AppendString(Green.Add(s))
}

// ColoredShortCallerEncoder for colored shortCallerEncoder
func ColoredShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	s := Blue.Add(caller.TrimmedPath())
	enc.AppendString(fmt.Sprintf("[%s]", s))
}

func main() {
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// encoder.EncodeCaller = zapcore.ShortCallerEncoder
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	// * Add Custom Encoder
	encoder.EncodeTime = ColoredTimeEncoder
	encoder.EncodeCaller = ColoredShortCallerEncoder

	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		),

		zap.AddCaller(),
	)

	logger.Info("TEST ZAP LOG Level Info")
	logger.Warn("TEST ZAP LOG Level Warn")
	logger.Error("TEST ZAP LOG Level Error")
	logger.Debug("TEST ZAP LOG Level Debug")
}

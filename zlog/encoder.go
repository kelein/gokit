package zlog

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// Terminal Colors
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

// Add adds the color to the given string
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

// Any adds the color to any type
func (c Color) Any(s interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", uint8(c), s)
}

// Bold adds a bold color to the given string
func (c Color) Bold(s string) string {
	return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", uint8(c), s)
}

var coloredEncoderOption = map[string]Color{
	"time":   Green,
	"level":  Black,
	"caller": Magenta,
}

var levelColor = map[zapcore.Level]Color{
	zapcore.DebugLevel:  Green,
	zapcore.InfoLevel:   White,
	zapcore.WarnLevel:   Yellow,
	zapcore.ErrorLevel:  Red,
	zapcore.DPanicLevel: Blue,
	zapcore.PanicLevel:  Blue,
	zapcore.FatalLevel:  Blue,
}

// ColoredTimeEncoder custom encoder for time field
func ColoredTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	s := coloredEncoderOption["time"].Add(t.Format(time.RFC3339))
	enc.AppendString(s)
}

// ColoredTimeEncoderWithLayout custom encoder for time field with a layout
func ColoredTimeEncoderWithLayout(layout string) zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		s := coloredEncoderOption["time"].Add(t.Format(layout))
		enc.AppendString(s)
	}
}

// ColoredShortCallerEncoder for colored shortCallerEncoder
func ColoredShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	ln := strconv.Itoa(caller.Line)
	fn := strings.TrimSuffix(caller.TrimmedPath(), ln)
	buf := bufPool.Get()
	buf.AppendString("[")
	buf.WriteString(Magenta.Add(fn))
	buf.WriteString(Cyan.Add(ln))
	buf.WriteString("]")
	s := buf.String()
	buf.Free()
	enc.AppendString(s)
}

// ColoredCapitalLevelEncoder for colored capital level
func ColoredCapitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	s := coloredEncoderOption["level"].Bold(l.CapitalString())
	enc.AppendString(s)
}

var bufPool = buffer.NewPool()

const coloredConsoleEncoder = "colored-console"

// ColoredConsoleEncoder encode logger entry with color
type ColoredConsoleEncoder struct {
	*zapcore.EncoderConfig
	zapcore.Encoder
}

// NewColoredConsoleEncoder create ColoredConsoleEncoder instance
func NewColoredConsoleEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return ColoredConsoleEncoder{
		EncoderConfig: &cfg,
		Encoder:       zapcore.NewConsoleEncoder(cfg),
	}
}

// EncodeEntry encode each log field
func (c ColoredConsoleEncoder) EncodeEntry(
	entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	colorFn := levelColor[entry.Level].Add
	if entry.Level > zapcore.ErrorLevel {
		colorFn = levelColor[entry.Level].Bold
	}
	entry.Message = colorFn(entry.Message)
	return c.Encoder.EncodeEntry(entry, fields)
}

// RegisterColorConsoleEncoder register an encoder constructor for zapcore
func RegisterColorConsoleEncoder() {
	zap.RegisterEncoder(coloredConsoleEncoder,
		func(config zapcore.EncoderConfig) (zapcore.Encoder, error) {
			return NewColoredConsoleEncoder(config), nil
		})
}

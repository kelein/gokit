package zlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/lumberjack.v2"
)

var std = New()

// Sugar for std logger
var Sugar = std.Sugar

// Logger for zlog
// type Logger struct {
// 	zlog  *zap.Logger
// 	level Level
// }

// New create a logger instance
func New() *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeCaller = ColoredShortCallerEncoder
	cfg.EncodeLevel = ColoredCapitalLevelEncoder
	cfg.EncodeTime = ColoredTimeEncoderWithLayout("2006-01-02 15:04:05")

	// * Custom Encoder With Color
	encoder := NewColoredConsoleEncoder(cfg)
	syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	levelFn := LevelEnablerFunc(func(l Level) bool { return l >= DebugLevel })
	core := zapcore.NewCore(encoder, syncer, levelFn)
	// zlog := zap.New(core, WithCaller(true), WithTrace(FatalLevel))
	// return &Logger{zlog: zlog, level: level}
	return zap.New(core, WithCaller(true), WithTrace(FatalLevel))
}

// NewRotateLogger create rotate logger instance
func NewRotateLogger(tops []TeeOption, ops ...Option) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeCaller = ColoredShortCallerEncoder
	cfg.EncodeLevel = ColoredCapitalLevelEncoder
	cfg.EncodeTime = ColoredTimeEncoderWithLayout("2006-01-02 15:04:05")

	// * Custom Encoder With Color
	encoder := NewColoredConsoleEncoder(cfg)

	cores := []zapcore.Core{}
	for i := range tops {
		opt := tops[i]
		syncer := zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(&opt.Rotation),
			zapcore.AddSync(os.Stdout))
		core := zapcore.NewCore(encoder, syncer, opt.LevelFunc)
		cores = append(cores, core)
	}

	return zap.New(
		zapcore.NewTee(cores...),
		WithCaller(true),
		WithTrace(FatalLevel),
	)
}

// Default create a default logger
func Default() *zap.Logger { return std }

// Sync flushing buffered log of standard logger
func Sync() error {
	if std == nil {
		return nil
	}
	return std.Sync()
}

// Sync flushing buffered log which must be called before exit
// func (l *Logger) Sync() error {
// 	return l.zlog.Sync()
// }

// Sugar create a SugaredLogger
// func (l *Logger) Sugar() *zap.SugaredLogger {
// 	return l.zlog.Sugar()
// }

// ResetDefault replace default logger
// func ResetDefault(l *Logger) {
// 	std = l
// 	Info = std.Info
// 	Warn = std.Warn
// 	Error = std.Error
// 	DPanic = std.DPanic
// 	Fatal = std.Fatal
// 	Debug = std.Debug
// }

// ResetDefault replace default logger
func ResetDefault(l *zap.Logger) {
	std = l
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Fatal = std.Fatal
	Debug = std.Debug
}

// Debug logs a message at DebugLevel
// func (l *Logger) Debug(msg string, fileds ...Field) {
// 	l.zlog.Debug(msg, fileds...)
// }

// Info logs a message at DebugLevel
// func (l *Logger) Info(msg string, fileds ...Field) {
// 	l.zlog.Info(msg, fileds...)
// }

// Warn logs a message at WarnLevel
// func (l *Logger) Warn(msg string, fileds ...Field) {
// 	l.zlog.Warn(msg, fileds...)
// }

// Error logs a message at ErrorLevel
// func (l *Logger) Error(msg string, fileds ...Field) {
// 	l.zlog.Error(msg, fileds...)
// }

// Panic logs a message at PanicLevel
// func (l *Logger) Panic(msg string, fileds ...Field) {
// 	l.zlog.Panic(msg, fileds...)
// }

// DPanic logs a message at DPanicLevel
// func (l *Logger) DPanic(msg string, fileds ...Field) {
// 	l.zlog.DPanic(msg, fileds...)
// }

// Fatal logs a message at FatalLevel
// func (l *Logger) Fatal(msg string, fileds ...Field) {
// 	l.zlog.Fatal(msg, fileds...)
// }

// Level zlog logger level
type Level = zapcore.Level

// Zlog Logger Level
const (
	DebugLevel  Level = zap.DebugLevel
	InfoLevel   Level = zap.InfoLevel
	WarnLevel   Level = zap.WarnLevel
	ErrorLevel  Level = zap.ErrorLevel
	DPanicLevel Level = zap.DPanicLevel
	PanicLevel  Level = zap.PanicLevel
	FatalLevel  Level = zap.PanicLevel
)

// Field is an alias for zlog Field
type Field = zap.Field

// Zlog Field Definition
var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any
)

// Logger function alias
var (
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Fatal  = std.Fatal
	Debug  = std.Debug
)

// Option configures a Logger
type Option = zap.Option

// Logger Option
var (
	WithCaller = zap.WithCaller
	WithTrace  = zap.AddStacktrace
)

// LevelEnablerFunc implement LevelEnabler
type LevelEnablerFunc = zap.LevelEnablerFunc

// Rotation option for logger rotate
type Rotation = lumberjack.Logger

// DefaultRotation with default rotate option.
// MaxSize: 2GB MaxBackups: 5 MaxAge: 7 days
var DefaultRotation = func(filename string) Rotation {
	return Rotation{
		Filename:   filename,
		MaxSize:    2 << 10,
		MaxAge:     7,
		MaxBackups: 5,
		Compress:   true,
		LocalTime:  true,
	}
}

// TeeOption for writer syncer
type TeeOption struct {
	Rotation  Rotation
	LevelFunc LevelEnablerFunc
}

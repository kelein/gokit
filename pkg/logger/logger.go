package logger

import (
	"fmt"

	"github.com/pkg/errors"
)

// Log Driver Type
const (
	ZAP    string = "zap"
	LOGRUS string = "logrus"
)

// Logger for common interface
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

// Init ...
func Init(opt *Option) error {
	log, err := Build(opt)
	if err != nil {
		return errors.Wrap(err, "failed to build logger")
	}
	setLogger(log)
	return nil
}

func setLogger(l Logger) {}

var logBuilders = make(map[string]LogBuilder)

// LogBuilder for build logger interface
type LogBuilder interface {
	Build(opt *Option) (Logger, error)
}

// RegisterLogBuilder register all kinds of log builder
func RegisterLogBuilder(builderType string, builder LogBuilder) {
	if _, ok := logBuilders[builderType]; ok {
		panic(fmt.Sprintf("logger builder %q already registered", builderType))
	}
	logBuilders[builderType] = builder
}

// Build get log factory and make logger instance
func Build(opt *Option) (Logger, error) {
	logger, err := getLogBuilder(opt.Code).Build(opt)
	if err != nil {
		return nil, errors.Wrap(err, "logger build failed")
	}
	return logger, nil
}

func getLogBuilder(key string) LogBuilder {
	return logBuilders[key]
}

// Option ...
type Option struct {
	Code      string `json:"code,omitempty" yaml:"code,omitempty"`
	Level     string `json:"level,omitempty" yaml:"level,omitempty"`
	AddCaller bool   `json:"addCaller,omitempty" yaml:"addCaller,omitempty"`
}

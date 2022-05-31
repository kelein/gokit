package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/kelein/gokit/pkg/logger"
)

var logrusBuilder Builder

const logrusBuilderType = "logrus"

func init() {
	logger.RegisterLogBuilder(logrusBuilderType, &logrusBuilder)
}

// Builder implements Logger Builder interface
type Builder struct{}

// Build make a logrus logger instance
func (b *Builder) Build(opt *logger.Option) (logger.Logger, error) {

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetLevel(logrus.InfoLevel)
	// log.SetOutput(os.Stdout)
	return log, nil
}

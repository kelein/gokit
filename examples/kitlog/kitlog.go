package main

import (
	stdlog "log"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/common/promlog"
)

var logger = promlog.New(&promlog.Config{})

var stdlogger = stdLogger{}

type stdLogger stdlog.Logger

var confLogger = log.With(logger, "component", "configuration")

func (s *stdLogger) Log(keyvals ...interface{}) error {
	s.Log(keyvals)
	return nil
}

func demo() {
	level.Info(logger).Log("msg", "go kit log info")
	level.Info(confLogger).Log("msg", "go kit log info with component")
}

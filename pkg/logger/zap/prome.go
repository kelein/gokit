package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// Prometheus Encoders
const (
	JSONEncoderName    = "json-prom"
	ConsoleEncoderName = "console-prom"
)

// func init() {
// 	registerEncoder(JSONEncoderName, zapcore.NewJSONEncoder)
// 	registerEncoder(ConsoleEncoderName, zapcore.NewConsoleEncoder)
// }

func registerEncoder(name string, ctor func(zapcore.EncoderConfig) zapcore.Encoder) error {
	return zap.RegisterEncoder(name, func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return &wrap{ctor(cfg)}, nil
	})
}

type wrap struct{ zapcore.Encoder }

func (w *wrap) Clone() zapcore.Encoder {
	return &wrap{w.Encoder.Clone()}
}

func (w *wrap) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b, err := w.Encoder.EncodeEntry(entry, fields)
	// if err != nil {
	// 	log.Errors.WithLabelValues(caller, entry.Caller.TrimmedPath()).Inc()
	// 	return nil, err
	// }
	// lbls := lvlLbl(entry.Level)
	// logEntries.With(lbls).Inc()
	// logBytes.With(lbls).Observe(float64(b.Len()))
	return b, err
}

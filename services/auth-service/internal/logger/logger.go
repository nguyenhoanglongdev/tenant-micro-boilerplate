package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "os"
)

func NewLogger(logToFile bool, logFilePath string) (*zap.Logger, error) {
	var cfg zap.Config
	if logToFile {
		cfg = zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: false,
			Encoding:    "json", // better for file logs
			OutputPaths: []string{logFilePath, "stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
		}
	} else {
		cfg = zap.NewProductionConfig()
		cfg.OutputPaths = []string{"stdout"}
	}

	return cfg.Build()
}

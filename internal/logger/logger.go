package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewLogger create new logger with channel
func NewLogger(channel string) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.Level = getLevel()

	logger, _ := cfg.Build()
	sugar := logger.Sugar()

	return sugar.Named(channel)
}

func getLevel() zap.AtomicLevel {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		l, err := zapcore.ParseLevel("INFO")
		if err != nil {
			panic(err)
		}

		return zap.NewAtomicLevelAt(l)
	}

	l, err := zapcore.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	return zap.NewAtomicLevelAt(l)
}

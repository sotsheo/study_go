package zaps

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initSugarLog() *zap.Logger {
	// logger, _ := zap.NewProduction()
	config := *&zap.Config{
		Encoding: "json",
		EncoderConfig: *zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
			TimeKey:    "time",
		},
	}
	log, _ := config.Build()
	return log.Default()
}

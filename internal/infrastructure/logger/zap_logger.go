package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()

	// ✨ Nonaktifkan stacktrace untuk level error biasa
	cfg.DisableStacktrace = true

	// ✨ Ubah format waktu & level agar lebih ringkas
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.MessageKey = "msg"

	logger, _ := cfg.Build()
	return logger.Sugar()
}

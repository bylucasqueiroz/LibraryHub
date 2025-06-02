package logger

import (
	"os"

	"github.com/bylucasqueiroz/libraryhub/pkg/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// defaultLogger is the default logger instance
	defaultLogger *zap.Logger

	// osExit is used to make the Fatal function testable
	osExit = os.Exit
)

// Config represents the logger configuration
type Config struct {
	Level      string `json:"level" yaml:"level"`
	Encoding   string `json:"encoding" yaml:"encoding"`
	OutputPath string `json:"output_path" yaml:"output_path"`
}

// init initializes the default logger
func init() {
	defaultLogger = newProductionLogger()
}

// newProductionLogger creates a new production-ready logger
func newProductionLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

// Configure sets up the logger with the provided configuration
func Configure(cfg Config) error {
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         cfg.Encoding,
		EncoderConfig:    zap.NewProductionConfig().EncoderConfig,
		OutputPaths:      []string{cfg.OutputPath},
		ErrorOutputPaths: []string{cfg.OutputPath},
	}

	logger, err := config.Build()
	if err != nil {
		return err
	}

	defaultLogger = logger
	return nil
}

// WithContext returns a logger with context fields
func WithContext(ctx *context.DefaultContext) *zap.Logger {
	return defaultLogger.With(
		zap.String("correlation_id", ctx.CorrelationID()),
		zap.Duration("runtime", ctx.Runtime()),
	)
}

// Debug logs a message at debug level
func Debug(ctx *context.DefaultContext, msg string, fields ...zap.Field) {
	WithContext(ctx).Debug(msg, fields...)
}

// Info logs a message at info level
func Info(ctx *context.DefaultContext, msg string, fields ...zap.Field) {
	WithContext(ctx).Info(msg, fields...)
}

// Warn logs a message at warn level
func Warn(ctx *context.DefaultContext, msg string, fields ...zap.Field) {
	WithContext(ctx).Warn(msg, fields...)
}

// Error logs a message at error level
func Error(ctx *context.DefaultContext, msg string, fields ...zap.Field) {
	WithContext(ctx).Error(msg, fields...)
}

// Fatal logs a message at fatal level and then calls os.Exit(1)
func Fatal(ctx *context.DefaultContext, msg string, fields ...zap.Field) {
	WithContext(ctx).Fatal(msg, fields...)
	osExit(1)
}

// GetLogger returns the default logger instance
func GetLogger() *zap.Logger {
	return defaultLogger
}

// Sync flushes any buffered log entries
func Sync() error {
	return defaultLogger.Sync()
}

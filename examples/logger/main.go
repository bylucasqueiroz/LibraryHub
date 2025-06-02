package main

import (
	"context"
	"errors"
	"time"

	appctx "github.com/bylucasqueiroz/libraryhub/pkg/context"
	"github.com/bylucasqueiroz/libraryhub/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Configure the logger
	err := logger.Configure(logger.Config{
		Level:      "debug",
		Encoding:   "console", // Use console for human-readable output
		OutputPath: "stdout",
	})
	if err != nil {
		panic(err)
	}

	// Create a context with correlation ID
	ctx := appctx.New(context.Background())
	ctx.WithCorrelationID("example-logger-correlation-id")

	// Demonstrate different log levels
	demonstrateLogLevels(ctx)

	// Demonstrate structured logging with additional fields
	demonstrateStructuredLogging(ctx)

	// Demonstrate error logging
	demonstrateErrorLogging(ctx)

	// Demonstrate logging with different contexts
	demonstrateMultipleContexts()

	// Ensure all logs are written
	logger.Sync()
}

func demonstrateLogLevels(ctx *appctx.DefaultContext) {
	logger.Debug(ctx, "This is a debug message")
	logger.Info(ctx, "This is an info message")
	logger.Warn(ctx, "This is a warning message")
	logger.Error(ctx, "This is an error message")
	// Note: We're not demonstrating Fatal as it would exit the program
}

func demonstrateStructuredLogging(ctx *appctx.DefaultContext) {
	// Log with additional structured fields
	logger.Info(ctx, "Processing user request",
		zap.String("user_id", "123"),
		zap.String("action", "login"),
		zap.Int("attempt", 1),
		zap.Duration("response_time", 50*time.Millisecond),
	)

	// Log with nested structured data
	logger.Info(ctx, "Order processed",
		zap.String("order_id", "ORDER123"),
		zap.Int("items_count", 3),
		zap.Float64("total_amount", 99.99),
		zap.Strings("categories", []string{"electronics", "accessories"}),
	)
}

func demonstrateErrorLogging(ctx *appctx.DefaultContext) {
	// Simulate an error
	err := errors.New("database connection failed")

	// Log error with additional context
	logger.Error(ctx, "Failed to process request",
		zap.Error(err),
		zap.String("component", "database"),
		zap.String("operation", "query"),
		zap.Int("retry_attempt", 3),
	)
}

func demonstrateMultipleContexts() {
	// Parent context
	parentCtx := appctx.New(context.Background())
	parentCtx.WithCorrelationID("parent-correlation-id")
	logger.Info(parentCtx, "Parent context operation started")

	// Child context
	childCtx := appctx.New(parentCtx.Context())
	childCtx.WithCorrelationID("child-correlation-id")
	logger.Info(childCtx, "Child context operation started",
		zap.String("operation", "child_task"),
	)
}

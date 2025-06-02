package logger

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	appctx "github.com/bylucasqueiroz/libraryhub/pkg/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewProductionLogger(t *testing.T) {
	logger := newProductionLogger()
	assert.NotNil(t, logger)
}

func TestLogger(t *testing.T) {
	// Test default logger initialization
	assert.NotNil(t, defaultLogger)

	// Create a temporary log file
	tmpfile, err := os.CreateTemp("", "test-*.log")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	// Configure logger to write to the temporary file
	err = Configure(Config{
		Level:      "debug",
		Encoding:   "json",
		OutputPath: tmpfile.Name(),
	})
	require.NoError(t, err)

	// Create a test context
	ctx := appctx.New(context.Background())
	ctx.WithCorrelationID("test-correlation-id")

	// Test all log levels except Fatal
	testAllLogLevels(t, ctx, tmpfile)

	// Test GetLogger
	assert.NotNil(t, GetLogger())
	assert.Equal(t, defaultLogger, GetLogger())
}

func testAllLogLevels(t *testing.T, ctx *appctx.DefaultContext, tmpfile *os.File) {
	tests := []struct {
		level   string
		logFunc func(ctx *appctx.DefaultContext, msg string, fields ...zap.Field)
	}{
		{"debug", Debug},
		{"info", Info},
		{"warn", Warn},
		{"error", Error},
	}

	for _, tt := range tests {
		t.Run("log level "+tt.level, func(t *testing.T) {
			// Clear file
			require.NoError(t, tmpfile.Truncate(0))
			_, err := tmpfile.Seek(0, 0)
			require.NoError(t, err)

			// Log message
			testMessage := tt.level + " message"
			testField := zap.String("test_field", "test_value")
			tt.logFunc(ctx, testMessage, testField)

			// Sync to ensure logs are written
			require.NoError(t, Sync())

			// Read and parse the log
			content, err := os.ReadFile(tmpfile.Name())
			require.NoError(t, err)

			var logEntry map[string]interface{}
			err = json.Unmarshal(content, &logEntry)
			require.NoError(t, err)

			// Verify log contents
			assert.Equal(t, testMessage, logEntry["msg"])
			assert.Equal(t, "test_value", logEntry["test_field"])
			assert.Equal(t, "test-correlation-id", logEntry["correlation_id"])
			assert.Equal(t, tt.level, logEntry["level"])
			assert.Contains(t, logEntry, "runtime")
		})
	}
}

func TestLoggerConfiguration(t *testing.T) {
	tests := []struct {
		name        string
		config      Config
		expectError bool
	}{
		{
			name: "valid configuration",
			config: Config{
				Level:      "info",
				Encoding:   "json",
				OutputPath: "stdout",
			},
			expectError: false,
		},
		{
			name: "valid console encoding",
			config: Config{
				Level:      "debug",
				Encoding:   "console",
				OutputPath: "stdout",
			},
			expectError: false,
		},
		{
			name: "invalid level",
			config: Config{
				Level:      "invalid",
				Encoding:   "json",
				OutputPath: "stdout",
			},
			expectError: true,
		},
		{
			name: "invalid output path",
			config: Config{
				Level:      "info",
				Encoding:   "json",
				OutputPath: "/nonexistent/path/log.txt",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Configure(tt.config)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithContext(t *testing.T) {
	// Create a temporary log file
	tmpfile, err := os.CreateTemp("", "test-*.log")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	// Configure logger
	err = Configure(Config{
		Level:      "info",
		Encoding:   "json",
		OutputPath: tmpfile.Name(),
	})
	require.NoError(t, err)

	// Create a test context with known values
	ctx := appctx.New(context.Background())
	ctx.WithCorrelationID("test-correlation-id-2")

	// Test WithContext directly
	logger := WithContext(ctx)
	assert.NotNil(t, logger)

	// Log a test message using the returned logger
	logger.Info("direct logger test")
	require.NoError(t, Sync())

	// Read and parse the log
	content, err := os.ReadFile(tmpfile.Name())
	require.NoError(t, err)

	var logEntry map[string]interface{}
	err = json.Unmarshal(content, &logEntry)
	require.NoError(t, err)

	// Verify context fields are present
	assert.Equal(t, "test-correlation-id-2", logEntry["correlation_id"])
	assert.Contains(t, logEntry, "runtime")
}

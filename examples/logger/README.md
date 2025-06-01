# Logger Example

This example demonstrates the usage of the structured logging package that integrates with our custom context implementation.

## Features Demonstrated

- Logger configuration
- Different log levels (Debug, Info, Warn, Error)
- Structured logging with additional fields
- Error logging with stack traces
- Context integration
- Parent-child context relationship in logs

## Running the Example

```bash
go run main.go
```

## Expected Output

The example will demonstrate:
1. Different log levels and their formatting
2. Structured data in logs
3. Error logging with stack traces
4. Context fields (correlation ID, worker ID, runtime) automatically included
5. Parent-child relationship in logs

## Key Concepts

### Logger Configuration
```go
logger.Configure(logger.Config{
    Level:      "debug",
    Encoding:   "console",
    OutputPath: "stdout",
})
```

### Basic Logging
```go
logger.Info(ctx, "This is an info message")
logger.Error(ctx, "This is an error message")
```

### Structured Logging
```go
logger.Info(ctx, "Processing user request",
    zap.String("user_id", "123"),
    zap.String("action", "login"),
    zap.Int("attempt", 1),
)
```

### Error Logging
```go
logger.Error(ctx, "Operation failed",
    zap.Error(err),
    zap.String("component", "database"),
)
```

## Log Output Format

When using JSON encoding, logs will look like:
```json
{
    "level": "info",
    "ts": "2024-03-14T15:04:05.000Z",
    "msg": "Processing user request",
    "correlation_id": "example-correlation-id",
    "worker_id": 1,
    "runtime": "0.123s",
    "user_id": "123",
    "action": "login",
    "attempt": 1
}
``` 
# Go Mono Repository

A collection of production-ready Go packages designed to enhance application development with robust, reusable components.

## About

This monorepo contains essential Go packages that provide common functionality for building reliable and maintainable applications. Each package is designed with a focus on simplicity, performance, and proper testing.

## Key Features

- **Context Package**: Enhanced context management
  - Metadata support with type-safe operations
  - Built-in correlation ID tracking
  - Runtime duration tracking
  - Clean context inheritance

- **Logger Package**: Structured logging with context awareness
  - Multiple log levels (Debug, Info, Warn, Error, Fatal)
  - JSON and console output formats
  - Automatic correlation ID inclusion
  - Runtime duration tracking
  - Comprehensive error logging

## Quick Start

```bash
go get github.com/bylucasqueiroz/go-mono
```

### Using Context

```go
// Create a new context with correlation ID
ctx := context.New(context.Background())
ctx.WithCorrelationID("request-123")

// Add metadata
ctx.SetMetadata("user_id", "user-456")

// Get metadata
if value, exists := ctx.GetMetadata("user_id"); exists {
    // Use the value
}
```

### Using Logger

```go
// Configure logger
logger.Configure(logger.Config{
    Level:      "debug",
    Encoding:   "json",
    OutputPath: "stdout",
})

// Log with context
logger.Info(ctx, "Processing request",
    zap.String("user_id", "123"),
    zap.String("action", "login"),
)
```

## Documentation

- [Context Package](examples/context/README.md)
- [Logger Package](examples/logger/README.md)

## Development

### Prerequisites

- Go 1.21 or higher
- Make (optional)

### Testing

```bash
go test ./...
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and development process.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

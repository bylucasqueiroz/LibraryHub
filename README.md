# Go Mono Repository

This monorepo contains various Go packages designed to provide common functionality across different projects.

## Packages

### Context Package (`pkg/context`)

A custom context implementation that extends Go's standard `context.Context` with additional features:

- **Metadata Management**: Flexible key-value storage for context-specific data
- **Correlation ID**: Built-in support for request tracing with correlation IDs
- **Runtime Tracking**: Automatic tracking of context lifetime
- **Context Inheritance**: Proper parent-child context relationship management

Example usage:

```go
// Create a new context
ctx := context.New(context.Background())

// Set correlation ID
ctx.WithCorrelationID("request-123")

// Add custom metadata
ctx.SetMetadata("user_id", "user-456")

// Get metadata
if value, exists := ctx.GetMetadata("user_id"); exists {
    // Use the value
}

// Get correlation ID
correlationID := ctx.CorrelationID()

// Get runtime duration
duration := ctx.Runtime()
```

## Installation

```bash
go get github.com/yourusername/libraryhub
```

## Development

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile commands)

### Testing

Run all tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Contributing

Please read our [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

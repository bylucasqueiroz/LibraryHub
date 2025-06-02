# Context Example

This example demonstrates the usage of the custom context package that extends the standard Go context with additional functionality.

## Features Demonstrated

- Creating a new context with correlation ID
- Setting and getting metadata
- Creating child contexts
- Context cancellation
- Context runtime tracking

## Running the Example

```bash
go run main.go
```

## Expected Output

The example will show:
1. Basic context usage with correlation ID
2. Metadata storage and retrieval
3. Child context inheritance
4. Context cancellation in action
5. Runtime duration tracking

## Key Concepts

### Context Creation
```go
ctx := appctx.New(context.Background())
ctx.WithCorrelationID("example-correlation-id")
```

### Metadata Management
```go
ctx.SetMetadata("environment", "production")
value, exists := ctx.GetMetadata("environment")
```

### Child Contexts
```go
childCtx := appctx.New(ctx.Context())
childCtx.WithCorrelationID("child-correlation-id")
```

### Context Cancellation
```go
ctx.Cancel()
<-ctx.Done() // Wait for cancellation
``` 
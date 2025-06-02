package main

import (
	"context"
	"fmt"
	"time"

	appctx "github.com/bylucasqueiroz/go-mono/pkg/context"
)

func main() {
	// Create a new context with background
	ctx := appctx.New(context.Background())
	ctx.WithCorrelationID("example-correlation-id")

	// Add some metadata
	ctx.SetMetadata("environment", "production")
	ctx.SetMetadata("service", "example-service")

	// Demonstrate context usage
	processWithContext(ctx)

	// Create a child context
	childCtx := appctx.New(ctx.Context())
	childCtx.WithCorrelationID("child-correlation-id")
	childCtx.SetMetadata("child", true)

	// Demonstrate child context usage
	processWithContext(childCtx)

	// Demonstrate context cancellation
	cancelExample()
}

func processWithContext(ctx *appctx.DefaultContext) {
	fmt.Printf("\nProcessing with context:\n")
	fmt.Printf("Correlation ID: %s\n", ctx.CorrelationID())

	// Get metadata
	if env, ok := ctx.GetMetadata("environment"); ok {
		fmt.Printf("Environment: %v\n", env)
	}
	if service, ok := ctx.GetMetadata("service"); ok {
		fmt.Printf("Service: %v\n", service)
	}
	if child, ok := ctx.GetMetadata("child"); ok {
		fmt.Printf("Is Child: %v\n", child)
	}

	// Show context runtime
	time.Sleep(100 * time.Millisecond) // Simulate some work
	fmt.Printf("Runtime: %v\n", ctx.Runtime())
}

func cancelExample() {
	fmt.Printf("\nDemonstrating context cancellation:\n")
	ctx := appctx.New(context.Background())

	// Start a goroutine that will be cancelled
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received cancellation signal")
			return
		case <-time.After(5 * time.Second):
			fmt.Println("Worker completed")
		}
	}()

	// Cancel the context after 1 second
	time.Sleep(1 * time.Second)
	ctx.Cancel()
	fmt.Println("Context cancelled")

	// Wait to see the cancellation effect
	time.Sleep(100 * time.Millisecond)
}

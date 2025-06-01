package context

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBridgeContext(t *testing.T) {
	t.Run("basic context operations", func(t *testing.T) {
		ctx := context.Background()
		bCtx := New(ctx)

		// Test metadata
		bCtx.SetMetadata("test", "value")
		val, ok := bCtx.GetMetadata("test")
		assert.True(t, ok)
		assert.Equal(t, "value", val)

		// Test non-existent metadata
		_, ok = bCtx.GetMetadata("nonexistent")
		assert.False(t, ok)

		// Test Context() method
		assert.Equal(t, bCtx.ctx, bCtx.Context())
	})

	t.Run("correlation ID operations", func(t *testing.T) {
		ctx := context.Background()
		bCtx := New(ctx)

		// Test auto-generated correlation ID
		assert.NotEmpty(t, bCtx.CorrelationID())
		initialID := bCtx.CorrelationID()

		// Verify correlation ID is in metadata
		metadataID, ok := bCtx.GetMetadata(MetadataKeyCorrelationID)
		assert.True(t, ok)
		assert.Equal(t, initialID, metadataID)

		// Test custom correlation ID
		customID := "custom-correlation-id"
		bCtx.WithCorrelationID(customID)
		assert.Equal(t, customID, bCtx.CorrelationID())

		// Verify custom ID is in metadata
		metadataID, ok = bCtx.GetMetadata(MetadataKeyCorrelationID)
		assert.True(t, ok)
		assert.Equal(t, customID, metadataID)

		// Test setting invalid correlation ID type
		bCtx.SetMetadata(MetadataKeyCorrelationID, 123) // Set non-string value
		assert.Empty(t, bCtx.CorrelationID())
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx := context.Background()
		bCtx := New(ctx)

		// Test initial state
		assert.NoError(t, bCtx.Err())
		select {
		case <-bCtx.Done():
			t.Fatal("Context should not be done")
		default:
			// Expected
		}

		// Test cancellation
		bCtx.Cancel()
		assert.Error(t, bCtx.Err())
		assert.Equal(t, context.Canceled, bCtx.Err())

		// Test Done channel
		select {
		case <-bCtx.Done():
			// Expected
		default:
			t.Fatal("Context should be done")
		}
	})

	t.Run("context deadline", func(t *testing.T) {
		deadline := time.Now().Add(10 * time.Millisecond)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()

		bCtx := New(ctx)
		gotDeadline, ok := bCtx.Deadline()
		assert.True(t, ok)
		assert.Equal(t, deadline, gotDeadline)
	})

	t.Run("context runtime", func(t *testing.T) {
		ctx := context.Background()
		bCtx := New(ctx)

		// Sleep to ensure measurable runtime
		time.Sleep(10 * time.Millisecond)
		runtime := bCtx.Runtime()
		assert.True(t, runtime >= 10*time.Millisecond)
	})

	t.Run("context inheritance", func(t *testing.T) {
		parentCtx := New(context.Background())
		parentCtx.WithCorrelationID("parent-id")
		parentCtx.SetMetadata("parent", true)

		childCtx := New(parentCtx.Context())
		childCtx.WithCorrelationID(parentCtx.CorrelationID())

		assert.Equal(t, "parent-id", childCtx.CorrelationID())

		// Verify parent cancellation affects child
		parentCtx.Cancel()
		assert.Error(t, childCtx.Err())
	})

	t.Run("context value", func(t *testing.T) {
		type key string
		testKey := key("test")
		ctx := context.WithValue(context.Background(), testKey, "test-value")

		bCtx := New(ctx)
		assert.Equal(t, "test-value", bCtx.Value(testKey))
	})
}

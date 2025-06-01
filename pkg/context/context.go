package context

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	// MetadataKeyCorrelationID is the key used for correlation ID in metadata
	MetadataKeyCorrelationID = "correlation_id"
)

// DefaultContext wraps the standard context.Context and adds Kafka bridge specific functionality
type DefaultContext struct {
	ctx       context.Context
	cancel    context.CancelFunc
	startTime time.Time
	metadata  map[string]interface{}
}

// New creates a new DefaultContext with the given parent context
func New(parent context.Context) *DefaultContext {
	ctx, cancel := context.WithCancel(parent)
	bc := &DefaultContext{
		ctx:       ctx,
		cancel:    cancel,
		startTime: time.Now(),
		metadata:  make(map[string]interface{}),
	}
	// Set default correlation ID
	bc.SetMetadata(MetadataKeyCorrelationID, uuid.New().String())
	return bc
}

// Context returns the underlying context.Context
func (bc *DefaultContext) Context() context.Context {
	return bc.ctx
}

// WithCorrelationID sets a custom correlation ID for this context
func (bc *DefaultContext) WithCorrelationID(correlationID string) *DefaultContext {
	bc.SetMetadata(MetadataKeyCorrelationID, correlationID)
	return bc
}

// CorrelationID returns the correlation ID associated with this context
func (bc *DefaultContext) CorrelationID() string {
	if id, ok := bc.GetMetadata(MetadataKeyCorrelationID); ok {
		if strID, ok := id.(string); ok {
			return strID
		}
	}
	return ""
}

// SetMetadata sets a metadata value in the context
func (bc *DefaultContext) SetMetadata(key string, value interface{}) {
	bc.metadata[key] = value
}

// GetMetadata retrieves a metadata value from the context
func (bc *DefaultContext) GetMetadata(key string) (interface{}, bool) {
	val, ok := bc.metadata[key]
	return val, ok
}

// Deadline returns the deadline of the underlying context
func (bc *DefaultContext) Deadline() (deadline time.Time, ok bool) {
	return bc.ctx.Deadline()
}

// Done returns the done channel of the underlying context
func (bc *DefaultContext) Done() <-chan struct{} {
	return bc.ctx.Done()
}

// Err returns the error of the underlying context
func (bc *DefaultContext) Err() error {
	return bc.ctx.Err()
}

// Value returns a value from the underlying context
func (bc *DefaultContext) Value(key interface{}) interface{} {
	return bc.ctx.Value(key)
}

// Cancel cancels this context
func (bc *DefaultContext) Cancel() {
	bc.cancel()
}

// Runtime returns the duration since this context was created
func (bc *DefaultContext) Runtime() time.Duration {
	return time.Since(bc.startTime)
}

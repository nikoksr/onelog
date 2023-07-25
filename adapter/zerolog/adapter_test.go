package zerologadapter

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/nikoksr/onelog/internal/testutils"
)

// TestNewAdapter tests if NewAdapter returns a non-nil *Adapter.
func TestNewAdapter(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	adapter := NewAdapter(&logger)

	assert.NotNil(t, adapter, "the returned adapter should not be nil")
}

// TestContexts tests if each log level returns a valid *Context.
func TestContexts(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	adapter := NewAdapter(&logger)

	// Debug
	logContext := adapter.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")

	// Info
	logContext = adapter.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")

	// Warn
	logContext = adapter.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")

	// Error
	logContext = adapter.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")

	// Fatal
	logContext = adapter.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
}

// TestMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestMethods(t *testing.T) {
	t.Parallel()

	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.DurationFieldInteger = true
	zerolog.DurationFieldUnit = time.Nanosecond
	zerolog.MessageFieldName = "msg"

	buff := new(bytes.Buffer)
	logger := zerolog.New(buff).With().Timestamp().Logger()
	adapter := NewAdapter(&logger)

	testutils.TestingMethods(t, adapter, buff)
}

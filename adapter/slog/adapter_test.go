package slogadapter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"

	"github.com/nikoksr/onelog/internal/testutils"
)

// TestNewAdapter tests if NewAdapter returns a non-nil *Adapter.
func TestNewAdapter(t *testing.T) {
	t.Parallel()

	slogLogger := slog.Default()
	logger := NewAdapter(slogLogger)

	assert.NotNil(t, logger, "the returned adapter should not be nil")
}

// TestContexts tests if each log level returns a valid *Context.
func TestContexts(t *testing.T) {
	t.Parallel()

	slogLogger := slog.Default()
	logger := NewAdapter(slogLogger)

	// Debug
	logContext := logger.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelDebug, "the returned context should have the correct log level")

	// Info
	logContext = logger.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelInfo, "the returned context should have the correct log level")

	// Warn
	logContext = logger.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelWarn, "the returned context should have the correct log level")

	// Error
	logContext = logger.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelError, "the returned context should have the correct log level")

	// Fatal; note that slog does not have a fatal level, so this should return an error level context
	logContext = logger.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelError, "the returned context should have the correct log level")
}

// TestMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestMethods(t *testing.T) {
	t.Parallel()

	buff := new(bytes.Buffer)
	handler := slog.NewJSONHandler(buff, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	adapter := NewAdapter(logger)

	testutils.TestingMethods(t, adapter, buff)
}

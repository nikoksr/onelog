package slogadapter

import (
	"bytes"
	"io"
	"testing"

	"github.com/nikoksr/onelog"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"

	"github.com/nikoksr/onelog/internal/testutils"
)

func newTestingAdapter(out io.Writer) onelog.Logger {
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	return NewAdapter(logger)
}

// TestNewAdapter tests if NewAdapter returns a non-nil *Adapter.
func TestNewAdapter(t *testing.T) {
	t.Parallel()

	adapter := newTestingAdapter(io.Discard)

	assert.NotNil(t, adapter, "the returned adapter should not be nil")
}

// TestContexts tests if each log level returns a valid *Context.
func TestContexts(t *testing.T) {
	t.Parallel()

	adapter := newTestingAdapter(io.Discard)

	// Debug
	logContext := adapter.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelDebug, "the returned context should have the correct log level")

	// Info
	logContext = adapter.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelInfo, "the returned context should have the correct log level")

	// Warn
	logContext = adapter.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelWarn, "the returned context should have the correct log level")

	// Error
	logContext = adapter.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelError, "the returned context should have the correct log level")

	// Fatal; note that slog does not have a fatal level, so this should return an error level context
	logContext = adapter.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, slog.LevelError, "the returned context should have the correct log level")
}

// TestMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestMethods(t *testing.T) {
	t.Parallel()

	buff := new(bytes.Buffer)
	adapter := newTestingAdapter(buff)

	testutils.TestingMethods(t, adapter, buff)
}

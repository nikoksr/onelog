package zapadapter

import (
	"bytes"
	"io"
	"testing"

	"github.com/nikoksr/onelog"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/nikoksr/onelog/internal/testutils"
)

func newSugarAdapter(out io.Writer) onelog.Logger {
	return NewSugarAdapter(newLogger(out).Sugar())
}

// TestNewSugarAdapter tests if NewSugarAdapter returns a non-nil *Adapter.
func TestNewSugarAdapter(t *testing.T) {
	t.Parallel()

	adapter := newSugarAdapter(io.Discard)

	assert.NotNil(t, adapter, "the returned adapter should not be nil")
}

// TestSugarContexts tests if each log level returns a valid *Context.
func TestSugarContexts(t *testing.T) {
	t.Parallel()

	adapter := newSugarAdapter(io.Discard)

	// Debug
	logContext := adapter.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.DebugLevel, "the returned context should have the correct log level")

	// Info
	logContext = adapter.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.InfoLevel, "the returned context should have the correct log level")

	// Warn
	logContext = adapter.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.WarnLevel, "the returned context should have the correct log level")

	// Error
	logContext = adapter.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.ErrorLevel, "the returned context should have the correct log level")

	// Fatal
	logContext = adapter.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.FatalLevel, "the returned context should have the correct log level")
}

// TestSugarMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestSugarMethods(t *testing.T) {
	t.Parallel()

	buff := new(bytes.Buffer)
	adapter := newSugarAdapter(buff)

	testutils.TestingMethods(t, adapter, buff)
}

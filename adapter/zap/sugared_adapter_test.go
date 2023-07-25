package zapadapter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/nikoksr/onelog/internal/testutils"
)

// TestNewSugarAdapter tests if NewSugarAdapter returns a non-nil *Adapter.
func TestNewSugarAdapter(t *testing.T) {
	t.Parallel()

	zapLogger, _ := zap.NewDevelopment()
	logger := NewSugarAdapter(zapLogger.Sugar())

	assert.NotNil(t, logger, "the returned adapter should not be nil")
}

// TestSugarContexts tests if each log level returns a valid *Context.
func TestSugarContexts(t *testing.T) {
	t.Parallel()

	zapLogger, _ := zap.NewDevelopment()
	logger := NewSugarAdapter(zapLogger.Sugar())

	// Debug
	logContext := logger.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.DebugLevel, "the returned context should have the correct log level")

	// Info
	logContext = logger.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.InfoLevel, "the returned context should have the correct log level")

	// Warn
	logContext = logger.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.WarnLevel, "the returned context should have the correct log level")

	// Error
	logContext = logger.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.ErrorLevel, "the returned context should have the correct log level")

	// Fatal
	logContext = logger.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(SugarContext), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*SugarContext).level, zap.FatalLevel, "the returned context should have the correct log level")
}

// TestSugarMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestSugarMethods(t *testing.T) {
	t.Parallel()

	buff := new(bytes.Buffer)
	logger := newTestingLogger(buff)
	adapter := NewSugarAdapter(logger.Sugar())

	testutils.TestingMethods(t, adapter, buff)
}

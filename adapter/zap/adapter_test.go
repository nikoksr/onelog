package zapadapter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nikoksr/onelog/internal/testutils"
)

// TestNewAdapter tests if NewAdapter returns a non-nil *Adapter.
func TestNewAdapter(t *testing.T) {
	t.Parallel()

	zapLogger, _ := zap.NewDevelopment()
	logger := NewAdapter(zapLogger)

	assert.NotNil(t, logger, "the returned adapter should not be nil")
}

// TestContexts tests if each log level returns a valid *Context.
func TestContexts(t *testing.T) {
	t.Parallel()

	zapLogger, _ := zap.NewDevelopment()
	logger := NewAdapter(zapLogger)

	// Debug
	logContext := logger.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.DebugLevel, "the returned context should have the correct log level")

	// Info
	logContext = logger.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.InfoLevel, "the returned context should have the correct log level")

	// Warn
	logContext = logger.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.WarnLevel, "the returned context should have the correct log level")

	// Error
	logContext = logger.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.ErrorLevel, "the returned context should have the correct log level")

	// Fatal
	logContext = logger.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.FatalLevel, "the returned context should have the correct log level")
}

func newTestingLogger(buff *bytes.Buffer) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey:     "msg",
				TimeKey:        "time",
				EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
				EncodeDuration: zapcore.NanosDurationEncoder,
			}),
			zapcore.AddSync(buff),
			zapcore.DebugLevel,
		),
		zap.ErrorOutput(zapcore.AddSync(buff)),
	)
}

// TestMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestMethods(t *testing.T) {
	t.Parallel()

	buff := new(bytes.Buffer)
	logger := newTestingLogger(buff)
	adapter := NewAdapter(logger)

	testutils.TestingMethods(t, adapter, buff)
}

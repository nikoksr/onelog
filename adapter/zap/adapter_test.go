package zapadapter

import (
	"bytes"
	"io"
	"testing"

	"github.com/nikoksr/onelog"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nikoksr/onelog/internal/testutils"
)

func newLogger(out io.Writer) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey:     "msg",
				TimeKey:        "time",
				EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
				EncodeDuration: zapcore.NanosDurationEncoder,
			}),
			zapcore.AddSync(out),
			zapcore.DebugLevel,
		),
		zap.ErrorOutput(zapcore.AddSync(out)),
	)
}

func newAdapter(out io.Writer) onelog.Logger {
	return NewAdapter(newLogger(out))
}

// TestNewAdapter tests if NewAdapter returns a non-nil *Adapter.
func TestNewAdapter(t *testing.T) {
	t.Parallel()

	adapter := newAdapter(io.Discard)

	assert.NotNil(t, adapter, "the returned adapter should not be nil")
}

// TestContexts tests if each log level returns a valid *Context.
func TestContexts(t *testing.T) {
	t.Parallel()

	adapter := newAdapter(io.Discard)

	// Debug
	logContext := adapter.Debug()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.DebugLevel, "the returned context should have the correct log level")

	// Info
	logContext = adapter.Info()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.InfoLevel, "the returned context should have the correct log level")

	// Warn
	logContext = adapter.Warn()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.WarnLevel, "the returned context should have the correct log level")

	// Error
	logContext = adapter.Error()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.ErrorLevel, "the returned context should have the correct log level")

	// Fatal
	logContext = adapter.Fatal()
	assert.NotNil(t, logContext, "the returned context should not be nil")
	assert.IsType(t, new(Context), logContext, "the returned context should be of type *Context")
	assert.Equal(t, logContext.(*Context).level, zap.FatalLevel, "the returned context should have the correct log level")
}

// TestMethods tests if each method returns a non-nil *Context and if the log is written correctly.
func TestMethods(t *testing.T) {
	t.Parallel()

	buff := new(bytes.Buffer)
	adapter := newAdapter(buff)

	testutils.TestingMethods(t, adapter, buff)
}

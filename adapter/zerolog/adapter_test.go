package zerologadapter

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/nikoksr/onelog"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/nikoksr/onelog/internal/testutils"
)

// Setting global zerolog settings to make sure the tests are deterministic.
func TestMain(m *testing.M) {
	originalTimeFieldFormat := zerolog.TimeFieldFormat
	originalDurationFieldInteger := zerolog.DurationFieldInteger
	originalDurationFieldUnit := zerolog.DurationFieldUnit
	originalMessageFieldName := zerolog.MessageFieldName

	defer func() {
		zerolog.TimeFieldFormat = originalTimeFieldFormat
		zerolog.DurationFieldInteger = originalDurationFieldInteger
		zerolog.DurationFieldUnit = originalDurationFieldUnit
		zerolog.MessageFieldName = originalMessageFieldName
	}()

	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.DurationFieldInteger = true
	zerolog.DurationFieldUnit = time.Nanosecond
	zerolog.MessageFieldName = "msg"

	os.Exit(m.Run())
}

func newAdapter(out io.Writer) onelog.Logger {
	logger := zerolog.New(out).With().Timestamp().Logger()
	return NewAdapter(&logger)
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

	buff := new(bytes.Buffer)
	adapter := newAdapter(buff)

	testutils.TestingMethods(t, adapter, buff)
}

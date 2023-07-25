package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nikoksr/onelog"
)

func assertEqualSlices[T comparable](t *testing.T, expected []T, actual any) {
	t.Helper()

	actualValues, ok := actual.([]any)
	require.True(t, ok, "Expected argument 'actual' to be a slice of numbers, however, received: %T", actual)

	require.Equal(t, len(expected), len(actualValues), "Expected slices to have same length. Expected: %d, Actual: %d", len(expected), len(actualValues))

	for i := range expected {
		assert.EqualValues(t, expected[i], actualValues[i], "Expected all values in slices to match. Mismatch at index: %d", i)
	}
}

func validateTimestamp(t *testing.T, expected time.Time, got any) {
	t.Helper()

	switch v := got.(type) {
	case string:
		gotTime, err := time.Parse(time.RFC3339Nano, v)
		require.NoError(t, err, "Expected log to contain a valid date time in RFC3339Nano format.")
		assert.EqualValues(t, expected.UTC(), gotTime.UTC(), "Expected date time in log to match the provided date time. Expected: %v, Got: %v", expected.UTC(), gotTime.UTC())
	case float64:
		// When Unix time is being sent in seconds
		gotTime := time.Unix(int64(v), 0)
		assert.EqualValues(t, expected.UTC(), gotTime.UTC(), "Expected log to contain the given Unix time. Expected: %v, Got: %v", expected.UTC(), gotTime.UTC())
	default:
		t.Errorf("Unexpected data type for date time in the log: %T", v)
	}
}

func validateTimestamps(t *testing.T, expected []time.Time, got any) {
	t.Helper()

	switch values := got.(type) {
	case []string:
		for i, v := range values {
			validateTimestamp(t, expected[i], v)
		}
	case []float64:
		for i, v := range values {
			validateTimestamp(t, expected[i], v)
		}
	case []any:
		for i, v := range values {
			validateTimestamp(t, expected[i], v)
		}
	default:
		t.Errorf("Unexpected type %T", values)
	}
}

func validateDuration(t *testing.T, expected time.Duration, got any) {
	t.Helper()

	switch v := got.(type) {
	case string:
		// Parse the duration from the string
		got, err := time.ParseDuration(v)
		require.NoError(t, err, "the log should contain a valid duration")
		assert.EqualValues(t, expected, got, "the log should contain the correct duration")
	case int64, float64:
		// When the duration is being sent in milliseconds
		assert.EqualValues(t, expected.Nanoseconds(), v, "the log should contain the correct duration")
	default:
		t.Errorf("Unexpected type %T", v)
	}
}

func validateErrors(t *testing.T, expected []error, got any) {
	t.Helper()

	switch v := got.(type) {
	case []any:
		for i, err := range v {
			// Errors are logged as slice of errors
			if errMap, ok := err.(map[string]any); ok {
				require.True(t, ok, "the log should contain a map of errors, but got %T", err)
				assert.Equal(t, expected[i].Error(), errMap["error"], "the log should contain the correct error message at index %d", i)
				// Errors are logged as slice of strings
			} else if errStr, ok := err.(string); ok {
				require.True(t, ok, "the log should contain a string error, but got %T", err)
				assert.Equal(t, expected[i].Error(), errStr, "the log should contain the correct error message at index %d", i)
			} else {
				t.Errorf("Unexpected type %T", err)
			}
		}
	default:
		t.Errorf("Unexpected type %T", v)
	}
}

func TestingMethods(t *testing.T, logger onelog.Logger, logSink *bytes.Buffer) {
	t.Helper()

	// Create a reusable context
	logContext := logger.Info()

	// Helper functions
	now := func() time.Time {
		return time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	stringer1 := bytes.NewBufferString("Value 1")
	stringer2 := bytes.NewBufferString("Value 2")

	// Test that the methods return a non-nil *Context
	tests := []struct {
		Name     string
		Fn       func() onelog.LoggerContext
		Validate func(t *testing.T, result map[string]any)
	}{
		{
			Name: "Str",
			Fn:   func() onelog.LoggerContext { return logContext.Str("Test", "Value") },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.Equal(t, "Value", value, "the log should contain the correct value")
			},
		},
		{
			Name: "Strs",
			Fn:   func() onelog.LoggerContext { return logContext.Strs("Test", []string{"Value1", "Value2"}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.Equal(t, []any{"Value1", "Value2"}, value, "the log should contain the correct value")
			},
		},
		{
			Name: "Stringer",
			Fn:   func() onelog.LoggerContext { return logContext.Stringer("Test", stringer1) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.Equal(t, "Value 1", value, "the log should contain the correct value")
			},
		},
		{
			Name: "Stringers",
			Fn: func() onelog.LoggerContext {
				return logContext.Stringers("Test", []fmt.Stringer{stringer1, stringer2})
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{"Value 1", "Value 2"}, value)
			},
		},
		{
			Name: "Int",
			Fn:   func() onelog.LoggerContext { return logContext.Int("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, 42, value, "the log should contain the correct value")
			},
		},
		{
			Name: "Ints",
			Fn:   func() onelog.LoggerContext { return logContext.Ints("Test", []int{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{1, 2, 3}, value)
			},
		},
		{
			Name: "Int8",
			Fn:   func() onelog.LoggerContext { return logContext.Int8("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, int8(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Ints8",
			Fn:   func() onelog.LoggerContext { return logContext.Ints8("Test", []int8{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{int8(1), int8(2), int8(3)}, value)
			},
		},
		{
			Name: "Int16",
			Fn:   func() onelog.LoggerContext { return logContext.Int16("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, int16(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Ints16",
			Fn:   func() onelog.LoggerContext { return logContext.Ints16("Test", []int16{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{int16(1), int16(2), int16(3)}, value)
			},
		},
		{
			Name: "Int32",
			Fn:   func() onelog.LoggerContext { return logContext.Int32("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, int32(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Ints32",
			Fn:   func() onelog.LoggerContext { return logContext.Ints32("Test", []int32{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{int32(1), int32(2), int32(3)}, value)
			},
		},
		{
			Name: "Int64",
			Fn:   func() onelog.LoggerContext { return logContext.Int64("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, int64(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Ints64",
			Fn:   func() onelog.LoggerContext { return logContext.Ints64("Test", []int64{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{int64(1), int64(2), int64(3)}, value)
			},
		},
		{
			Name: "Uint",
			Fn:   func() onelog.LoggerContext { return logContext.Uint("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, uint(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Uints",
			Fn:   func() onelog.LoggerContext { return logContext.Uints("Test", []uint{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{uint(1), uint(2), uint(3)}, value)
			},
		},
		{
			Name: "Uint8",
			Fn:   func() onelog.LoggerContext { return logContext.Uint8("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, uint8(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Uints8",
			Fn:   func() onelog.LoggerContext { return logContext.Uints8("Test", []uint8{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				t.Log(result)
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []uint8{1, 2, 3}, value)
			},
		},
		{
			Name: "Uint16",
			Fn:   func() onelog.LoggerContext { return logContext.Uint16("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, uint16(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Uints16",
			Fn:   func() onelog.LoggerContext { return logContext.Uints16("Test", []uint16{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{uint16(1), uint16(2), uint16(3)}, value)
			},
		},
		{
			Name: "Uint32",
			Fn:   func() onelog.LoggerContext { return logContext.Uint32("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, uint32(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Uints32",
			Fn:   func() onelog.LoggerContext { return logContext.Uints32("Test", []uint32{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{uint32(1), uint32(2), uint32(3)}, value)
			},
		},
		{
			Name: "Uint64",
			Fn:   func() onelog.LoggerContext { return logContext.Uint64("Test", 42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, uint64(42), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Uints64",
			Fn:   func() onelog.LoggerContext { return logContext.Uints64("Test", []uint64{1, 2, 3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{uint64(1), uint64(2), uint64(3)}, value)
			},
		},
		{
			Name: "Float32",
			Fn:   func() onelog.LoggerContext { return logContext.Float32("Test", 42.42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, 42.42, value, "the log should contain the correct value")
			},
		},
		{
			Name: "Floats32",
			Fn:   func() onelog.LoggerContext { return logContext.Floats32("Test", []float32{1.1, 2.2, 3.3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{1.1, 2.2, 3.3}, value)
			},
		},
		{
			Name: "Float64",
			Fn:   func() onelog.LoggerContext { return logContext.Float64("Test", 42.42) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, 42.42, value, "the log should contain the correct value")
			},
		},
		{
			Name: "Floats64",
			Fn:   func() onelog.LoggerContext { return logContext.Floats64("Test", []float64{1.1, 2.2, 3.3}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{1.1, 2.2, 3.3}, value)
			},
		},
		{
			Name: "Bool",
			Fn:   func() onelog.LoggerContext { return logContext.Bool("Test", true) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.Equal(t, true, value, "the log should contain the correct value")
			},
		},
		{
			Name: "Bools",
			Fn:   func() onelog.LoggerContext { return logContext.Bools("Test", []bool{true, false, true}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.Equal(t, []any{true, false, true}, value, "the log should contain the correct value")
			},
		},
		{
			Name: "Bytes",
			Fn:   func() onelog.LoggerContext { return logContext.Bytes("Test", []byte("Test")) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, []byte("Test"), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Time",
			Fn:   func() onelog.LoggerContext { return logContext.Time("Test", now()) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				validateTimestamp(t, now(), value)
			},
		},
		{
			Name: "Times",
			Fn:   func() onelog.LoggerContext { return logContext.Times("Test", []time.Time{now(), now()}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				validateTimestamps(t, []time.Time{now(), now()}, value)
			},
		},
		{
			Name: "Dur",
			Fn:   func() onelog.LoggerContext { return logContext.Dur("Test", time.Second) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				validateDuration(t, time.Second, value)
			},
		},
		{
			Name: "Durs",
			Fn:   func() onelog.LoggerContext { return logContext.Durs("Test", []time.Duration{time.Second, time.Second}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assertEqualSlices(t, []any{time.Second.Nanoseconds(), time.Second.Nanoseconds()}, value)
			},
		},
		{
			Name: "TimeDiff",
			Fn:   func() onelog.LoggerContext { return logContext.TimeDiff("Test", now(), now()) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				assert.EqualValues(t, time.Duration(0), value, "the log should contain the correct value")
			},
		},
		{
			Name: "IPAddr",
			Fn:   func() onelog.LoggerContext { return logContext.IPAddr("Test", net.IP{127, 0, 0, 1}) },
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				expected := net.ParseIP("127.0.0.1")
				require.NotNil(t, expected, "the log should contain a valid IP address")
				assert.EqualValues(t, expected.String(), value, "the log should contain the correct value")
			},
		},
		{
			Name: "IPPrefix",
			Fn: func() onelog.LoggerContext {
				return logContext.IPPrefix("Test", net.IPNet{IP: net.IP{127, 0, 0, 1}, Mask: net.IPMask{255, 255, 255, 0}})
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				expected := net.IPNet{IP: net.IP{127, 0, 0, 1}, Mask: net.IPMask{255, 255, 255, 0}}
				require.NotNil(t, expected, "the log should contain a valid IP prefix")
				assert.EqualValues(t, expected.String(), value, "the log should contain the correct value")
			},
		},
		{
			Name: "MACAddr",
			Fn: func() onelog.LoggerContext {
				return logContext.MACAddr("Test", net.HardwareAddr{0, 0, 0, 0, 0, 0})
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["Test"]
				require.True(t, ok, "the log should contain the key 'Test'")
				expected, err := net.ParseMAC("00:00:00:00:00:00")
				require.NoError(t, err, "the log should contain a valid MAC address")
				assert.EqualValues(t, expected.String(), value, "the log should contain the correct value")
			},
		},
		{
			Name: "Err",
			Fn: func() onelog.LoggerContext {
				return logContext.Err(fmt.Errorf("test error"))
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["error"]
				require.True(t, ok, "the log should contain the key 'error'")
				assert.Equal(t, "test error", value, "the log should contain the correct value")
			},
		},
		{
			Name: "Errs",
			Fn: func() onelog.LoggerContext {
				return logContext.Errs("errors", []error{fmt.Errorf("test error1"), fmt.Errorf("test error2")})
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["errors"]
				require.True(t, ok, "the log should contain the key 'errors'")
				validateErrors(t, []error{fmt.Errorf("test error1"), fmt.Errorf("test error2")}, value)
			},
		},
		{
			Name: "AnErr",
			Fn: func() onelog.LoggerContext {
				return logContext.AnErr("my_error", fmt.Errorf("test error"))
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["my_error"]
				require.True(t, ok, "the log should contain the key 'my_error'")
				assert.Equal(t, "test error", value, "the log should contain the correct value")
			},
		},
		{
			Name: "Any",
			Fn: func() onelog.LoggerContext {
				return logContext.Any("my_any", "test any")
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["my_any"]
				require.True(t, ok, "the log should contain the key 'my_any'")
				assert.Equal(t, "test any", value, "the log should contain the correct value")
			},
		},
		{
			Name: "Fields",
			Fn: func() onelog.LoggerContext {
				return logContext.Fields(map[string]any{"my_field": "test field"})
			},
			Validate: func(t *testing.T, result map[string]any) {
				t.Helper()
				value, ok := result["my_field"]
				require.True(t, ok, "the log should contain the key 'my_field'")
				assert.Equal(t, "test field", value, "the log should contain the correct value")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			defer logSink.Reset()

			// Check if the returned context is non-nil
			assert.NotNil(t, tc.Fn(), "the returned context should not be nil")

			// Send the log
			tc.Fn().Msg("Test message")

			// Check if the log was written
			assert.NotEmpty(t, logSink.String(), "the log should have been written")

			// Some logger prefix the log with a timestamp, so we need to remove it before unmarshalling the log
			msg := logSink.String()
			require.NotEmpty(t, msg, "the log should have been written")

			idx := strings.Index(msg, "{")
			msg = msg[idx:]

			// Validate the log
			result := make(map[string]any)

			err := json.Unmarshal([]byte(msg), &result)
			assert.NoError(t, err, "the log should be valid json")

			tc.Validate(t, result)
		})
	}
}

package onelog

import (
	"fmt"
	"net"
	"time"
)

// Fields type is an alias for a map that stores key-value pairs.
type Fields map[string]any

// Logger interface provides methods for logging at various levels.
type Logger interface {
	// Debug returns a LoggerContext for a debug log.
	Debug() LoggerContext

	// Info returns a LoggerContext for an info log.
	Info() LoggerContext

	// Warn returns a LoggerContext for a warn log.
	Warn() LoggerContext

	// Error returns a LoggerContext for an error log.
	Error() LoggerContext

	// Fatal returns a LoggerContext for a fatal log.
	Fatal() LoggerContext
}

// LoggerContext interface provides methods for adding context to logs.
type LoggerContext interface {
	// Str adds the field key with val as a string to the logger context.
	Str(key, value string) LoggerContext

	// Strs adds the field key with val as a []string to the logger context.
	Strs(key string, value []string) LoggerContext

	// Stringer adds the field key with val as a fmt.Stringer to the logger context.
	Stringer(key string, val fmt.Stringer) LoggerContext

	// Stringers adds the field key with val as a []fmt.Stringer to the logger context.
	Stringers(key string, vals []fmt.Stringer) LoggerContext

	// Int adds the field key with val as an int to the logger context.
	Int(key string, value int) LoggerContext

	// Ints adds the field key with val as a []int to the logger context.
	Ints(key string, value []int) LoggerContext

	// Int8 adds the field key with val as an int8 to the logger context.
	Int8(key string, value int8) LoggerContext

	// Ints8 adds the field key with val as a []int8 to the logger context.
	Ints8(key string, value []int8) LoggerContext

	// Int16 adds the field key with val as an int16 to the logger context.
	Int16(key string, value int16) LoggerContext

	// Ints16 adds the field key with val as a []int16 to the logger context.
	Ints16(key string, value []int16) LoggerContext

	// Int32 adds the field key with val as an int32 to the logger context.
	Int32(key string, value int32) LoggerContext

	// Ints32 adds the field key with val as a []int32 to the logger context.
	Ints32(key string, value []int32) LoggerContext

	// Int64 adds the field key with val as an int64 to the logger context.
	Int64(key string, value int64) LoggerContext

	// Ints64 adds the field key with val as a []int64 to the logger context.
	Ints64(key string, value []int64) LoggerContext

	// Uint adds the field key with val as a uint to the logger context.
	Uint(key string, value uint) LoggerContext

	// Uints adds the field key with val as a []uint to the logger context.
	Uints(key string, value []uint) LoggerContext

	// Uint8 adds the field key with val as a uint8 to the logger context.
	Uint8(key string, value uint8) LoggerContext

	// Uints8 adds the field key with val as a []uint8 to the logger context.
	Uints8(key string, value []uint8) LoggerContext

	// Uint16 adds the field key with val as a uint16 to the logger context.
	Uint16(key string, value uint16) LoggerContext

	// Uints16 adds the field key with val as a []uint16 to the logger context.
	Uints16(key string, value []uint16) LoggerContext

	// Uint32 adds the field key with val as a uint32 to the logger context.
	Uint32(key string, value uint32) LoggerContext

	// Uints32 adds the field key with val as a []uint32 to the logger context.
	Uints32(key string, value []uint32) LoggerContext

	// Uint64 adds the field key with val as a uint64 to the logger context.
	Uint64(key string, value uint64) LoggerContext

	// Uints64 adds the field key with val as a []uint64 to the logger context.
	Uints64(key string, value []uint64) LoggerContext

	// Float32 adds the field key with val as a float32 to the logger context.
	Float32(key string, value float32) LoggerContext

	// Floats32 adds the field key with val as a []float32 to the logger context.
	Floats32(key string, value []float32) LoggerContext

	// Float64 adds the field key with val as a float64 to the logger context.
	Float64(key string, value float64) LoggerContext

	// Floats64 adds the field key with val as a []float64 to the logger context.
	Floats64(key string, value []float64) LoggerContext

	// Bool adds the field key with val as a bool to the logger context.
	Bool(key string, value bool) LoggerContext

	// Bools adds the field key with val as a []bool to the logger context.
	Bools(key string, value []bool) LoggerContext

	// Time adds the field key with val as a time.Time to the logger context.
	Time(key string, value time.Time) LoggerContext

	// Times adds the field key with val as a []time.Time to the logger context.
	Times(key string, value []time.Time) LoggerContext

	// Dur adds the field key with val as a time.Duration to the logger context.
	Dur(key string, value time.Duration) LoggerContext

	// Durs adds the field key with val as a []time.Duration to the logger context.
	Durs(key string, value []time.Duration) LoggerContext

	// TimeDiff adds the field key with val as duration between t and start to the logger context.
	TimeDiff(key string, t time.Time, start time.Time) LoggerContext

	// IPAddr adds the field key with val as a net.IP to the logger context.
	IPAddr(key string, value net.IP) LoggerContext

	// IPPrefix adds the field key with val as a net.IPNet to the logger context.
	IPPrefix(key string, value net.IPNet) LoggerContext

	// MACAddr adds the field key with val as a net.HardwareAddr to the logger context.
	MACAddr(key string, value net.HardwareAddr) LoggerContext

	// Err adds the key "error" with val as an error to the logger context.
	Err(err error) LoggerContext

	// Errs adds the field key with val as a []error to the logger context.
	Errs(key string, errs []error) LoggerContext

	// AnErr adds the field key with val as an error to the logger context.
	AnErr(key string, err error) LoggerContext

	// Any adds the field key with val as an interface{} to the logger context.
	Any(key string, value any) LoggerContext

	// Fields adds the field key with val as a Fields to the logger context.
	Fields(fields Fields) LoggerContext

	// Msg sends the LoggerContext with msg to the logger.
	Msg(msg string)

	// Msgf sends the LoggerContext with formatted msg to the logger.
	Msgf(format string, v ...any)
}

package onelog

import (
	"net"
	"time"
)

// Fields is a helper type for providing extra context in log messages.
type Fields map[string]any

type Logger interface {
	Debug() LoggerContext
	Info() LoggerContext
	Warn() LoggerContext
	Error() LoggerContext
	Fatal() LoggerContext
}

type LoggerContext interface {
	Str(key, value string) LoggerContext
	Strs(key string, value []string) LoggerContext

	Int(key string, value int) LoggerContext
	Ints(key string, value []int) LoggerContext
	Int8(key string, value int8) LoggerContext
	Ints8(key string, value []int8) LoggerContext
	Int16(key string, value int16) LoggerContext
	Ints16(key string, value []int16) LoggerContext
	Int32(key string, value int32) LoggerContext
	Ints32(key string, value []int32) LoggerContext
	Int64(key string, value int64) LoggerContext
	Ints64(key string, value []int64) LoggerContext
	Uint(key string, value uint) LoggerContext
	Uints(key string, value []uint) LoggerContext
	Uint8(key string, value uint8) LoggerContext
	Uints8(key string, value []uint8) LoggerContext
	Uint16(key string, value uint16) LoggerContext
	Uints16(key string, value []uint16) LoggerContext
	Uint32(key string, value uint32) LoggerContext
	Uints32(key string, value []uint32) LoggerContext
	Uint64(key string, value uint64) LoggerContext
	Uints64(key string, value []uint64) LoggerContext
	Float32(key string, value float32) LoggerContext
	Floats32(key string, value []float32) LoggerContext
	Float64(key string, value float64) LoggerContext
	Floats64(key string, value []float64) LoggerContext

	Bool(key string, value bool) LoggerContext
	Bools(key string, value []bool) LoggerContext

	Time(key string, value time.Time) LoggerContext
	Times(key string, value []time.Time) LoggerContext
	Dur(key string, value time.Duration) LoggerContext
	Durs(key string, value []time.Duration) LoggerContext
	TimeDiff(key string, t time.Time, start time.Time) LoggerContext

	IPAddr(key string, value net.IP) LoggerContext
	IPPrefix(key string, value net.IPNet) LoggerContext
	MACAddr(key string, value net.HardwareAddr) LoggerContext

	Err(err error) LoggerContext
	Errs(key string, errs []error) LoggerContext
	AnErr(key string, err error) LoggerContext

	Any(key string, value any) LoggerContext
	Fields(fields Fields) LoggerContext

	Msg(msg string)
	Msgf(format string, v ...any)
}

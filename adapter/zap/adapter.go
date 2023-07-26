package zapadapter

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nikoksr/onelog"
)

// Compile-time check that Adapter and Context implements onelog.Logger and onelog.LoggerContext respectively
var (
	_ onelog.Logger        = (*Adapter)(nil)
	_ onelog.LoggerContext = (*Context)(nil)
)

type (
	// Adapter is a zap adapter for onelog. It implements the onelog.Logger interface.
	Adapter struct {
		logger *zap.Logger
	}

	// Context is the zap logging context. It implements the onelog.LoggerContext interface.
	Context struct {
		level  zapcore.Level
		logger *zap.Logger
		fields []zapcore.Field
	}
)

// NewAdapter creates a new zap adapter for onelog.
func NewAdapter(l *zap.Logger) onelog.Logger {
	return &Adapter{
		logger: l,
	}
}

func (a *Adapter) newContext(level zapcore.Level) onelog.LoggerContext {
	return &Context{
		level:  level,
		logger: a.logger,
		fields: make([]zapcore.Field, 0),
	}
}

// With returns the logger with the given fields.
func (a *Adapter) With(fields ...any) onelog.Logger {
	return &Adapter{logger: a.logger.Sugar().With(fields...).Desugar()}
}

// Debug returns a LoggerContext for a debug log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Debug() onelog.LoggerContext {
	return a.newContext(zap.DebugLevel)
}

// Info returns a LoggerContext for a info log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Info() onelog.LoggerContext {
	return a.newContext(zap.InfoLevel)
}

// Warn returns a LoggerContext for a warn log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Warn() onelog.LoggerContext {
	return a.newContext(zap.WarnLevel)
}

// Error returns a LoggerContext for a error log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Error() onelog.LoggerContext {
	return a.newContext(zap.ErrorLevel)
}

// Fatal returns a LoggerContext for a fatal log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Fatal() onelog.LoggerContext {
	return a.newContext(zap.FatalLevel)
}

func (c *Context) reset() {
	c.fields = make([]zapcore.Field, 0)
}

// Bytes adds the field key with val as a []byte to the logger context.
func (c *Context) Bytes(key string, value []byte) onelog.LoggerContext {
	c.fields = append(c.fields, zap.ByteString(key, value))

	return c
}

// Hex adds the field key with val as a hex string to the logger context.
func (c *Context) Hex(key string, value []byte) onelog.LoggerContext {
	c.fields = append(c.fields, zap.String(key, fmt.Sprintf("%x", value)))

	return c
}

// RawJSON adds the field key with val as a raw json string to the logger context.
func (c *Context) RawJSON(key string, value []byte) onelog.LoggerContext {
	c.fields = append(c.fields, zap.ByteString(key, value))

	return c
}

// Str adds the field key with val as a string to the logger context.
func (c *Context) Str(key string, value string) onelog.LoggerContext {
	c.fields = append(c.fields, zap.String(key, value))

	return c
}

// Strs adds the field key with val as a []string to the logger context.
func (c *Context) Strs(key string, value []string) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Strings(key, value))

	return c
}

// Stringer adds the field key with val as a fmt.Stringer to the logger context.
func (c *Context) Stringer(key string, val fmt.Stringer) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Stringer(key, val))

	return c
}

// Stringers adds the field key with val as a []fmt.Stringer to the logger context.
func (c *Context) Stringers(key string, vals []fmt.Stringer) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Stringers(key, vals))

	return c
}

// Int adds the field key with val as a int to the logger context.
func (c *Context) Int(key string, value int) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int(key, value))

	return c
}

// Ints adds the field key with val as a []int to the logger context.
func (c *Context) Ints(key string, value []int) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Ints(key, value))

	return c
}

// Int8 adds the field key with val as a int8 to the logger context.
func (c *Context) Int8(key string, value int8) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int8(key, value))

	return c
}

// Ints8 adds the field key with val as a []int8 to the logger context.
func (c *Context) Ints8(key string, value []int8) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int8s(key, value))

	return c
}

// Int16 adds the field key with val as a int16 to the logger context.
func (c *Context) Int16(key string, value int16) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int16(key, value))

	return c
}

// Ints16 adds the field key with val as a []int16 to the logger context.
func (c *Context) Ints16(key string, value []int16) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int16s(key, value))

	return c
}

// Int32 adds the field key with val as a int32 to the logger context.
func (c *Context) Int32(key string, value int32) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int32(key, value))

	return c
}

// Ints32 adds the field key with val as a []int32 to the logger context.
func (c *Context) Ints32(key string, value []int32) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int32s(key, value))

	return c
}

// Int64 adds the field key with val as a int64 to the logger context.
func (c *Context) Int64(key string, value int64) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int64(key, value))

	return c
}

// Ints64 adds the field key with val as a []int64 to the logger context.
func (c *Context) Ints64(key string, value []int64) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Int64s(key, value))

	return c
}

// Uint adds the field key with val as a uint to the logger context.
func (c *Context) Uint(key string, value uint) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint(key, value))

	return c
}

// Uints adds the field key with val as a []uint to the logger context.
func (c *Context) Uints(key string, value []uint) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uints(key, value))

	return c
}

// Uint8 adds the field key with val as a uint8 to the logger context.
func (c *Context) Uint8(key string, value uint8) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint8(key, value))

	return c
}

// Uints8 adds the field key with val as a []uint8 to the logger context.
func (c *Context) Uints8(key string, value []uint8) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint8s(key, value))

	return c
}

// Uint16 adds the field key with val as a uint16 to the logger context.
func (c *Context) Uint16(key string, value uint16) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint16(key, value))

	return c
}

// Uints16 adds the field key with val as a []uint16 to the logger context.
func (c *Context) Uints16(key string, value []uint16) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint16s(key, value))

	return c
}

// Uint32 adds the field key with val as a uint32 to the logger context.
func (c *Context) Uint32(key string, value uint32) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint32(key, value))

	return c
}

// Uints32 adds the field key with val as a []uint32 to the logger context.
func (c *Context) Uints32(key string, value []uint32) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint32s(key, value))

	return c
}

// Uint64 adds the field key with val as a uint64 to the logger context.
func (c *Context) Uint64(key string, value uint64) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint64(key, value))

	return c
}

// Uints64 adds the field key with val as a []uint64 to the logger context.
func (c *Context) Uints64(key string, value []uint64) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Uint64s(key, value))

	return c
}

// Float32 adds the field key with val as a float32 to the logger context.
func (c *Context) Float32(key string, value float32) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Float32(key, value))

	return c
}

// Floats32 adds the field key with val as a []float32 to the logger context.
func (c *Context) Floats32(key string, value []float32) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Float32s(key, value))

	return c
}

// Float64 adds the field key with val as a float64 to the logger context.
func (c *Context) Float64(key string, value float64) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Float64(key, value))

	return c
}

// Floats64 adds the field key with val as a []float64 to the logger context.
func (c *Context) Floats64(key string, value []float64) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Float64s(key, value))

	return c
}

// Bool adds the field key with val as a bool to the logger context.
func (c *Context) Bool(key string, value bool) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Bool(key, value))

	return c
}

// Bools adds the field key with val as a []bool to the logger context.
func (c *Context) Bools(key string, value []bool) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Bools(key, value))

	return c
}

// Time adds the field key with val as a time.Time to the logger context.
func (c *Context) Time(key string, value time.Time) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Time(key, value))

	return c
}

// Times adds the field key with val as a []time.Time to the logger context.
func (c *Context) Times(key string, value []time.Time) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Times(key, value))

	return c
}

// Dur adds the field key with val as a time.Duration to the logger context.
func (c *Context) Dur(key string, value time.Duration) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Duration(key, value))

	return c
}

// Durs adds the field key with val as a []time.Duration to the logger context.
func (c *Context) Durs(key string, value []time.Duration) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Durations(key, value))

	return c
}

// TimeDiff adds the field key with begin and end as a time.Time to the logger context.
func (c *Context) TimeDiff(key string, begin, end time.Time) onelog.LoggerContext {
	diff := end.Sub(begin)
	c.fields = append(c.fields, zap.Duration(key, diff))

	return c
}

// IPAddr adds the field key with val as a net.IP to the logger context.
func (c *Context) IPAddr(key string, value net.IP) onelog.LoggerContext {
	c.fields = append(c.fields, zap.String(key, value.String()))

	return c
}

// IPPrefix adds the field key with val as a net.IPNet to the logger context.
func (c *Context) IPPrefix(key string, value net.IPNet) onelog.LoggerContext {
	c.fields = append(c.fields, zap.String(key, value.String()))

	return c
}

// MACAddr adds the field key with val as a net.HardwareAddr to the logger context.
func (c *Context) MACAddr(key string, value net.HardwareAddr) onelog.LoggerContext {
	c.fields = append(c.fields, zap.String(key, value.String()))

	return c
}

// AnErr adds the field key with val as a error to the logger context.
func (c *Context) AnErr(key string, err error) onelog.LoggerContext {
	c.fields = append(c.fields, zap.NamedError(key, err))

	return c
}

// Err adds the field key with val as a error to the logger context.
func (c *Context) Err(err error) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Error(err))

	return c
}

// Errs adds the field key with val as a []error to the logger context.
func (c *Context) Errs(key string, errs []error) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Errors(key, errs))

	return c
}

// Any adds the field key with val as a arbitrary value to the logger context.
func (c *Context) Any(key string, value any) onelog.LoggerContext {
	c.fields = append(c.fields, zap.Any(key, value))

	return c
}

func (c *Context) Fields(fields onelog.Fields) onelog.LoggerContext {
	for k, v := range fields {
		c.fields = append(c.fields, zap.Any(k, v))
	}

	return c
}

// Msg sends the LoggerContext with msg to the logger.
func (c *Context) Msg(msg string) {
	c.logger.Log(c.level, msg, c.fields...)
	c.reset()
}

// Msgf sends the LoggerContext with formatted msg to the logger.
func (c *Context) Msgf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	c.Msg(msg)
}

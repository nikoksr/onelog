package slogadapter

import (
	"fmt"
	"net"
	"time"

	"github.com/shopspring/decimal"

	"golang.org/x/exp/slog"

	"github.com/nikoksr/onelog"
)

// Compile-time check that Adapter and Context implements onelog.Logger and onelog.LoggerContext respectively
var (
	_ onelog.Logger        = (*Adapter)(nil)
	_ onelog.LoggerContext = (*Context)(nil)
)

type (
	// Adapter is a slog adapter for onelog. It implements the onelog.Logger interface.
	Adapter struct {
		logger *slog.Logger
	}

	// Context is the slog logging context. It implements the onelog.LoggerContext interface.
	Context struct {
		level  slog.Level
		logger *slog.Logger
		fields []any
	}
)

// NewAdapter creates a new slog adapter for onelog.
func NewAdapter(l *slog.Logger) onelog.Logger {
	return &Adapter{
		logger: l,
	}
}

func (a *Adapter) newContext(level slog.Level) *Context {
	return &Context{
		level:  level,
		logger: a.logger,
		fields: make([]any, 0),
	}
}

// With returns the logger with the given fields.
func (a *Adapter) With(fields ...any) onelog.Logger {
	return &Adapter{logger: a.logger.With(fields...)}
}

// Debug returns a LoggerContext for a debug log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Debug() onelog.LoggerContext {
	return a.newContext(slog.LevelDebug)
}

// Info returns a LoggerContext for an info log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Info() onelog.LoggerContext {
	return a.newContext(slog.LevelInfo)
}

// Warn returns a LoggerContext for a warn log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Warn() onelog.LoggerContext {
	return a.newContext(slog.LevelWarn)
}

// Error returns a LoggerContext for an error log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Error() onelog.LoggerContext {
	return a.newContext(slog.LevelError)
}

// Fatal returns a LoggerContext for a fatal log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Fatal() onelog.LoggerContext {
	return a.newContext(slog.LevelError) // Using Error level here because Fatal is not supported by slog
}

// Bytes adds the field key with val as a []byte to the logger context.
func (c *Context) Bytes(key string, value []byte) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, string(value)))

	return c
}

// Hex adds the field key with val as a hex string to the logger context.
func (c *Context) Hex(key string, value []byte) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, fmt.Sprintf("%x", value)))

	return c
}

// RawJSON adds the field key with val as a raw JSON string to the logger context.
func (c *Context) RawJSON(key string, value []byte) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, string(value)))

	return c
}

// Str adds the field key with val as a string to the logger context.
func (c *Context) Str(key string, value string) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, value))

	return c
}

// Strs adds the field key with val as a []string to the logger context.
func (c *Context) Strs(key string, value []string) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Stringer adds the field key with val as a fmt.Stringer to the logger context.
func (c *Context) Stringer(key string, value fmt.Stringer) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, value.String()))

	return c
}

// Stringers adds the field key with val as a []fmt.Stringer to the logger context.
func (c *Context) Stringers(key string, value []fmt.Stringer) onelog.LoggerContext {
	// Todo: Better way to do this?
	strs := make([]string, len(value))
	for i, str := range value {
		strs[i] = str.String()
	}
	c.fields = append(c.fields, slog.Any(key, strs))

	return c
}

// Int adds the field key with val as a int to the logger context.
func (c *Context) Int(key string, value int) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Int(key, value))

	return c
}

// Ints adds the field key with val as a []int to the logger context.
func (c *Context) Ints(key string, value []int) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Int8 adds the field key with val as a int8 to the logger context.
func (c *Context) Int8(key string, value int8) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Int64(key, int64(value)))

	return c
}

// Ints8 adds the field key with val as a []int8 to the logger context.
func (c *Context) Ints8(key string, value []int8) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Int16 adds the field key with val as a int16 to the logger context.
func (c *Context) Int16(key string, value int16) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Int64(key, int64(value)))

	return c
}

// Ints16 adds the field key with val as a []int16 to the logger context.
func (c *Context) Ints16(key string, value []int16) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Int32 adds the field key with val as a int32 to the logger context.
func (c *Context) Int32(key string, value int32) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Int64(key, int64(value)))

	return c
}

// Ints32 adds the field key with val as a []int32 to the logger context.
func (c *Context) Ints32(key string, value []int32) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Int64 adds the field key with val as a int64 to the logger context.
func (c *Context) Int64(key string, value int64) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Int64(key, value))

	return c
}

// Ints64 adds the field key with val as a []int64 to the logger context.
func (c *Context) Ints64(key string, value []int64) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Uint adds the field key with val as a uint to the logger context.
func (c *Context) Uint(key string, value uint) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Uint64(key, uint64(value)))

	return c
}

// Uints adds the field key with val as a []uint to the logger context.
func (c *Context) Uints(key string, value []uint) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Uint8 adds the field key with val as a uint8 to the logger context.
func (c *Context) Uint8(key string, value uint8) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Uint64(key, uint64(value)))

	return c
}

// Uints8 adds the field key with val as a []uint8 to the logger context.
func (c *Context) Uints8(key string, value []uint8) onelog.LoggerContext {
	// Todo: Better way to do this?
	// Convert []uint8 to []uint64
	uints := make([]uint64, len(value))
	for i, v := range value {
		uints[i] = uint64(v)
	}

	c.fields = append(c.fields, slog.Any(key, uints))

	return c
}

// Uint16 adds the field key with val as a uint16 to the logger context.
func (c *Context) Uint16(key string, value uint16) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Uint64(key, uint64(value)))

	return c
}

// Uints16 adds the field key with val as a []uint16 to the logger context.
func (c *Context) Uints16(key string, value []uint16) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Uint32 adds the field key with val as a uint32 to the logger context.
func (c *Context) Uint32(key string, value uint32) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Uint64(key, uint64(value)))

	return c
}

// Uints32 adds the field key with val as a []uint32 to the logger context.
func (c *Context) Uints32(key string, value []uint32) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Uint64 adds the field key with val as a uint64 to the logger context.
func (c *Context) Uint64(key string, value uint64) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Uint64(key, value))

	return c
}

// Uints64 adds the field key with val as a []uint64 to the logger context.
func (c *Context) Uints64(key string, value []uint64) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Float32 adds the field key with val as a float32 to the logger context.
func (c *Context) Float32(key string, value float32) onelog.LoggerContext {
	d, _ := decimal.NewFromFloat32(value).Float64()

	c.fields = append(c.fields, slog.Float64(key, d))

	return c
}

// Floats32 adds the field key with val as a []float32 to the logger context.
func (c *Context) Floats32(key string, value []float32) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Float64 adds the field key with val as a float64 to the logger context.
func (c *Context) Float64(key string, value float64) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Float64(key, value))

	return c
}

// Floats64 adds the field key with val as a []float64 to the logger context.
func (c *Context) Floats64(key string, value []float64) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Bool adds the field key with val as a bool to the logger context.
func (c *Context) Bool(key string, value bool) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Bool(key, value))

	return c
}

// Bools adds the field key with val as a []bool to the logger context.
func (c *Context) Bools(key string, value []bool) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Time adds the field key with val as a time.Time to the logger context.
func (c *Context) Time(key string, value time.Time) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Time(key, value))

	return c
}

// Times adds the field key with val as a []time.Time to the logger context.
func (c *Context) Times(key string, value []time.Time) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// Dur adds the field key with val as a time.Duration to the logger context.
func (c *Context) Dur(key string, value time.Duration) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Duration(key, value))

	return c
}

// Durs adds the field key with val as a []time.Duration to the logger context.
func (c *Context) Durs(key string, value []time.Duration) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

// TimeDiff adds the field key with begin and end as a time.Time to the logger context.
func (c *Context) TimeDiff(key string, begin, end time.Time) onelog.LoggerContext {
	diff := end.Sub(begin)
	c.fields = append(c.fields, slog.Duration(key, diff))

	return c
}

// IPAddr adds the field key with val as a net.IPAddr to the logger context.
func (c *Context) IPAddr(key string, value net.IP) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, value.String()))

	return c
}

// IPPrefix adds the field key with val as a net.IPPrefix to the logger context.
func (c *Context) IPPrefix(key string, value net.IPNet) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, value.String()))

	return c
}

// MACAddr adds the field key with val as a net.HardwareAddr to the logger context.
func (c *Context) MACAddr(key string, value net.HardwareAddr) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, value.String()))

	return c
}

// AnErr adds the field key with val as a error to the logger context.
func (c *Context) AnErr(key string, value error) onelog.LoggerContext {
	c.fields = append(c.fields, slog.String(key, value.Error()))

	return c
}

// Err adds the field "error" with val as a error to the logger context.
func (c *Context) Err(value error) onelog.LoggerContext {
	c.AnErr("error", value)

	return c
}

// Errs adds the field "error" with val as a []error to the logger context.
func (c *Context) Errs(key string, value []error) onelog.LoggerContext {
	// Todo: Better way to do this?
	// Convert []error to []string. If we don't do this, slog prints empty objects
	errs := make([]string, len(value))
	for i, err := range value {
		errs[i] = err.Error()
	}

	c.fields = append(c.fields, slog.Any(key, errs))

	return c
}

// Any adds the field key with val as a arbitrary value to the logger context.
func (c *Context) Any(key string, value any) onelog.LoggerContext {
	c.fields = append(c.fields, slog.Any(key, value))

	return c
}

func (c *Context) Fields(fields onelog.Fields) onelog.LoggerContext {
	for key, value := range fields {
		c.fields = append(c.fields, slog.Any(key, value))
	}

	return c
}

// Msg sends the LoggerContext with msg to the logger.
func (c *Context) Msg(msg string) {
	//nolint:staticcheck // passing a nil context is fine, check slog.Logger.Info implementation for example
	c.logger.Log(nil, c.level, msg, c.fields...)
	c.fields = make([]any, 0) // reset fields
}

// Msgf sends the LoggerContext with formatted msg to the logger.
func (c *Context) Msgf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	c.Msg(msg)
}

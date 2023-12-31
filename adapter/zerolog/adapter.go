package zerologadapter

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"

	"github.com/nikoksr/onelog"
)

// Compile-time check that Adapter and Context implements onelog.Logger and onelog.LoggerContext respectively
var (
	_ onelog.Logger        = (*Adapter)(nil)
	_ onelog.LoggerContext = (*Context)(nil)
)

type (
	// Adapter is a zerolog adapter for onelog. It implements the onelog.Logger interface.
	Adapter struct {
		logger *zerolog.Logger
	}

	// Context is the zerolog logging context. It implements the onelog.LoggerContext interface.
	Context struct {
		logger       *zerolog.Logger
		event        *zerolog.Event
		resetEventFn func() *zerolog.Event
	}
)

// NewAdapter creates a new zerolog adapter for onelog.
func NewAdapter(l *zerolog.Logger) onelog.Logger {
	return &Adapter{
		logger: l,
	}
}

// With returns the logger with the given fields.
func (a *Adapter) With(fields ...any) onelog.Logger {
	logger := a.logger.With().Fields(fields).Logger()
	return &Adapter{logger: &logger}
}

// Debug returns a LoggerContext for a debug log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Debug() onelog.LoggerContext {
	return &Context{
		logger:       a.logger,
		event:        a.logger.Debug(),
		resetEventFn: a.logger.Debug,
	}
}

// Info returns a LoggerContext for a info log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Info() onelog.LoggerContext {
	return &Context{
		logger:       a.logger,
		event:        a.logger.Info(),
		resetEventFn: a.logger.Info,
	}
}

// Warn returns a LoggerContext for a warn log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Warn() onelog.LoggerContext {
	return &Context{
		logger:       a.logger,
		event:        a.logger.Warn(),
		resetEventFn: a.logger.Warn,
	}
}

// Error returns a LoggerContext for a error log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Error() onelog.LoggerContext {
	return &Context{
		logger:       a.logger,
		event:        a.logger.Error(),
		resetEventFn: a.logger.Error,
	}
}

// Fatal returns a LoggerContext for a fatal log. To send the log, use the Msg or Msgf methods.
func (a *Adapter) Fatal() onelog.LoggerContext {
	return &Context{
		logger:       a.logger,
		event:        a.logger.Fatal(),
		resetEventFn: a.logger.Fatal,
	}
}

func (c *Context) reset() {
	c.event = c.resetEventFn()
}

// Bytes adds the field key with val as a []byte to the logger context.
func (c *Context) Bytes(key string, value []byte) onelog.LoggerContext {
	c.event.Bytes(key, value)

	return c
}

// Hex adds the field key with val as a hex string to the logger context.
func (c *Context) Hex(key string, value []byte) onelog.LoggerContext {
	c.event.Hex(key, value)

	return c
}

// RawJSON adds the field key with val as a json.RawMessage to the logger context.
func (c *Context) RawJSON(key string, value []byte) onelog.LoggerContext {
	c.event.RawJSON(key, value)

	return c
}

// Str adds the field key with val as a string to the logger context.
func (c *Context) Str(key, value string) onelog.LoggerContext {
	c.event.Str(key, value)

	return c
}

// Strs adds the field key with val as a []string to the logger context.
func (c *Context) Strs(key string, value []string) onelog.LoggerContext {
	c.event.Strs(key, value)

	return c
}

// Stringer adds the field key with val as a fmt.Stringer to the logger context.
func (c *Context) Stringer(key string, val fmt.Stringer) onelog.LoggerContext {
	c.event.Stringer(key, val)

	return c
}

// Stringers adds the field key with val as a []fmt.Stringer to the logger context.
func (c *Context) Stringers(key string, vals []fmt.Stringer) onelog.LoggerContext {
	c.event.Stringers(key, vals)

	return c
}

// Int adds the field key with i as a int to the logger context.
func (c *Context) Int(key string, value int) onelog.LoggerContext {
	c.event.Int(key, value)

	return c
}

// Ints adds the field key with i as a []int to the logger context.
func (c *Context) Ints(key string, value []int) onelog.LoggerContext {
	c.event.Ints(key, value)

	return c
}

// Int8 adds the field key with i as a int8 to the logger context.
func (c *Context) Int8(key string, value int8) onelog.LoggerContext {
	c.event.Int8(key, value)

	return c
}

// Ints8 adds the field key with i as a []int8 to the logger context.
func (c *Context) Ints8(key string, value []int8) onelog.LoggerContext {
	c.event.Ints8(key, value)

	return c
}

// Int16 adds the field key with i as a int16 to the logger context.
func (c *Context) Int16(key string, value int16) onelog.LoggerContext {
	c.event.Int16(key, value)

	return c
}

// Ints16 adds the field key with i as a []int16 to the logger context.
func (c *Context) Ints16(key string, value []int16) onelog.LoggerContext {
	c.event.Ints16(key, value)

	return c
}

// Int32 adds the field key with i as a int32 to the logger context.
func (c *Context) Int32(key string, value int32) onelog.LoggerContext {
	c.event.Int32(key, value)

	return c
}

// Ints32 adds the field key with i as a []int32 to the logger context.
func (c *Context) Ints32(key string, value []int32) onelog.LoggerContext {
	c.event.Ints32(key, value)

	return c
}

// Int64 adds the field key with i as a int64 to the logger context.
func (c *Context) Int64(key string, value int64) onelog.LoggerContext {
	c.event.Int64(key, value)

	return c
}

// Ints64 adds the field key with i as a []int64 to the logger context.
func (c *Context) Ints64(key string, value []int64) onelog.LoggerContext {
	c.event.Ints64(key, value)

	return c
}

// Uint adds the field key with i as a uint to the logger context.
func (c *Context) Uint(key string, value uint) onelog.LoggerContext {
	c.event.Uint(key, value)

	return c
}

// Uints adds the field key with i as a []uint to the logger context.
func (c *Context) Uints(key string, value []uint) onelog.LoggerContext {
	c.event.Uints(key, value)

	return c
}

// Uint8 adds the field key with i as a uint8 to the logger context.
func (c *Context) Uint8(key string, value uint8) onelog.LoggerContext {
	c.event.Uint8(key, value)

	return c
}

// Uints8 adds the field key with i as a []uint8 to the logger context.
func (c *Context) Uints8(key string, value []uint8) onelog.LoggerContext {
	c.event.Uints8(key, value)

	return c
}

// Uint16 adds the field key with i as a uint16 to the logger context.
func (c *Context) Uint16(key string, value uint16) onelog.LoggerContext {
	c.event.Uint16(key, value)

	return c
}

// Uints16 adds the field key with i as a []uint16 to the logger context.
func (c *Context) Uints16(key string, value []uint16) onelog.LoggerContext {
	c.event.Uints16(key, value)

	return c
}

// Uint32 adds the field key with i as a uint32 to the logger context.
func (c *Context) Uint32(key string, value uint32) onelog.LoggerContext {
	c.event.Uint32(key, value)

	return c
}

// Uints32 adds the field key with i as a []uint32 to the logger context.
func (c *Context) Uints32(key string, value []uint32) onelog.LoggerContext {
	c.event.Uints32(key, value)

	return c
}

// Uint64 adds the field key with i as a uint64 to the logger context.
func (c *Context) Uint64(key string, value uint64) onelog.LoggerContext {
	c.event.Uint64(key, value)

	return c
}

// Uints64 adds the field key with i as a []uint64 to the logger context.
func (c *Context) Uints64(key string, value []uint64) onelog.LoggerContext {
	c.event.Uints64(key, value)

	return c
}

// Float32 adds the field key with f as a float32 to the logger context.
func (c *Context) Float32(key string, value float32) onelog.LoggerContext {
	c.event.Float32(key, value)

	return c
}

// Floats32 adds the field key with f as a []float32 to the logger context.
func (c *Context) Floats32(key string, value []float32) onelog.LoggerContext {
	c.event.Floats32(key, value)

	return c
}

// Float64 adds the field key with f as a float64 to the logger context.
func (c *Context) Float64(key string, value float64) onelog.LoggerContext {
	c.event.Float64(key, value)

	return c
}

// Floats64 adds the field key with f as a []float64 to the logger context.
func (c *Context) Floats64(key string, value []float64) onelog.LoggerContext {
	c.event.Floats64(key, value)

	return c
}

// Bool adds the field key with b as a bool to the logger context.
func (c *Context) Bool(key string, value bool) onelog.LoggerContext {
	c.event.Bool(key, value)

	return c
}

// Bools adds the field key with b as a []bool to the logger context.
func (c *Context) Bools(key string, value []bool) onelog.LoggerContext {
	c.event.Bools(key, value)

	return c
}

// Time adds the field key with t as a time.Time to the logger context.
func (c *Context) Time(key string, value time.Time) onelog.LoggerContext {
	c.event.Time(key, value)

	return c
}

// Times adds the field key with t as a []time.Time to the logger context.
func (c *Context) Times(key string, value []time.Time) onelog.LoggerContext {
	c.event.Times(key, value)

	return c
}

// Dur adds the field key with d as a time.Duration to the logger context.
func (c *Context) Dur(key string, value time.Duration) onelog.LoggerContext {
	c.event.Dur(key, value)

	return c
}

// Durs adds the field key with d as a []time.Duration to the logger context.
func (c *Context) Durs(key string, value []time.Duration) onelog.LoggerContext {
	c.event.Durs(key, value)

	return c
}

// TimeDiff adds the field key with begin and end as a time.Time to the logger context.
func (c *Context) TimeDiff(key string, begin, end time.Time) onelog.LoggerContext {
	c.event.TimeDiff(key, begin, end)

	return c
}

// IPAddr adds the field key with ip as a net.IP to the logger context.
func (c *Context) IPAddr(key string, value net.IP) onelog.LoggerContext {
	c.event.IPAddr(key, value)

	return c
}

// IPPrefix adds the field key with ip as a net.IPNet to the logger context.
func (c *Context) IPPrefix(key string, value net.IPNet) onelog.LoggerContext {
	c.event.IPPrefix(key, value)

	return c
}

// MACAddr adds the field key with ip as a net.HardwareAddr to the logger context.
func (c *Context) MACAddr(key string, value net.HardwareAddr) onelog.LoggerContext {
	c.event.MACAddr(key, value)

	return c
}

// Err adds the field "error" with err as a error to the logger context.
func (c *Context) Err(err error) onelog.LoggerContext {
	c.event.Err(err)

	return c
}

// Errs adds the field key with errs as a []error to the logger context.
func (c *Context) Errs(key string, errs []error) onelog.LoggerContext {
	c.event.Errs(key, errs)

	return c
}

// AnErr adds the field key with err as a error to the logger context.
func (c *Context) AnErr(key string, err error) onelog.LoggerContext {
	c.event.AnErr(key, err)

	return c
}

// Any adds the field key with val as a arbitrary value to the logger context.
func (c *Context) Any(key string, value any) onelog.LoggerContext {
	c.event.Any(key, value)

	return c
}

// Fields adds the fields to the logger context.
func (c *Context) Fields(fields onelog.Fields) onelog.LoggerContext {
	c.event.Fields(fields)

	return c
}

// Msg sends the LoggerContext with msg to the logger.
func (c *Context) Msg(msg string) {
	c.event.Msg(msg)
	c.reset()
}

// Msgf sends the LoggerContext with formatted msg to the logger.
func (c *Context) Msgf(format string, v ...any) {
	c.event.Msgf(format, v...)
	c.reset()
}

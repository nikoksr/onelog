package zerologadapter

import (
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
	Adapter struct {
		logger *zerolog.Logger
	}

	Context struct {
		logger *zerolog.Logger
		event  *zerolog.Event
	}
)

func NewAdapter(l *zerolog.Logger) onelog.Logger {
	return &Adapter{
		logger: l,
	}
}

func (a *Adapter) Debug() onelog.LoggerContext {
	return &Context{
		logger: a.logger,
		event:  a.logger.Debug(),
	}
}

func (a *Adapter) Info() onelog.LoggerContext {
	return &Context{
		logger: a.logger,
		event:  a.logger.Info(),
	}
}

func (a *Adapter) Warn() onelog.LoggerContext {
	return &Context{
		logger: a.logger,
		event:  a.logger.Warn(),
	}
}

func (a *Adapter) Error() onelog.LoggerContext {
	return &Context{
		logger: a.logger,
		event:  a.logger.Error(),
	}
}

func (a *Adapter) Fatal() onelog.LoggerContext {
	return &Context{
		logger: a.logger,
		event:  a.logger.Fatal(),
	}
}

// Str adds the field key with val as a string to the logger context.
func (c *Context) Str(key, value string) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Str(key, value)

	return c
}

// Strs adds the field key with val as a []string to the logger context.
func (c *Context) Strs(key string, value []string) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Strs(key, value)

	return c
}

// Int adds the field key with i as a int to the logger context.
func (c *Context) Int(key string, value int) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Int(key, value)

	return c
}

// Ints adds the field key with i as a []int to the logger context.
func (c *Context) Ints(key string, value []int) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Ints(key, value)

	return c
}

// Int8 adds the field key with i as a int8 to the logger context.
func (c *Context) Int8(key string, value int8) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Int8(key, value)

	return c
}

// Ints8 adds the field key with i as a []int8 to the logger context.
func (c *Context) Ints8(key string, value []int8) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Ints8(key, value)

	return c
}

// Int16 adds the field key with i as a int16 to the logger context.
func (c *Context) Int16(key string, value int16) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Int16(key, value)

	return c
}

// Ints16 adds the field key with i as a []int16 to the logger context.
func (c *Context) Ints16(key string, value []int16) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Ints16(key, value)

	return c
}

// Int32 adds the field key with i as a int32 to the logger context.
func (c *Context) Int32(key string, value int32) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Int32(key, value)

	return c
}

// Ints32 adds the field key with i as a []int32 to the logger context.
func (c *Context) Ints32(key string, value []int32) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Ints32(key, value)

	return c
}

// Int64 adds the field key with i as a int64 to the logger context.
func (c *Context) Int64(key string, value int64) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Int64(key, value)

	return c
}

// Ints64 adds the field key with i as a []int64 to the logger context.
func (c *Context) Ints64(key string, value []int64) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Ints64(key, value)

	return c
}

// Uint adds the field key with i as a uint to the logger context.
func (c *Context) Uint(key string, value uint) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uint(key, value)

	return c
}

// Uints adds the field key with i as a []uint to the logger context.
func (c *Context) Uints(key string, value []uint) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uints(key, value)

	return c
}

// Uint8 adds the field key with i as a uint8 to the logger context.
func (c *Context) Uint8(key string, value uint8) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uint8(key, value)

	return c
}

// Uints8 adds the field key with i as a []uint8 to the logger context.
func (c *Context) Uints8(key string, value []uint8) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uints8(key, value)

	return c
}

// Uint16 adds the field key with i as a uint16 to the logger context.
func (c *Context) Uint16(key string, value uint16) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uint16(key, value)

	return c
}

// Uints16 adds the field key with i as a []uint16 to the logger context.
func (c *Context) Uints16(key string, value []uint16) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uints16(key, value)

	return c
}

// Uint32 adds the field key with i as a uint32 to the logger context.
func (c *Context) Uint32(key string, value uint32) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uint32(key, value)

	return c
}

// Uints32 adds the field key with i as a []uint32 to the logger context.
func (c *Context) Uints32(key string, value []uint32) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uints32(key, value)

	return c
}

// Uint64 adds the field key with i as a uint64 to the logger context.
func (c *Context) Uint64(key string, value uint64) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uint64(key, value)

	return c
}

// Uints64 adds the field key with i as a []uint64 to the logger context.
func (c *Context) Uints64(key string, value []uint64) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Uints64(key, value)

	return c
}

// Float32 adds the field key with f as a float32 to the logger context.
func (c *Context) Float32(key string, value float32) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Float32(key, value)

	return c
}

// Floats32 adds the field key with f as a []float32 to the logger context.
func (c *Context) Floats32(key string, value []float32) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Floats32(key, value)

	return c
}

// Float64 adds the field key with f as a float64 to the logger context.
func (c *Context) Float64(key string, value float64) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Float64(key, value)

	return c
}

// Floats64 adds the field key with f as a []float64 to the logger context.
func (c *Context) Floats64(key string, value []float64) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Floats64(key, value)

	return c
}

// Bool adds the field key with b as a bool to the logger context.
func (c *Context) Bool(key string, value bool) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Bool(key, value)

	return c
}

// Bools adds the field key with b as a []bool to the logger context.
func (c *Context) Bools(key string, value []bool) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Bools(key, value)

	return c
}

// Time adds the field key with t as a time.Time to the logger context.
func (c *Context) Time(key string, value time.Time) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Time(key, value)

	return c
}

// Times adds the field key with t as a []time.Time to the logger context.
func (c *Context) Times(key string, value []time.Time) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Times(key, value)

	return c
}

// Dur adds the field key with d as a time.Duration to the logger context.
func (c *Context) Dur(key string, value time.Duration) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Dur(key, value)

	return c
}

// Durs adds the field key with d as a []time.Duration to the logger context.
func (c *Context) Durs(key string, value []time.Duration) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Durs(key, value)

	return c
}

// IPAddr adds the field key with ip as a net.IP to the logger context.
func (c *Context) IPAddr(key string, value net.IP) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.IPAddr(key, value)

	return c
}

// IPPrefix adds the field key with ip as a net.IPNet to the logger context.
func (c *Context) IPPrefix(key string, value net.IPNet) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.IPPrefix(key, value)

	return c
}

// MACAddr adds the field key with ip as a net.HardwareAddr to the logger context.
func (c *Context) MACAddr(key string, value net.HardwareAddr) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.MACAddr(key, value)

	return c
}

// Err adds the field "error" with err as a error to the logger context.
func (c *Context) Err(err error) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Err(err)

	return c
}

// Errs adds the field key with errs as a []error to the logger context.
func (c *Context) Errs(key string, errs []error) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Errs(key, errs)

	return c
}

// AnErr adds the field key with err as a error to the logger context.
func (c *Context) AnErr(key string, err error) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.AnErr(key, err)

	return c
}

// Any adds the field key with val as a arbitrary value to the logger context.
func (c *Context) Any(key string, value any) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Any(key, value)

	return c
}

// Fields adds the fields to the logger context.
func (c *Context) Fields(fields onelog.Fields) onelog.LoggerContext {
	if c == nil {
		return nil
	}

	c.event.Fields(fields)

	return c
}

// Msg sends the logger context with level DEBUG.
func (c *Context) Msg(msg string) {
	if c == nil {
		return
	}

	c.event.Msg(msg)
}

// Msgf sends the logger context with level DEBUG.
func (c *Context) Msgf(format string, v ...any) {
	if c == nil {
		return
	}

	c.event.Msgf(format, v...)
}

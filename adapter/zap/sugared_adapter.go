package zapadapter

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/nikoksr/onelog"
)

// Compile-time check that Adapter and Context implements onelog.Logger and onelog.LoggerContext respectively
var (
	_ onelog.Logger        = (*Adapter)(nil)
	_ onelog.LoggerContext = (*Context)(nil)
)

type (
	// SugarAdapter is a zap-sugared adapter for onelog. It implements the onelog.Logger interface.
	SugarAdapter struct {
		logger *zap.SugaredLogger
	}

	// SugarContext is the zap-sugared logging context. It implements the onelog.LoggerContext interface.
	SugarContext struct {
		level  zapcore.Level
		logger *zap.SugaredLogger
		fields []any
	}
)

// NewSugarAdapter creates a new zap-sugared adapter for onelog.
func NewSugarAdapter(l *zap.SugaredLogger) onelog.Logger {
	return &SugarAdapter{
		logger: l,
	}
}

func (a *SugarAdapter) newContext(level zapcore.Level) onelog.LoggerContext {
	return &SugarContext{
		level:  level,
		logger: a.logger,
		fields: make([]any, 0),
	}
}

// With returns the logger with the given fields.
func (a *SugarAdapter) With(fields ...any) onelog.Logger {
	return &SugarAdapter{logger: a.logger.With(fields...)}
}

// Debug returns a LoggerContext for a debug log. To send the log, use the Msg or Msgf methods.
func (a *SugarAdapter) Debug() onelog.LoggerContext {
	return a.newContext(zapcore.DebugLevel)
}

// Info returns a LoggerContext for an info log. To send the log, use the Msg or Msgf methods.
func (a *SugarAdapter) Info() onelog.LoggerContext {
	return a.newContext(zapcore.InfoLevel)
}

// Warn returns a LoggerContext for a warning log. To send the log, use the Msg or Msgf methods.
func (a *SugarAdapter) Warn() onelog.LoggerContext {
	return a.newContext(zapcore.WarnLevel)
}

// Error returns a LoggerContext for an error log. To send the log, use the Msg or Msgf methods.
func (a *SugarAdapter) Error() onelog.LoggerContext {
	return a.newContext(zapcore.ErrorLevel)
}

// Fatal returns a LoggerContext for a fatal log. To send the log, use the Msg or Msgf methods.
func (a *SugarAdapter) Fatal() onelog.LoggerContext {
	return a.newContext(zapcore.FatalLevel)
}

func (c *SugarContext) addField(key string, value any) {
	c.fields = append(c.fields, key)
	c.fields = append(c.fields, value)
}

func (c *SugarContext) addFields(fields onelog.Fields) {
	for key, value := range fields {
		c.addField(key, value)
	}
}

// Bytes adds the field key with val as a []byte to the logger context.
func (c *SugarContext) Bytes(key string, value []byte) onelog.LoggerContext {
	c.addField(key, string(value))

	return c
}

// Hex adds the field key with val as a hex string to the logger context.
func (c *SugarContext) Hex(key string, value []byte) onelog.LoggerContext {
	c.addField(key, fmt.Sprintf("%x", value))

	return c
}

// RawJSON adds the field key with val as a json.RawMessage to the logger context.
func (c *SugarContext) RawJSON(key string, value []byte) onelog.LoggerContext {
	c.addField(key, string(value))

	return c
}

// Str adds the field key with val as a string to the logger context.
func (c *SugarContext) Str(key string, value string) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Strs adds the field key with val as a []string to the logger context.
func (c *SugarContext) Strs(key string, value []string) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Stringer adds the field key with val as a fmt.Stringer to the logger context.
func (c *SugarContext) Stringer(key string, val fmt.Stringer) onelog.LoggerContext {
	return c.Str(key, val.String())
}

// Stringers adds the field key with val as a []fmt.Stringer to the logger context.
func (c *SugarContext) Stringers(key string, vals []fmt.Stringer) onelog.LoggerContext {
	// Todo: Better way to do this?
	strings := make([]string, len(vals))
	for i, val := range vals {
		strings[i] = val.String()
	}

	return c.Strs(key, strings)
}

// Int adds the field key with val as a int to the logger context.
func (c *SugarContext) Int(key string, value int) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Ints adds the field key with val as a []int to the logger context.
func (c *SugarContext) Ints(key string, value []int) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Int8 adds the field key with val as a int8 to the logger context.
func (c *SugarContext) Int8(key string, value int8) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Ints8 adds the field key with val as a []int8 to the logger context.
func (c *SugarContext) Ints8(key string, value []int8) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Int16 adds the field key with val as a int16 to the logger context.
func (c *SugarContext) Int16(key string, value int16) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Ints16 adds the field key with val as a []int16 to the logger context.
func (c *SugarContext) Ints16(key string, value []int16) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Int32 adds the field key with val as a int32 to the logger context.
func (c *SugarContext) Int32(key string, value int32) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Ints32 adds the field key with val as a []int32 to the logger context.
func (c *SugarContext) Ints32(key string, value []int32) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Int64 adds the field key with val as a int64 to the logger context.
func (c *SugarContext) Int64(key string, value int64) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Ints64 adds the field key with val as a []int64 to the logger context.
func (c *SugarContext) Ints64(key string, value []int64) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uint adds the field key with val as a uint to the logger context.
func (c *SugarContext) Uint(key string, value uint) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uints adds the field key with val as a []uint to the logger context.
func (c *SugarContext) Uints(key string, value []uint) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uint8 adds the field key with val as a uint8 to the logger context.
func (c *SugarContext) Uint8(key string, value uint8) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uints8 adds the field key with val as a []uint8 to the logger context.
func (c *SugarContext) Uints8(key string, value []uint8) onelog.LoggerContext {
	// Todo: Better way to do this?
	// Convert []uint8 to []uint64
	uints := make([]uint64, len(value))
	for i, v := range value {
		uints[i] = uint64(v)
	}

	c.addField(key, uints)

	return c
}

// Uint16 adds the field key with val as a uint16 to the logger context.
func (c *SugarContext) Uint16(key string, value uint16) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uints16 adds the field key with val as a []uint16 to the logger context.
func (c *SugarContext) Uints16(key string, value []uint16) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uint32 adds the field key with val as a uint32 to the logger context.
func (c *SugarContext) Uint32(key string, value uint32) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uints32 adds the field key with val as a []uint32 to the logger context.
func (c *SugarContext) Uints32(key string, value []uint32) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uint64 adds the field key with val as a uint64 to the logger context.
func (c *SugarContext) Uint64(key string, value uint64) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Uints64 adds the field key with val as a []uint64 to the logger context.
func (c *SugarContext) Uints64(key string, value []uint64) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Float32 adds the field key with val as a float32 to the logger context.
func (c *SugarContext) Float32(key string, value float32) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Floats32 adds the field key with val as a []float32 to the logger context.
func (c *SugarContext) Floats32(key string, value []float32) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Float64 adds the field key with val as a float64 to the logger context.
func (c *SugarContext) Float64(key string, value float64) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Floats64 adds the field key with val as a []float64 to the logger context.
func (c *SugarContext) Floats64(key string, value []float64) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Bool adds the field key with val as a bool to the logger context.
func (c *SugarContext) Bool(key string, value bool) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Bools adds the field key with val as a []bool to the logger context.
func (c *SugarContext) Bools(key string, value []bool) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Time adds the field key with val as a time.Time to the logger context.
func (c *SugarContext) Time(key string, value time.Time) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Times adds the field key with val as a []time.Time to the logger context.
func (c *SugarContext) Times(key string, value []time.Time) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Dur adds the field key with val as a time.Duration to the logger context.
func (c *SugarContext) Dur(key string, value time.Duration) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Durs adds the field key with val as a []time.Duration to the logger context.
func (c *SugarContext) Durs(key string, value []time.Duration) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// TimeDiff adds the field key with begin and end as a time.Time to the logger context.
func (c *SugarContext) TimeDiff(key string, begin, end time.Time) onelog.LoggerContext {
	diff := end.Sub(begin)
	c.addField(key, diff)

	return c
}

// IPAddr adds the field key with val as a net.IP to the logger context.
func (c *SugarContext) IPAddr(key string, value net.IP) onelog.LoggerContext {
	c.addField(key, value.String())

	return c
}

// IPPrefix adds the field key with val as a net.IPNet to the logger context.
func (c *SugarContext) IPPrefix(key string, value net.IPNet) onelog.LoggerContext {
	c.addField(key, value.String())

	return c
}

// MACAddr adds the field key with val as a net.HardwareAddr to the logger context.
func (c *SugarContext) MACAddr(key string, value net.HardwareAddr) onelog.LoggerContext {
	c.addField(key, value.String())

	return c
}

// AnErr adds the field "error" with err as a string to the logger context.
func (c *SugarContext) AnErr(key string, err error) onelog.LoggerContext {
	c.addField(key, err.Error())

	return c
}

// Err adds the field "error" with err as a string to the logger context.
func (c *SugarContext) Err(err error) onelog.LoggerContext {
	c.AnErr("error", err)

	return c
}

// Errs adds the field key with val as a []error to the logger context.
func (c *SugarContext) Errs(key string, errs []error) onelog.LoggerContext {
	c.addField(key, errs)

	return c
}

// Any adds the field key with val as a any to the logger context.
func (c *SugarContext) Any(key string, value any) onelog.LoggerContext {
	c.addField(key, value)

	return c
}

// Fields adds the field key with val as a Fields to the logger context.
func (c *SugarContext) Fields(fields onelog.Fields) onelog.LoggerContext {
	c.addFields(fields)

	return c
}

// Msg sends the LoggerContext with msg to the logger.
func (c *SugarContext) Msg(msg string) {
	switch c.level {
	case zapcore.DebugLevel:
		c.logger.Debugw(msg, c.fields...)
	case zapcore.InfoLevel:
		c.logger.Infow(msg, c.fields...)
	case zapcore.WarnLevel:
		c.logger.Warnw(msg, c.fields...)
	case zapcore.ErrorLevel:
		c.logger.Errorw(msg, c.fields...)
	case zapcore.FatalLevel:
		c.logger.Fatalw(msg, c.fields...)
	}
	c.fields = make([]any, 0) // reset fields
}

// Msgf sends the LoggerContext with formatted msg to the logger.
func (c *SugarContext) Msgf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	c.Msg(msg)
}

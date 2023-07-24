package nopadapter

import (
	"fmt"
	"net"
	"time"

	"github.com/nikoksr/onelog"
)

// Compile-time check that Adapter and Context implements onelog.Logger and onelog.LoggerContext respectively
var (
	_ onelog.Logger        = (*Adapter)(nil)
	_ onelog.LoggerContext = (*Context)(nil)
)

type (
	Adapter struct{}

	Context struct{}
)

func NewAdapter() onelog.Logger { return &Adapter{} }

func (a *Adapter) Debug() onelog.LoggerContext { return &Context{} }
func (a *Adapter) Info() onelog.LoggerContext  { return &Context{} }
func (a *Adapter) Warn() onelog.LoggerContext  { return &Context{} }
func (a *Adapter) Error() onelog.LoggerContext { return &Context{} }
func (a *Adapter) Fatal() onelog.LoggerContext { return &Context{} }

func (c *Context) Str(_, _ string) onelog.LoggerContext                             { return c }
func (c *Context) Strs(_ string, _ []string) onelog.LoggerContext                   { return c }
func (c *Context) Stringer(_ string, _ fmt.Stringer) onelog.LoggerContext           { return c }
func (c *Context) Stringers(_ string, _ []fmt.Stringer) onelog.LoggerContext        { return c }
func (c *Context) Int(_ string, _ int) onelog.LoggerContext                         { return c }
func (c *Context) Ints(_ string, _ []int) onelog.LoggerContext                      { return c }
func (c *Context) Int8(_ string, _ int8) onelog.LoggerContext                       { return c }
func (c *Context) Ints8(_ string, _ []int8) onelog.LoggerContext                    { return c }
func (c *Context) Int16(_ string, _ int16) onelog.LoggerContext                     { return c }
func (c *Context) Ints16(_ string, _ []int16) onelog.LoggerContext                  { return c }
func (c *Context) Int32(_ string, _ int32) onelog.LoggerContext                     { return c }
func (c *Context) Ints32(_ string, _ []int32) onelog.LoggerContext                  { return c }
func (c *Context) Int64(_ string, _ int64) onelog.LoggerContext                     { return c }
func (c *Context) Ints64(_ string, _ []int64) onelog.LoggerContext                  { return c }
func (c *Context) Uint(_ string, _ uint) onelog.LoggerContext                       { return c }
func (c *Context) Uints(_ string, _ []uint) onelog.LoggerContext                    { return c }
func (c *Context) Uint8(_ string, _ uint8) onelog.LoggerContext                     { return c }
func (c *Context) Uints8(_ string, _ []uint8) onelog.LoggerContext                  { return c }
func (c *Context) Uint16(_ string, _ uint16) onelog.LoggerContext                   { return c }
func (c *Context) Uints16(_ string, _ []uint16) onelog.LoggerContext                { return c }
func (c *Context) Uint32(_ string, _ uint32) onelog.LoggerContext                   { return c }
func (c *Context) Uints32(_ string, _ []uint32) onelog.LoggerContext                { return c }
func (c *Context) Uint64(_ string, _ uint64) onelog.LoggerContext                   { return c }
func (c *Context) Uints64(_ string, _ []uint64) onelog.LoggerContext                { return c }
func (c *Context) Float32(_ string, _ float32) onelog.LoggerContext                 { return c }
func (c *Context) Floats32(_ string, _ []float32) onelog.LoggerContext              { return c }
func (c *Context) Float64(_ string, _ float64) onelog.LoggerContext                 { return c }
func (c *Context) Floats64(_ string, _ []float64) onelog.LoggerContext              { return c }
func (c *Context) Bool(_ string, _ bool) onelog.LoggerContext                       { return c }
func (c *Context) Bools(_ string, _ []bool) onelog.LoggerContext                    { return c }
func (c *Context) Time(_ string, _ time.Time) onelog.LoggerContext                  { return c }
func (c *Context) Times(_ string, _ []time.Time) onelog.LoggerContext               { return c }
func (c *Context) Dur(_ string, _ time.Duration) onelog.LoggerContext               { return c }
func (c *Context) Durs(_ string, _ []time.Duration) onelog.LoggerContext            { return c }
func (c *Context) TimeDiff(_ string, _ time.Time, _ time.Time) onelog.LoggerContext { return c }
func (c *Context) IPAddr(_ string, _ net.IP) onelog.LoggerContext                   { return c }
func (c *Context) IPPrefix(_ string, _ net.IPNet) onelog.LoggerContext              { return c }
func (c *Context) MACAddr(_ string, _ net.HardwareAddr) onelog.LoggerContext        { return c }
func (c *Context) Err(_ error) onelog.LoggerContext                                 { return c }
func (c *Context) Errs(_ string, _ []error) onelog.LoggerContext                    { return c }
func (c *Context) AnErr(_ string, _ error) onelog.LoggerContext                     { return c }
func (c *Context) Any(_ string, _ any) onelog.LoggerContext                         { return c }
func (c *Context) Fields(_ onelog.Fields) onelog.LoggerContext                      { return c }

func (c *Context) Msg(_ string)            {}
func (c *Context) Msgf(_ string, _ ...any) {}

/*
Package nopadapter implements a no-operation (nop) adapter for the Go onelog library.

The nopadapter package is useful when you want to integrate with the onelog library, but do not require any backend
processing for the logs. It can be used as a placeholder for a real logger, or as a way to disable logging in your
application.

Example:

	log := nopadapter.NewAdapter()
	log.Debug().Str("debug", "message").Msg("This debug message will not be logged anywhere")
	log.Info().Str("info", "message").Msg("This info message will not be logged anywhere")

Note that as it is a no-operation (nop) implementation, no actual logging will be performed.
*/
package nopadapter

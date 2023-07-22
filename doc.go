// Package onelog is a general-purpose logging interface heavily inspired by the zerolog API.
// It is designed to provide a user-friendly API for diverse logging requirements. This package
// supports a wide range of data types and log levels, creating flexibility for various use cases.
//
// onelog includes adapters for several commonly used loggers, enabling easy integration
// and compatibility with existing logging methodologies. It reduces the friction associated
// with logging setup and promotes consistency in logging across different parts of a project or across different projects.
//
// Here is a brief example using the zapadapter and slogadapter:

//			import (
//		        "go.uber.org/zap"
//		        "github.com/nikoksr/onelog"
//		        "github.com/nikoksr/onelog/zapadapter"
//		        "github.com/nikoksr/onelog/slogadapter"
//			)
//
//			type superTracker struct {
//		     superEventLogger onelog.Logger
//			}
//
//			func main() {
//		        // Let's use zap's development logger as our superhero event logger
//		        logger, _ := zap.NewDevelopment()
//
//		        tracker := &superTracker{
//		            superEventLogger: zapadapter.NewAdapter(logger),
//		        }
//
//		        // Now let's log a superhero event
//		        tracker.superEventLogger.Info().Msg("Superman spotted in New York!")
//
//		        // Or perhaps we'd rather use slog for logging our superhero sightings
//		        logger := slog.Default()
//		        tracker.superEventLogger = slogadapter.NewAdapter(logger)
//
//	         // And now we can log another sighting
//	         tracker.superEventLogger.Info().Msg("Wonder Woman seen flying over Paris!")
//			}
package onelog

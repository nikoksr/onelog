// Package onelog is a general-purpose logging interface heavily inspired by the zerolog API.
// It is designed to provide a user-friendly API for diverse logging requirements. This package
// supports a wide range of data types and log levels, creating flexibility for various use cases.
//
// onelog includes adapters for several commonly used loggers, enabling easy integration
// and compatibility with existing logging methodologies. It reduces the friction associated
// with logging setup and promotes consistency in logging across different parts of a project or across different projects.
//
// Here is a brief example using the zapadapter and slogadapter:
//
// package main
//
// import (
//
//	"time"
//
//	"go.uber.org/zap"
//	"golang.org/x/exp/slog"
//
//	slogadapter "github.com/nikoksr/onelog/adapter/slog"
//	zapadapter "github.com/nikoksr/onelog/adapter/zap"
//
// )
//
//	func main() {
//	    // Let's use zap's production logger as our superhero event logger
//	    logger, _ := zap.NewProduction()
//
//	    // Use the zapadapter to create a onelog.Logger compatible logger
//	    superheroTracker := zapadapter.NewAdapter(logger)
//
//	    // Start logging
//	    superheroTracker.Debug().Msg("Tracking superheroes...")
//
//	    // Now let's log a superhero event
//	    superheroTracker.Info().
//	        Str("superhero", "Superman").
//	        Str("location", "New York").
//	        Time("time", time.Now()).
//	        Msg("Superman seen flying over New York!")
//
//	    // Or perhaps we'd rather use slog for logging our superhero sightings
//	    superheroTracker = slogadapter.NewAdapter(slog.Default())
//
//	    // And now we can log another sighting
//	    superheroTracker.Info().
//	        Str("superhero", "Batman").
//	        Str("location", "Gotham").
//	        Time("time", time.Now()).
//	        Msg("Batman seen driving through Gotham!")
//
//	    // Output:
//	    // {"level":"info","ts":1690213152.0569847,"caller":"zap/adapter.go:547","msg":"Superman seen flying over New York!","superhero":"Superman","location":"New York","time":1690213152.0569835}
//	    // 2023/07/24 17:39:12 INFO Batman seen driving through Gotham! superhero=Batman location=Gotham time=2023-07-24T17:39:12.057+02:00
//	    //
//	    // Note: The lines above look differently because we switched the logger in between.
//	}
package onelog

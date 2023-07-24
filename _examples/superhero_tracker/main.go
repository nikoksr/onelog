package main

import (
	"go.uber.org/zap"
	"golang.org/x/exp/slog"

	"github.com/nikoksr/onelog"
	slogadapter "github.com/nikoksr/onelog/adapter/slog"
	zapadapter "github.com/nikoksr/onelog/adapter/zap"
)

type superheroTracker struct {
	logger onelog.Logger
}

func main() {
	// Let's use zap's development logger as our superhero event logger
	logger, _ := zap.NewDevelopment()

	heroes := &superheroTracker{
		logger: zapadapter.NewAdapter(logger),
	}

	// Now let's log a superhero event
	heroes.logger.Info().Msg("Superman spotted in New York!")

	// Or perhaps we'd rather use slog for logging our superhero sightings
	heroes.logger = slogadapter.NewAdapter(slog.Default())

	// And now we can log another sighting
	heroes.logger.Info().Msg("Wonder Woman seen flying over Paris!")
}

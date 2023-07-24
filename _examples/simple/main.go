package simple

import (
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/nikoksr/onelog"
	zerologadapter "github.com/nikoksr/onelog/adapter/zerolog"
)

func main() {
	var logger onelog.Logger

	// Create a zerolog logger
	zlogger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Assign the zerolog logger to the onelog logger
	logger = zerologadapter.NewAdapter(&zlogger)

	// Now we can use the onelog logger
	logger.Info().
		Str("foo", "bar").
		Int("n", 42).
		Time("time", time.Now()).
		Msg("Hello world!")
}

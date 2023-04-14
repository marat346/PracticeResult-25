package cities

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "02.01.2006 15:04:05"}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
}

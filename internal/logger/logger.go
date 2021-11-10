package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log = zerolog.New(os.Stdout).With().Logger().Level(zerolog.InfoLevel)

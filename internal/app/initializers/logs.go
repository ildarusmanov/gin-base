package initializers

import (
	"os"
	"strconv"

	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	LogLevelEnv       = "LOG_LEVEL"
	EnableJSONLogsEnv = "ENABLE_JSON_LOGS"
	DefaultLogLevel   = 0
)

func InitializeLogs() error {
	logLevel, err := strconv.Atoi(envy.Get(LogLevelEnv, "0"))
	if err != nil {
		logLevel = DefaultLogLevel
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	// json or human readable output
	if envy.Get(EnableJSONLogsEnv, "false") == "false" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// add filepath+row num to log(app/cmd/serve.go:37)
	log.Logger = log.With().Caller().Logger()

	return nil
}

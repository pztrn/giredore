package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	Logger         zerolog.Logger
	SuperVerbosive bool
)

// Initialize initializes zerolog with proper formatting and log level.
// nolint:forbidigo
func Initialize() {
	// Check environment for logger level.
	// Defaulting to INFO.
	loggerLevel, loggerLevelFound := os.LookupEnv("LOGGER_LEVEL")
	if loggerLevelFound {
		fmt.Println("Setting logger level to:", loggerLevel)

		switch strings.ToUpper(loggerLevel) {
		case "DEBUG":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "INFO":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "WARN":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "ERROR":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "FATAL":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		default:
			fmt.Println("Invalid logger level passed:", loggerLevel)
			fmt.Println("Forcing INFO")
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	} else {
		fmt.Println("Setting logger level to: info")
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// nolint:exhaustruct
	output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.RFC3339}
	output.FormatLevel = func(lvlRaw interface{}) string {
		var formattedLvl string

		if lvl, ok := lvlRaw.(string); ok {
			lvl = strings.ToUpper(lvl)
			switch lvl {
			case "DEBUG":
				formattedLvl = fmt.Sprintf("\x1b[30m%-5s\x1b[0m", lvl)
			case "ERROR":
				formattedLvl = fmt.Sprintf("\x1b[31m%-5s\x1b[0m", lvl)
			case "FATAL":
				formattedLvl = fmt.Sprintf("\x1b[35m%-5s\x1b[0m", lvl)
			case "INFO":
				formattedLvl = fmt.Sprintf("\x1b[32m%-5s\x1b[0m", lvl)
			case "PANIC":
				formattedLvl = fmt.Sprintf("\x1b[36m%-5s\x1b[0m", lvl)
			case "WARN":
				formattedLvl = fmt.Sprintf("\x1b[33m%-5s\x1b[0m", lvl)
			default:
				formattedLvl = lvl
			}
		}

		return fmt.Sprintf("| %s |", formattedLvl)
	}

	Logger = zerolog.New(output).With().Timestamp().Logger()
}

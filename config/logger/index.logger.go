package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-fiber-api/config/config"
	"os"
)

// NewLogger : initialize logger
func NewLogger(cfg *config.Config) zerolog.Logger {
	zerolog.TimeFieldFormat = cfg.Logger.TimeFormat

	if cfg.Logger.Prettier {
		//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

		logFile, _ := os.OpenFile(
			"debug.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)

		multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
		log.Logger = zerolog.New(multi).With().Logger()
	}

	zerolog.SetGlobalLevel(cfg.Logger.Level)

	return log.Hook(PreforkHook{})
}

// PreforkHook : prefer hook for zerologger
type PreforkHook struct{}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	//if fiber.IsChild() {
	//	e.Discard()
	//}
}

type StringerFunc func() string

func (f StringerFunc) String() string {
	return f()
}

package infrastructure

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx/fxevent"
	"os"
	"time"
)

func NewLogger() Logger {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})
	return Logger{
		logger: log.Logger,
	}
}

type Logger struct {
	logger zerolog.Logger
}

func (log *Logger) Log() zerolog.Logger {
	return log.logger
}

func NewFxLogger(logger Logger) fxevent.Logger {
	return &FxLogger{
		Log: logger.Log().With().Str("module", "fx").Logger(),
	}
}

type FxLogger struct {
	Log zerolog.Logger
}

func (fx *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
	case *fxevent.OnStartExecuted:
	case *fxevent.OnStopExecuting:
	case *fxevent.OnStopExecuted:
	case *fxevent.Supplied:
		if e.Err != nil {
			fx.Log.Error().Err(e.Err).Msgf("Cannot supply %s", e.TypeName)
		} else {
			fx.Log.Debug().Msgf("Supplied %s", e.TypeName)
		}
	case *fxevent.Provided:
		if e.Err != nil {
			fx.Log.Error().
				Err(e.Err).
				Msgf("Cannot provide %v with %s", e.OutputTypeNames, e.ConstructorName)
		} else {
			fx.Log.Debug().Msgf("Provided %v <= %s", e.OutputTypeNames, e.ConstructorName)
		}
	case *fxevent.Invoking:
		fx.Log.Debug().Msgf("Invoking %s", e.FunctionName)
	case *fxevent.Invoked:
		if e.Err != nil {
			fx.Log.Error().Err(e.Err).Msgf("%s exits with error", e.FunctionName)
		}
	case *fxevent.Stopping:
		fx.Log.Info().Msgf("Signal %s received. Stopping application...", e.Signal)
	case *fxevent.Stopped:
		if e.Err != nil {
			fx.Log.Info().Err(e.Err).Msgf("Failed to graceful stop application")
		}
	case *fxevent.RollingBack:
	case *fxevent.RolledBack:
	case *fxevent.Started:
		if e.Err != nil {
			fx.Log.Error().Err(e.Err).Msgf("Cannot start application")
		} else {
			fx.Log.Info().Msg("Application started")
		}
	case *fxevent.LoggerInitialized:
	}
}

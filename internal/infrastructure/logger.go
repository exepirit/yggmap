package infrastructure

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx/fxevent"
)

func NewLogger() Logger {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})
	return Logger{logger: log.Logger}
}

type Logger struct {
	logger zerolog.Logger
	fx     *FxLogger
	gin    *GinLogger
}

func (log *Logger) Log() zerolog.Logger {
	return log.logger
}

func (log *Logger) GetFxLogger() fxevent.Logger {
	if log.fx == nil {
		log.fx = &FxLogger{
			Log: log.Log().With().Str("module", "fx").Logger(),
		}
	}
	return log.fx
}

func (log *Logger) GetGinLogger() gin.HandlerFunc {
	if log.gin == nil {
		log.gin = &GinLogger{
			Log:       log.Log().With().Str("module", "server").Logger(),
			SkipPaths: nil,
		}
	}
	return log.gin.HandleRequest
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

type GinLogger struct {
	Log       zerolog.Logger
	SkipPaths map[string]bool
}

func (log *GinLogger) HandleRequest(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery

	start := time.Now()

	// Process request
	ctx.Next()

	latency := time.Now().Sub(start)

	// Log only when path is not being skipped
	if _, ok := log.SkipPaths[path]; !ok {
		if raw != "" {
			path = path + "?" + raw
		}

		respSize := "-"
		if size := ctx.Writer.Size(); size > 0 {
			respSize = strconv.Itoa(size)
		}

		errors := ctx.Errors.ByType(gin.ErrorTypePrivate)
		for _, err := range errors {
			log.Log.Error().
				Err(err).
				Msgf("Error occurred while handling request %s %s", ctx.Request.Method, path)
		}

		log.Log.Info().Msgf(
			"%s %s %d %s %s", ctx.Request.Method, path, ctx.Writer.Status(),
			respSize, latency.String())
	}
}

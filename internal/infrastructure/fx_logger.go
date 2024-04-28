package infrastructure

import (
	"go.uber.org/fx/fxevent"
	"log/slog"
)

// FxLogger is used for logging events related to the Fx application lifecycle.
type FxLogger struct {
	Logger *slog.Logger
}

// LogEvent handles different types of fxevent.Event and logs them based on their type.
func (fx *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
	case *fxevent.OnStartExecuted:
	case *fxevent.OnStopExecuting:
	case *fxevent.OnStopExecuted:
	case *fxevent.Supplied:
		if e.Err != nil {
			fx.Logger.Error("Cannot supply dependency", "error", e.Err, "type", e.TypeName)
		} else {
			fx.Logger.Debug("Dependency supplied", "type", e.TypeName)
		}
	case *fxevent.Provided:
		if e.Err != nil {
			fx.Logger.Error(
				"Cannot provide dependency",
				"error", e.Err,
				"from", e.ConstructorName,
				"to", e.OutputTypeNames,
			)
		} else {
			fx.Logger.Debug(
				"Dependency provided",
				"from", e.ConstructorName,
				"to", e.OutputTypeNames,
			)
		}
	case *fxevent.Invoking:
		fx.Logger.Debug(
			"Invoking function",
			"func", e.FunctionName,
		)
	case *fxevent.Invoked:
		if e.Err != nil {
			fx.Logger.Error(
				"Function returns error",
				"error", e.Err,
				"func", e.FunctionName,
			)
		}
	case *fxevent.Stopping:
		fx.Logger.Info("Signal received. Stopping application", "signal", e.Signal)
	case *fxevent.Stopped:
		if e.Err != nil {
			fx.Logger.Error("Failed to graceful stop the application", "error", e.Err)
		}
	case *fxevent.RollingBack:
	case *fxevent.RolledBack:
	case *fxevent.Started:
		if e.Err != nil {
			fx.Logger.Error("Cannot start application", "error", e.Err)
		} else {
			fx.Logger.Info("Application started")
		}
	case *fxevent.LoggerInitialized:
	}
}

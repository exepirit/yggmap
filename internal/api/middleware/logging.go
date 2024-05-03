package middleware

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"log/slog"
	"time"
)

// Logging is a middleware that log every request.
func Logging(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	operation := graphql.GetOperationContext(ctx)
	responseHandler := next(ctx)
	return func(ctx context.Context) *graphql.Response {
		start := time.Now()
		r := responseHandler(ctx)

		attrs := []any{
			slog.String("latency", time.Now().Sub(start).String()),
			slog.String("operation", operation.OperationName),
			slog.Int("responseSize", len(r.Data)),
		}
		if r.Errors != nil {
			attrs = append(attrs, slog.String("error", r.Errors.Error()))
		}

		slog.InfoContext(ctx, "GraphQL query handled", attrs...)
		return r
	}
}

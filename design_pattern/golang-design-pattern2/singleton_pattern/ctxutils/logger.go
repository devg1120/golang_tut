package ctxutils

import "golang.org/x/net/context"

type contextKey string

const traceIDKey = contextKey("trace_id")

func NewContextWithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceIDFromContext(ctx context.Context) (string, bool) {
	traceID, ok := ctx.Value(traceIDKey).(string)
	return traceID, ok
}

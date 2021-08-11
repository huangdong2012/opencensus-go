package trace

import (
	"context"
	"go.opencensus.io/trace/propagation"
)

func extractSpanContext(ctx context.Context) *SpanContext {
	if val := ctx.Value(contextKey{}); val != nil {
		if sp, ok := val.(*Span); ok {
			sc := sp.SpanContext()
			return &sc
		}
	}
	return nil
}

func Inject(ctx context.Context) string {
	spCtx := extractSpanContext(ctx)
	if spCtx == nil {
		return ""
	}

	return string(propagation.Binary(*spCtx))
}

func Extract(str string) *SpanContext {
	if len(str) == 0 {
		return nil
	}
	sp, ok := propagation.FromBinary([]byte(str))
	if !ok {
		return nil
	}
	return &sp
}

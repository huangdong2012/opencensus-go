package propagation

import (
	"context"
	"go.opencensus.io/trace"
)

func Inject(ctx context.Context) string {
	spCtx := trace.ExtractSpanContext(ctx)
	if spCtx == nil {
		return ""
	}

	return string(Binary(*spCtx))
}

func Extract(str string) *trace.SpanContext {
	if len(str) == 0 {
		return nil
	}
	sp, ok := FromBinary([]byte(str))
	if !ok {
		return nil
	}
	return &sp
}

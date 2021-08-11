package propagation

import (
	"context"
	"encoding/hex"
	"go.opencensus.io/trace"
)

func Inject(ctx context.Context) string {
	spCtx := trace.ExtractSpanContext(ctx)
	if spCtx == nil {
		return ""
	}

	data := Binary(*spCtx)
	return hex.EncodeToString(data)
}

func Extract(str string) *trace.SpanContext {
	if len(str) == 0 {
		return nil
	}
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil
	}
	sp, ok := FromBinary(data)
	if !ok {
		return nil
	}
	return &sp
}

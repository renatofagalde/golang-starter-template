package constants

import (
	"context"
)

type Headers struct {
	Stage   string `json:"stage"`
	Journey string `form:"journey" binding:"required"`
	TraceID string `form:"trace_id" binding:"required"`
}

var H = Headers{
	Stage:   "Stage",
	Journey: "X-Request-Journey",
	TraceID: "X-Request-ID",
}

func (h *Headers) GetJourney(ctx context.Context) string {
	if journeyVal := ctx.Value(H.Journey); journeyVal != nil {
		if jStr, ok := journeyVal.(string); ok {
			return jStr
		}
	}
	return "unknown"
}

func (h *Headers) GetTraceID(ctx context.Context) string {
	if traceIDVal := ctx.Value(H.TraceID); traceIDVal != nil {
		if tStr, ok := traceIDVal.(string); ok {
			return tStr
		}
	}
	return "unknown"
}

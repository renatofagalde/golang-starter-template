package constants

type Headers struct {
	TraceID string `form:"trace_id" binding:"required"`
	Journey string `form:"journey" binding:"required"`
}

var H = Headers{
	TraceID: "X-Request-ID",
	Journey: "X-Request-Journey",
}

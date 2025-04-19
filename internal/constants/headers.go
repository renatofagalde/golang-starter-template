package constants

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

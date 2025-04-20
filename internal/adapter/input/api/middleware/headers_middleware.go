package middleware

import (
	"bootstrap/internal/constants"
	"context"
	"github.com/gin-gonic/gin"
)

// HeadersMiddleware extracts specific headers and adds them to the request context
func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		journeyHeader := c.GetHeader(constants.H.Journey)
		traceIDHeader := c.GetHeader(constants.H.TraceID)

		requestContext := c.Request.Context()
		if journeyHeader != "" {
			requestContext = context.WithValue(requestContext, constants.H.Journey, journeyHeader)
		}
		if traceIDHeader != "" {
			requestContext = context.WithValue(requestContext, constants.H.TraceID, traceIDHeader)
		}

		c.Request = c.Request.WithContext(requestContext)
		c.Next()
	}
}

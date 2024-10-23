package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware(c *gin.Context) {
	requestId := c.GetHeader("X-Request-Id")
	if requestId == "" {
		requestId = uuid.New().String()
	}

	c.Set("request_id", requestId)
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "request_id", requestId))
	c.Writer.Header().Set("X-Request-Id", requestId)

	c.Next()
}

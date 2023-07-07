package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/global"
	"go.uber.org/zap"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		global.Logger.Info(
			zap.String("Method", c.Request.Method),
			zap.String("Path", c.Request.URL.Path),
			zap.Int("Status", c.Writer.Status()),
			zap.Duration("Cost", endTime.Sub(startTime)),
		)
	}
}

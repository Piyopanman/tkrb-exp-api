package middleware

import (
	"tkrb-exp-api/src/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//AccessLog アクセスログを出力
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		if status < 400 {
			logging.Logger.Info("SuccessResponse",
				zap.Int("status", status),
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("host", c.Request.Host),
				zap.String("remoteAddress", c.Request.RemoteAddr),
			)
		} else {
			logging.Logger.Error("ErrorResponse",
				zap.Int("status", status),
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("host", c.Request.Host),
				zap.String("remoteAddress", c.Request.RemoteAddr),
			)
		}
	}
}

package middlewares

import (
	"github.com/SyarifKA/himbara/logs/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogWithMessage(message string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.RotateLogIfNeeded()
		logger.Logger.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"path":    c.FullPath(),
			"ip":      c.ClientIP(),
			"headers": c.Request.Header,
		}).Info(message)

		handler(c)
	}
}

package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Context handles all errors occured in TodoController.
func HandleErrors(c *gin.Context) {
	c.Next()
	logger, _ := zap.NewDevelopment()
	requestID := c.Request.Header.Get("X-Request-Id")
	publicError := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if publicError != nil && publicError.Meta != nil {
		statusCode := publicError.Meta.(int)
		errMsg := publicError.Error()
		if er, ok := publicError.Err.(interface{ StackTrace() errors.StackTrace }); ok {
			logger.Warn(errMsg, zap.String("requestID", requestID), zap.String("stackTrace", fmt.Sprintf("%+v\n\n", er.StackTrace())), zap.Time("now", time.Now()))
			c.JSON(statusCode, gin.H{
				"message": errMsg,
			})
		}
	}
	privateError := c.Errors.ByType(gin.ErrorTypePrivate).Last()
	if privateError != nil && privateError.Meta != nil {
		statusCode := privateError.Meta.(int)
		errMsg := privateError.Error()
		if er, ok := privateError.Err.(interface{ StackTrace() errors.StackTrace }); ok {
			logger.Error(errMsg, zap.String("requestID", requestID), zap.String("stackTrace", fmt.Sprintf("%+v\n\n", er.StackTrace())), zap.Time("now", time.Now()))
			c.JSON(statusCode, gin.H{
				"message": "システムエラーが発生しました",
			})
		}
	}

}

package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandleErrors(c *gin.Context) {
	c.Next()
	errorToPrint := c.Errors.ByType(gin.ErrorTypeAny).Last()
	if errorToPrint != nil {
		statusCode := errorToPrint.Meta.(int)
		errMsg := errorToPrint.Error()
		c.JSON(statusCode, gin.H{
			"message": errMsg,
		})
	}
}

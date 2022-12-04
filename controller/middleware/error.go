package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandleErrors(c *gin.Context) {
	c.Next()
	errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if errorToPrint != nil && errorToPrint.Meta != nil {
		statusCode := errorToPrint.Meta.(int)
		errMsg := errorToPrint.Error()
		c.JSON(statusCode, gin.H{
			"message": errMsg,
		})
	}
}

package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Context handles all errors occured in TodoController.
func HandleErrors(c *gin.Context) {
	c.Next()
	publicError := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if publicError != nil && publicError.Meta != nil {
		statusCode := publicError.Meta.(int)
		errMsg := publicError.Error()
		c.JSON(statusCode, gin.H{
			"message": errMsg,
		})
	}
	privateError := c.Errors.ByType(gin.ErrorTypePrivate).Last()
	if privateError != nil && privateError.Meta != nil {
		statusCode := privateError.Meta.(int)
		errMsg := privateError.Error()
		log.Print(errMsg)
		c.JSON(statusCode, gin.H{
			"message": "システムエラーが発生しました",
		})
	}

}

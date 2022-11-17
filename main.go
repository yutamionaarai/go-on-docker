package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		fmt.Fprintf(w, "Hello, World")
	})

	r.GET("todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		log.Printf(id)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run()
}

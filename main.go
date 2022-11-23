package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        int64
	Name      string
	Todos     []Todo
	CreatedAT time.Time
	UpdatedAt time.Time
}

type Todo struct {
	ID             int64
	Title          string
	Description    string
	Status         string
	Priority       int64
	ExpirationDate time.Time
	UserID         int64
	CreatedAT      time.Time
	UpdatedAt      time.Time
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("failed to load ENV", err)
	}
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to open DB", err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		fmt.Fprintf(w, "Hello, World")
	})

	r.GET("todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		var todo Todo
		err := db.First(&todo, id).Error
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(todo)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run()
}

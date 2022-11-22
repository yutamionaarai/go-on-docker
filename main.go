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
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserRefer"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todo struct {
	ID             int64     `json:"id" gorm:"primaryKey"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	Priority       int64     `json:"priority"`
	ExpirationDate time.Time `json:"expiration_date"`
	UserID         int64     `json:"user_id"`
	CreatedAT      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
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
		var todos []map[string]interface{}
		db.Table("todos").Find(&todos, "id = ?", id)
		fmt.Print(todos)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run()
}

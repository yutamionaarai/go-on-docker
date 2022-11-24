package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Todos     []Todo    `json:"todos"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todo struct {
	ID             int64     `json:"id"`
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

	// todoリストを全件取得
	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		err := db.Find(&todos).Error
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": todos,
		})
	})
	// 該当のIDのtodoリストを取得
	r.GET("todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo
		err := db.First(&todo, id).Error
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	})
	// todoリストの作成
	r.POST("todos/", func(c *gin.Context) {
		var user User
		var todoRequest Todo
		err := c.BindJSON(&todoRequest)
		if err != nil {
			fmt.Print(err)
		}
		err = db.First(&user, todoRequest.UserID).Error
		if err != nil {
			fmt.Print(err)
		}

		todo := Todo{UserID: user.ID, Title: todoRequest.Title, Description: todoRequest.Status, Status: todoRequest.Status, Priority: todoRequest.Priority, ExpirationDate: todoRequest.ExpirationDate}
		err = db.Create(&todo).Error
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"id": todo.ID,
		},
		})
	})
	// 該当のIDのtodoリストの更新
	r.PATCH("todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todoRequest Todo
		err := c.BindJSON(&todoRequest)
		if err != nil {
			fmt.Print(err)
		}
		var todo Todo
		err = db.First(&todo, id).Error
		if err != nil {
			fmt.Print(err)
		}
		todo.Title = todoRequest.Title
		todo.Description = todoRequest.Description
		todo.Status = todoRequest.Status
		todo.Priority = todoRequest.Priority
		todo.ExpirationDate = todoRequest.ExpirationDate
		err = db.Save(&todo).Error
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"id": todo.ID,
		},
		})
	})
	r.DELETE("todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo
		err := db.Delete(&todo, id)
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.Run()
}

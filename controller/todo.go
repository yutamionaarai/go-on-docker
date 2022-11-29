package controller

import (
	"app/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		fmt.Fprintf(w, "Hello, World")
	}
}

// todoリストを全件取得
func FindTodosController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos *[]model.Todo
		todos, err := service.FindTodos(db, todos)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": todos,
		})
	}
}

func FindTodoController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var todo *model.Todo
		todo, err := service.FirstTodo(db, todo, id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": todo,
		})
	}
}

func CreateTodoController(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		var todoRequest model.Todo
		err := c.BindJSON(&todoRequest)
		if err != nil {
			fmt.Print(err)
		}
		if err := db.First(&user, todoRequest.UserID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			return
		}
		todo := model.Todo{UserID: user.ID, Title: todoRequest.Title, Description: todoRequest.Description,
			Status: todoRequest.Status, Priority: todoRequest.Priority, ExpirationDate: todoRequest.ExpirationDate}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"id": todo.ID,
		},
		})
	}
}

func D(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var todoRequest model.Todo
		err := c.BindJSON(&todoRequest)
		if err != nil {
			fmt.Print(err)
		}
		var todo model.Todo
		if err = db.First(&todo, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			return
		}
		updateTodo := model.Todo{Title: todoRequest.Title, Description: todo.Description,
			Status: todo.Status, Priority: todoRequest.Priority, ExpirationDate: todo.ExpirationDate}
		if err := db.Model(&todo).Updates(updateTodo).Error; err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"id": todo.ID,
		},
		})

	}
}

func E(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var todo model.Todo
		err := db.Delete(&todo, id)
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusNoContent, gin.H{})

	}
}

package controller

import (
	"app/model"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{
		db: db,
	}
}

func (t *TodoController) HelloController(c *gin.Context) {
	w := c.Writer
	fmt.Fprintf(w, "Hello, World")
}

// todoリストを全件取得
func (t *TodoController) FindTodosController(c *gin.Context) {
	var todos []model.Todo
	err := t.db.Find(&todos).Error
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	c.JSON(200, gin.H{
		"data": todos,
	})
}

// 該当のIDのtodoリストを取得
func (t *TodoController) FindTodoController(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	c.JSON(200, gin.H{
		"data": todo,
	})
}

// todoリストの作成
func (t *TodoController) CreateTodoController(c *gin.Context) {
	var user model.User
	var todoRequest model.Todo
	err := c.ShouldBindJSON(&todoRequest)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	if err := todoRequest.TodoValidate(c); err != nil {
		return
	}

	if err := t.db.First(&user, todoRequest.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	todo := model.Todo{UserID: user.ID, Title: todoRequest.Title, Description: todoRequest.Description,
		Status: todoRequest.Status, Priority: todoRequest.Priority, ExpirationDate: todoRequest.ExpirationDate}

	if err := t.db.Create(&todo).Error; err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	c.JSON(200, gin.H{"data": gin.H{
		"id": todo.ID,
	},
	})
}

// 該当のIDのtodoリストの更新
func (t *TodoController) UpdateTodoController(c *gin.Context) {
	id := c.Param("id")
	var todoRequest model.Todo
	err := c.ShouldBindJSON(&todoRequest)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	if err := todoRequest.TodoValidate(c); err != nil {
		return
	}

	var todo model.Todo
	if err := t.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	updateTodo := model.Todo{Title: todoRequest.Title, Description: todo.Description,
		Status: todo.Status, Priority: todoRequest.Priority, ExpirationDate: todo.ExpirationDate}
	if err := t.db.Model(&todo).Updates(updateTodo).Error; err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	c.JSON(200, gin.H{"data": gin.H{
		"id": todo.ID,
	},
	})

}

// 該当のIDのtodoリストの削除
func (t *TodoController) DeleteTodoController(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := t.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	err := t.db.Delete(&todo, id).Error
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}
	c.JSON(204, gin.H{})
}

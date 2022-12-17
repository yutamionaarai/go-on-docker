package controller

import (
	"app/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

// TodoController returns Hello World.
func (t *TodoController) HelloController(c *gin.Context) {
	w := c.Writer
	fmt.Fprintf(w, "Hello, World")
}

// TodoController fetchs All Todos from DB and returns it.
func (t *TodoController) FindTodosController(c *gin.Context) {
	var todos []*model.Todo
	err := t.db.Find(&todos).Error
	if err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	todosResponse := model.FindTodosResponse{Todos: todos}
	c.JSON(200, gin.H{
		"data": todosResponse,
	})
}

// TodoController fetchs One Todo from DB and returns it.
func (t *TodoController) FindTodoController(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := errors.New(err.Error())
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	todoResponse := model.FindTodoResponse{Todo: &todo}

	c.JSON(200, gin.H{
		"data": todoResponse,
	})
}

// A TodoController Creates One Todo on DB.
func (t *TodoController) CreateTodoController(c *gin.Context) {
	var user model.User
	var todoRequest model.TodoRequest
	err := c.ShouldBindJSON(&todoRequest)
	if err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	if err := todoRequest.TodoValidate(c); err != nil {
		return
	}
	if err := t.db.First(&user, todoRequest.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := errors.New(err.Error())
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	todo := model.Todo{UserID: todoRequest.UserID, Title: todoRequest.Title, Description: todoRequest.Description,
		Status: todoRequest.Status, Priority: todoRequest.Priority, ExpirationDate: todoRequest.ExpirationDate}

	if err := t.db.Omit("created_at", "updated_at").Create(&todo).Error; err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	createResponse := model.CreateTodoResponse{ID: todo.ID}
	c.JSON(200, gin.H{
		"data": createResponse,
	})
}

// A TodoController Updates One Todo on DB.
func (t *TodoController) UpdateTodoController(c *gin.Context) {
	id := c.Param("id")
	var todoRequest model.TodoRequest
	err := c.ShouldBindJSON(&todoRequest)
	if err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	if err := todoRequest.TodoValidate(c); err != nil {
		return
	}
	var fetchedTodo model.Todo
	if err := t.db.First(&fetchedTodo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := errors.New(err.Error())
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	updateTodo := model.Todo{Title: todoRequest.Title, Description: todoRequest.Description,
		Status: todoRequest.Status, Priority: todoRequest.Priority, ExpirationDate: todoRequest.ExpirationDate}
	if err := t.db.Omit("created_at", "updated_at").Model(&fetchedTodo).Updates(updateTodo).Error; err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	updateResponse := model.UpdateTodoResponse{ID: fetchedTodo.ID}
	c.JSON(200, gin.H{
		"data": updateResponse,
	})
}

// A TodoController Deletes One Todo.
func (t *TodoController) DeleteTodoController(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := t.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := errors.New(err.Error())
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		}
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	err := t.db.Delete(&todo, id).Error
	if err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	deleteResponse := model.DeleteTodoResponse{}
	c.JSON(200, gin.H{
		"data": deleteResponse,
	})
}

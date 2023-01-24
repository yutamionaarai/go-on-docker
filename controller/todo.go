package controller

import (
	"app/model"
	"app/repository"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type TodoController struct {
	repo repository.TodoRepository
}

func NewTodoController(repo repository.TodoRepository) *TodoController {
	return &TodoController{
		repo: repo,
	}
}

// TodoController returns Hello World.
func (t *TodoController) HelloController(c *gin.Context) {
	w := c.Writer
	fmt.Fprintf(w, "Hello, World")
}

// TodoController fetchs All Todos from DB and returns it.
func (t *TodoController) FindTodosController(c *gin.Context) {
	todosResponse, err := t.repo.FindTodos()
	if err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return
	}
	c.JSON(200, gin.H{
		"data": todosResponse,
	})
}

// TodoController fetchs One Todo from DB and returns it.
func (t *TodoController) FindTodoController(c *gin.Context) {
	idInt64, err := convertStringIDToInt64(c, "id")
	if err != nil {
		return
	}
	todoResponse, err := t.repo.FindTodo(idInt64)
	if err := t.handleErrorResponse(c, err); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data": todoResponse,
	})
}

// A TodoController Creates One Todo on DB.
func (t *TodoController) CreateTodoController(c *gin.Context) {
	var todoRequest *model.TodoRequest
	err := c.ShouldBindJSON(&todoRequest)
	if err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	if err := todoRequest.TodoValidate(c); err != nil {
		return
	}
	createTodoResponse, err := t.repo.CreateTodo(todoRequest)
	if err := t.handleErrorResponse(c, err); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data": createTodoResponse,
	})
}

// A TodoController Updates One Todo on DB.
func (t *TodoController) UpdateTodoController(c *gin.Context) {
	idInt64, err := convertStringIDToInt64(c, "id")
	if err != nil {
		return
	}
	var todoRequest *model.TodoRequest
	if err := c.ShouldBindJSON(&todoRequest); err != nil {
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	if err := todoRequest.TodoValidate(c); err != nil {
		return
	}
	updateTodoResponse, err := t.repo.UpdateTodo(todoRequest, idInt64)
	if err := t.handleErrorResponse(c, err); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data": updateTodoResponse,
	})
}

// A TodoController Deletes One Todo.
func (t *TodoController) DeleteTodoController(c *gin.Context) {
	idInt64, err := convertStringIDToInt64(c, "id")
	if err != nil {
		return
	}
	deleteTodoResponse, err := t.repo.DeleteTodo(idInt64)
	if err := t.handleErrorResponse(c, err); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data": deleteTodoResponse,
	})
}

// handleError is a function to handle error response.
func (t *TodoController) handleErrorResponse(c *gin.Context, err error) error {
	if err != nil {
		if err.Error() == "record not found" {
			err := errors.New(err.Error())
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return err
		}
		err := errors.New(err.Error())
		c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
		return err
	}
	return nil
}

// convertStringIDToInt64 is a function to convert ID of string received outside to Int64.
func convertStringIDToInt64(c *gin.Context, param string) (int64, error) {
	id := c.Param(param)
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
		return 0, err
	}
	return idInt64, nil
}

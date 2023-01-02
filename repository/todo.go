package repository

import (
	"app/model"
	"errors"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindTodo(id int64) (model.FindTodoResponse, error)
	FindTodos() (model.FindTodosResponse, error)
	CreateTodo(t *model.TodoRequest) (model.CreateTodoResponse, error)
	UpdateTodo(t *model.TodoRequest, id int64) (model.UpdateTodoResponse, error)
	DeleteTodo(t *model.Todo, id int64) (model.DeleteTodoResponse, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) FindTodo(id int64) (model.FindTodoResponse, error) {
	var todo *model.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		handleNotFoundError(err)
		return model.FindTodoResponse{}, err
	}
	return model.FindTodoResponse{Todo: todo}, nil
}

func (r *todoRepository) FindTodos() (model.FindTodosResponse, error) {
	var todos []*model.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return model.FindTodosResponse{}, err
	}
	return model.FindTodosResponse{Todos: todos}, nil
}

func (r *todoRepository) CreateTodo(t *model.TodoRequest) (model.CreateTodoResponse, error) {
	var user *model.User
	if err := r.db.First(&user, t.UserID).Error; err != nil {
		return model.CreateTodoResponse{}, err
	}
	todo := model.Todo{
		UserID:         t.UserID,
		Title:          t.Title,
		Description:    t.Description,
		Status:         t.Status,
		Priority:       t.Priority,
		ExpirationDate: t.ExpirationDate,
	}

	if err := r.db.Omit("created_at", "updated_at").Create(&todo).Error; err != nil {
		return model.CreateTodoResponse{}, err
	}
	return model.CreateTodoResponse{ID: todo.ID}, nil
}

func (r *todoRepository) UpdateTodo(t *model.TodoRequest, id int64) (model.UpdateTodoResponse, error) {
	var todo *model.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		handleNotFoundError(err)
		return model.UpdateTodoResponse{}, err
	}
	updateTodo := model.Todo{
		Title:          t.Title,
		Description:    t.Description,
		Status:         t.Status,
		Priority:       t.Priority,
		ExpirationDate: t.ExpirationDate,
	}
	if err := r.db.Omit("created_at", "updated_at").Model(&todo).Updates(updateTodo).Error; err != nil {
		return model.UpdateTodoResponse{}, err
	}
	return model.UpdateTodoResponse{ID: id}, nil

}

func (r *todoRepository) DeleteTodo(t *model.Todo, id int64) (model.DeleteTodoResponse, error) {
	err := r.db.First(&t, id).Error
	if err != nil {
		handleNotFoundError(err)
		return model.DeleteTodoResponse{}, err
	}

	if err := r.db.Delete(&t, id).Error; err != nil {
		return model.DeleteTodoResponse{}, err
	}
	return model.DeleteTodoResponse{}, nil
}

// handleNotFoundError is a function to name the error of gorm.ErrRecordNotFound for getting gormDB  out of Controller layer
func handleNotFoundError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	return err
}

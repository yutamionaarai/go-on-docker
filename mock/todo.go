package mock

import (
	"app/model"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (m *TodoRepositoryMock) FindTodo(id int64) (model.FindTodoResponse, error) {
	if id < 0 {
        return model.FindTodoResponse{},fmt.Errorf("id is negative")
    }
	result := m.Called(id)
	return result.Get(0).(model.FindTodoResponse), result.Error(1)
}

func (m *TodoRepositoryMock) FindTodos() (model.FindTodosResponse, error) {
	result := m.Called()
	return result.Get(0).(model.FindTodosResponse), result.Error(1)
}

func (m *TodoRepositoryMock) CreateTodo(t *model.TodoRequest) (model.CreateTodoResponse, error) {
	if err := t.Validate(); err != nil {
        return model.CreateTodoResponse{}, fmt.Errorf("titleは必須項目です。")
    }
	result := m.Called(t)
	return result.Get(0).(model.CreateTodoResponse), result.Error(1)
}

func (m *TodoRepositoryMock) UpdateTodo(t *model.TodoRequest, id int64) (model.UpdateTodoResponse, error) {
	if err := t.Validate(); err != nil {
		fmt.Print(err)
        return model.UpdateTodoResponse{}, err
    }
    if id < 0{
        return model.UpdateTodoResponse{},fmt.Errorf("id is negative")
    }
	result := m.Called(t, id)
	return result.Get(0).(model.UpdateTodoResponse), result.Error(1)
}

func (m *TodoRepositoryMock) DeleteTodo(id int64) (model.DeleteTodoResponse, error) {
    if id < 0{
        return model.DeleteTodoResponse{},fmt.Errorf("id is negative")
    }
	result := m.Called(id)
	return result.Get(0).(model.DeleteTodoResponse), result.Error(1)
}

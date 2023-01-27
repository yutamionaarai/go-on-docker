package mock

import (
	"app/model"

	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (m *TodoRepositoryMock) FindTodo(id int64) (model.FindTodoResponse, error) {
	result := m.Called(id)
	return result.Get(0).(model.FindTodoResponse), result.Error(1)
}

func (m *TodoRepositoryMock) FindTodos() (model.FindTodosResponse, error) {
	result := m.Called()
	return result.Get(0).(model.FindTodosResponse), result.Error(1)
}

func (m *TodoRepositoryMock) CreateTodo(t *model.TodoRequest) (model.CreateTodoResponse, error) {
	result := m.Called(t)
	return result.Get(0).(model.CreateTodoResponse), result.Error(1)
}

func (m *TodoRepositoryMock) UpdateTodo(t *model.TodoRequest, id int64) (model.UpdateTodoResponse, error) {
	result := m.Called(id, t)
	return result.Get(0).(model.UpdateTodoResponse), result.Error(1)
}

func (m *TodoRepositoryMock) DeleteTodo(id int64) (model.DeleteTodoResponse, error) {
	result := m.Called(id)
	return result.Get(0).(model.DeleteTodoResponse), result.Error(1)
}

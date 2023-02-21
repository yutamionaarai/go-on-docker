package testdata

import (
	"app/model"
	"time"
)

var Todo = &model.Todo{
	ID:             1,
	Title:          "title",
	Description:    "Description",
	Status:         "Status",
	Priority:       1,
	UserID:         1,
	ExpirationDate: &time.Time{},
	CreatedAt:      time.Now(),
	UpdatedAt:      time.Now(),
}

var Todos = []*model.Todo{
	&model.Todo{
		ID:             1,
		Title:          "title",
		Description:    "Description",
		Status:         "Status",
		Priority:       1,
		UserID:         1,
		ExpirationDate: &time.Time{},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	&model.Todo{
		ID:             2,
		Title:          "title2",
		Description:    "Description2",
		Status:         "Status2",
		Priority:       2,
		UserID:         1,
		ExpirationDate: &time.Time{},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
}

var TodoRequest = &model.TodoRequest{
	Title:          "title",
	Description:    "Description",
	Status:         "Status",
	Priority:       1,
	ExpirationDate: &time.Time{},
	UserID:         1,
}

// カラムが不足
var InvalidTodoRequest = &model.TodoRequest{Title: "Title"}

// FindTodo関数の正常のレスポンス
var expirationDate = time.Date(2023, 12, 25, 3, 0, 0, 0, time.UTC)
var CreateTodoRequest = &model.TodoRequest{
	UserID:         1,
	Title:          "study",
	Description:    "Go",
	ExpirationDate: &expirationDate,
	Status:         "pending",
	Priority:       1,
}

// ユーザーIDが存在しない
var InvalidCreateTodoRequest = &model.TodoRequest{
	UserID:         99,
	Title:          "study",
	Description:    "Go",
	ExpirationDate: &expirationDate,
	Status:         "pending",
	Priority:       1,
}

var FindTodoResponse = model.FindTodoResponse{
	Todo: &model.Todo{
		ID:             1,
		UserID:         1,
		Title:          "study",
		Description:    "Go",
		ExpirationDate: &expirationDate,
		Status:         "pending",
		Priority:       1,
	},
}

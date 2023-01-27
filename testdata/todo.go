package testdata

import (
	"app/model"
	"time"
)


var Todo = &model.Todo{
	ID:          1,
	Title:       "title",
	Description: "Description",
	Status:      "Status",
	Priority:    1,
	UserID:      1,
	ExpirationDate: &time.Time{},
	CreatedAT: time.Now(),
	UpdatedAt: time.Now(),
}

var Todos = []*model.Todo{
	&model.Todo {
		ID: 1,
		Title: "title",
		Description: "Description",
		Status: "Status",
		Priority: 1,
		UserID: 1,
		ExpirationDate: &time.Time{},
		CreatedAT: time.Now(),
		UpdatedAt: time.Now(),
	},
	&model.Todo{
		ID: 2,
		Title: "title2",
		Description: "Description2",
		Status: "Status2",
		Priority: 2,
		UserID: 1,
		ExpirationDate: &time.Time{},
		CreatedAT: time.Now(),
		UpdatedAt: time.Now(),
	},
}

var TodoRequest = &model.TodoRequest{
	Title: "title",
	Description: "Description",
	Status: "Status",
	Priority: 1,
	ExpirationDate: &time.Time{},
	UserID: 1,
}

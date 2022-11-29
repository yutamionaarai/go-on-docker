package service

import (
	"app/model"

	"gorm.io/gorm"
)

func FindTodos(db *gorm.DB, todos *[]model.Todo) (*[]model.Todo, error) {
	err := db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func FirstTodo(db *gorm.DB, todos *model.Todo, id int64) (*model.Todo, error) {
	err := db.First(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

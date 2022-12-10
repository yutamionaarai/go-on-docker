package model

import (
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

// A Todo is ...
type Todo struct {
	ID             int64      `json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Status         string     `json:"status"`
	Priority       int64      `json:"priority"`
	ExpirationDate *time.Time `json:"expiration_date"`
	UserID         int64      `json:"user_id"`
	CreatedAT      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type TodoRequest struct {
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Status         string     `json:"status"`
	Priority       int64      `json:"priority"`
	ExpirationDate *time.Time `json:"expiration_date"`
	UserID         int64      `json:"user_id"`
}

type FindTodoResponse struct {
	Todo *Todo `json:"todos"`
}

type FindTodosResponse struct {
	Todos []*Todo `json:"todos"`
}

type CreateTodoResponse struct {
	ID int64 `json:"id"`
}
type UpdateTodoResponse struct {
	ID int64 `json:"id"`
}
type DeleteTodoResponse struct {
}

// Decides Ozzo-Validation Rules for Todo
func (t *TodoRequest) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Title, validation.Required.Error("必須項目です。"), validation.Length(0, 32).Error("32文字以下にしてください。")),
		validation.Field(&t.Description, validation.Length(0, 256).Error("256文字以下にしてください。")),
		validation.Field(&t.Status, validation.Length(0, 32).Error("32文字以下にしてください。")),
		validation.Field(&t.Priority, validation.Min(0).Error("0以上の数字にしてください。"), validation.Max(1000).Error("1000以下の数字にしてください。")),
		validation.Field(&t.ExpirationDate, validation.Min(time.Now()).Error("現在よりも後の日時を選択してください。")),
		validation.Field(&t.UserID, validation.Required.Error("必須項目です。")),
	)
}

// TodoValidate implements Validate Decided with Above Rule
func (t *TodoRequest) TodoValidate(c *gin.Context) error {
	if err := t.Validate(); err != nil {
		if err, ok := err.(validation.InternalError); ok {
			// バリデーション処理中のInternal Server Errorを切り分け
			// 参考にした箇所 https://qiita.com/gold-kou/items/201a19d9d0c760cc2104
			c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
			return err
		}
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return err
	}
	return nil
}

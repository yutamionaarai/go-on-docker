package model

import (
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Todo struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	Priority       int64     `json:"priority"`
	ExpirationDate time.Time `json:"expiration_date"`
	UserID         int64     `json:"user_id"`
	CreatedAT      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (t Todo) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Title, validation.Required.Error("必須項目です。"), validation.Length(0, 32).Error("32文字以下にしてください。")),
		validation.Field(&t.Description, validation.Length(0, 256).Error("256文字以下にしてください。")),
		validation.Field(&t.Status, validation.Length(0, 32).Error("32文字以下にしてください。")),
		validation.Field(&t.Priority, validation.Required, is.Int.Error("数字で入力して下さい"), validation.Min(0).Error("0以上の数字にしてください。"), validation.Max(100).Error("100以下の数字にしてください。")),
		validation.Field(&t.UserID, validation.Required.Error("必須項目です。"), is.Int.Error("数字で入力して下さい")),
	)
}

func (todoRequest *Todo) TodoValidate(c *gin.Context) error {
	if err := todoRequest.Validate(); err != nil {
		if err, ok := err.(validation.InternalError); ok {
			// バリデーション処理中のInternal Server Errorを切り分け
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
			return err
		}
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return err
	}
	return nil
}

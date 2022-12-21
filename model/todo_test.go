package model

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTodoValidate(t *testing.T) {
	// 正常系データ
	t.Parallel()

	testNormalTodoRequest := TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: 1, UserID: 1}
	if ok := assert.NoError(t, testNormalTodoRequest.TodoValidate(&gin.Context{})); !ok {
		t.Error("想定とは異なり、バリデーションエラーが発生しました。")
	}

	// 異常系データ(UserIDが存在しない)
	testMissTodoRequest := TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: 1}
	if ok := assert.Error(t, testMissTodoRequest.TodoValidate(&gin.Context{})); !ok {
		t.Error("想定とは異なり、バリデーションエラーが発生しませんでした。")
	}
}

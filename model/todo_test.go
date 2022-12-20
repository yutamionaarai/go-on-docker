package model

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTodoValidate(t *testing.T) {
	// 正常系データ
	testNormalTodoRequest := TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: 1, UserID: 1}
	assert.NoError(t, testNormalTodoRequest.TodoValidate(&gin.Context{}))

	// 異常系データ(UserIDが存在しない)
	testMissTodoRequest := TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: 1}
	assert.Error(t, testMissTodoRequest.TodoValidate(&gin.Context{}))
}

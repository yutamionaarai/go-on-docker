package model

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTodoValidate(t *testing.T) {
	// 正常系データ
	t.Parallel()
	c, _ := gin.CreateTestContext(nil)

	testCases := map[string]struct {
		request TodoRequest
		wantErr bool
	}{
		"正常系データ": {
			TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: 1, UserID: 1},
			false,
		},
		"異常系データ(UserIDが存在しない)": {
			TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: 1},
			true,
		},
		"異常系データ(Priorityが負)": {
			TodoRequest{Title: "goの勉強", Description: "テストコードを書く", Status: "始めたばかり", Priority: -1, UserID: 1},
			true,
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotErr := tc.request.TodoValidate(c)
			if tc.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}

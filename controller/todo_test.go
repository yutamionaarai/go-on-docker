package controller_test

import (
	"app/controller"
	"app/mock"
	"app/model"
	"app/testdata"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)


type TodoControllerTestSuite struct {
	suite.Suite
	mock   *mock.TodoRepositoryMock
	ctrl   *controller.TodoController
   }

// テスト開始前の共通処理
func (s *TodoControllerTestSuite) SetupTest() {
    s.mock = new(mock.TodoRepositoryMock)
    s.ctrl = controller.NewTodoController(s.mock)
}

// FindTodoに対するテスト
func (s *TodoControllerTestSuite) TestFindTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		id      int64
		wantErr error
		wantTodoResponse model.FindTodoResponse
	}{
		"正常系データ": {
			id:      int64(1),
			wantErr: nil,
			wantTodoResponse: model.FindTodoResponse{Todo: testdata.Todo},
 	},
		"異常系データ(IDが負)": {
			id:      int64(-1),
			wantErr: fmt.Errorf("id is negative"),
			wantTodoResponse: model.FindTodoResponse{},
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("FindTodo", tc.id).Return(tc.wantTodoResponse, tc.wantErr)
			gotResponse, gotErr := s.ctrl.FindTodo(tc.id)
			// 期待するエラーと返却されたエラーが一致するか
			s.Equal(tc.wantErr, gotErr)
			// 期待するレスポンスと返却されたエラーが一致するか
			s.Equal(tc.wantTodoResponse, gotResponse)
		})
	}
}


func (s *TodoControllerTestSuite) TestFindTodos() {
	s.T().Parallel()
	testCases := map[string]struct {
		wantErr error
		wantTodosResponse model.FindTodosResponse
	}{
		"正常系データ": {
			wantErr: nil,
			wantTodosResponse: model.FindTodosResponse{Todos: testdata.Todos},
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("FindTodos").Return(tc.wantTodosResponse, tc.wantErr)
			gotResponse, gotErr := s.ctrl.FindTodos()
			s.Equal(tc.wantErr, gotErr)
			s.Equal(tc.wantTodosResponse, gotResponse)
		})
	}
}

func (s *TodoControllerTestSuite) TestCreateTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		id      int64
		wantErr error
		todoRequest *model.TodoRequest
		wantTodoResponse model.CreateTodoResponse
	}{
		"正常系データ": {
			id:      int64(1),
			wantErr: nil,
			todoRequest: testdata.TodoNormalRequest,
			wantTodoResponse: model.CreateTodoResponse{ID: 1},
		},
		"異常系データ(項目タイトルが存在しない)": {
			id:          int64(-1),
			wantErr: fmt.Errorf("titleは必須項目です。"),
			todoRequest: testdata.TodoAbnormalRequest,
			wantTodoResponse: model.CreateTodoResponse{},
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("CreateTodo", tc.todoRequest).Return(tc.wantTodoResponse, tc.wantErr)
			gotResponse, gotErr := s.ctrl.CreateTodo(tc.todoRequest)
			s.Equal(tc.wantErr, gotErr)
			s.Equal(tc.wantTodoResponse, gotResponse)
		})
	}
}


func TestTodoControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TodoControllerTestSuite))
}

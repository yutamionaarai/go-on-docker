package controller_test

import (
	"app/controller"
	"app/mock"
	"app/model"
	"app/router"
	"app/testdata"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TodoControllerTestSuite struct {
	suite.Suite
	mock   *mock.TodoRepositoryMock
	router *gin.Engine
}

func (s *TodoControllerTestSuite) SetupTest() {
	s.mock = new(mock.TodoRepositoryMock)
	todoController := controller.NewTodoController(
		s.mock,
	)
	s.router = router.NewRouter(todoController)
}

func (s *TodoControllerTestSuite) TestFindTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		wantStatusCode int
	}{
		"正常系データ": {
			wantStatusCode: http.StatusOK,
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("FindTodo", int64(1)).Return(model.FindTodoResponse{Todo: testdata.Todo}, nil)
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/todos/1", nil)
			s.NoError(err)
			s.router.ServeHTTP(w, req)
			s.Equal(tc.wantStatusCode, w.Code)
		})
	}
}

func (s *TodoControllerTestSuite) TestFindsTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		wantStatusCode int
	}{
		"正常系データ": {
			wantStatusCode: http.StatusOK,
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("FindTodos").Return(model.FindTodosResponse{Todos: testdata.Todos}, nil)
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/todos/", nil)
			s.NoError(err)
			s.router.ServeHTTP(w, req)
			s.Equal(tc.wantStatusCode, w.Code)
		})
	}
}

func (s *TodoControllerTestSuite) TestCreateTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		todoRequest         *model.TodoRequest
		wantStatusCode      int
		wantTodoResponse    model.CreateTodoResponse
		wantTodoResponseErr error
	}{
		"正常系データ": {
			todoRequest:         testdata.TodoRequest,
			wantStatusCode:      http.StatusOK,
			wantTodoResponse:    model.CreateTodoResponse{ID: 0},
			wantTodoResponseErr: nil,
		},
		"異常系：Internal Server Error nilが送信された場合": {
			todoRequest:         nil,
			wantStatusCode:      http.StatusInternalServerError,
			wantTodoResponse:    model.CreateTodoResponse{},
			wantTodoResponseErr: fmt.Errorf("Internal Server Error"),
		},
		"異常系：Bad Request Error 不正なリクエストパラメータが送信された場合": {
			todoRequest:         testdata.InvalidTodoRequest,
			wantStatusCode:      http.StatusBadRequest,
			wantTodoResponse:    model.CreateTodoResponse{},
			wantTodoResponseErr: fmt.Errorf("Bad Request Error"),
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("CreateTodo", tc.todoRequest).Return(tc.wantTodoResponse, tc.wantTodoResponseErr)
			w := httptest.NewRecorder()
			jsonValue, err := json.Marshal(tc.todoRequest)
			s.NoError(err)
			req, err := http.NewRequest("POST", "/todos/", bytes.NewBuffer(jsonValue))
			s.NoError(err)
			s.router.ServeHTTP(w, req)
			s.Equal(tc.wantStatusCode, w.Code)
			// Internal Server Errorの時は以降のテストは行わない
			if tc.wantStatusCode == http.StatusInternalServerError {
				return
			}
			body, err := ioutil.ReadAll(w.Result().Body)
			s.NoError(err)
			var response model.CreateTodoResponse
			json.Unmarshal(body, &response)
			s.Equal(tc.wantTodoResponse, response)
		})
	}
}
func (s *TodoControllerTestSuite) TestUpdateTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		todoRequest      *model.TodoRequest
		wantStatusCode   int
		wantTodoResponse model.UpdateTodoResponse
	}{
		"正常系データ": {
			todoRequest:      testdata.TodoRequest,
			wantStatusCode:   http.StatusOK,
			wantTodoResponse: model.UpdateTodoResponse{ID: 0},
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			w := httptest.NewRecorder()
			s.mock.On("UpdateTodo", int64(1), tc.todoRequest).Return(tc.wantTodoResponse, nil)
			jsonValue, err := json.Marshal(tc.todoRequest)
			s.NoError(err)
			req, err := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonValue))
			s.NoError(err)
			s.router.ServeHTTP(w, req)
			body, err := ioutil.ReadAll(w.Result().Body)
			s.NoError(err)
			var response model.UpdateTodoResponse
			json.Unmarshal(body, &response)
			s.Equal(tc.wantStatusCode, w.Code)
			s.Equal(tc.wantTodoResponse, response)
		})
	}
}
func (s *TodoControllerTestSuite) TestDeleteTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		wantStatusCode int
	}{
		"正常系データ": {
			wantStatusCode: http.StatusOK,
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			w := httptest.NewRecorder()
			s.mock.On("DeleteTodo", int64(1)).Return(model.DeleteTodoResponse{}, nil)
			req, err := http.NewRequest("DELETE", "/todos/1", nil)
			s.NoError(err)
			s.router.ServeHTTP(w, req)
			s.Equal(tc.wantStatusCode, w.Code)
		})
	}
}
func TestTodoControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TodoControllerTestSuite))
}

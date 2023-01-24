package controller_test

import (
	"app/controller"
	"app/mock"
	"app/model"
	"app/router"
	"app/testdata"
	"bytes"
	"encoding/json"
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
			s.mock.On("FindTodo", int64(1)).Return(model.FindTodoResponse{Todo:testdata.Todo}, nil)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/todos/1", nil)
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
			req, _ := http.NewRequest("GET", "/todos/", nil)
			s.router.ServeHTTP(w, req)
			s.Equal(tc.wantStatusCode, w.Code)
		})
	}
}

func (s *TodoControllerTestSuite) TestCreateTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		todoRequest *model.TodoRequest
		wantStatusCode int
		wantTodoResponse model.CreateTodoResponse
	}{
		"正常系データ": {
			todoRequest: testdata.TodoRequest,
			wantStatusCode: http.StatusOK,
			wantTodoResponse: model.CreateTodoResponse{ID:0},
	    },
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			s.T().Parallel()
			s.mock.On("CreateTodo", tc.todoRequest).Return(tc.wantTodoResponse, nil)
			w := httptest.NewRecorder()
			jsonValue, _ := json.Marshal(tc.todoRequest)
			req, _ := http.NewRequest("POST", "/todos/", bytes.NewBuffer(jsonValue))
			s.router.ServeHTTP(w, req)
			body, _ := ioutil.ReadAll(w.Result().Body)
			var response model.CreateTodoResponse
			json.Unmarshal(body, &response)
			s.Equal(tc.wantStatusCode, w.Code)
			s.Equal(tc.wantTodoResponse, response)
		})
	}
}


func (s *TodoControllerTestSuite) TestUpdateTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		todoRequest *model.TodoRequest
		wantStatusCode int
		wantTodoResponse model.UpdateTodoResponse
	}{
    	"正常系データ": {
			todoRequest: testdata.TodoRequest,
			wantStatusCode: http.StatusOK,
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
			jsonValue, _ := json.Marshal(tc.todoRequest)
			req, _ := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonValue))
			s.router.ServeHTTP(w, req)
			body, _ := ioutil.ReadAll(w.Result().Body)
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
			req, _ := http.NewRequest("DELETE", "/todos/1", nil)
			s.router.ServeHTTP(w, req)
			s.Equal(tc.wantStatusCode, w.Code)
		})
	}
}

func TestTodoControllerTestSuite(t *testing.T) {
    suite.Run(t, new(TodoControllerTestSuite))
}

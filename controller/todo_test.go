package controller_test

import (
	"app/controller"
	"app/controller/middleware"
	"app/mock"
	"app/model"
	"app/testdata"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)


func NewRouter(todoController *controller.TodoController) *gin.Engine {
    r := gin.Default()
    r.Use(requestid.New())
    r.Use(middleware.HandleErrors)

    todos := r.Group("/todos")
    {
        todos.GET("/hello", todoController.HelloController)
        // todoリストを全件取得
        todos.GET("/", todoController.FindTodosController)
        // 該当のIDのtodoリストを取得
        todos.GET("/:id", todoController.FindTodoController)
        // todoリストの作成
        todos.POST("/", todoController.CreateTodoController)
        // 該当のIDのtodoリストの更新
        todos.PUT("/:id", todoController.UpdateTodoController)
        // 該当のIDのtodoリストの削除
        todos.DELETE("/:id", todoController.DeleteTodoController)
    }
    return r
}

type TodoControllerTestSuite struct {
    suite.Suite
    mock   *mock.TodoRepositoryMock
    router   *gin.Engine
}

func (s *TodoControllerTestSuite) SetupTest() {
    s.mock = new(mock.TodoRepositoryMock)
	s.mock.On("FindTodo").Return(model.FindTodoResponse{Todo: testdata.Todo}, nil)
	s.mock.On("FindTodos").Return(model.FindTodosResponse{Todos: testdata.Todos}, nil)
	s.mock.On("CreateTodo").Return(model.CreateTodoResponse{ID: 1}, nil)
	s.mock.On("UpdateTodo").Return(model.UpdateTodoResponse{ID: 1}, nil)
	s.mock.On("DeleteTodo").Return(model.DeleteTodoResponse{}, nil)
    todoController := controller.NewTodoController(
        s.mock,
    )
    s.router = NewRouter(todoController)
}


func (s *TodoControllerTestSuite) TestFindTodo() {
    s.T().Parallel()
	w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/todos/:id", nil)
    s.router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *TodoControllerTestSuite) TestFindsTodo() {
    s.T().Parallel()
	w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/todos/", nil)
    s.router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *TodoControllerTestSuite) TestCreateTodo() {
    s.T().Parallel()
	w := httptest.NewRecorder()
    jsonValue, _ := json.Marshal(testdata.TodoNormalRequest)
    req, _ := http.NewRequest("POST", "/todos/", bytes.NewBuffer(jsonValue))
    s.router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}


func (s *TodoControllerTestSuite) TestUpdateTodo() {
    s.T().Parallel()
	w := httptest.NewRecorder()
    jsonValue, _ := json.Marshal(testdata.TodoNormalRequest)
    req, _ := http.NewRequest("PUT", "/todos/:id", bytes.NewBuffer(jsonValue))
    s.router.ServeHTTP(w, req)
    var response model.UpdateTodoResponse
    json.NewDecoder(w.Body).Decode(&response)
    
	s.Equal(http.StatusOK, w.Code)
    s.Equal(model.UpdateTodoResponse{ID: 1}, response)
}

func (s *TodoControllerTestSuite) TestDeleteTodo() {
    s.T().Parallel()
	w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/todos/:id", nil)
    s.router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func TestTodoControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TodoControllerTestSuite))
}

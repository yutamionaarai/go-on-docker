package router

import (
	"app/controller"
	"app/controller/middleware"
	"app/repository"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewRouter implement various Endpoints.
func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(requestid.New())
	todoController := controller.NewTodoController(
		repository.NewTodoRepository(
			db,
		),
	)
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

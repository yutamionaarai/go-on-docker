package router

import (
	"app/controller"
	"app/controller/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	todoController := controller.NewTodoController(db)
	r.Use(middleware.HandleErrors)
	r.GET("/", todoController.HelloController)

	todos := r.Group("/todos")
	{
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

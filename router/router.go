package router

import (
	"app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.Hello())
	// todoリストを全件取得
	r.GET("/todos", controller.FindTodosController(db))
	// 該当のIDのtodoリストを取得
	r.GET("todos/:id", controller.FindTodoController(db))
	// todoリストの作成
	r.POST("todos/", controller.CreateTodoController(db))
	// 該当のIDのtodoリストの更新
	r.PUT("todos/:id", controller.D(db))
	// 該当のIDのtodoリストの削除
	r.DELETE("todos/:id", controller.E(db))
	return r
}

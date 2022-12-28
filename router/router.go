package router

import (
	"app/controller"
	"app/controller/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewRouter implement various Endpoints.
func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3001",
		}, // アクセス元のドメインを許可する
		AllowMethods: []string{
			"GET",
			"POST",
			"POST",
			"PUT",
			"DELETE",
		}, // アクセス元の HTTP メソッドを許可する
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization", // アクセス元の HTTP ヘッダーを許可する
		},
	}
	// r.Use(cors.Default())
	r.Use(cors.New(config))

	r.Use(requestid.New())
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

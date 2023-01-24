package main

import (
	"log"
	"os"

	"app/controller"
	"app/db"
	"app/repository"
	"app/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("failed to load ENV", err)
	}
	dsn := os.Getenv("DB_DSN")
	todoDB, err := db.NewDB(dsn)
	if err != nil {
		log.Fatalln("failed to open DB", err)
	}
	defer func() {
		err := db.CloseDB(todoDB)
		if err != nil {
			log.Fatalln("failed to close DB", err)
		}
	}()
	todoController := controller.NewTodoController(
		repository.NewTodoRepository(
			todoDB,
		),
	)


	r := router.NewRouter(todoController)
	r.Run()
}

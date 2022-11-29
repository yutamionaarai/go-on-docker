package main

import (
	"log"
	"os"

	"app/db"
	"app/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("failed to load ENV", err)
	}
	dsn := os.Getenv("DB_DSN")
	db, err := db.NewDB(dsn)
	if err != nil {
		log.Fatalln("failed to open DB", err)
	}
	r := router.NewRouter(db)
	r.Run()
}

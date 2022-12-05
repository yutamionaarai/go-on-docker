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
	postgresDB, err := db.NewDB(dsn)
	if err != nil {
		log.Fatalln("failed to open DB", err)
	}
	defer func() {
		err := db.CloseDB(postgresDB)
		if err != nil {
			log.Fatalln("failed to close DB", err)
		}
	}()

	r := router.NewRouter(postgresDB)
	r.Run()
}

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/chilipizdrick/muzek-server/api"
	"github.com/chilipizdrick/muzek-server/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("[INFO] No .env file found.")
	}
	if os.Getenv("POSTGRES_DSN") == "" {
		log.Fatalln("[FATAL] POSTGRESQL_DSN environment variable not specified.")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	log.Printf("%s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("[FATAL] Failed to connect to database.")
	}

	database.AutoMigrateSchemas(db)

	r := gin.New()

	api.AssignRouteHandlers(r, db)

	r.Run()
}

package main

import (
	"log"
	"userdata-redis-example-go/config"
	"userdata-redis-example-go/db"
	"userdata-redis-example-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewConfig()
	redisConfig := config.LoadRedisConfig()

	redisDB := db.NewRedisDB(redisConfig)

	defer redisDB.Close()

	_, err := db.Initialize(cfg)

	if err != nil {
		log.Fatalf("Failed top initialize database: %v", err)

	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	r := gin.Default()

	handlers.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

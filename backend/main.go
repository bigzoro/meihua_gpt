package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/example/meihua_gpt/backend/database"
	"github.com/example/meihua_gpt/backend/handlers"
	"github.com/example/meihua_gpt/backend/models"
)

func main() {
	dbPath := os.Getenv("MEIHUA_DB_PATH")
	if dbPath == "" {
		dbPath = "meihua.db"
	}

	db, err := database.Init(dbPath, &models.Divination{})
	if err != nil {
		log.Fatalf("failed to initialise database: %v", err)
	}

	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

	handlers.RegisterDivinationRoutes(router.Group("/api"), db)

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	log.Printf("starting server on %s", addr)
	if err := router.Run(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

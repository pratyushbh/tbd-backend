package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	r "tbd-backend/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found.")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := r.InitRouter()
	router.Run(":" + port)
}

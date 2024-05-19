package main

import (
	db "final-project/database"
	"final-project/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db.ConnectDatabase()
	r := router.RouterSetup()
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	if err := r.Run(addr); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

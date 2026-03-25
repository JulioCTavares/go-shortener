package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julioctavares/go-shortener/internal/handler"
	"github.com/julioctavares/go-shortener/pkg/config"
)

func main() {
	if err := config.RunMigrations(); err != nil {
		log.Fatal("Failed to apply migrations: ", err)
	}
	config.ConnectDB()
	config.ConnectRedis()

	router := handler.SetupRoutes()

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

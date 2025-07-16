package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/imhasandl/subscription-manager/database"
	"github.com/imhasandl/subscription-manager/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("set port in .env file")
	}

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("set database url in .env file")
	}

	db, err := database.InitDatabase(db_url)
	if err != nil {
		log.Fatalf("can't start database: %v", err)
	}

	apiConfig := handlers.NewConfig(db)

	_ = apiConfig

	mux := http.NewServeMux()

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port: %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("can't start server: %v", err)
	}
}

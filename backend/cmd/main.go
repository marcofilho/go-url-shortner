package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcofilho/go-url-shortner/backend/internals/configs"
	"github.com/marcofilho/go-url-shortner/backend/internals/database"
	"github.com/marcofilho/go-url-shortner/backend/internals/handlers"
	"github.com/marcofilho/go-url-shortner/backend/internals/service"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	mongoURI := "mongodb://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Url + ":" +
		fmt.Sprintf("%d", cfg.Port)
	repo, err := database.NewMongoRepository(mongoURI, cfg.DatabaseName, cfg.CollectionName)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	svc := service.NewService(repo)
	h := &handlers.Handler{Service: svc}

	// Rotas
	http.HandleFunc("/shorten", h.ShortenURL)
	http.HandleFunc("/", h.GetURL)

	log.Println("Servidor rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

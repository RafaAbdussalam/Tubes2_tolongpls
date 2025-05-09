package main

import (
	"little_alchemy_backend/internal/handler"
	"little_alchemy_backend/internal/repo"
	"log"
)

func main() {

	// Initialize repository
	repo, err := repo.NewRepository("data/alchemy.db", "data/alchemy.csv")
	if err != nil {
		log.Fatalf("failed to initialize repository")
	}

	// Initialize router
	router := handler.NewRouter(repo)
	router.Run("0.0.0.0:8080")

}

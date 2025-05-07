package main

import (
	"fmt"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"little_alchemy_backend/internal/tree"
	"log"
)

func main() {
	// Initialize dependencies
	repo, err := repo.NewRepository("data/alchemy.db", "data/alchemy.csv") 
	if err != nil {
		log.Fatalf("shi not workin")
	}

	builder, err := tree.NewBuilder(repo, model.BFS, 5)
	if err != nil {
		log.Fatalf("shi not workin")
	}

	tree, err := builder.BuildTree("Pressure", 3)
	if err != nil {
		log.Fatalf("shi not workin")
	}

	fmt.Println(tree.String())

}
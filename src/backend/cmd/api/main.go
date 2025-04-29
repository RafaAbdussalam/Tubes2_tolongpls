package main

import (
	"fmt"
	"little_alchemy_backend/internal/repo"
	"log"
)

func main() {
	repo, err := repo.NewRepository("data/alchemy.db", "data/alchemy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Get all recipes for "Fire"
	recipes, err := repo.GetRecipes("Bucket")
	if err != nil {
		log.Fatal(err)
	}

	for i, r := range recipes {
		fmt.Print(i+1 , ". ")
		r.PrintR()
	}
}
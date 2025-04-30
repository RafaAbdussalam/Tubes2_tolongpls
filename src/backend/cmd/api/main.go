package main

import (
	"fmt"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"little_alchemy_backend/internal/tree"
	"log"
)

func main() {
	repo, err := repo.NewRepository("data/alchemy.db", "data/alchemy.csv")
	if err != nil {
		log.Fatal(err)
	}

	builder, err := tree.NewBuilder(repo, model.BFS, 1)
	if err != nil {
		log.Fatal(err)
	}

	tree, err := builder.BuildTree("Mud")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(tree.String())

}
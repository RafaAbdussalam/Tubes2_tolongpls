package tree

import (
	"fmt"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
)

type TreeBuilder interface {
	BuildTree(rootElement string, amount int) (*model.RecipeTree, error)
}

func NewBuilder(repo *repo.RecipeRepository, traversal model.Traversal) (TreeBuilder, error) {
	switch traversal {

		// BFS
		case model.BFS:
			return &BFSBuilder{repo: *repo}, nil

		// DFS
		case model.DFS:
			return &DFSBuilder{repo: *repo}, nil // nanti ganti ke DFS

		// Invalid
		default:
			return nil, fmt.Errorf("mode tidak valid")
	}
}

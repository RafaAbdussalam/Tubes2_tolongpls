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
		case model.BFS:
			return &BFSBuilder{repo: *repo}, nil
			// if amount > 1 {
			// 	return &ParallelBuilder{
			// 			base: &BFSBuilder{repo: repo},
			// 			amount: amount,
			// 	}
			// }
		case model.DFS:
			// if amount > 1 {
			// 	return &ParallelBuilder{
			// 			base: &DFSBuilder{repo: repo},
			// 			amount: amount,
			// 	}
			// }
			return &DFSBuilder{repo: *repo}, nil // nanti ganti ke DFS
		default:
			return nil, fmt.Errorf("mode tidak valid")
	}
}

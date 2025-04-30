package tree

import (
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
)

type DFSBuilder struct {
   repo repo.RecipeRepository
}

func (b *DFSBuilder) BuildTree(rootElement string) (*model.RecipeTree, error) {
	
	return nil, nil
	
}


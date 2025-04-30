package service

import (
	"little_alchemy_backend/internal/repo"
	"little_alchemy_backend/internal/tree"
)

type RecipeService struct {
	repo repo.RecipeRepository
	builder tree.TreeBuilder
}

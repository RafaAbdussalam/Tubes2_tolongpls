package tree

import (
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"time"
)

type BFSBuilder struct {
   repo repo.RecipeRepository
}

func (b *BFSBuilder) BuildTree(rootElement string, amount int) (*model.RecipeTree, error) {
	
	start := time.Now()
	tree := model.NewTree(rootElement, model.BFS) // Start tree
	queue := model.NewQueue(tree.Root) // Add root to queue
	
	// Loop through queue
	for !queue.IsEmpty() && tree.Root.RecipeCount < amount{
		current := queue.Pop()

		if current.IsPrimary {
			continue
		}

		recipes, err := b.repo.GetRecipesFor(current.Element)
		if err != nil {
			return nil, err
		}

		// Make new recipe node for each recipe
		for _, recipe := range recipes {

			// New element node
			item1 := model.NewElementNode(recipe.Item1, current.Depth + 1)
			item2 := model.NewElementNode(recipe.Item2, current.Depth + 1)		
			
			// New recipe node
			recipeNode := &model.RecipeNode{
				ParentElement: current,
				Item1: item1,
				Item2: item2,
				RecipeCount: item1.RecipeCount * item2.RecipeCount,
			}

			current.Ingredients = append(current.Ingredients, recipeNode)

			// Update NodeCount
			tree.NodeCount += 2
			if item1.Depth > tree.Depth {
				tree.Depth = item1.Depth
			}

			// Update parent recipe node
			item1.Parent = recipeNode
    		item2.Parent = recipeNode
			model.BubbleCount(nil, recipeNode)

			// Stop if found enough recipes
			if tree.Root.RecipeCount == amount {
				model.PruneTree(tree.Root)
				elapsed := time.Since(start)
				tree.Time = int(elapsed)
				return tree, nil
			}

			queue.Push(item1)
			queue.Push(item2)
		}
	}

	// Furbish tree
	model.PruneTree(tree.Root)
	elapsed := time.Since(start)
	tree.Time = int(elapsed)

	return tree, nil
	
}	
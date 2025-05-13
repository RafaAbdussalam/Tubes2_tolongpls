package tree

import (
	"context"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"sync"
	"time"
)

type DFSBuilder struct {
	repo   repo.RecipeRepository
}

func (d *DFSBuilder) BuildTree(rootElement string, amount int) (*model.RecipeTree, error) {

	start := time.Now()
	
	// Multithreading tools
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tree := model.NewTree(rootElement, model.DFS) // Start tree
	stack := model.NewStack(tree.Root)            // Add root to queue

	// Loop through queue
	for !stack.IsEmpty() {

		// Stop if recipe count exceeded
		if ctx.Err() != nil || tree.RecipeCount >= uint64(amount) {
			break
		}

		// Pop node from stack
		current := stack.Pop()

		// Skip primary element
		if current.IsPrimary {
			continue
		}

		// Get recipes for current element
		recipes, err := d.repo.GetRecipesFor(current.Element)
		if err != nil {
			cancel()
			wg.Wait()
			return nil, err
		}

		// Make new recipe node for each recipe
		for _, recipe := range recipes {

			wg.Add(1)

			go func(recipe *model.Recipe) {
				defer wg.Done()
				
				if ctx.Err() != nil {
					return
				}

				// New element node
				item1 := model.NewElementNode(recipe.Item1, nil, current.Depth + 1)
				item2 := model.NewElementNode(recipe.Item2, nil, current.Depth + 1)

				// New recipe node
				recipeNode := model.NewRecipeNode(item1, item2, current)
				item1.ParentRecipe = recipeNode
				item2.ParentRecipe = recipeNode

				mu.Lock()
				defer mu.Unlock()

				// New ingredient
				current.Ingredients = append(current.Ingredients, recipeNode)

				// Update node count
				tree.NodeCount += 2

				// Update depth
				if item1.Depth > tree.Depth {
					tree.Depth = item1.Depth
				}

				// Recount recipes
				if item1.IsPrimary && item2.IsPrimary {
					tree.CountRecipes(current)
				}

				// Stop if found enough recipes
				if tree.RecipeCount >= uint64(amount) {
					cancel()
					return
				}

				stack.Push(item1)
				stack.Push(item2)

			}(recipe)
		}

		wg.Wait()
	}

	// Polish tree
	tree.TrimTree(amount)
	tree.PruneTree()

	// Time tree
	elapsed := time.Since(start)
	tree.Time = int(elapsed.Milliseconds())

	return tree, nil

}
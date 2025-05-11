package tree

import (
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"sync"
	"time"
)

type BFSBuilder struct {
	repo repo.RecipeRepository
}

func (b *BFSBuilder) BuildTree(rootElement string, amount int) (*model.RecipeTree, error) {

	start := time.Now()
	
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	tree := model.NewTree(rootElement, model.BFS) // Start tree
	queue := model.NewQueue(tree.Root)            // Add root to queue

	// Loop through queue
	for !queue.IsEmpty() && tree.RecipeCount < uint64(amount) {
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
			wg.Add(1)

			go func(recipe *model.Recipe) {
				defer wg.Done()

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

				// Update Node Count
				tree.NodeCount += 2
				if item1.Depth > tree.Depth {
					tree.Depth = item1.Depth
				}

				// Recount recipes
				tree.CountRecipes(current)

				// Stop if found enough recipes
				if tree.RecipeCount == uint64(amount) {
					return
				}

				queue.Push(item1)
				queue.Push(item2)

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

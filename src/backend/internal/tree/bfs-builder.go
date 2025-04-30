package tree

import (
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
)

type BFSBuilder struct {
   repo repo.RecipeRepository
}

func (b *BFSBuilder) BuildTree(rootElement string) (*model.RecipeTree, error) {
	
	tree := model.NewTree(rootElement, "bfs")
	queue := model.NewQueue(tree.Root)

	for !queue.IsEmpty() && tree.Depth < 5 {
		current := queue.Pop()
		// fmt.Println(current.Element)

		if current.IsPrimary {
			continue
		}

		recipes, err := b.repo.GetRecipesFor(current.Element)
		if err != nil {
			return nil, err
		}

		for _, recipe := range recipes {
			item1 := model.NewElementNode(recipe.Item1, current.Depth + 1)
			item2 := model.NewElementNode(recipe.Item2, current.Depth + 1)
			
			current.Ingredients = append(current.Ingredients, &model.RecipeNode{
				Item1: item1,
				Item2: item2,
			})

			tree.Count += 2
			if item1.Depth > tree.Depth {
				tree.Depth = item1.Depth
			}

			queue.Push(item1)
			queue.Push(item2)
		}
	}

	return tree, nil
	
}
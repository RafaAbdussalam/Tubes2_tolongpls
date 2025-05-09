package tree

import (
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"sync"
)

type DFSBuilder struct {
	repo            repo.RecipeRepository
	visitedElements map[string]bool
	mutex           sync.Mutex
	nodeCount       uint8
	maxDepth        uint8
}

func NewDFSBuilder(repo repo.RecipeRepository) *DFSBuilder {
	return &DFSBuilder{
		repo:            repo,
		visitedElements: make(map[string]bool),
		nodeCount:       0,
		maxDepth:        0,
	}
}

func (b *DFSBuilder) BuildTree(rootElement string) (*model.RecipeTree, error) {
	// Inisialisasi pohon resep
	tree := model.NewTree(rootElement, model.DFS)

	// Inisialisasi map elemen yang sudah dikunjungi
	b.visitedElements = make(map[string]bool)
	b.nodeCount = 1 // Mulai dengan node root
	b.maxDepth = 0

	// Tandai root sebagai dikunjungi
	b.visitedElements[rootElement] = true

	// Mulai DFS dari elemen root
	err := b.buildDFSRecursive(tree.Root, 0)

	// Perbarui statistik pohon
	tree.Count = b.nodeCount
	tree.Depth = b.maxDepth

	return tree, err
}

func (b *DFSBuilder) buildDFSRecursive(node *model.ElementNode, depth uint8) error {
	// Perbarui kedalaman maksimum jika diperlukan
	if depth > b.maxDepth {
		b.maxDepth = depth
	}

	// Jika ini adalah elemen dasar, kita selesai dengan cabang ini
	if node.IsPrimary {
		return nil
	}

	// Dapatkan resep yang mungkin untuk elemen ini
	recipes, err := b.repo.GetRecipesFor(node.Element)
	if err != nil || len(recipes) == 0 {
		return err
	}

	// Proses setiap resep
	for _, recipe := range recipes {
		// Buat node resep baru
		recipeNode := &model.RecipeNode{}

		// Buat node elemen untuk bahan-bahan
		item1Node := model.NewElementNode(recipe.Item1, depth+1)
		item2Node := model.NewElementNode(recipe.Item2, depth+1)

		// Hubungkan node resep ke node bahan
		recipeNode.Item1 = item1Node
		recipeNode.Item2 = item2Node

		// Tambahkan node resep ke bahan elemen saat ini
		node.Ingredients = append(node.Ingredients, recipeNode)

		// Tambahkan jumlah node
		b.mutex.Lock()
		b.nodeCount += 2 // Dua node elemen baru
		b.mutex.Unlock()

		// Lanjutkan DFS pada setiap bahan, tetapi hanya jika belum dikunjungi
		b.mutex.Lock()
		visitedItem1 := b.visitedElements[recipe.Item1]
		visitedItem2 := b.visitedElements[recipe.Item2]
		b.mutex.Unlock()

		if !visitedItem1 && !item1Node.IsPrimary {
			b.mutex.Lock()
			b.visitedElements[recipe.Item1] = true
			b.mutex.Unlock()

			if err := b.buildDFSRecursive(item1Node, depth+1); err != nil {
				return err
			}
		}

		if !visitedItem2 && !item2Node.IsPrimary {
			b.mutex.Lock()
			b.visitedElements[recipe.Item2] = true
			b.mutex.Unlock()

			if err := b.buildDFSRecursive(item2Node, depth+1); err != nil {
				return err
			}
		}
	}

	return nil
}

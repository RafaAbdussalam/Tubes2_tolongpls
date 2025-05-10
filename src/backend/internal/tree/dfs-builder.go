package tree

import (
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"sync"
	"time"
	"fmt"
)

type DFSBuilder struct {
	repo            repo.RecipeRepository
	visitedElements map[string]bool
	mutex           sync.Mutex
	nodeCount       int
	maxDepth        int
}

func NewDFSBuilder(repo repo.RecipeRepository) *DFSBuilder {
	return &DFSBuilder{
		repo:            repo,
		visitedElements: make(map[string]bool),
		nodeCount:       0,
		maxDepth:        0,
	}
}

func (b *DFSBuilder) BuildTree(rootElement string, amount int) (*model.RecipeTree, error) {
	tree := model.NewTree(rootElement, model.DFS)
	
	// Catat waktu mulai
	startTime := time.Now()

	b.visitedElements = make(map[string]bool)
	b.nodeCount = 1 
	b.maxDepth = 0

	b.visitedElements[rootElement] = true

	if amount <= 1 {
		err := b.buildDFSRecursive(tree.Root, 0)
		if err != nil {
			return nil, err
		}
	} else {
		// Gunakan pendekatan multithreading untuk banyak resep
		err := b.buildMultipleRecipes(tree.Root, amount)
		if err != nil {
			return nil, err
		}
	}

	// Hitung recipe count untuk semua node
    b.calculateRecipeCount(tree.Root)
    
    if tree.Root.RecipeCount == 0 {
        fmt.Println("Warning: Root element memiliki recipe count 0.")
        fmt.Println("Pruning dinonaktifkan untuk menampilkan semua jalur yang ditemukan.")
    } else {
        // Lakukan pruning hanya jika root memiliki recipe count positif
        model.PruneTree(tree.Root)
    }
	
	tree.NodeCount = b.countTreeNodes(tree.Root)
	tree.Depth = b.maxDepth
	
	tree.Time = int(time.Since(startTime).Milliseconds())
	
	// Set RecipeCount di RecipeTree
	model.SetRecipeCount(tree)
	
	return tree, nil
}

// buildDFSRecursive adalah implementasi DFS rekursif standar (untuk satu resep)
func (b *DFSBuilder) buildDFSRecursive(node *model.ElementNode, depth int) error {
	// Perbarui kedalaman maksimum jika diperlukan
	if depth > b.maxDepth {
		b.maxDepth = depth
	}

	// Jika ini adalah elemen dasar, kita selesai dengan cabang ini
	if node.IsPrimary {
		node.RecipeCount = 1 // Elemen dasar memiliki recipe count 1
		return nil
	}

	// Dapatkan resep yang mungkin untuk elemen ini
	recipes, err := b.repo.GetRecipesFor(node.Element)
	if err != nil {
		return err
	}
	
	// Jika tidak ada resep, kembalikan tanpa error
	if len(recipes) == 0 {
		return nil
	}

	// Proses setiap resep
	for _, recipe := range recipes {
		// Buat node resep baru
		recipeNode := &model.RecipeNode{
			ParentElement: node,
		}

		// Buat node elemen untuk bahan-bahan
		item1Node := model.NewElementNode(recipe.Item1, depth+1)
		item2Node := model.NewElementNode(recipe.Item2, depth+1)

		// Hubungkan node resep ke node bahan
		recipeNode.Item1 = item1Node
		recipeNode.Item2 = item2Node
		item1Node.Parent = recipeNode
		item2Node.Parent = recipeNode

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
		} else if item1Node.IsPrimary {
			// Set recipe count untuk elemen primer
			item1Node.RecipeCount = 1
		}

		if !visitedItem2 && !item2Node.IsPrimary {
			b.mutex.Lock()
			b.visitedElements[recipe.Item2] = true
			b.mutex.Unlock()

			if err := b.buildDFSRecursive(item2Node, depth+1); err != nil {
				return err
			}
		} else if item2Node.IsPrimary {
			// Set recipe count untuk elemen primer
			item2Node.RecipeCount = 1
		}
	}

	return nil
}

// processIngredientThreadSafe memproses bahan secara rekursif dengan thread-safety
func (b *DFSBuilder) processIngredientThreadSafe(node *model.ElementNode, visited map[string]bool, depth int, resultMutex *sync.Mutex) {
	// Update kedalaman maksimum dengan aman
	resultMutex.Lock()
	if depth > b.maxDepth {
		b.maxDepth = depth
	}
	resultMutex.Unlock()
	
	// Jika elemen primer, selesai
	if node.IsPrimary {
		node.RecipeCount = 1 // Elemen primer memiliki recipe count 1
		return
	}
	
	// Tandai sebagai dikunjungi
	visited[node.Element] = true
	
	// Dapatkan resep untuk elemen ini
	recipes, err := b.repo.GetRecipesFor(node.Element)
	if err != nil || len(recipes) == 0 {
		return
	}
	
	// Pilih resep pertama saja untuk menghindari eksplorasi berlebihan
	recipe := recipes[0]
	
	// Buat node resep baru
	recipeNode := &model.RecipeNode{
		ParentElement: node,
	}
	
	// Buat node elemen untuk bahan-bahan
	item1Node := model.NewElementNode(recipe.Item1, depth+1)
	item2Node := model.NewElementNode(recipe.Item2, depth+1)
	
	// Hubungkan node resep ke node bahan
	recipeNode.Item1 = item1Node
	recipeNode.Item2 = item2Node
	item1Node.Parent = recipeNode
	item2Node.Parent = recipeNode
	
	// Tambahkan node resep ke bahan elemen saat ini
	node.Ingredients = append(node.Ingredients, recipeNode)
	
	// Lanjutkan secara rekursif, hindari siklus
	if !visited[recipe.Item1] && !item1Node.IsPrimary {
		newVisited := b.copyMap(visited)
		newVisited[recipe.Item1] = true
		b.processIngredientThreadSafe(item1Node, newVisited, depth+1, resultMutex)
	} else if item1Node.IsPrimary {
		// Set recipe count untuk elemen primer
		item1Node.RecipeCount = 1
	}
	
	if !visited[recipe.Item2] && !item2Node.IsPrimary {
		newVisited := b.copyMap(visited)
		newVisited[recipe.Item2] = true
		b.processIngredientThreadSafe(item2Node, newVisited, depth+1, resultMutex)
	} else if item2Node.IsPrimary {
		// Set recipe count untuk elemen primer
		item2Node.RecipeCount = 1
	}
}

// buildMultipleRecipes mencari banyak resep dengan multithreading
func (b *DFSBuilder) buildMultipleRecipes(rootNode *model.ElementNode, amount int) error {
	// Dapatkan resep untuk elemen root
	recipes, err := b.repo.GetRecipesFor(rootNode.Element)
	if err != nil {
		return err
	}
	
	// Jika tidak ada resep, kembalikan tanpa error
	if len(recipes) == 0 {
		return nil
	}

	// Batasi jumlah resep yang akan diproses
	recipeCount := len(recipes)
	if amount < recipeCount {
		recipeCount = amount
	}

	// WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Mutex untuk sinkronisasi
	var resultMutex sync.Mutex

	// Proses setiap resep dalam goroutine terpisah
	for i := 0; i < recipeCount; i++ {
		recipe := recipes[i]
		
		wg.Add(1)
		go func(r *model.Recipe) {
			defer wg.Done()
			
			// Buat node resep baru
			recipeNode := &model.RecipeNode{
				ParentElement: rootNode,
			}
			
			// Buat node elemen untuk bahan-bahan
			item1Node := model.NewElementNode(r.Item1, 1)
			item2Node := model.NewElementNode(r.Item2, 1)
			
			// Hubungkan node resep ke node bahan
			recipeNode.Item1 = item1Node
			recipeNode.Item2 = item2Node
			item1Node.Parent = recipeNode
			item2Node.Parent = recipeNode
			
			// Buat salinan visited map untuk thread ini
			localVisited := make(map[string]bool)
			localVisited[rootNode.Element] = true
			
			// Proses setiap bahan secara rekursif
			b.processIngredientThreadSafe(item1Node, localVisited, 1, &resultMutex)
			b.processIngredientThreadSafe(item2Node, localVisited, 1, &resultMutex)
			
			// Tambahkan hasil ke tree dengan aman
			resultMutex.Lock()
			rootNode.Ingredients = append(rootNode.Ingredients, recipeNode)
			b.nodeCount += b.countNodesRecursive(recipeNode) // Hitung total node baru
			resultMutex.Unlock()
		}(recipe)
	}
	
	// Tunggu semua goroutine selesai
	wg.Wait()
	
	return nil
}

// calculateRecipeCount menghitung recipe count untuk node secara rekursif
func (b *DFSBuilder) calculateRecipeCount(node *model.ElementNode) int {
	if node == nil {
		return 0
	}
	
	// Jika elemen dasar, recipeCount = 1
	if node.IsPrimary {
		node.RecipeCount = 1
		return 1
	}
	
	// Jika tidak ada resep, recipeCount = 0
	if len(node.Ingredients) == 0 {
		node.RecipeCount = 0
		return 0
	}
	
	totalCount := 0
	
	// Hitung untuk setiap resep
	for _, recipe := range node.Ingredients {
		// Periksa null pointer
		if recipe == nil || recipe.Item1 == nil || recipe.Item2 == nil {
			continue
		}
		
		item1Count := b.calculateRecipeCount(recipe.Item1)
		item2Count := b.calculateRecipeCount(recipe.Item2)
		
		// Recipe count adalah perkalian dari recipe count bahan-bahannya
		recipeCount := item1Count * item2Count
		recipe.RecipeCount = recipeCount
		
		// Total count adalah penjumlahan dari semua resep
		totalCount += recipeCount
	}
	
	node.RecipeCount = totalCount
	return totalCount
}

// countTreeNodes menghitung jumlah node dalam pohon setelah pruning
func (b *DFSBuilder) countTreeNodes(node *model.ElementNode) int {
	if node == nil {
		return 0
	}
	
	count := 1 // Count this node
	
	for _, recipe := range node.Ingredients {
		// Count recipe node
		count++

		// Count item1 subtree
		if recipe.Item1 != nil {
			count += b.countTreeNodes(recipe.Item1)
		}
		// Count item2 subtree
		if recipe.Item2 != nil {
			count += b.countTreeNodes(recipe.Item2)
		}
	}
	
	return count
}

// countNodesRecursive menghitung jumlah node dalam subpohon recipe
func (b *DFSBuilder) countNodesRecursive(recipeNode *model.RecipeNode) int {
	if recipeNode == nil {
		return 0
	}

	count := 2 // Untuk item1 dan item2
	// Tambahkan jumlah node dari subtree item1
	if recipeNode.Item1 != nil && !recipeNode.Item1.IsPrimary {
		for _, r := range recipeNode.Item1.Ingredients {
			count += b.countNodesRecursive(r)
		}
	}
	// Tambahkan jumlah node dari subtree item2
	if recipeNode.Item2 != nil && !recipeNode.Item2.IsPrimary {
		for _, r := range recipeNode.Item2.Ingredients {
			count += b.countNodesRecursive(r)
		}
	}
	return count
}

// copyMap membuat salinan dari map
func (b *DFSBuilder) copyMap(original map[string]bool) map[string]bool {
	copied := make(map[string]bool)
	for k, v := range original {
		copied[k] = v
	}
	return copied
}
package tree


import (
    "fmt"
    "little_alchemy_backend/internal/model"
    "little_alchemy_backend/internal/repo"
    "sync"
    "time"
)


type DFSBuilder struct {
    repo            repo.RecipeRepository
    visitedElements map[string]bool
    Mutex           sync.Mutex
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
    b.nodeCount = 1 // Mulai dengan node root
    b.maxDepth = 0

    b.visitedElements[rootElement] = true

    if amount <= 1 {
        // Gunakan DFS yang memastikan satu jalur lengkap
        err := b.buildSingleFullPath(tree.Root, 0)
        if err != nil {
            return nil, err
        }
    } else {
        // Gunakan multithreading untuk banyak resep
        err := b.buildMultipleFullPaths(tree.Root, amount)
        if err != nil {
            return nil, err
        }
    }
    b.calculateRecipeCount(tree.Root)
   
    // Cek jika recipe count 0, jangan terapkan pruning
    if tree.Root.RecipeCount == 0 {
        fmt.Println("Warning: Recipe count 0 - menampilkan semua jalur yang ditemukan tanpa pruning")
    } else {
        // Terapkan pruning hanya jika recipe count > 0
        tree.PruneTree()
    }

    tree.NodeCount =  b.nodeCount
    tree.Depth = b.maxDepth
   
    // Set waktu eksekusi
    tree.Time = int(time.Since(startTime).Milliseconds())

    tree.SetRecipeCount()

    return tree, nil
}

func (b *DFSBuilder) buildSingleFullPath(node *model.ElementNode, depth int) error {
    // Perbarui kedalaman maksimum
    if depth > b.maxDepth {
        b.maxDepth = depth
    }
    if node.IsPrimary {
        node.RecipeCount = 1
        return nil
    }

    recipes, err := b.repo.GetRecipesFor(node.Element)
    if err != nil {
        return err
    }
    if len(recipes) == 0 {
        fmt.Printf("Warning: Tidak ada resep untuk elemen '%s'\n", node.Element)
        return nil
    }

    recipe := recipes[0]

    recipeNode := &model.RecipeNode{
        ParentElement: node,
    }

    item1Node := model.NewElementNode(recipe.Item1, recipeNode, depth+1)
    item2Node := model.NewElementNode(recipe.Item2, recipeNode, depth+1)

    recipeNode.Item1 = item1Node
    recipeNode.Item2 = item2Node
    item1Node.ParentRecipe = recipeNode
    item2Node.ParentRecipe = recipeNode

    node.Ingredients = append(node.Ingredients, recipeNode)

	b.nodeCount += 2
	b.maxDepth += 1
    if !item1Node.IsPrimary {
        // Reset visited untuk memastikan jalur lengkap
        if err := b.buildSingleFullPath(item1Node, depth+1); err != nil {
            return err
        }
    } else {
        item1Node.RecipeCount = 1
    }

    if !item2Node.IsPrimary {
        // Reset visited untuk memastikan jalur lengkap
        if err := b.buildSingleFullPath(item2Node, depth+1); err != nil {
            return err
        }
    } else {
        item2Node.RecipeCount = 1
    }
    return nil
}

func (b *DFSBuilder) processFullPath(node *model.ElementNode, depth int, resultMutex *sync.Mutex) {
    resultMutex.Lock()
    if depth > b.maxDepth {
        b.maxDepth = depth
    }
    resultMutex.Unlock()

    if node.IsPrimary {
        node.RecipeCount = 1
        return
    }

    recipes, err := b.repo.GetRecipesFor(node.Element)
    if err != nil || len(recipes) == 0 {
        return
    }

    recipe := recipes[0]

    recipeNode := &model.RecipeNode{
        ParentElement: node,
    }

    item1Node := model.NewElementNode(recipe.Item1, recipeNode, depth+1)
    item2Node := model.NewElementNode(recipe.Item2, recipeNode, depth+1)

    recipeNode.Item1 = item1Node
    recipeNode.Item2 = item2Node
    item1Node.ParentRecipe = recipeNode
    item2Node.ParentRecipe = recipeNode

    node.Ingredients = append(node.Ingredients, recipeNode)

    if !item1Node.IsPrimary {
        b.processFullPath(item1Node, depth+1, resultMutex)
    } else {
        item1Node.RecipeCount = 1
    }
   
    if !item2Node.IsPrimary {
        b.processFullPath(item2Node, depth+1, resultMutex)
    } else {
        item2Node.RecipeCount = 1
    }
}

func (b *DFSBuilder) buildMultipleFullPaths(rootNode *model.ElementNode, amount int) error {
    recipes, err := b.repo.GetRecipesFor(rootNode.Element)
    if err != nil {
        return err
    }
   
    if len(recipes) == 0 {
        fmt.Printf("Warning: Tidak ada resep untuk elemen '%s'\n", rootNode.Element)
        return nil
    }

    recipeCount := len(recipes)
    if amount < recipeCount {
        recipeCount = amount
    }

    var wg sync.WaitGroup
    var resultMutex sync.Mutex

    for i := 0; i < recipeCount; i++ {
        recipe := recipes[i]
       
        wg.Add(1)
        go func(r *model.Recipe) {
            defer wg.Done()

            recipeNode := &model.RecipeNode{
                ParentElement: rootNode,
            }

            item1Node := model.NewElementNode(r.Item1, recipeNode, 1)
            item2Node := model.NewElementNode(r.Item2, recipeNode, 1)

            recipeNode.Item1 = item1Node
            recipeNode.Item2 = item2Node
            item1Node.ParentRecipe = recipeNode
            item2Node.ParentRecipe = recipeNode

            b.processFullPath(item1Node, 1, &resultMutex)
            b.processFullPath(item2Node, 1, &resultMutex)
           
            // Tambahkan hasil ke tree dengan aman
            resultMutex.Lock()
            rootNode.Ingredients = append(rootNode.Ingredients, recipeNode)
            b.nodeCount += 2 + b.countNodesRecursive(recipeNode) // Hitung total node baru
            resultMutex.Unlock()
        }(recipe)
    }

    wg.Wait()
   
    return nil
}

func (b *DFSBuilder) calculateRecipeCount(node *model.ElementNode) int {
    if node == nil {
        return 0
    }

    if node.IsPrimary {
        node.RecipeCount = 1
        return 1
    }

    if len(node.Ingredients) == 0 {
        node.RecipeCount = 0
        return 0
    }
   
    totalCount := 0

    for _, recipe := range node.Ingredients {
        if recipe == nil || recipe.Item1 == nil || recipe.Item2 == nil {
            continue
        }

        item1Count := b.calculateRecipeCount(recipe.Item1)
        item2Count := b.calculateRecipeCount(recipe.Item2)

        var recipeCount int
        if item1Count > 0 && item2Count > 0 {
            recipeCount = item1Count * item2Count
        } else {
            recipeCount = 0
        }
       
		recipe.RecipeCount = uint64(recipeCount)
        totalCount += recipeCount
    }
   
    node.RecipeCount = uint64(totalCount)
    return totalCount
}

func (b *DFSBuilder) countNodesRecursive(recipeNode *model.RecipeNode) int {
    if recipeNode == nil {
        return 0
    }
   
    count := 0
    if recipeNode.Item1 != nil {
        count += 1 + b.countSubtreeNodes(recipeNode.Item1)
    }
   
    // Count item2 subtree
    if recipeNode.Item2 != nil {
        count += 1 + b.countSubtreeNodes(recipeNode.Item2)
    }
   
    return count
}

func (b *DFSBuilder) countSubtreeNodes(node *model.ElementNode) int {
    if node == nil || node.IsPrimary {
        return 0
    }
   
    count := 0
   
    for _, recipe := range node.Ingredients {
        count++

        if recipe.Item1 != nil {
            count += 1 + b.countSubtreeNodes(recipe.Item1)
        }
       
        if recipe.Item2 != nil {
            count += 1 + b.countSubtreeNodes(recipe.Item2)
        }
    }
   
    return count
}
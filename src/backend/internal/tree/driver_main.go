package tree

import (
	"fmt"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"strings"
	"time"
)

// TestDFSSingle menjalankan pengujian algoritma DFS untuk pencarian resep tunggal
func TestDFSSingle(repository repo.RecipeRepository, targetElements []string) {
	fmt.Println("\n=== PENGUJIAN ALGORITMA DFS UNTUK RESEP TUNGGAL ===")
	
	for _, element := range targetElements {
		fmt.Printf("\n----------------------------------------\n")
		fmt.Printf("Mencari resep untuk: %s\n", element)
		fmt.Printf("----------------------------------------\n")
		
		// Buat DFS builder
		builder := NewDFSBuilder(repository)
		
		// Catat waktu mulai
		startTime := time.Now()
		
		// Bangun pohon resep
		recipeTree, err := builder.BuildTree(element)
		
		// Catat waktu selesai
		duration := time.Since(startTime)
		
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		
		// Tampilkan hasil
		fmt.Printf("Waktu pencarian: %v\n", duration)
		fmt.Printf("Jumlah node: %d\n", recipeTree.Count)
		fmt.Printf("Kedalaman pohon: %d\n", recipeTree.Depth)
		
		// Tampilkan pohon resep (ringkas)
		fmt.Println("\nRingkasan Pohon Resep:")
		PrintCompactTree(recipeTree.Root, 0, 20) // Cetak hingga 3 level kedalaman
	}
}

// PrintCompactTree mencetak versi ringkas dari pohon resep
func PrintCompactTree(node *model.ElementNode, level int, maxLevel int) {
	if node == nil {
		return
	}
	
	indent := strings.Repeat("  ", level)
	
	// Cetak elemen saat ini
	if level == 0 {
		fmt.Printf("%sTarget: %s\n", indent, node.Element)
	} else {
		if node.IsPrimary {
			fmt.Printf("%s- %s (PRIMER)\n", indent, node.Element)
			return
		} else {
			fmt.Printf("%s- %s\n", indent, node.Element)
		}
	}
	
	// Jika mencapai batas kedalaman, hentikan
	if level >= maxLevel {
		if len(node.Ingredients) > 0 {
			fmt.Printf("%s  ...\n", indent)
		}
		return
	}
	
	// Cetak resep (maksimal 3 resep pertama)
	maxRecipes := 5
	if len(node.Ingredients) > maxRecipes {
		fmt.Printf("%s  (Menampilkan %d dari %d resep)\n", 
			indent, maxRecipes, len(node.Ingredients))
	}
	
	for i, recipe := range node.Ingredients {
		if i >= maxRecipes {
			break
		}
		
		fmt.Printf("%s  Resep %d: %s + %s\n", 
			indent, 
			i+1, 
			recipe.Item1.Element, 
			recipe.Item2.Element)
		
		// Cetak bahan secara rekursif
		PrintCompactTree(recipe.Item1, level+2, maxLevel)
		PrintCompactTree(recipe.Item2, level+2, maxLevel)
	}
}

// GetShortestRecipe mengembalikan resep dengan jalur terpendek
func GetShortestRecipe(node *model.ElementNode) *model.RecipeNode {
	if node == nil || len(node.Ingredients) == 0 {
		return nil
	}
	
	shortestRecipe := node.Ingredients[0]
	shortestDepth := getRecipeDepth(shortestRecipe)
	
	for _, recipe := range node.Ingredients {
		depth := getRecipeDepth(recipe)
		if depth < shortestDepth {
			shortestRecipe = recipe
			shortestDepth = depth
		}
	}
	
	return shortestRecipe
}

// getRecipeDepth menghitung kedalaman maksimal dari sebuah resep
func getRecipeDepth(recipe *model.RecipeNode) int {
	if recipe == nil {
		return 0
	}
	
	depth1 := getNodeMaxDepth(recipe.Item1)
	depth2 := getNodeMaxDepth(recipe.Item2)
	
	if depth1 > depth2 {
		return depth1
	}
	return depth2
}

// getNodeMaxDepth menghitung kedalaman maksimal dari sebuah node
func getNodeMaxDepth(node *model.ElementNode) int {
	if node == nil {
		return 0
	}
	
	if node.IsPrimary || len(node.Ingredients) == 0 {
		return 1
	}
	
	shortestRecipe := GetShortestRecipe(node)
	return 1 + getRecipeDepth(shortestRecipe)
}
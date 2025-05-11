package main

import (
	"bufio"
	"flag"
	"fmt"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"little_alchemy_backend/internal/tree"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dbPath := flag.String("db", "./data/alchemy.db", "Path ke database SQLite")
	csvPath := flag.String("csv", "./data/alchemy.csv", "Path ke file CSV")
	
	flag.Parse()

	fmt.Println("=== PENGUJIAN DFS BUILDER LITTLE ALCHEMY 2 ===")
	fmt.Printf("DB Path: %s\n", *dbPath)
	fmt.Printf("CSV Path: %s\n", *csvPath)
	
	// Inisialisasi repository
	repository, err := repo.NewRepository(*dbPath, *csvPath)
	if err != nil {
		fmt.Printf("Error: Tidak dapat membuat repository: %v\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	var targetElement string
	var recipeCount int
	var displayDepth int
	var searchElement string

	fmt.Print("\nMasukkan elemen target (default: Brick): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		targetElement = "Brick"
	} else {
		targetElement = input
	}

	var mode string
	for {
		fmt.Print("Pilih mode (1: Single Recipe, 2: Multiple Recipe): ")
		mode, _ = reader.ReadString('\n')
		mode = strings.TrimSpace(mode)
		
		if mode == "1" || mode == "2" {
			break
		}
		fmt.Println("Input tidak valid. Silakan pilih 1 atau 2.")
	}

	if mode == "2" {
		for {
			fmt.Print("Masukkan jumlah resep yang ingin dicari (2-10): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			
			count, err := strconv.Atoi(input)
			if err == nil && count >= 2 && count <= 10 {
				recipeCount = count
				break
			}
			fmt.Println("Input tidak valid. Masukkan angka antara 2 dan 10.")
		}
	} else {
		recipeCount = 1
	}

	fmt.Print("Masukkan kedalaman tampilan pohon (default: 20): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		displayDepth = 20
	} else {
		depth, err := strconv.Atoi(input)
		if err == nil && depth > 0 {
			displayDepth = depth
		} else {
			fmt.Println("Input tidak valid. Menggunakan nilai default: 20")
			displayDepth = 20
		}
	}
	
	fmt.Println("\n=== RINGKASAN PENGUJIAN ===")
	fmt.Printf("Elemen Target: %s\n", targetElement)
	if recipeCount > 1 {
		fmt.Printf("Mode: Multiple Recipe DFS (%d resep)\n", recipeCount)
	} else {
		fmt.Println("Mode: Single Recipe DFS")
	}
	fmt.Printf("Kedalaman Tampilan: %d\n", displayDepth)
	if searchElement != "" {
		fmt.Printf("Pencarian Elemen: %s\n", searchElement)
	}
	fmt.Println("=========================")

	dfsBuilder := tree.NewDFSBuilder(*repository)

	startTime := time.Now()

	fmt.Printf("\nMembangun pohon resep untuk '%s'...\n", targetElement)
	recipeTree, err := dfsBuilder.BuildTree(targetElement, recipeCount)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	duration := time.Since(startTime)
	
	// Tampilkan hasil
	fmt.Printf("\n=== HASIL PENCARIAN ===\n")
	fmt.Printf("Waktu pencarian: %v\n", duration)
	fmt.Printf("Jumlah node: %d\n", recipeTree.NodeCount)
	fmt.Printf("Kedalaman pohon: %d\n", recipeTree.Depth)
	fmt.Printf("Recipe count untuk root: %d\n", recipeTree.RecipeCount)
	fmt.Printf("Jumlah resep ditemukan: %d\n", len(recipeTree.Root.Ingredients))
	
	// Tampilkan pohon resep (ringkas)
	fmt.Println("\nRingkasan Pohon Resep:")
	PrintCompactTree(recipeTree.Root, 0, displayDepth)

	if searchElement != "" {
		fmt.Printf("\n=== MENCARI ELEMEN: %s ===\n", searchElement)
		foundNode := FindNodeByName(recipeTree.Root, searchElement)
		if foundNode != nil {
			fmt.Printf("Elemen '%s' ditemukan dalam pohon!\n", searchElement)
			PrintNodeInfo(foundNode)
			
			// Cetak jalur dari elemen ini ke root
			fmt.Printf("\nJalur dari '%s' ke root:\n", searchElement)
			PrintPathToRoot(foundNode)
		} else {
			fmt.Printf("Elemen '%s' tidak ditemukan dalam pohon.\n", searchElement)
		}
	}
}

func PrintCompactTree(node *model.ElementNode, level int, maxLevel int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", level)
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
	if level >= maxLevel {
		if len(node.Ingredients) > 0 {
			fmt.Printf("%s  ...\n", indent)
		}
		return
	}
	if len(node.Ingredients) == 0 && !node.IsPrimary {
		fmt.Printf("%s  (Tidak ada resep tersedia)\n", indent)
		return
	}
	maxRecipes := 1
	if len(node.Ingredients) > maxRecipes {
		fmt.Printf("%s  (Menampilkan %d dari %d resep)\n",
			indent, maxRecipes, len(node.Ingredients))
	}

	for i, recipe := range node.Ingredients {
		if i >= maxRecipes {
			break
		}

		if recipe == nil {
			fmt.Printf("%s  Resep %d: NULL RECIPE\n", indent, i+1)
			continue
		}
		if recipe.Item1 == nil || recipe.Item2 == nil {
			fmt.Printf("%s  Resep %d: INVALID RECIPE (missing ingredients)\n", indent, i+1)
			continue
		}
		fmt.Printf("%s  Resep %d: %s + %s\n",
			indent, i+1, recipe.Item1.Element, recipe.Item2.Element)

		PrintCompactTree(recipe.Item1, level+2, maxLevel)
		PrintCompactTree(recipe.Item2, level+2, maxLevel)
	}
}

func FindNodeByName(root *model.ElementNode, targetName string) *model.ElementNode {
	if root == nil {
		return nil
	}
	if strings.EqualFold(root.Element, targetName) {
		return root
	}
	for _, recipe := range root.Ingredients {
		// Periksa null pointers
		if recipe == nil || recipe.Item1 == nil || recipe.Item2 == nil {
			continue
		}
		if found := FindNodeByName(recipe.Item1, targetName); found != nil {
			return found
		}
		if found := FindNodeByName(recipe.Item2, targetName); found != nil {
			return found
		}
	}
	return nil
}

func PrintNodeInfo(node *model.ElementNode) {
	fmt.Printf("\nInformasi Node: %s\n", node.Element)
	fmt.Printf("Primary: %v\n", node.IsPrimary)
	fmt.Printf("Depth: %d\n", node.Depth)
	fmt.Printf("Recipe Count: %d\n", node.RecipeCount)
	
	if node.IsPrimary {
		fmt.Println("Elemen ini adalah elemen dasar (primer).")
		return
	}

	fmt.Printf("Jumlah Resep: %d\n", len(node.Ingredients))
	
	if len(node.Ingredients) > 0 {
		fmt.Println("\nResep yang tersedia:")
		for i, recipe := range node.Ingredients {
			if recipe != nil && recipe.Item1 != nil && recipe.Item2 != nil {
				fmt.Printf("%d. %s + %s (Count: %d)\n", 
					i+1, recipe.Item1.Element, recipe.Item2.Element, recipe.RecipeCount)
			}
		}
	} else {
		fmt.Println("Tidak ada resep yang tersedia untuk elemen ini.")
	}
}

func PrintPathToRoot(node *model.ElementNode) {
	if node == nil {
		return
	}

	path := []string{node.Element}
	current := node
	
	for current.ParentRecipe != nil && current.ParentRecipe.ParentElement != nil {
		parentElement := current.ParentRecipe.ParentElement

		if current.ParentRecipe.Item1 == current {
			path = append(path, fmt.Sprintf("%s + %s -> %s", 
				current.Element, current.ParentRecipe.Item2.Element, parentElement.Element))
		} else {
			path = append(path, fmt.Sprintf("%s + %s -> %s", 
				current.ParentRecipe.Item1.Element, current.Element, parentElement.Element))
		}
		current = parentElement
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%s\n", path[i])
	}
}
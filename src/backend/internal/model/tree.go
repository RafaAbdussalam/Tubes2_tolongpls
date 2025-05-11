package model

import (
	"fmt"
	"sort"
	"strings"
)

// Tree
type RecipeTree struct {
	Mode      	Traversal    	`json:"algorithm"`
	Depth     	int          	`json:"depth"`
	NodeCount 	int        		`json:"node_count"`
	RecipeCount uint64 			`json:"recipe_count"`
	Time 			int 				`json:"time"`
	Root      	*ElementNode 	`json:"tree_data"`
}

func NewTree(root string, mode Traversal) *RecipeTree {
	return &RecipeTree{
		Root:      		NewElementNode(root, nil, 0),
		Depth:     		0,
		NodeCount: 		1,
		RecipeCount: 	0,
		Mode:      		mode,
	}
}

func (tree *RecipeTree) SetRecipeCount() {
	tree.RecipeCount = uint64(tree.Root.RecipeCount)
}

// Recount recipes in tree
func (tree *RecipeTree) CountRecipes(node *ElementNode) uint64 {
	bubbleNodes(node)
	tree.SetRecipeCount()
	return tree.RecipeCount
}

// Recount recipes starting from recipe node
func bubbleRecipes(node *RecipeNode) {
	if node == nil {
		return
	}	

	// Recount recipe node
	node.RecipeCount = node.Item1.RecipeCount * node.Item2.RecipeCount

	// Recount parent element node
	if node.ParentElement != nil {
		bubbleNodes(node.ParentElement)
	}
}

// Recount recipes starting from element node
func bubbleNodes(node *ElementNode) {
	if node == nil {
		return
	}

	// Recount element node
	count := uint64(0)
	for _, recipe := range node.Ingredients {
		count += recipe.RecipeCount
	}
	node.RecipeCount = count

	// Recount parent recipe node
	if node.ParentRecipe != nil {
		bubbleRecipes(node.ParentRecipe)
	}
}

// Prune tree to have no zero recipe count nodes
func (tree RecipeTree) PruneTree() {
	pruneNode(tree.Root)
}

// Prune starting from this node
func pruneNode(node *ElementNode) {
	if node == nil {
		return
	}

	trimmed := make([]*RecipeNode, 0, len(node.Ingredients))
	for _, recipe := range node.Ingredients {

		// Trim useless recipes nodes
		if recipe == nil || recipe.RecipeCount == 0 {
			continue
		}

		pruneNode(recipe.Item1)
		pruneNode(recipe.Item2)

		// Keep node otherwise
		trimmed = append(trimmed, recipe)
	}

	node.Ingredients = trimmed
}

// Trims tree to have correct recipe count
func (tree *RecipeTree) TrimTree(amount int) {
	if tree.RecipeCount <= uint64(amount) {
		return
	}
	
	// Get leaf nodes 
	leaves := getLeafRecipes(tree.Root)
	
	// Sort leaves from deepest 
	sort.SliceStable(leaves, func(i, j int) bool {
		return leaves[i].Depth > leaves[j].Depth // deeper first
	})
	
	// Remove leaves until correct count is reached
	for _, leaf := range leaves {
		if tree.RecipeCount == uint64(amount) {
			break 
		}
		
		// Simulate removing leaf 
		leaf.RecipeCount = 0
		tree.CountRecipes(leaf.ParentElement)

		// Undo removal if undershoot  
		if tree.RecipeCount < uint64(amount) {
			bubbleRecipes(leaf)
			tree.SetRecipeCount() 
		}

	}
}

// Get leaf recipe nodes (with primary children) from given node
func getLeafRecipes(node *ElementNode) []*RecipeNode {
	var leaves []*RecipeNode

	for _, recipe := range node.Ingredients {
		if recipe.RecipeCount > 0 && recipe.Item1.IsPrimary && recipe.Item2.IsPrimary {

			// Found a leaf 
			leaves = append(leaves, recipe)

		} else {

			// Recursively get leaves 
			leaves = append(leaves, getLeafRecipes(recipe.Item1)...)
			leaves = append(leaves, getLeafRecipes(recipe.Item2)...)
			
		}
	}
	
	return leaves
}

// Type of Traversal
type Traversal string

const (
	BFS           Traversal = "bfs"
	DFS           Traversal = "dfs"
	Bidirectional Traversal = "bd"
)

// Print recipe tree
func (t *RecipeTree) String() string {
	var sb strings.Builder
	t.printNode(&sb, t.Root, 0)
	sb.WriteString(fmt.Sprintf(
		"\nTree (Mode: %s, Depth: %d, Nodes: %d, Recipes: %v, Time: %vms)\n\n", 
		t.Mode, t.Depth, t.NodeCount, t.Root.RecipeCount, t.Time,
	))
	return sb.String()
}

// Print recipe tree data 
func (t *RecipeTree) printNode(sb *strings.Builder, node *ElementNode, indentLevel int) {
	indent := strings.Repeat("  ", indentLevel)

	sb.WriteString(fmt.Sprintf("%s%s", indent, node.Element))
	if node.IsPrimary {
		sb.WriteString(" (PRIMARY),\n")
		return
	} else {
		sb.WriteString(fmt.Sprintf(" (%v)", node.RecipeCount))
	}
	sb.WriteString(" {\n")

	for i, recipe := range node.Ingredients {
		sb.WriteString(fmt.Sprintf("%s  %v (%v) {\n", indent, i+1, recipe.RecipeCount))
		t.printNode(sb, recipe.Item1, indentLevel+2)
		t.printNode(sb, recipe.Item2, indentLevel+2)
		sb.WriteString(fmt.Sprintf("%s  },\n", indent))
	}
	sb.WriteString(fmt.Sprintf("%s},\n", indent))
}
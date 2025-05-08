package model

import (
	"fmt"
	"strings"
)

// Tree
type RecipeTree struct {
	Mode      	Traversal    	`json:"algorithm"`
	Depth     	int          	`json:"depth"`
	NodeCount 	int        		`json:"node_count"`
	Time 			int 				`json:"time"`
	Root      	*ElementNode 	`json:"tree_data"`
}

func NewTree(rootElement string, mode Traversal) *RecipeTree {
	return &RecipeTree{
		Root:      NewElementNode(rootElement, 0),
		Depth:     0,
		NodeCount: 1,
		Mode:      mode,
	}
}

// Update recipe count 
func BubbleCount(elementNode *ElementNode, recipeNode *RecipeNode) {
	if elementNode == nil && recipeNode == nil {
		return
	}
	if elementNode != nil {
		newCount := int(0)
		for _, ingredient := range elementNode.Ingredients {
			newCount += ingredient.RecipeCount
		}
		elementNode.RecipeCount = newCount
		BubbleCount(nil, elementNode.Parent)
	} else if recipeNode != nil {
		recipeNode.RecipeCount = recipeNode.Item1.RecipeCount * recipeNode.Item2.RecipeCount
		BubbleCount(recipeNode.ParentElement, nil)
	}
}

// Trim nodes with zero count
func PruneTree(elementNode *ElementNode) {
	if elementNode == nil {
		return
	}

	trimmed := make([]*RecipeNode, 0, len(elementNode.Ingredients))
	for _, recipe := range elementNode.Ingredients {
		if recipe == nil || recipe.RecipeCount == 0 {
			continue
		}

		PruneTree(recipe.Item1)
		PruneTree(recipe.Item2)

		trimmed = append(trimmed, recipe)
	}

	elementNode.Ingredients = trimmed
}

// Type of Traversal
type Traversal string

const (
	BFS           Traversal = "bfs"
	DFS           Traversal = "dfs"
	Bidirectional Traversal = "bd"
)

// Print Tree Method
func (t *RecipeTree) String() string {
	var sb strings.Builder
	t.printNode(&sb, t.Root, 0)
	sb.WriteString(fmt.Sprintf("\nTree (Mode: %s, Depth: %d, Nodes: %d, Recipes: %v, Time: %v)\n\n", t.Mode, t.Depth, t.NodeCount, t.Root.RecipeCount, t.Time))
	return sb.String()
}

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
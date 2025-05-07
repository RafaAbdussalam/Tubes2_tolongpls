package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ElementNode struct {
	Element     string         `json:"element"`
	Ingredients []*RecipeNode  `json:"ingredients,omitempty"`
	IsPrimary   bool           `json:"isPrimary"`
	Depth       int          	`json:"depth"`
	RecipeCount int  				`json:"recipe_count"`
	Parent		*RecipeNode 	`json:"parent_node,omitempty"`
}

type RecipeNode struct {
	ParentElement *ElementNode 	`json:"parent_element"`
	Item1 *ElementNode 				`json:"element_1"`
	Item2 *ElementNode 				`json:"element_2"`
	RecipeCount int  					`json:"recipe_count"`
}

// Constructor
func NewElementNode(name string, depth int) *ElementNode {
	isPrimary := isPrimary(name)
	return &ElementNode{
		Element: name,
		IsPrimary: isPrimary,
		Depth: depth,
		RecipeCount: boolToInt(isPrimary),
	}
}

type RecipeTree struct {
	Root  		*ElementNode	`json:"root"`
	Depth 		int        		`json:"depth"` 
	NodeCount 	uint8    		`json:"node_count"` 
	Mode  		Traversal		`json:"mode"` 
}

// Add MarshalJSON to handle Traversal enum properly
func (t Traversal) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

// Constructor
func NewTree(rootElement string, mode Traversal) *RecipeTree {
	return &RecipeTree{
		Root: NewElementNode(rootElement, 0),
		Depth: 0,
		NodeCount: 1,
		Mode: mode,
	}
}

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

// Type of Traversal
type Traversal string

const (
	BFS          	Traversal = "bfs"
	DFS          	Traversal = "dfs"
	Bidirectional 	Traversal = "bd"
)

// Print Tree Method
func (t *RecipeTree) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nTree (Mode: %s, Depth: %d, Nodes: %d)\n\n", t.Mode, t.Depth, t.NodeCount))
	t.printNode(&sb, t.Root, 0)
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
		if recipe.RecipeCount != 0 {
			sb.WriteString(fmt.Sprintf("%s  %v (%v) {\n", indent, i+1, recipe.RecipeCount))
			t.printNode(sb, recipe.Item1, indentLevel+2)
			t.printNode(sb, recipe.Item2, indentLevel+2)
			sb.WriteString(fmt.Sprintf("%s  },\n", indent))
		}
	}
	sb.WriteString(fmt.Sprintf("%s},\n", indent))
}

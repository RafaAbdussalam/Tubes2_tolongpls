package model

import (
	"fmt"
	"strings"
)

type ElementNode struct {
	Element     string
	Ingredients []*RecipeNode // anak simpul
	IsPrimary   bool
	Depth       uint8
}

type RecipeNode struct {
	Item1 *ElementNode
	Item2 *ElementNode
}

// Constructor
func NewElementNode(name string, depth uint8) *ElementNode {
	return &ElementNode{
		Element:   name,
		IsPrimary: isPrimary(name),
		Depth:     depth,
	}
}

type RecipeTree struct {
	Root  *ElementNode
	Depth uint8 // kedalaman tercapai
	Count uint8 // jumlah simpul
	Mode  Traversal
}

// Constructor
func NewTree(rootElement string, mode Traversal) *RecipeTree {
	return &RecipeTree{
		Root:  NewElementNode(rootElement, 0),
		Depth: 0,
		Count: 1,
		Mode:  mode,
	}
}

type Traversal string

const (
   BFS Traversal = "bfs"
   DFS Traversal = "dfs"
	Bidirectional Traversal = "bd"
)

func (t *RecipeTree) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nTree (Mode: %s, Depth: %d, Nodes: %d)\n\n", t.Mode, t.Depth, t.Count))
	t.printNode(&sb, t.Root, 0)
	return sb.String()
}

func (t *RecipeTree) printNode(sb *strings.Builder, node *ElementNode, indentLevel int) {
	indent := strings.Repeat("  ", indentLevel)

	sb.WriteString(fmt.Sprintf("%s%s", indent, node.Element))
	if node.IsPrimary {
		sb.WriteString(" (PRIME),\n")
		return
	}
	sb.WriteString(" {\n")

	for i, recipe := range node.Ingredients {
		sb.WriteString(fmt.Sprintf("%s  %v {\n", indent, i+1))
		t.printNode(sb, recipe.Item1, indentLevel+2)
		t.printNode(sb, recipe.Item2, indentLevel+2)
		sb.WriteString(fmt.Sprintf("%s  },\n", indent))
	}
	sb.WriteString(fmt.Sprintf("%s},\n", indent))
}
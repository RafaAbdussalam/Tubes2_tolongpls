package model

type ElementNode struct {
	Element     string        `json:"element"`
	Ingredients []*RecipeNode `json:"children"`
	IsPrimary   bool          `json:"-"`
	Depth       int           `json:"-"`
	RecipeCount int           `json:"-"`
	Parent      *RecipeNode   `json:"-"`
}

type RecipeNode struct {
	ParentElement *ElementNode `json:"-"`
	Item1         *ElementNode `json:"item_1"`
	Item2         *ElementNode `json:"item_2"`
	RecipeCount   int          `json:"-"`
}

// Constructor
func NewElementNode(name string, depth int) *ElementNode {
	isPrimary := isPrimary(name)
	return &ElementNode{
		Element:     name,
		IsPrimary:   isPrimary,
		Depth:       depth,
		RecipeCount: boolToInt(isPrimary),
	}
}

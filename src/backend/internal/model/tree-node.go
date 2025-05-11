package model

// Structures
type ElementNode struct {
	Element      string        `json:"element"`
	Ingredients  []*RecipeNode `json:"children"`
	ParentRecipe *RecipeNode   `json:"-"`
	IsPrimary    bool          `json:"-"`
	Depth        int           `json:"-"`
	RecipeCount  uint64        `json:"-"`
}

type RecipeNode struct {
	Item1         *ElementNode `json:"item_1"`
	Item2         *ElementNode `json:"item_2"`
	ParentElement *ElementNode `json:"-"`
	Depth         int          `json:"-"`
	RecipeCount   uint64       `json:"-"`
}

// Constructor
func NewElementNode(name string, ParentRecipe *RecipeNode, depth int) *ElementNode {
	isPrimary := isPrimary(name)
	return &ElementNode{
		Element:      name,
		ParentRecipe: ParentRecipe,
		IsPrimary:    isPrimary,
		Depth:        depth,
		RecipeCount:  uint64(boolToInt(isPrimary)),
	}
}

func NewRecipeNode(item1, item2, ParentElement *ElementNode) *RecipeNode {
	return &RecipeNode{
		Item1: 			item1,
		Item2: 			item2,
		ParentElement: ParentElement,
		Depth: 			item1.Depth,
		RecipeCount: 	item1.RecipeCount * item2.RecipeCount,
	}
}

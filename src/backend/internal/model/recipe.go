package model

import "fmt"

type Recipe struct {
	Element   string
	Item1     string
	Item2     string
	IsPrimary bool
}

// Constructor
func NewRecipe(element, item1, item2 string) *Recipe {
	return &Recipe{
		Element:   element,
		Item1:     item1,
		Item2:     item2,
		IsPrimary: isPrimary(element),
	}
}

// Cek apakah elemen primer
func isPrimary(element string) bool {
	primes := map[string]bool{
		"Earth": true,
		"Air":   true,
		"Fire":  true,
		"Water": true,
	}
	return primes[element]
}

func (r *Recipe) PrintR() {
	fmt.Printf("%s = %s + %s\n", r.Element, r.Item1, r.Item2)
}
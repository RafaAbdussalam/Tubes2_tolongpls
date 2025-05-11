package scraper

import (
	"fmt"
	"little_alchemy_backend/internal/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(doc *goquery.Document) ([]*model.Recipe, ElementTierMap, error) {
	var recipes []*model.Recipe
	elementTiers := make(ElementTierMap)
	currentTier := -1

	// Cari setiap jenis elemen
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		headline := s.Find(".mw-headline")

		sectionID, exists := headline.Attr("id")
		if !exists {
			return
		}

		// Catat tier dari section 
		if strings.Contains(sectionID, "Starting_elements") || strings.Contains(sectionID, "Special_element") {
			currentTier = 0
		} else if strings.HasPrefix(sectionID, "Tier_") {

			// Ambil tier dari judul 
			tierParts := strings.Split(sectionID, "_")
			if len(tierParts) >= 2 {
				tierNum, err := strconv.Atoi(tierParts[1])
				if err == nil {
					currentTier = tierNum
				}
			}

		} else {
			return
		}

		// Cari tabel untuk jenis elemen
		table := s.NextAllFiltered("table").First()
		if table.Length() == 0 {
			return
		}

		// Cari setiap baris pada tabel
		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if i == 0 { // skip header
				return
			}

			tds := tr.Find("td")
			if tds.Length() < 2 {
				return
			}

			// Cari nama elemen (kolom pertama)
			element := strings.TrimSpace(tds.Eq(0).Text())
			var item1, item2 string

			// Catat tier elemen
			if currentTier >= 0 {
				elementTiers[element] = currentTier
			}

			// Cari recipe elemen (kolom kedua)
			liFound := false
			tds.Eq(1).Find("li").Each(func(_ int, li *goquery.Selection) {
				text := strings.TrimSpace(li.Text())
				if strings.Contains(text, "+") {

					// Pisahkan berdasarkan '+' dalam <li>
					parts := strings.Split(text, "+")

					// Catat recipe
					if len(parts) == 2 {
						item1 = strings.TrimSpace(parts[0])
						item2 = strings.TrimSpace(parts[1])
						recipes = append(recipes, model.NewRecipe(element, item1, item2))
					}

					liFound = true
				}
			})

			// Tidak ada '+' dalam <li>, berarti elemen spesial
			if !liFound || (element == "Earth" || element == "Air" || element == "Fire" || element == "Water" || element == "Time") {
				recipes = append(recipes, model.NewRecipe(element, "", ""))
			}
		})

		fmt.Println("Parsed:", sectionID)

	})

	fmt.Println("HTML terbaca")
	return recipes, elementTiers, nil
}

package repo

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/scraper"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

type RecipeRepository struct {
	db      *sql.DB
	csvPath string
	tiers   scraper.ElementTierMap
	mode    string
}

func NewRepository(dbPath, csvPath, jsonPath string) (*RecipeRepository, error) {

	// Mencoba file json
	tiers, err := getTierMap(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("tidak bisa mengakses json: %w", err)
	}

	// Mencoba database SQLite
	if _, err := os.Stat(dbPath); err == nil {
		db, err := sql.Open("sqlite", dbPath)
		if err == nil {
			if err = db.Ping(); err == nil {
				log.Println("Menggunakan database SQLite")
				return &RecipeRepository{
					db:    db,
					tiers: tiers,
					mode:  "db",
				}, nil
			} else {
				log.Println("Gagal menghubungkan ke database SQLite:", err)
			}
		} else {
			log.Println("Gagal membuka database SQLite:", err)
		}
	} else {
		log.Println("Tidak ada file database:", err)
	}

	// Mencoba file csv
	if _, err := os.Stat(csvPath); err == nil {
		log.Println("Menggunakan file CSV")
		return &RecipeRepository{
			csvPath: csvPath,
			tiers:   tiers,
			mode:    "csv",
		}, nil
	} else {
		log.Println("Tidak ada file CSV:", err)
	}
	
	// Keduanya gagal
	return nil, fmt.Errorf("tidak bisa mengakses db dan csv: %w", err)

}

func (repo *RecipeRepository) GetRecipesFor(element string) ([]*model.Recipe, error) {
	switch repo.mode {
	case "db":
		return repo.getFromDB(element)
	case "csv":
		return repo.getFromCSV(element)
	default:
		return nil, fmt.Errorf("mode tidak ada: %s", repo.mode)
	}
}

func (repo *RecipeRepository) getFromDB(element string) ([]*model.Recipe, error) {

	// Ambil semua recipe elemen dari DB
	rows, err := repo.db.Query(`
		SELECT element, item1, item2 
		FROM elements 
		WHERE element = ?`,
		element)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Simpan dalam suatu slice recipe
	var recipes []*model.Recipe
	for rows.Next() {
		
		// Buat unsur recipe
		var element, item1, item2 string
		if err := rows.Scan(&element, &item1, &item2); err != nil {
			return nil, err
		}

		// Ambil tier setiap unsur
		elemTier, okElem := repo.tiers[element]
		tier1, ok1 := repo.tiers[item1]
		tier2, ok2 := repo.tiers[item2]

		// Append jika tier element lebih besar dari tier items
		if okElem && ok1 && ok2 && tier1 < elemTier && tier2 < elemTier {
			recipes = append(recipes, model.NewRecipe(element, item1, item2))
		}

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (repo *RecipeRepository) getFromCSV(element string) ([]*model.Recipe, error) {

	// Buka file csv
	file, err := os.Open(repo.csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Baca seluruh csv
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Cari dan simpan baris dengan recipe elemen dalam slice
	var recipes []*model.Recipe
	for i, record := range records {

		// Skip header
		if i == 0 {
			continue
		}

		// Simpan recipe element
		if len(record) == 3 && strings.EqualFold(record[0], element) {

			// Ambil tier setiap unsur
			elemTier, okElem := repo.tiers[record[0]]
			tier1, ok1 := repo.tiers[record[1]]
			tier2, ok2 := repo.tiers[record[2]]

			// Append jika tier element lebih besar dari tier items
			if okElem && ok1 && ok2 && tier1 < elemTier && tier2 < elemTier {
				recipes = append(recipes, model.NewRecipe(record[0], record[1], record[2]))
			}

		}
	}

	return recipes, nil
}

func getTierMap(jsonPath string) (scraper.ElementTierMap, error) {

	// Buka file json
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Konversi json menjadi hashmap
	var tiers scraper.ElementTierMap
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tiers)
	if err != nil {
		return nil, err
	}

	return tiers, nil
}
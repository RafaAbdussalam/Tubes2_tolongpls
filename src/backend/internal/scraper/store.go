package scraper

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"little_alchemy_backend/internal/model"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

// Constructor
type DataStore struct{}

type ElementTierMap map[string]int

// Simpan recipes ke CSV
func (ds *DataStore) SaveToCSV(recipes []*model.Recipe, csvPath string) error {

	// Setup file CSV
	file, err := os.Create(csvPath)
	if err != nil {
		log.Printf("Gagal membuat file %s: %v \n", csvPath, err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Tulis header CSV
	if err := writer.Write([]string{"Element", "Item1", "Item2"}); err != nil {
		return err
	}

	// Tulis data ke CSV
	for _, r := range recipes {
		record := []string{r.Element, r.Item1, r.Item2}
		if err := writer.Write(record); err != nil {
			return err
		}
		// fmt.Println("Saved in CSV: ", r.Element)
	}

	fmt.Println("Data tersimpan di", csvPath)
	return nil
}

// Simpan recipes ke DB
func (ds *DataStore) SaveToDB(recipes []*model.Recipe, dbPath string) error {

	// Setup SQLite database
	db, err := sql.Open("sqlite", dbPath+"?_journal=OFF&_sync=OFF&_locking_mode=EXCLUSIVE")
	if err != nil {
		log.Printf("Gagal membuat file %s: %v \n", dbPath, err)
		return err
	}
	defer db.Close()

	// Buat tabel database (jika belum ada)
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS elements (
			element TEXT,
			item1 TEXT,
			item2 TEXT
		);
	`); err != nil {
		return err
	}

	// Hapus tabel lama
	if _, err := db.Exec("DELETE FROM elements"); err != nil {
		log.Printf("Gagal menghapus data lama: %v\n", err)
		return err
	}

	// Tulis data ke database
	for _, r := range recipes {
		_, err := db.Exec(
			"INSERT INTO elements (element, item1, item2) VALUES (?, ?, ?)",
			r.Element, r.Item1, r.Item2,
		)
		if err != nil {
			log.Printf("Gagal memasukkan data ke SQLite %s: %v \n", r.Element, err)
			continue
		}
		// fmt.Println("Saved in SQLite Database: ", r.Element)
	}

	fmt.Println("Data tersimpan di", dbPath)
	return nil
}

// Simpan tiers ke json
func (ds *DataStore) SaveMap(tiers ElementTierMap, jsonPath string) error {

	// Setup file json
	file, err := os.Create(jsonPath)
	if err != nil {
		log.Printf("Gagal membuat file %s: %v \n", jsonPath, err)
		return err
	}
	defer file.Close()

	// Tulis map sebagai json
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	err = encoder.Encode(tiers)
	if err != nil {
		return err
	}

	fmt.Println("Data tersimpan di", jsonPath)
	return nil
}
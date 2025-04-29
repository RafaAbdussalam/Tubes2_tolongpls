package scraper

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	store *DataStore 
}

// Consructor
func NewScraper() *Scraper {
	return &Scraper{
		store: &DataStore{},
	}
}

func (ws *Scraper) Scrape(htmlPath, csvPath, dbPath string) error {

	// Open file HTML lokal
	file, err := os.Open(htmlPath)
	if err != nil {
		log.Printf("Gagal membuka file HTML: %v", err)
		return err
	}
	defer file.Close()

	// Buat dokument dari file reader
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Printf("Gagal membuat dokumen dari HTML: %v", err)
		return err
	}

	// Ambil recipe
	recipes, err := ParseHTML(doc)
	if err != nil {
		return err
	}

	// Simpan ke CSV
	if err := ws.store.SaveToCSV(recipes, csvPath); err != nil {
		return err
	}
	
	// Simpan ke SQLite database
	if err := ws.store.SaveToDB(recipes, dbPath); err != nil {
		return err
	}

	return nil
}
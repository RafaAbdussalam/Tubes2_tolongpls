package main

import (
	"little_alchemy_backend/internal/scraper"
	"log"
)

	func main() {
		scraper := scraper.NewScraper()
		
		err := scraper.Scrape(
			"data/Elements (Little Alchemy 2) _ Little Alchemy Wiki _ Fandom.html",
			"data/alchemy.csv",
			"data/alchemy.db",
		)
		
		if err != nil {
			log.Fatalf("Scraping gagal: %v", err)
		}
		
		log.Println("Scraping selesai")
	}
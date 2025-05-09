package main

import (
	"flag"
	"fmt"
	"little_alchemy_backend/internal/repo"
	"little_alchemy_backend/internal/tree"
	"strings"
)

func main() {
	// Path ke database atau file CSV
	dbPath := flag.String("db", "./data/alchemy.db", "Path ke database SQLite")
	csvPath := flag.String("csv", "./data/alchemy.csv", "Path ke file CSV")
	
	// Parameter untuk elemen target
	elementsFlag := flag.String("elements", "", "Daftar elemen target yang dipisahkan koma (contoh: Brick,Life,Human)")
	
	// Parse flag
	flag.Parse()
	
	fmt.Println("=== PENGUJIAN ALGORITMA DFS UNTUK RESEP TUNGGAL ===")
	
	// Inisialisasi repositori
	repository, err := repo.NewRepository(*dbPath, *csvPath)
	if err != nil {
		fmt.Printf("Error: Gagal membuat repositori: %v\n", err)
		return
	}
	
	// Tentukan elemen target
	var targetElements []string
	
	if *elementsFlag == "" {
		// Gunakan elemen target default
		targetElements = []string{"Brick"}
		fmt.Println("Menggunakan elemen target default:", strings.Join(targetElements, ", "))
	} else {
		// Gunakan elemen target kustom
		targetElements = strings.Split(*elementsFlag, ",")
		
		// Hapus spasi di awal dan akhir
		for i := range targetElements {
			targetElements[i] = strings.TrimSpace(targetElements[i])
		}
		
		fmt.Println("Menggunakan elemen target kustom:", strings.Join(targetElements, ", "))
	}
	
	// Jalankan pengujian DFS untuk resep tunggal
	// Perhatikan bahwa kita harus menyesuaikan parameter repository
	// sesuai dengan definisi fungsi TestDFSSingle
	
	// Jika TestDFSSingle menerima pointer:
	// tree.TestDFSSingle(&repository, targetElements)
	
	// Jika TestDFSSingle menerima nilai:
	tree.TestDFSSingle(*repository, targetElements)
	
	// Catatan: Gunakan salah satu dari dua baris di atas,
	// bergantung pada definisi fungsi TestDFSSingle
}
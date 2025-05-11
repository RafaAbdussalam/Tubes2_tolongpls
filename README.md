# Tugas Besar 2 | Strategi Algoritma (IF2211)

Ini adalah Tugas Besar 2 mata kuliah Strategi Algoritma (IF2211) yang berupa aplikasi berbasis web untuk membantu menyelesaikan permainan Little Alchemy 2. Aplikasi mencari _recipe_ semua elemen-elemen yang ada pada Little Alchemy 2 menggunakan strategi _Breadth First Search_ (BFS) maupun _Depth First Search_ (DFS). Secara teknis, ini melibatkan pembentukan pohon dinamis dari suatu elemen dengan BFS dan DFS untuk merepresentasikan _recipe_ dari suatu elemen.

## BFS dan DFS

_Breadth First Search_ (BFS) adalah strategi traversal di mana setiap node pada suatu level ditelusuri sepenuhnya sebelum melanjutkan ke level berikutnya. _Depth First Search_ (DFS) adalah strategi traversal di mana penelusuran bergerak dari akar hingga node yang sedekatnya dengan tujuan sebelum kembali menelusuri node lainnya. Kedua BFS dan DFS dapat digunakan untuk membentuk pohon ruang status dari suatu persoalan secara dinamis, dengan simpulnya merepresentasikan suatu keadaan pada persoalan. Dalam persoalan mencari _recipe_ elemen di Little Alchemy 2, dapat dibangun suatu pohon dinamis mulai dari suatu elemen untuk merepresentasikan suatu _recipe_, dengan setiap simpulnya adalah suatu elemen pembangun dan daun adalah elemen primer. Jadi, pendekatan BFS dan DFS dapat digunakan menyelesaikan masalah ini: BFS akan membangun simpul elemen berdasarkan tier, sedangkan DFS akan membangun pohon menuju simpul daun.

## Persyaratan Program

Aplikasi menggunakan bahasa JavaScript dengan _framework_ React,js untuk _frontend_ dan bahasa Go untuk _backend_, sehingga diperlukan hal berikut.

1. ...
2. ...

## Menjalankan Program

Berikut langkah-langkah untuk menjalankan program.

1. Clone repository ini.

   ```sh
   git clone https://github.com/RafaAbdussalam/Tubes2_tolongpls
   ```

2. Masuk direktori `Tubes2_tolongpls`.

   ```sh
   cd Tubes2_tolongpls
   ```

3. Jalankan ...

   ```sh
   bla bla bla
   ```

4. ...

   ```sh
   bla bla bla
   ```

5. ...

```sh
bla bla bla
```

## Struktur Program

Berikut struktur dari aplikasi ini.

```
├── doc/
│   └── tolong pls.pdf
├── src/
│   ├── backend
│   │   ├── cmd
│   │   │   ├── api/main.go
│   │   │   └── scraper/main.go
│   │   ├── data
│   │   │   ├── alchemy.csv
│   │   │   ├── alchemy.db
│   │   │   └── tiers.json
│   │   ├── internal
│   │   │   ├── handler/http.go
│   │   │   ├── model
│   │   │   │   ├── data-structures.go
│   │   │   │   ├── recipe.go
│   │   │   │   ├── tree-node.go
│   │   │   │   └── tree.go
│   │   │   ├── repo/repo.go
│   │   │   ├── scraper
│   │   │   │   ├── parser.go
│   │   │   │   ├── scrape.go
│   │   │   │   └── store.go
│   │   │   └── tree
│   │   │       ├── bfs-builder.go
│   │   │       ├── builder.go
│   │   │       └── dfs-builder.go
│   │   ├── go.mod
│   │   └── go.sum
│   ├── frontend
│   │   ├── public
│   │   ├── src
│   │   │   ├── components
│   │   │   │   ├── RecipeResults
│   │   │   │   │   ├── RecipeTree
│   │   │   │   │   ├── ResetButton
│   │   │   │   │   └── StatsPanel
│   │   │   │   └── SearchControls
│   │   │   │       ├── AlgorithmToggle
│   │   │   │       ├── MaxPathsInput
│   │   │   │       ├── ModeToggle
│   │   │   │       └── SearchBar
│   │   │   ├── pages
│   │   │   ├── styles
│   │   │   ├── utils
│   │   │   ├── App.js
│   │   │   └── index.js
│   │   ├── package.json
│   │   └── package-lock.json
├── Makefile
└── README.md
```

## How It Works

Jelasin di sini...

## Author

1. [Timothy Niels Ruslim](https://github.com/timoruslim) (10123053)
2. [Rafa Abdussalam Danadyaksa](https://github.com/RafaAbdussalam) (13523133)
3. [Muhammad Nazih Najumudin](https://github.com/nazihstei) (13523144)

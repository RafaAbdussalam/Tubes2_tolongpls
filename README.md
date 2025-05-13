# Tubes2_tolongpls

🎯 **Tugas Besar 2 | Stategi Algoritma**  
📌 **Institut Teknologi Bandung, Informatika, K-3, Kelompok tolong pls**  

---

## **Tentang Proyek Ini**  
Proyek ini merupakan bagian dari Tugas Besar 2 mata kuliah Strategi Algoritma, yang mengharuskan mahasiswa untuk mengembangkan sebuah sistem pencarian berbasis Depth-First Search (DFS) guna membangun pohon resep (recipe tree) pada permainan Little Alchemy 2.

### BFS dan DFS
_Breadth First Search_ (BFS) adalah strategi traversal di mana setiap node pada suatu level ditelusuri sepenuhnya sebelum melanjutkan ke level berikutnya. _Depth First Search_ (DFS) adalah strategi traversal di mana penelusuran bergerak dari akar hingga node yang sedekatnya dengan tujuan sebelum kembali menelusuri node lainnya. Kedua BFS dan DFS dapat digunakan untuk membentuk pohon ruang status dari suatu persoalan secara dinamis, dengan simpulnya merepresentasikan suatu keadaan pada persoalan. Dalam persoalan mencari _recipe_ elemen di Little Alchemy 2, dapat dibangun suatu pohon dinamis mulai dari suatu elemen untuk merepresentasikan suatu _recipe_, dengan setiap simpulnya adalah suatu elemen pembangun dan daun adalah elemen primer. Jadi, pendekatan BFS dan DFS dapat digunakan menyelesaikan masalah ini: BFS akan membangun simpul elemen berdasarkan tier, sedangkan DFS akan membangun pohon menuju simpul daun.

### **Apa Itu Game Little Alchemy 2?**  
Little Alchemy 2 merupakan permainan berbasis web / aplikasi yang dikembangkan oleh Recloak yang dirilis pada tahun 2017, permainan ini bertujuan untuk membuat 720 elemen dari 4 elemen dasar yang tersedia yaitu air, earth, fire, dan water. Permainan ini merupakan sekuel dari permainan sebelumnya yakni Little Alchemy 1 yang dirilis tahun 2010.

## Komponen-komponen dari permainan ini antara lain:
✔️ **Elemen dasar**
Dalam permainan Little Alchemy 2, terdapat 4 elemen dasar yang tersedia yaitu water, fire, earth, dan air,
4 elemen dasar tersebut nanti akan di-combine menjadi elemen turunan yang berjumlah 720 elemen.

✔️ **Elemen turunan**
Terdapat 720 elemen turunan yang dibagi menjadi beberapa tier tergantung tingkat kesulitan dan banyak langkah yang harus dilakukan.
Setiap elemen turunan memiliki recipe yang terdiri atas elemen lainnya atau elemen itu sendiri.

✔️ **Combine Mechanism**
Untuk mendapatkan elemen turunan pemain dapat melakukan combine antara 2 elemen untuk menghasilkan elemen baru. Elemen turunan yang telah didapatkan dapat digunakan kembali oleh pemain untuk membentuk elemen lainnya.

---

## 🚀 **Fitur Program**  
✔️ **Algoritma** pencarian recipe elemen dalam permainan Little Alchemy 2 dengan menggunakan strategi BFS dan DFS.    

---

## 🛠 **Cara Menjalankan Program**  

### 1️⃣ **Persiapan Awal**  
Pastikan Anda memiliki: 

- ✅ **Docker** `https://www.docker.com/products/docker-desktop/`
- ✅  atau **Go** dan **Node.js** jika ingin menjalankan dev mode

### 2️⃣ **Clone Repository**  
```sh
git clone https://github.com/RafaAbdussalam/Tubes2_tolongpls.git
```

### 3️⃣ **Masuk Direktori**  
```sh
cd .\Tubes2_tolongpls\
```

### 4️⃣ **Compile dan Jalankan Program**  
```sh
docker-compose up
```

Program dimulai ! 

---

## How It Works
Baca aja pdf nya yaa...

---

## **Tampilan Permainan**  
![Gambar 1 : Landing Page ](https://github.com/user-attachments/assets/d3240b95-6639-439a-be80-cb02e3945611)
![Gambar 2 : BFS Tree ](https://github.com/user-attachments/assets/de9c97bf-5aeb-41f8-9761-51835ff1c1af)
![Gambar 3 : DFS Tree ](https://github.com/user-attachments/assets/d6ee63d3-e1a1-4144-b98d-fba39754b676)

---

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

---

## 👨‍💻 **Tim Pengembang**  

<p align="center">
  <table>
    <tr align="center">
      <td>
        <img src="https://github.com/timoruslim.png" width="100" height="100"><br>
        <b>Timothy Niels Ruslim</b><br>
        10123053
      </td>
      <td>
        <img src="https://github.com/RafaAbdussalam.png" width="100" height="100"><br>
        <b>Rafa Abdussalam Danadyaksa</b><br>
        13523133
      </td>
      <td>
        <img src="https://github.com/nazihstei.png" width="100" height="100"><br>
        <b>Muhammad Nazih Najmudin</b><br>
        13523144
      </td>
    </tr>
  </table>
</p>

---

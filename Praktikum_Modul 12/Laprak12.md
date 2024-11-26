# <h1 align="center">Laporan Praktikum Alpro 2</h1>
<p align="center">Arvel Mahsa Athallah Firdaus/2311102188/IF-11-01</p>

# Selection Sort
## No 1
![Selection sort no 1](https://github.com/user-attachments/assets/a7b89d25-7d0e-4883-9890-7e43ef14f0e0)

### Source Code 
```go
// Selection Sort
// No 1
package main

import (
	"fmt"
)

func rumahkerabat(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari indeks elemen terkecil di sisa array
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil dengan elemen di indeks i
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukan banyak daerah : ")
	fmt.Scan(&n)

	// Simpan semua data rumah kerabat di slice dua dimensi
	allRegions := make([][]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukan jumlah kerabat : ")
		var m int
		fmt.Scan(&m)

		// Baca nomor rumah
		arr := make([]int, m)
		fmt.Print("Masukan nomor kerabat : ")
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// Sort array menggunakan Selection Sort
		rumahkerabat(arr)
		allRegions[i] = arr
	}

	// Cetak hasil setelah semua pengurutan selesai
	fmt.Println("Cetak Hasil Setelah Pengurutan")
	for i, region := range allRegions {
		fmt.Printf("Daerah %d : ", i+1)
		for _, num := range region {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```
### Output 
![Output selection no 1](https://github.com/user-attachments/assets/3a9053b9-98b1-452d-b0b7-ffd4d0590cae)

## No 2
![Selection Sort soal no 2](https://github.com/user-attachments/assets/cd9fb1ee-b7e0-412a-9853-69e7e8ab5fa7)

### Source Code 
```go
// No 2
package main

import (
	"fmt"
)

// Selection Sort untuk urutkan dari kecil ke besar
func rumahkerabatAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Selection Sort untuk urutkan dari besar ke kecil
func rumahkerabatDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&n)

	// Simpan semua data rumah kerabat di slice dua dimensi
	allRegionsOdd := make([][]int, n) // Untuk nomor ganjil
	allRegionsEven := make([][]int, n) // Untuk nomor genap

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah %d: ", i+1)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah %d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// Pisahkan nomor ganjil dan genap
		var odd, even []int
		for _, num := range arr {
			if num%2 == 0 {
				even = append(even, num)
			} else {
				odd = append(odd, num)
			}
		}

		// Urutkan ganjil secara menurun dan genap secara menaik
		rumahkerabatDesc(even)
		rumahkerabatAsc(odd)

		// Simpan hasil ke slice terpisah
		allRegionsOdd[i] = odd
		allRegionsEven[i] = even
	}

	// Cetak hasil setelah semua pengurutan selesai
	fmt.Println("\nHasil setelah pengurutan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Daerah %d (terurut): ", i+1)
		for _, num := range allRegionsOdd[i] {
			fmt.Printf("%d ", num)
		}
		for _, num := range allRegionsEven[i] {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```
### Output 
![Output no 2 Selection Sort](https://github.com/user-attachments/assets/69100c26-e124-4213-b121-2713ae9db63c)

# Insertion Sort
## No 1
![Insertion Sort Soal No 1](https://github.com/user-attachments/assets/56c7b2a6-ac6f-47be-abe2-ba7df03a35ce)

### Source Code 
```go
// Insertion Sort
//No 1
package main
import "fmt"

type arrInt [4321]int

func insertionSort1(T *arrInt, n int) {
	var temp, i, j int

	for i = 1; i < n; i++ {
		temp = (*T)[i]
		j = i

		for j > 0 && temp < (*T)[j-1] {
			T[j] = T[j-1]
			j--
		}

		T[j] = temp
	}
}

func distance(T *arrInt, n int) {
	var padaOra = true
	var distance = T[1] - T[0]
	for i := 1 ; i < n - 1; i++{
		if T[i+1] - T[i] != distance{
			padaOra = false
			break
		}
	}
	if padaOra == true {
		fmt.Printf("Data Berjarak %d ", distance)
	} else {
		fmt.Println("Data Berjarak Tidak Tepat")
	}

}

func main() {
	var bil,n int
	var T arrInt
	var i = 0
	for {
		fmt.Print("Input bilangan bunder yang diakhiri oleh bilangan negatif : ")
		fmt.Scanln(&bil)
		if bil < 0 {
			break
		}
		T[i] = bil
		i++
		n++
	}
	insertionSort1(&T, n)
	fmt.Print("Data Setelah Di urutkan : ")
	for i := 0 ; i < n ; i++ {
		fmt.Print(T[i]," ")
	}
	distance(&T, n)
}

```
### Output 
![Output no 1 Insertion sort](https://github.com/user-attachments/assets/37d00fdc-bcc1-4a27-83d7-e235fd8d4e2b)

## No 2
![Insertion Sort Soal No 2 part 1](https://github.com/user-attachments/assets/267c0f07-cdd3-4803-b9fc-73250ee8ff83)

![Insertion Sort Soal No 2 part 2](https://github.com/user-attachments/assets/06bd6ee5-f78e-4da5-b169-4ced11cc6969)

### Source Code 
```go
// No2
package main

import "fmt"

const nMax int = 7919

type Bukuid struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku = [nMax]Bukuid

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\nMasukan ID buku : ")
		fmt.Scan(&pustaka[i].id)

		fmt.Print("Masukan judul buku : ")
		fmt.Scan(&pustaka[i].judul)

		fmt.Print("Masukan penulis buku : ")
		fmt.Scan(&pustaka[i].penulis)

		fmt.Print("Masukan penerbit buku : ")
		fmt.Scan(&pustaka[i].penerbit)

		fmt.Print("Masukan eksemplar buku : ")
		fmt.Scan(&pustaka[i].eksemplar)

		fmt.Print("Masukan tahun buku : ")
		fmt.Scan(&pustaka[i].tahun)

		fmt.Print("Masukan rating buku : ")
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerFavorit(pustaka *DaftarBuku, n int) {
	indexMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[indexMax].rating {
			indexMax = i
		}
	}
	fmt.Println("\nBuku dengan rating terduwur : ", pustaka[indexMax].judul)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Bukuid

	for i := 1; i < n; i++ {
		temp = (*pustaka)[i]
		j := i

		for j > 0 && temp.rating > (*pustaka)[j-i].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func CetakTerbaru(pustaka *DaftarBuku, n int) {
	// laporan 5 judul buku dengan rating tertinggi
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println("\nJudul Berdasarkan Rating Terduwur : ", pustaka[i].judul)
	}
}

func CariBuku(pustaka *DaftarBuku, n int, r int) {
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high/2)

		if pustaka[mid].rating == r {
			fmt.Println(">>> Data Buku <<<\n")
			fmt.Println("Judul Buku   : ", pustaka[mid].judul)
			fmt.Println("Penulis Buku : ", pustaka[mid].penulis)
			fmt.Println("Tahun 	 	  : ", pustaka[mid].tahun)
			fmt.Println("Penerbit 	  : ", pustaka[mid].penerbit)
			fmt.Println("Eksemplar    : ", pustaka[mid].eksemplar)
			fmt.Println("Rating 	  : ", pustaka[mid].rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Data Empty")
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	fmt.Print("Masukan Jumlah Buku Yang Mau didaftarkan : ")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)
	CetakTerFavorit(&pustaka, n)
	UrutBuku(&pustaka, n)
	CetakTerbaru(&pustaka, n)

	fmt.Print("Input Rating Buku Yang Di Golet : ")
	fmt.Scan(&rating)

	CariBuku(&pustaka, n, rating)
}

```
### Output 
![Output Insertion No2](https://github.com/user-attachments/assets/a9c7111a-7c10-4e48-b326-4ee8b41140ae)



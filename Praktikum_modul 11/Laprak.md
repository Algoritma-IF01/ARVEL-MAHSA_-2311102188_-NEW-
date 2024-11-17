# <h1 align="center">Laporan Praktikum Alpro 2</h1>
<p align="center">Arvel Mahsa Athallah Firdaus/2311102188/IF-11-01</p>

# Modul Pencarian Nilai Ekstrim Pada Himpunan Data
## No 2
![soal 11 no 2](https://github.com/user-attachments/assets/c830c8f5-91a5-4cce-9bb9-250c683d8ede)
### Source Code
```go
package main
import "fmt"
func main (){
	// Menginputkan jumlah ikan yang akan dijual (x) dan jumlah ikan per wadah (y)
	var x, y int
	fmt.Print("Inputkan Jumlah ikan (x) dan jumlah ikan perwadah (y) : ")
	fmt.Scan(&x,&y)

	// Menginputkan berat ikan sesuai jumlah ikan (x)
	berat := make([]float64, x)
	fmt.Printf("Berat masing-masing dari %d ikan = \n",x)
	for i := 0; i < x ; i++{
		fmt.Scan(&berat[i])
	}

	// Menghitung total berat setiap wadah
	itungwadah := (x + y - 1) / y // untuk menghitung jumlah wadah
	totalberat := make([]float64,itungwadah)
	for i := 0; i < x ; i++{
		indexWadah := i / y
		totalberat[indexWadah] += berat[i]
	}

	// Menghitung rata rata setiap wadah
	rata_rata_berat := make([]float64, itungwadah)
	for i := 0 ; i < itungwadah ; i++{
		hitung := y

	if i == itungwadah-1 && x%y !=0 {
		hitung = x % y
	}
	rata_rata_berat[i] = totalberat[i] / float64(hitung)
	}

	// Output total berat disetiap wadah
	fmt.Println("\nTotal berat di setiap wadah = ")
	for _, total := range totalberat {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	// Output rata rata di setiap wadah
	fmt.Println("\nTotal berat di setiap wadah = ")
	for _, rata2 := range rata_rata_berat {
		fmt.Printf("%.2f ", rata2)
	}
	fmt.Println()
}
```
### Output
![output no 2](https://github.com/user-attachments/assets/6708fa81-7a2a-4f7b-bef7-3c372bd784e6)

## No 3
![soal 11 no 3](https://github.com/user-attachments/assets/4c18db28-3995-404e-9707-a33cb146b4d1)

### Source Code
```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	// Terdefinisi array dinamis arrBerat
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Menghitung berat minimum dan maksimum dalam array
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	// Menghitung dan mengembalikan rerata berat balita dalam array
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var arrBerat arrBalita
	var n int
	var bMin, bMax float64

	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scan(&n)

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	hitungMinMax(arrBerat, n, &bMin, &bMax)

	rataRata := rerata(arrBerat, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg", bMin)
	fmt.Printf("\nBerat balita maksimum: %.2f kg", bMax)
	fmt.Printf("\nRerata berat balita: %.2f kg", rataRata)
}

```
### Output
![output no 3](https://github.com/user-attachments/assets/5b49da64-63e9-4964-b511-8ddd1aba4f12)



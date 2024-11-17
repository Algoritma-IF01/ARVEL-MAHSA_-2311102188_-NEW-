// No 1
// package main
// import "fmt"

// func main() {
// 	var N int
// 	fmt.Print("masukan jumlah anak kelinci : ")
// 	fmt.Scan(&N)

// 	if N <= 0 || N > 1000 {
// 		fmt.Println("jumlah anak kelinci harus antara 1 dan 1000 ")
// 		return
// 	}

// 	weights := make([]float64, N)
// 	fmt.Println("Masukan berat anak kelinci : ")
// 	for i := 0 ; i < N ; i++{
// 		fmt.Scan(&weights[i])
// 	}

// 	minWeight, maxWeight := weights[0], weights[0]
// 	for _, weight := range weights[1:]{
// 		if weight < minWeight {
// 			minWeight = weight
// 		}
// 		if weight > maxWeight {
// 			maxWeight = weight
// 		}

// 	}

// 	fmt.Printf("Berat kelinci terkecil : %.2f\n", minWeight)
// 	fmt.Printf("Berat kelinci terbesar : %.2f\n", maxWeight)
// }

// No 2
// package main
// import "fmt"
// func main (){
// 	// Menginputkan jumlah ikan yang akan dijual (x) dan jumlah ikan per wadah (y)
// 	var x, y int
// 	fmt.Print("Inputkan Jumlah ikan (x) dan jumlah ikan perwadah (y) : ")
// 	fmt.Scan(&x,&y)

// 	// Menginputkan berat ikan sesuai jumlah ikan (x)
// 	berat := make([]float64, x)
// 	fmt.Printf("Berat masing-masing dari %d ikan = \n",x)
// 	for i := 0; i < x ; i++{
// 		fmt.Scan(&berat[i])
// 	}

// 	// Menghitung total berat setiap wadah
// 	itungwadah := (x + y - 1) / y // untuk menghitung jumlah wadah
// 	totalberat := make([]float64,itungwadah)
// 	for i := 0; i < x ; i++{
// 		indexWadah := i / y
// 		totalberat[indexWadah] += berat[i]
// 	}

// 	// Menghitung rata rata setiap wadah
// 	rata_rata_berat := make([]float64, itungwadah)
// 	for i := 0 ; i < itungwadah ; i++{
// 		hitung := y

// 	if i == itungwadah-1 && x%y !=0 {
// 		hitung = x % y
// 	}
// 	rata_rata_berat[i] = totalberat[i] / float64(hitung)
// 	}

// 	// Output total berat disetiap wadah
// 	fmt.Println("\nTotal berat di setiap wadah = ")
// 	for _, total := range totalberat {
// 		fmt.Printf("%.2f ", total)
// 	}
// 	fmt.Println()

// 	// Output rata rata di setiap wadah
// 	fmt.Println("\nTotal berat di setiap wadah = ")
// 	for _, rata2 := range rata_rata_berat {
// 		fmt.Printf("%.2f ", rata2)
// 	}
// 	fmt.Println()
// }

// No 2
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

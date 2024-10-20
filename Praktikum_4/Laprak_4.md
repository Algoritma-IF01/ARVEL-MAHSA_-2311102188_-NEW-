# <h1 align="center">Laporan Praktikum Alpro 2</h1>
<p align="center">Arvel Mahsa Athallah Firdaus/2311102188/IF-11-01</p>

## Latihan Soal Modul 3 (Fungsi)
### Nomor 1
![Soal Nomor 1 Modul 3](https://github.com/user-attachments/assets/965b966c-4be9-4fae-8a9d-7128e83b4442)
### Source Code
```go
package main

import ("fmt")

// Fungsi untuk menghitung faktorial
func faktorial(n int) int{
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++{
		result *= i
	}
	return result
}

// Fungsi menghitung permutasi
func permutation(n, r int) int{
	return faktorial(n) / faktorial(n-r)
}

// fungsi menghitung kombinasi
func combination(n, r int) int{
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main (){
	var a, b, c, d int

	// Meminta input dari penggunaan
	fmt.Println("Masukan Nilai a, b, c, d : ")
	fmt.Scan(&a,&b, &c,&d)

	// Menghitung permutasi dan kombinasi untuk a terhadap b
	p1 := permutation(a,c)
	c1 := combination(a,c)

	// Menghitung permutasi dan kombinasi untuk c terhadap d
	p2 := permutation(b,d)
	c2 := combination(b,d)

	//Output hasil
	fmt.Printf("P(%d, %d) = %d\n", a, c, p1)
	fmt.Printf("C(%d, %d) = %d\n", a, c, c1)
	fmt.Printf("P(%d, %d) = %d\n", b, d, p2)
	fmt.Printf("C(%d, %d) = %d\n", b, d, c2)
}

```
### Output
![Output Nomor 1 Modul 3](https://github.com/user-attachments/assets/970cf77c-4494-4b72-890f-cde9aa773235)

### Nomor 2
![Soal Nomor 2 Modul 3](https://github.com/user-attachments/assets/895009c3-3640-44d0-908e-611f1a5ae853)
### Source Code
```go
package main
import ("fmt")

// Fungsi untuk fungsi matematika
// f(x) = x^2
func fx(x int)int{
	return x * x
}

// g(x) = x - 2
func gx(x int)int {
	return x - 2
}

// h(x) = x + 1
func hx(x int)int{
	return x + 1
}

// Fungsi komposisi
// f(g(h(x)))
func fogoh(x int)int {
	return fx(gx(hx(x)))
}

// g(h(f(x)))
func gohof(x int)int {
	return gx(hx(fx(x)))
}

// h(f(g(x)))
func hofog(x int)int {
	return hx(fx(gx(x)))
}

func main(){

	//Deklarasi variabel
	var a,b,c int
	fmt.Scan(&a, &b, &c)

	// Definisikan fungsi
	hasilfogoh := fogoh(a)
	hasilgohof := gohof(b)
	hasilhofog := hofog(c)

		// Print Hasil
		fmt.Printf("(fogoh)(%d) = %d", a, hasilfogoh)
		fmt.Printf("\n(gohof)(%d) = %d", b, hasilgohof)
		fmt.Printf("\n(hofog)(%d) = %d", c, hasilhofog)
}
```
### Output
![output nomor 2 Modul 3](https://github.com/user-attachments/assets/7f809f5a-72fa-4d61-8814-1b929ee930b0)

### Nomor 3 
![Soal Nomor 3 Modul 3](https://github.com/user-attachments/assets/7946c97a-2c55-4482-991e-6386ea064160)
### Source Code
```go
package main

import (
 "fmt"
 "math"
)
// Fungsi Untuk Mengecek apakah titik berada di dalam lingkaran atau tidak
func titik_dalam_lingkaran(cx, cy, radius, x, y float64) bool {
 return rumus_jarak(cx, cy, x, y) <= radius
}
// Fungsi untuk menghitung jarak titik (a,b) dan (c,d)
func rumus_jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c)+(b-d)*(b-d))
   }

func main() {
 
// Deklarasi Variabel
 var cx1, cx2, cy1, cy2, radius1, radius2, x, y float64

 // Input titik dari user
 fmt.Scan(&cx1, &cy1, &radius1)
 fmt.Scan(&cx2, &cy2, &radius2)
 fmt.Scan(&x, &y)

 // Pendefinisian apabila titik berada di lingkaran 1 dan 2
 dalam1 := titik_dalam_lingkaran(cx1, cy1, radius1, x, y)
 dalam2 := titik_dalam_lingkaran(cx2, cy2, radius2, x, y)

 if dalam1 && dalam2 {
  fmt.Println("Titik di dalam lingkaran 1 dan 2")
 } else if dalam1 {
  fmt.Println("Titik di dalam lingkaran 1")
 } else if dalam2 {
  fmt.Println("Titik di dalam lingkaran 2")
 } else {
  fmt.Println("Titik di luar lingkaran 1 dan 2")
 }
}
```
### Output
![output nomor 3 Modul 3](https://github.com/user-attachments/assets/f82f6c8c-ddea-410e-8fe3-4b4b7f20e35b)

## Latihan Soal Modul 4 (Prosedur)
### Nomor 1
### Source Code
```go
package main

import ("fmt")

//Fungsi rekrusif untuk menghitung deret fibonacci
func fibonacci(n int)int{
	if n == 0 {
		return 0
	} else if n == 1{
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main (){
	// Menampilkan deret fibonacci hingga suku ke-10
	fmt.Println("Deret Fibonacci hingga suku ke-10 : ")
	for i := 0 ; i <= 10; i++{
		fmt.Printf("Fibonacci(%d) = (%d)\n", i, fibonacci(i))
	}
}
```
### Output
![Output nomor 1 Modul 4](https://github.com/user-attachments/assets/e4dcde1d-b7a7-43b4-a262-1e6fc9af97c3)

### Nomor 2
![Soal Nomor 2 Modul 4](https://github.com/user-attachments/assets/7ddff449-44e8-4c38-952d-2ccbca3798aa)
### Source Code
```go
package main

import (
 "fmt"
)
// Prosedur Untuk Menghitung Skor Peserta
func hitungSkor(soal_dikerjakan *int, score *int) {

// Deklarasi Jumlah Soal
 var s1, s2, s3, s4, s5, s6, s7, s8 int

// Pendefinisian dengan pointer untuk soal dan skor yang diperoleh
// Data yang di inputkan akan masuk kedalam pointer 
 *soal_dikerjakan = 0
 *score = 0

 // Inputan Soal dari pointer akan di Scan. Nominal inputan akan masuk kedalam
 // pointer diatas kemudian akan di scan
 fmt.Scan(&s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8)

 // Array untuk menampung nominal soal
 jumlahSoal := []int{s1, s2, s3, s4, s5, s6, s7, s8}

 // Perulangan untuk mendeteksi jika nominal jika tidak melebihi 301, maka akan terhitung
 // mengerjakan 1 soal
 for a := 0; a < 8; a++ {
  if jumlahSoal[a] < 301 {
   *soal_dikerjakan++ // Untuk mencatat jumlah soal yang dikerjakan
   *score += jumlahSoal[a] // Skor didapat dari penjumlahan nominal soal pada array
  }
 }
}

func main() {

	// Deklarasi Variabel
  var score, soal_dikerjakan int
  var nama, pemenang string
  var soalTerbanyak, scoreKecil int

  // Perulangan untuk menginputkan data nama peserta dan soal yang dikerjakan
 for {
	// Input nama peserta
	// dibawah inputan nama peserta adalah jumlah soal yang dikerjakan
  fmt.Scan(&nama)

  // Jika user menginputkan "selesai" maka program akan berhenti
  if nama == "Selesai" || nama == "selesai" {
   break
  }

  // Dari inputan soal diatas akan dicek dengan prosedur yang dipanggil dibawah ini
  hitungSkor(&soal_dikerjakan, &score)

  // Mencari pemenang dengan jumlah pengerjaan soal terbanyak, jika jumlah soal yang
  // dikerjakan sama, maka akan ditentukan melalui waktu paling sedikit 
  if soal_dikerjakan > soalTerbanyak || (soal_dikerjakan == soalTerbanyak && score < scoreKecil) {
   soalTerbanyak = soal_dikerjakan
   scoreKecil = score
   pemenang = nama
  }
 }

 // Print Pemenang
 fmt.Println(pemenang, soalTerbanyak, scoreKecil)
}

```
### Output
![output nomor 2 Modul 4](https://github.com/user-attachments/assets/c58caba5-8bec-4ac7-82ff-1f611953d816)










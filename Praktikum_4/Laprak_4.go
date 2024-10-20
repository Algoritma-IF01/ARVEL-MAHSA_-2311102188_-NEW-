//Latihan Soal Modul 3
// Nomor 1
// package main

// import ("fmt")

// // Fungsi untuk menghitung faktorial
// func faktorial(n int) int{
// 	if n == 0 || n == 1 {
// 		return 1
// 	}
// 	result := 1
// 	for i := 2; i <= n; i++{
// 		result *= i
// 	}
// 	return result
// }

// // Fungsi menghitung permutasi
// func permutation(n, r int) int{
// 	return faktorial(n) / faktorial(n-r)
// }

// // fungsi menghitung kombinasi
// func combination(n, r int) int{
// 	return faktorial(n) / (faktorial(r) * faktorial(n-r))
// }

// func main (){
// 	var a, b, c, d int

// 	// Meminta input dari penggunaan
// 	fmt.Println("Masukan Nilai a, b, c, d : ")
// 	fmt.Scan(&a,&b, &c,&d)

// 	// Menghitung permutasi dan kombinasi untuk a terhadap b
// 	p1 := permutation(a,c)
// 	c1 := combination(a,c)

// 	// Menghitung permutasi dan kombinasi untuk c terhadap d
// 	p2 := permutation(b,d)
// 	c2 := combination(b,d)

// 	//Output hasil
// 	fmt.Printf("P(%d, %d) = %d\n", a, c, p1)
// 	fmt.Printf("C(%d, %d) = %d\n", a, c, c1)
// 	fmt.Printf("P(%d, %d) = %d\n", b, d, p2)
// 	fmt.Printf("C(%d, %d) = %d\n", b, d, c2)
// }

// Latihan Modul 4
// Nomor 1
// package main

// import ("fmt")

// //Fungsi rekrusif untuk menghitung deret fibonacci
// func fibonacci(n int)int{
// 	if n == 0 {
// 		return 0
// 	} else if n == 1{
// 		return 1
// 	} else {
// 		return fibonacci(n-1) + fibonacci(n-2)
// 	}
// }

// func main (){
// 	// Menampilkan deret fibonacci hingga suku ke-10
// 	fmt.Println("Deret Fibonacci hingga suku ke-10 : ")
// 	for i := 0 ; i <= 10; i++{
// 		fmt.Printf("Fibonacci(%d) = (%d)\n", i, fibonacci(i))
// 	}
// }

// Soal yang buat laprak
// soal latihan modul 3
// nomor 2
// package main
// import ("fmt")

// // Fungsi untuk fungsi matematika
// // f(x) = x^2
// func fx(x int)int{
// 	return x * x
// }

// // g(x) = x - 2
// func gx(x int)int {
// 	return x - 2
// }

// // h(x) = x + 1
// func hx(x int)int{
// 	return x + 1
// }

// // Fungsi komposisi
// // f(g(h(x)))
// func fogoh(x int)int {
// 	return fx(gx(hx(x)))
// }

// // g(h(f(x)))
// func gohof(x int)int {
// 	return gx(hx(fx(x)))
// }

// // h(f(g(x)))
// func hofog(x int)int {
// 	return hx(fx(gx(x)))
// }

// func main(){

// 	//Deklarasi variabel
// 	var a,b,c int
// 	fmt.Scan(&a, &b, &c)

// 	// Definisikan fungsi
// 	hasilfogoh := fogoh(a)
// 	hasilgohof := gohof(b)
// 	hasilhofog := hofog(c)

// 		// Print Hasil
// 		fmt.Printf("(fogoh)(%d) = %d", a, hasilfogoh)
// 		fmt.Printf("\n(gohof)(%d) = %d", b, hasilgohof)
// 		fmt.Printf("\n(hofog)(%d) = %d", c, hasilhofog)
// 	}

// nomor 3
// package main

// import (
//  "fmt"
//  "math"
// )
// // Fungsi Untuk Mengecek apakah titik berada di dalam lingkaran atau tidak
// func titik_dalam_lingkaran(cx, cy, radius, x, y float64) bool {
//  return rumus_jarak(cx, cy, x, y) <= radius
// }
// // Fungsi untuk menghitung jarak titik (a,b) dan (c,d)
// func rumus_jarak(a, b, c, d float64) float64 {
// 	return math.Sqrt((a-c)*(a-c)+(b-d)*(b-d))
//    }

// func main() {
 
// // Deklarasi Variabel
//  var cx1, cx2, cy1, cy2, radius1, radius2, x, y float64

//  // Input titik dari user
//  fmt.Scan(&cx1, &cy1, &radius1)
//  fmt.Scan(&cx2, &cy2, &radius2)
//  fmt.Scan(&x, &y)

//  // Pendefinisian apabila titik berada di lingkaran 1 dan 2
//  dalam1 := titik_dalam_lingkaran(cx1, cy1, radius1, x, y)
//  dalam2 := titik_dalam_lingkaran(cx2, cy2, radius2, x, y)

//  if dalam1 && dalam2 {
//   fmt.Println("Titik di dalam lingkaran 1 dan 2")
//  } else if dalam1 {
//   fmt.Println("Titik di dalam lingkaran 1")
//  } else if dalam2 {
//   fmt.Println("Titik di dalam lingkaran 2")
//  } else {
//   fmt.Println("Titik di luar lingkaran 1 dan 2")
//  }
// }

// Soal Nomor 2 Latihan 4
// nomor 2
// package main

// import (
//  "fmt"
// )
// // Prosedur Untuk Menghitung Skor Peserta
// func hitungSkor(soal_dikerjakan *int, score *int) {

// // Deklarasi Jumlah Soal
//  var s1, s2, s3, s4, s5, s6, s7, s8 int

// // Pendefinisian dengan pointer untuk soal dan skor yang diperoleh
// // Data yang di inputkan akan masuk kedalam pointer 
//  *soal_dikerjakan = 0
//  *score = 0

//  // Inputan Soal dari pointer akan di Scan. Nominal inputan akan masuk kedalam
//  // pointer diatas kemudian akan di scan
//  fmt.Scan(&s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8)

//  // Array untuk menampung nominal soal
//  jumlahSoal := []int{s1, s2, s3, s4, s5, s6, s7, s8}

//  // Perulangan untuk mendeteksi jika nominal jika tidak melebihi 301, maka akan terhitung
//  // mengerjakan 1 soal
//  for a := 0; a < 8; a++ {
//   if jumlahSoal[a] < 301 {
//    *soal_dikerjakan++ // Untuk mencatat jumlah soal yang dikerjakan
//    *score += jumlahSoal[a] // Skor didapat dari penjumlahan nominal soal pada array
//   }
//  }
// }

// func main() {

// 	// Deklarasi Variabel
//   var score, soal_dikerjakan int
//   var nama, pemenang string
//   var soalTerbanyak, scoreKecil int

//   // Perulangan untuk menginputkan data nama peserta dan soal yang dikerjakan
//  for {
// 	// Input nama peserta
// 	// dibawah inputan nama peserta adalah jumlah soal yang dikerjakan
//   fmt.Scan(&nama)

//   // Jika user menginputkan "selesai" maka program akan berhenti
//   if nama == "Selesai" || nama == "selesai" {
//    break
//   }

//   // Dari inputan soal diatas akan dicek dengan prosedur yang dipanggil dibawah ini
//   hitungSkor(&soal_dikerjakan, &score)

//   // Mencari pemenang dengan jumlah pengerjaan soal terbanyak, jika jumlah soal yang
//   // dikerjakan sama, maka akan ditentukan melalui waktu paling sedikit 
//   if soal_dikerjakan > soalTerbanyak || (soal_dikerjakan == soalTerbanyak && score < scoreKecil) {
//    soalTerbanyak = soal_dikerjakan
//    scoreKecil = score
//    pemenang = nama
//   }
//  }

//  // Print Pemenang
//  fmt.Println(pemenang, soalTerbanyak, scoreKecil)
// }

// Nomor 4
package main

import (
 "fmt"
)

// Prosedur untuk mencetak deret
func cetakDeret(bil int) {
	// Inputan tidak boleh lebih dari 1jt
 if bil > 1000000 {
  fmt.Println("Bilangan harus lebih kecil dari 1000000")
 }
 fmt.Print(bil, " ")

 // Perulangan untuk mencetak deret
 for {
  if bil == 1 {
   break
  }

  // Percabangan digunakan untuk mendeteksi bilangan yang genap atau ganjil
  // apabila bilangan yang di inputkan genap, maka rumus yang digunakan
  // untuk mencari deret selanjutnya adalah "bilangan dibagi dengan 2"
  if bil%2 == 0 {
   bil = bil / 2
   fmt.Print(bil, " ")
  } else { // jika ganjil maka rumus nya adalah "bilangan dikali 3 ditambah 1"
   bil = 3*bil + 1
   fmt.Print(bil, " ")
  }
 }
}

func main() {

// Deklarasi variabel
 var bil int

 // Meminta Inputan dari pengguna
 fmt.Scanln(&bil)
 cetakDeret(bil)
 
}

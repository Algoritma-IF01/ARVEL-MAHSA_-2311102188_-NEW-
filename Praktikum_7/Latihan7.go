// No 1
// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Definisi tipe bentukan untuk titik
// type Titik struct {
// 	x int
// 	y int
// }

// // Definisi tipe bentukan untuk lingkaran
// type Lingkaran struct {
// 	center Titik
// 	radius int
// }

// // Fungsi untuk menghitung jarak antara dua titik
// func jarak(p Titik, q Titik) float64 {
// 	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
// }

// // Fungsi untuk menentukan apakah titik berada di dalam lingkaran
// func didalam(c Lingkaran, p Titik) bool {
// 	return jarak(p, c.center) < float64(c.radius)
// }

// func main() {
// 	var (
// 		// Mengambil input untuk lingkaran 1
// 		lingkaran1 Lingkaran
// 		// Mengambil input untuk lingkaran 2
// 		lingkaran2 Lingkaran
// 		// Mengambil input untuk titik sembarang
// 		point Titik
// 	)

// 	// Input untuk lingkaran 1 (cx, cy, r)
// 	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1 (cx cy r):")
// 	fmt.Scan(&lingkaran1.center.x, &lingkaran1.center.y, &lingkaran1.radius)

// 	// Input untuk lingkaran 2 (cx, cy, r)
// 	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2 (cx cy r):")
// 	fmt.Scan(&lingkaran2.center.x, &lingkaran2.center.y, &lingkaran2.radius)

// 	// Input untuk titik sembarang (x, y)
// 	fmt.Println("Masukkan koordinat titik sembarang (x y):")
// 	fmt.Scan(&point.x, &point.y)

// 	// Mengecek posisi titik terhadap kedua lingkaran
// 	inLingkaran1 := didalam(lingkaran1, point)
// 	inLingkaran2 := didalam(lingkaran2, point)

// 	if inLingkaran1 && inLingkaran2 {
// 		fmt.Println("Titik di dalam lingkaran 1 dan 2")
// 	} else if inLingkaran1 {
// 		fmt.Println("Titik di dalam lingkaran 1")
// 	} else if inLingkaran2 {
// 		fmt.Println("Titik di dalam lingkaran 2")
// 	} else {
// 		fmt.Println("Titik di luar lingkaran 1 dan 2")
// 	}
// }

// // NO 2
// package main
// import ("fmt"
// "math"
// )
// func main (){
// 	var index int
// 	fmt.Print("Inputkan jumlah index : ")
// 	fmt.Scan(&index)

// 	// membuat array
// 	array := make([]int, index)

// 	// menginputkan angka sesuai dengan jumlah index yang di tentukan
// 	// kemudian angka disimpan didalam array
// 	fmt.Print("Inputkan angka : ")
// 	for i := 0 ; i < index ; i++{
// 		fmt.Scan(&array[i])
// 	}

// 	for {
// 		// menu operasi
// 		fmt.Println("\nOperasi")
// 		fmt.Println("1. Menampilkan Keseluruhan Array ")
// 		fmt.Println("2. Menampilkan Elemen Array Dengan Index Ganjil")
// 		fmt.Println("3. Menampilkan Elemen Array Dengan Index Genap ")
// 		fmt.Println("4. Menampilkan Elemen Array Dengan Dengan Index Kelipatan Bilangan X")
// 		fmt.Println("5. Menghapus Elemen Array Pada Index Tertentu")
// 		fmt.Println("6. Menampilkan Rata - Rata")
// 		fmt.Println("7. Menampilkan Standar Deviasi Atau Simpangan Baku")
// 		fmt.Println("8. Menampilkan Frekuensi Dari Suatu Bilangan Tertentu")
// 		var pilih string
// 		fmt.Println("\nInput Nomor Untuk Pilih Operasi : ")
// 		fmt.Scan(&pilih)

// 		switch pilih {
// 		case "1" : 
// 		fmt.Println("Elemen Array : ", array[:index])

// 		case "2" :
// 		fmt.Println("Elemen Array Dengan Index Ganjil : ")
// 		// index i diberi value 1, kurang dari jumlah index yang di inputkan, i adalah hasil penjumlahan dari 2
// 		// dan menampilkan bilangan dengan index ganjil
// 		for i:= 1 ; i < index ; i += 2 {
// 			fmt.Print(array[i], " ")
// 		}
// 		fmt.Println()

// 		case "3" : 
// 		// index i diberi value 0, kurang dari jumlah index yang di inputkan, i adalah hasil penjumlahan dari 2
// 		// supaya dapat menampilkan bilangan dengan index genap
// 		fmt.Println("Elemen Array Dengan Index Genap : ")
// 		for i := 0 ; i < index ; i+=2 {
// 			fmt.Print(array[i]," ")
// 		}
// 		fmt.Println()	

// 		case "4" :
// 			var x int
// 			fmt.Print("Inputkan Bilangan X : ")
// 			fmt.Scan(&x)
// 			fmt.Printf("Elemen Dengan Index Kelipatan %d : ", x)
// 			for i := x; i < index; i += x {
// 				fmt.Print(array[i], " ")
// 			}
// 			fmt.Println()

// 		case "5" :
// 			var n int
// 			fmt.Print("Inputkan Index Yang Ingin Dihapus : ")
// 			fmt.Scan(&n)
// 			if n >= 0 && n < index {
// 				for i := n; i < index-1; i++ {
// 					array[i] = array[i+1]
// 				}
// 				index-- // Hapus Array
// 				fmt.Println("Array Terkini :", array[:index])
// 			} else {
// 				fmt.Println("Index Tidak Valid")
// 			}

// 		case "6" :
// 			total := 0
// 			for _, nilai := range array[:index] {
// 				total += nilai
// 			}
// 			fmt.Printf("Rata-rata : %.2f\n", float64(total)/float64(index))

// 		case "7" :
// 			total, totalsq := 0, 0.0
// 			for _, nilai := range array[:index] {
// 				total += nilai // Hasil Penjumlahan Data Dalam Array
// 				totalsq += float64(nilai* nilai) // Hasil Penjumlahan Di kuadratkan
// 			}

// 			rata_rata := float64(total) / float64(index)
// 			// Rumus Mencari Simpangan Baku : hasil penjumlahan data pangkat dua dibagi jumlah data
// 			// dikurangi hasil kali dari rata rata
// 			stndrDev := math.Sqrt(totalsq/float64(index) - rata_rata*rata_rata)
// 			fmt.Printf("Standard deviasi / Simpangan Baku : %.2f\n", stndrDev)

// 		case "8" :
// 			frekuensi := make(map[int]int)
// 			for i := 0; i < index; i++ {
// 				frekuensi[array[i]]++
// 			}
// 			fmt.Println("Frekuensi dari suatu bilangan adalah :")
// 			for nilai, hitung := range frekuensi {
// 				fmt.Printf("%d: %d kali\n", nilai, hitung)
// 			}
// 		default:
// 			fmt.Println("opsi tidak ada")
// 			return
// 		}
// 	}
// }

// NO 3
// package main
// import "fmt"
// func main (){
// 	var club1, club2 string
// 	fmt.Print("Klub 1 : ")
// 	fmt.Scan(&club1)
// 	fmt.Print("Klub 2 : ")
// 	fmt.Scan(&club2)
// 	fmt.Println()

// 	//Array untuk menyimpan pemenang
// 	var pemenang[100]string
// 	hitung_pemenang := 0
// 	hitung_pertandingan := 1

// 	for{
// 		var skor_club1, skor_club2 int
// 		fmt.Printf("Pertandingan %d : ", hitung_pertandingan)
// 		fmt.Scan(&skor_club1,&skor_club2)

// 		if skor_club1 < 0 || skor_club2 < 0 {
// 			break
// 		}
// 	 // Jika Inputan Skor dibawah nol atau negatif maka program akan dihentikan
// 		if skor_club1 > skor_club2 {
// 			pemenang[hitung_pemenang] = fmt.Sprintf("Hasil %d : %s", hitung_pertandingan, club1)
// 		} else if skor_club2 > skor_club1 {
// 			pemenang[hitung_pemenang] = fmt.Sprintf("Hasil %d : %s", hitung_pertandingan, club2)
// 		} else {
// 			pemenang[hitung_pemenang] = fmt.Sprintf("Hasil %d : Draw", hitung_pertandingan)
// 		}
// 		hitung_pemenang++
// 		hitung_pertandingan++
// 	}
// 	for i := 0 ; i < hitung_pemenang; i++{
// 		fmt.Println(pemenang[i])
// 	}
// 	fmt.Println("Pertandingan Selesai")
// }

// NO 4
// package main
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )
// const NMAX int = 127
// type Tabel [NMAX]rune
// func isiArray(t *Tabel, n *int) {
//     // Data Tersedia dalam piranti masukan
//     // Array t berisi sejumlah n karakter yang di inputkan user. 
//     // Proses input selama karakter bukan TITIK (".") dan n <= NMAX
// 	fmt.Print("Teks	: ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	text := scanner.Text()

// 	for _, ch := range text {
// 		if ch == '.' || *n >= NMAX {
// 			break
// 		}
// 		t[*n] = ch
// 		(*n)++
// 	}
// }
// func cetakArray(t Tabel, n int) {
//     // terdefinisi array t yang berisi sejumlah n karakter
//     // n karakter array muncul dilayar 
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("%c", t[i])
// 	}
// 	fmt.Println()
// }
// func balikanArray(t *Tabel, n int) {
//     // Terdefinisi Array t yang berisi sejumlah n karakter 
//     // Urutan isi array t terbalik
// 	for i := 0; i < n/2; i++ {
// 		t[i], t[n-1-i] = t[n-1-i], t[i]
// 	}
// }
// func palindrom(t Tabel, n int) bool {
//     // Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom
//     // dan false apabila sebaliknya. Manfaatkan prosedur balikanArray 
// 	for i := 0; i < n/2; i++ {
// 		if t[i] != t[n-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var t Tabel
// 	var n int

// 	isiArray(&t, &n)

// 	balikanArray(&t, n)
// 	fmt.Print("Reverse teks	: ")
// 	cetakArray(t, n)

// 	cek := palindrom(t, n)
// 	fmt.Println("Palindrom	?", cek)

// }




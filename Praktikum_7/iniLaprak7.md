# <h1 align="center">Laporan Praktikum Alpro 2</h1>
<p align="center">Arvel Mahsa Athallah Firdaus/2311102188/IF-11-01</p>

# Latihan Soal 7 
## No 2
![Latihan7 nomor 2](https://github.com/user-attachments/assets/d959132f-9c7a-4bdd-aea3-40bf1a74fb93)

### Code Program
```go
package main
import ("fmt"
"math"
)
func main (){
	var index int
	fmt.Print("Inputkan jumlah index : ")
	fmt.Scan(&index)

	// membuat array
	array := make([]int, index)

	// menginputkan angka sesuai dengan jumlah index yang di tentukan
	// kemudian angka disimpan didalam array
	fmt.Print("Inputkan angka : ")
	for i := 0 ; i < index ; i++{
		fmt.Scan(&array[i])
	}

	for {
		// menu operasi
		fmt.Println("\nOperasi")
		fmt.Println("1. Menampilkan Keseluruhan Array ")
		fmt.Println("2. Menampilkan Elemen Array Dengan Index Ganjil")
		fmt.Println("3. Menampilkan Elemen Array Dengan Index Genap ")
		fmt.Println("4. Menampilkan Elemen Array Dengan Dengan Index Kelipatan Bilangan X")
		fmt.Println("5. Menghapus Elemen Array Pada Index Tertentu")
		fmt.Println("6. Menampilkan Rata - Rata")
		fmt.Println("7. Menampilkan Standar Deviasi Atau Simpangan Baku")
		fmt.Println("8. Menampilkan Frekuensi Dari Suatu Bilangan Tertentu")
		var pilih string
		fmt.Println("\nInput Nomor Untuk Pilih Operasi : ")
		fmt.Scan(&pilih)

		switch pilih {
		case "1" : 
		fmt.Println("Elemen Array : ", array[:index])

		case "2" :
		fmt.Println("Elemen Array Dengan Index Ganjil : ")
		// index i diberi value 1, kurang dari jumlah index yang di inputkan, i adalah hasil penjumlahan dari 2
		// dan menampilkan bilangan dengan index ganjil
		for i:= 1 ; i < index ; i += 2 {
			fmt.Print(array[i], " ")
		}
		fmt.Println()

		case "3" : 
		// index i diberi value 0, kurang dari jumlah index yang di inputkan, i adalah hasil penjumlahan dari 2
		// supaya dapat menampilkan bilangan dengan index genap
		fmt.Println("Elemen Array Dengan Index Genap : ")
		for i := 0 ; i < index ; i+=2 {
			fmt.Print(array[i]," ")
		}
		fmt.Println()	

		case "4" :
			var x int
			fmt.Print("Inputkan Bilangan X : ")
			fmt.Scan(&x)
			fmt.Printf("Elemen Dengan Index Kelipatan %d : ", x)
			for i := x; i < index; i += x {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case "5" :
			var n int
			fmt.Print("Inputkan Index Yang Ingin Dihapus : ")
			fmt.Scan(&n)
			if n >= 0 && n < index {
				for i := n; i < index-1; i++ {
					array[i] = array[i+1]
				}
				index-- // Hapus Array
				fmt.Println("Array Terkini :", array[:index])
			} else {
				fmt.Println("Index Tidak Valid")
			}

		case "6" :
			total := 0
			for _, nilai := range array[:index] {
				total += nilai
			}
			fmt.Printf("Rata-rata : %.2f\n", float64(total)/float64(index))

		case "7" :
			total, totalsq := 0, 0.0
			for _, nilai := range array[:index] {
				total += nilai // Hasil Penjumlahan Data Dalam Array
				totalsq += float64(nilai* nilai) // Hasil Penjumlahan Di kuadratkan
			}

			rata_rata := float64(total) / float64(index)
			// Rumus Mencari Simpangan Baku : hasil penjumlahan data pangkat dua dibagi jumlah data
			// dikurangi hasil kali dari rata rata
			stndrDev := math.Sqrt(totalsq/float64(index) - rata_rata*rata_rata)
			fmt.Printf("Standard deviasi / Simpangan Baku : %.2f\n", stndrDev)

		case "8" :
			frekuensi := make(map[int]int)
			for i := 0; i < index; i++ {
				frekuensi[array[i]]++
			}
			fmt.Println("Frekuensi dari suatu bilangan adalah :")
			for nilai, hitung := range frekuensi {
				fmt.Printf("%d: %d kali\n", nilai, hitung)
			}
		default:
			fmt.Println("opsi tidak ada")
			return
		}
	}
}

```
### Output Program
![output nomor 2 part 1](https://github.com/user-attachments/assets/37a0b145-d55e-4ab4-bf09-51d4eb4e3b01)

![output nomor 2 part 2](https://github.com/user-attachments/assets/f5633582-c432-40a1-acac-68289cf09a4a)

![output nomor 2 part 3](https://github.com/user-attachments/assets/ddba6b3d-a55c-4ae2-a7d8-e4c9a6364412)

![output nomor 2 part 4](https://github.com/user-attachments/assets/251e69d5-48f5-4778-874f-6b87545a58a3)

![output nomor 2 part 5](https://github.com/user-attachments/assets/0075c00e-cc96-4b5f-aba9-07089418ae98)

![output nomor 2 part 6](https://github.com/user-attachments/assets/7ee3f5e4-d5b4-4355-a027-84566ac3af63)

![output nomor 2 part 7](https://github.com/user-attachments/assets/f9ad06d2-f46e-4d04-beb5-6a9558a23037)

![output nomor 2 part 8](https://github.com/user-attachments/assets/168e1472-a34b-4281-a5d3-eb90bbdfef9b)

## No 3 
![soal nomor 3 part 1](https://github.com/user-attachments/assets/45eda0bb-d6a3-4265-8f79-0a3e0c9eff2c)

![soal nomor 3 part 2](https://github.com/user-attachments/assets/b5845c59-143f-458a-a54d-e161095bbf54)

### Code Program
```go
package main
import "fmt"
func main (){
	var club1, club2 string
	fmt.Print("Klub 1 : ")
	fmt.Scan(&club1)
	fmt.Print("Klub 2 : ")
	fmt.Scan(&club2)
	fmt.Println()

	//Array untuk menyimpan pemenang
	var pemenang[100]string
	hitung_pemenang := 0
	hitung_pertandingan := 1

	for{
		var skor_club1, skor_club2 int
		fmt.Printf("Pertandingan %d : ", hitung_pertandingan)
		fmt.Scan(&skor_club1,&skor_club2)

		if skor_club1 < 0 || skor_club2 < 0 {
			break
		}
	 // Jika Inputan Skor dibawah nol atau negatif maka program akan dihentikan
		if skor_club1 > skor_club2 {
			pemenang[hitung_pemenang] = fmt.Sprintf("Hasil %d : %s", hitung_pertandingan, club1)
		} else if skor_club2 > skor_club1 {
			pemenang[hitung_pemenang] = fmt.Sprintf("Hasil %d : %s", hitung_pertandingan, club2)
		} else {
			pemenang[hitung_pemenang] = fmt.Sprintf("Hasil %d : Draw", hitung_pertandingan)
		}
		hitung_pemenang++
		hitung_pertandingan++
	}
	for i := 0 ; i < hitung_pemenang; i++{
		fmt.Println(pemenang[i])
	}
	fmt.Println("Pertandingan Selesai")
}
```
### Output Program
![output nomor 3](https://github.com/user-attachments/assets/9a4ef557-c12f-41fb-9283-f846d1780bc6)

## No 4 
![Soal Latihan 7 no 4](https://github.com/user-attachments/assets/47f8be7f-8658-44a4-b193-fe3e7bf72212)

![Soal Latihan 7 no 4 part 1](https://github.com/user-attachments/assets/a06bce10-cd62-4bb3-a951-778a031ef8f4)

### Code Program
```go
package main
import (
	"bufio"
	"fmt"
	"os"
)
const NMAX int = 127
type Tabel [NMAX]rune
func isiArray(t *Tabel, n *int) {
    // Data Tersedia dalam piranti masukan
    // Array t berisi sejumlah n karakter yang di inputkan user. 
    // Proses input selama karakter bukan TITIK (".") dan n <= NMAX
	fmt.Print("Teks	: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	for _, ch := range text {
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		(*n)++
	}
}
func cetakArray(t Tabel, n int) {
    // terdefinisi array t yang berisi sejumlah n karakter
    // n karakter array muncul dilayar 
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}
func balikanArray(t *Tabel, n int) {
    // Terdefinisi Array t yang berisi sejumlah n karakter 
    // Urutan isi array t terbalik
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}
func palindrom(t Tabel, n int) bool {
    // Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom
    // dan false apabila sebaliknya. Manfaatkan prosedur balikanArray 
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var t Tabel
	var n int

	isiArray(&t, &n)

	balikanArray(&t, n)
	fmt.Print("Reverse teks	: ")
	cetakArray(t, n)

	cek := palindrom(t, n)
	fmt.Println("Palindrom	?", cek)

}
```
### Output Program
![output nomor 4](https://github.com/user-attachments/assets/d2d1ed6d-3ea1-4449-a89f-3b6d8c516081)


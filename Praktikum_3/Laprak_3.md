# <h1 align="center">Laporan Praktikum Alpro 2</h1>
<p align="center">Arvel Mahsa Athallah Firdaus/2311102188/IF-11-01</p>

# Soal Latihan Modul 2B
### Nomor 3
![Nomor 3](https://github.com/user-attachments/assets/20416946-09a0-4ecc-b652-2878311da1a3)

![Nomor 3 part 2](https://github.com/user-attachments/assets/fb6d1ccd-d0a4-493f-8b37-5d70b45964de)

#### Source Code
```go
package main
import ("fmt")
func selisih(a, b float64) float64{
	if a > b {
		return a - b
	}
	return b - a
}
func main (){
	var berat1, berat2 float64
	for berat1 > 0 || berat2 > 0 || berat1 + berat2 != 150 {
		fmt.Print("Masukan Berat Belanjaan Di Kedua Kantong : ")
		fmt.Scanln(&berat1, &berat2)

		if selisih(berat1,berat2) < 9 {
		fmt. Println("Sepeda Pak Andi Akan Oleng : false \n")
		} else if selisih(berat1,berat2) > 9 {
			fmt.Println("Sepeda Pak Andi Akan Oleng : true \n")
		}
		if berat1 + berat2 > 150{
			fmt.Println("Program Selesai...")
			break
		}
	}
}
```
#### Output Nomor 3
![Output No 3](https://github.com/user-attachments/assets/d0f49392-07fd-4f5e-98a1-f578c99075ad)

### Nomor 4
![Nomor 4](https://github.com/user-attachments/assets/93585c60-f5c6-4bc6-98f6-ba245e6ef152)

#### Source Code
```go
package main

import (
	"fmt"
	"math"
)

func countF(k int) float64 {
	Pembilang := math.Pow(float64(4*k+2), 2)
	Penyebut := float64((4*k + 1) * (4*k + 3))
	return Pembilang / Penyebut
}

func countAkarDua(k int) float64 {
	var akarDua float64 = 1
	for a := 0; a <= k; a++ {
		akarDua *= countF(a)
	}
	return akarDua
}

func main() {
	var k int
	fmt.Print("Nilai k : ")
	fmt.Scan(&k)

	fmt.Printf("Nilai f(%d): %f\n", k, countF(k))

	akarloro := countAkarDua(k)
	fmt.Println("Nilai Akar 2 : ", akarloro)
}
```
#### Output Nomor 4
![output nomor 4](https://github.com/user-attachments/assets/958fa6a6-44b9-488d-a4ae-a3f5557152af)

# Soal Latihan Modul 2C
### Nomor 1
![Soal Nomor 1](https://github.com/user-attachments/assets/928df602-b447-48e8-bd7f-5b650725a701)

#### Source Code
```go
package main
import ("fmt")

func main(){
	// Deklarasi Variabel
	var berat int

	// Inputan Berat Dari User
	fmt.Print("Berat Parcel : ")
	fmt.Scanf("%d",&berat)

	beratdalamkilo := berat / 1000 // Berat Dalam Kg
	biayajasa := beratdalamkilo * 10000 // Jasa Pengiriman Berdasarkan Berat Kg
	beratdalamgram := berat % 1000 // Sisa Berat Dalam Gram
	fmt.Printf("Detail Berat : %d Kg + %d gram ", beratdalamkilo, beratdalamgram)

	// Kondisi Bila Sisa Berat Lebih Dari 500, Maka Dikenakan biaya Rp.5 per-gram
	if beratdalamgram >= 500 {
		beratdalamgram_jika_lebih_dari_500g := beratdalamgram * 5
		fmt.Printf("\nDetail Biaya : Rp.%d + Rp.%d", biayajasa, beratdalamgram_jika_lebih_dari_500g)
		totalbiaya_jika_lebih_dari500g := biayajasa + beratdalamgram_jika_lebih_dari_500g
		fmt.Printf("\nTotal Biaya : Rp.%d", totalbiaya_jika_lebih_dari500g )

	// Kondisi Bila Sisa Berat Lebih Dari 500, Maka Dikenakan biaya Rp.15 per-gram
	} else if beratdalamgram <= 500 {
		beratdalamgram_jika_kurang_dari_500g := beratdalamgram * 15
		fmt.Printf("\nDetail Biaya : Rp.%d + Rp.%d", biayajasa, beratdalamgram_jika_kurang_dari_500g)
		totalbiaya_jika_kurang_dari500g := biayajasa + beratdalamgram_jika_kurang_dari_500g
		fmt.Printf("\nTotal Biaya  : Rp.%d", totalbiaya_jika_kurang_dari500g )
	}
}
```
### Output Nomor 1
![Output No 1 #1](https://github.com/user-attachments/assets/49170ad5-203d-4fc3-9987-9d13953ff557)

![Output 1 #2](https://github.com/user-attachments/assets/314e4cd1-4477-40cc-a306-907a5fdcf584)

![Output 1 #3](https://github.com/user-attachments/assets/669da484-b4b2-4170-9a5a-6104d99e43d0)

### Nomor 2
![Soal Nomor 2](https://github.com/user-attachments/assets/d3edcd7d-5116-4e75-a1be-bbfd14f15e40)

#### Source Code
```go
// Codingan Setelah Dibetulkan
package main
import ("fmt")
func main (){
	var nam float64
	var nmk string
	fmt.Print("Nilai Akhir Mata Kuliah : ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	}else if nam > 72.5 {
		nmk = "AB"
	}else if nam > 65 {
		nmk = "B"
	}else if nam > 57.7 {
		nmk = "BC"
	}else if nam > 50 {
		nmk = "C"
	}else if nam > 40 {
		nmk = "D"
	}else if nam <= 40 {
		nmk = "E"
	}
	fmt.Println("Nilai Mata Kuliah : ", nmk)
}
```
#### Output
![Output 2 #a](https://github.com/user-attachments/assets/72985cb4-31bd-4fee-ba9c-eacb4542aae7)

#### Jawaban A : Program tidak mengeluarkan output sesuai apa yang sudah di program alias error, Eksekusi tidak sesuai dengan soal

#### Jawaban B : Kesalahan program terletak pada ekspresi percabangan seharusnya menggunakan variabel nmk sebagai variabel yang memiliki tipe data string. Kesalahan kedua terletak pada penggunaan percabangan itu sendiri

#### Jawaban C :

![Nilai Akhir A](https://github.com/user-attachments/assets/2c437c42-7cbb-4b02-877c-0dcf1d2d23cf)

![Nilai Akhir B](https://github.com/user-attachments/assets/cbf7caa8-9624-42a3-ba7d-765c7dd05770)

![Nilai Akhir D](https://github.com/user-attachments/assets/f548f0ff-33bb-4af2-b423-86e022787ff0)

### Nomor 3
![Nomor 3](https://github.com/user-attachments/assets/85fafdcd-4d62-418f-b7ca-c79958f3e902)

#### Source Code 
```go
package main

import (
	"fmt"
)

func main() {
	var bil int
	fmt.Print("Bilangan : ")
	fmt.Scan(&bil)

	fmt.Printf("Faktor %d : ", bil)
	for i := 1; i <= bil; i++ {
		if bil % i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	isPrime := true
	for i := 2; i*i <= bil; i++{
		if bil %2 == 0{
			isPrime = false
			break
		}
	}
	if isPrime{
		fmt.Println("Prima : ", isPrime)
	}else {
		fmt.Println("Prima : ", isPrime)
	}
}

```
#### Output Nomor 3
![output nomor 3 soal 2c](https://github.com/user-attachments/assets/477b43eb-6e5f-4f26-a78a-613acd95a187)

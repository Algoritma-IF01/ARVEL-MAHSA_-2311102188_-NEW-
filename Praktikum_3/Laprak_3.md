# <h1 align="center">Laporan Praktikum Alpro 2</h1>
<p align="center">Arvel Mahsa Athallah Firdaus/2311102188/IF-11-01</p>

# Soal Latihan Modul 2B
### Nomor 3
![Nomor 3](https://github.com/user-attachments/assets/20416946-09a0-4ecc-b652-2878311da1a3)

![Nomor 3 part 2](https://github.com/user-attachments/assets/fb6d1ccd-d0a4-493f-8b37-5d70b45964de)

#### Source Code
```
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
```
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
### Nomor 2
### Nomor 3

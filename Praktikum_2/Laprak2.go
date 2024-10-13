// LAPRAK 2
//2311102188_ARVELMAHSA
// NO 3
// package main 

// import ("fmt"
//  "math")

// func main(){
// 	//Deklarasi Phi
// 	const phi = 3.1415926535

// 	//Inputan dari pengguna
// 	var jari2 float64
// 	fmt.Print("masukan jari-jari = ")
// 	fmt.Scan(&jari2)

// 	//Volume bola
// 	volume := (4.0 / 3.0) * phi * math.Pow(jari2,3)

// 	//Luas Kulit Bola
// 	LK_bola := 4 * phi * math.Pow(jari2, 2)

// 	// Cetak Hasil
// 	fmt.Printf(" Bola dengan jari jari %.0f \n memiliki volume %.4f dan luas Kulit %.4f\n", 
// jari2, volume, LK_bola)
// }

//NO 4
// package main 

// import "fmt"

// func main(){
// 	var derajat_celcius float64

// 	// Inputan Suhu Dari User
// 	fmt.Print("Masukan Suhu Celcius : ")
// 	fmt.Scan(&derajat_celcius)

// 	// Hitung Fahrenheit
// 	Fahrenheit := ( 9.0 / 5.0 ) * derajat_celcius + 32

// 	// Hitung Reamur
// 	Reamur := ( 4.0 / 5.0 ) * derajat_celcius

// 	// Hitung Kelvin 
// 	Kelvin := derajat_celcius + 273.15

// 	// Print Hasil
// 	fmt.Printf("Derajat Celcius : %.0f\n", derajat_celcius)
// 	fmt.Printf("Derajat Fahrenheit : %.0f\n", Fahrenheit)
// 	fmt.Printf("Derajat Reamur : %.0f\n", Reamur)
// 	fmt.Printf("Derajat Kelvin : %.0f\n", Kelvin)
// }

// NO 5
package main 

import "fmt"

func main (){

	// Deklarasi untuk menampung inputan integer dan karakter
	var v,w,x,y,z int16
	// kar1,kar2,kar3
	var karakter string

	// Baca Input 5 Karakter
	fmt.Print("Masukan 5 Data Integer (Berdasarkan Tabel ASCII 32 - 127): ")
	fmt.Scanln(&v,&w,&x,&y,&z)

	// Baca Input 3 Karakter
	fmt.Print("Masukan 3 Data Karakter : ")
	fmt.Scanln(&karakter)

	// Cetak 5 Nilai Integer
	fmt.Printf("%c%c%c%c%c \n ", v,w,x,y,z)

	// Cetak 3 Karakter Tanpa Spasi
	fmt.Printf("%c %c %c \n", karakter[0]+1, karakter[1]+1, karakter[2]+1)
}

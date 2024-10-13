// CEMUNGUTTT 😒

// package main
// import "fmt"
// func main(){
// 	var greetings = "Selamat Datang Di dunia DAP"
// 	var a , b int

// 	fmt. Println(greetings)
// 	fmt.Scanln(&a,&b)
// 	fmt.Printf("%v + %v = %v\n", a, b, a+b)
// }

// package main
// import "fmt"

// func main(){
// 	var a, b, c float64
// 	var hipotenusa bool

// 	fmt.Scanln( &a, &b, &c)
// 	hipotenusa = (c*c) == (a*a + b*b)
// 	fmt.Println("Sisi c adalah hipotenusa segitiga a, b, c : ", hipotenusa)
// }

// package main
// import "fmt"

// func main(){
// 	var a, b, c float64
// 	var hipotenusa bool

// 	fmt.Print("Input nilai A : ")
// 	fmt.Scanln(&a)
// 	fmt.Print("Input nilai B : ")
// 	fmt.Scanln(&b)
// 	fmt.Print("Input nilai C : ")
// 	fmt.Scanln(&c)
// 	hipotenusa = (c*c) == (a*a + b*b)
// 	fmt.Println("Sisi c adalah hipotenusa segitiga a, b, c : ", hipotenusa)
// }

// LATIHAN MODUL 2A
// No 1
// package main
// import "fmt"

// func main(){
// 	var (
// 		satu, dua, tiga string
// 		temp string
// 	)
// 	fmt.Print("Masukan input string : ")
// 	fmt.Scanln(&satu)
// 	fmt.Print("Masukan input string : ")
// 	fmt.Scanln(&dua)
// 	fmt.Print("Masukan input string : ")
// 	fmt.Scanln(&tiga)
// 	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
// 	temp = satu
// 	satu = dua
// 	dua = tiga 
// 	tiga = temp
// 	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
// } 

// No 2
// package main
// import "fmt"

// func main (){
// 	var tahun int

// 	fmt.Print("Input Tahun : ")
// 	fmt.Scanln(&tahun)

// 	if tahun%4 == 0 {
// 		fmt.Println("Tahun yang anda inputkan adalah kabisat ? : true")
// 	}else {
// 		fmt.Println("Tahun yang anda adalah kabisat ? : false")
// 	}
// }

// Cara 2
// package main
// import "fmt"

// func main (){
// 	var tahun int
// 	var cektahun bool

// 	fmt.Print("Input Tahun : ")
// 	fmt.Scanln(&tahun)

// 	cektahun = (tahun %4 == 0)
// 	fmt.Println("Tahun ini adalah kabisat :", cektahun)
// }
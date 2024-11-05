// package main

// import "fmt"

// type bilangan int
// type pecahan float64

// func main() {
// 	var a, b bilangan
// 	var hasil pecahan
// 	a = 9
// 	b = 5
// 	hasil = pecahan(a) / pecahan(b)
// 	fmt.Println(hasil)

// }

//STRUCT 
// package main
// import "fmt"
// type waktu struct {
// 	jam, menit, detik int
// }

// func main () {
// 	var wParkir, wPulang, durasi waktu
// 	var dParkir, dPulang, lParkir int

// 	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
// 	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

// 	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
// 	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600

// 	lParkir = dPulang - dParkir

// 	durasi.jam = lParkir / 3600
// 	durasi.menit = lParkir % 3600 / 60
// 	durasi.detik = lParkir % 3600 % 60

// 	fmt.Printf("Lama parkir : %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)

// }

// ARRAY 
// package main

// import "fmt"

// // Definisi tipe CircType
// type CircType struct {
// 	Radius float64
// }

// // Definisi tipe NewType
// type NewType struct {
// 	Name string
// }

// func main() {
// 	var (
// 		// Array arr mempunyai 73 elemen, masing-masing bertipe CircType
// 		arr [73]CircType

// 		// Array buf dengan 5 elemen, dengan nilai awal 7,3,5,2,11
// 		buf = [5]byte{7, 3, 5, 2, 11}

// 		// Array mhs berisi 2000 elemen NewType
// 		mhs [2000]NewType

// 		// Array dua dimensi rec berisi nilai float64
// 		rec [20][40]float64
// 	)

// 	// Mengisi beberapa elemen contoh
// 	arr[0] = CircType{Radius: 5.5}
// 	mhs[0] = NewType{Name: "Alice"}
// 	rec[0][0] = 3.14

// 	// Contoh penggunaan variabel
// 	fmt.Println("arr[0]:", arr[0])
// 	fmt.Println("buf:", buf)
// 	fmt.Println("mhs[0]:", mhs[0])
// 	fmt.Println("rec[0][0]:", rec[0][0])
// }

//MAP
// package main

// import "fmt"

// func main() {
// 	// Membuat map dengan tipe string sebagai kunci dan integer sebagai nilai
// 	scores := map[string]int{
// 		"John": 90,
// 		"Anne": 85,
// 	}

// 	// Mengambil nilai dari kunci
// 	johnScore := scores["John"]
// 	fmt.Println("Nilai John:", johnScore)

// 	// Mengganti nilai dari kunci yang ada
// 	scores["John"] = 95
// 	fmt.Println("Nilai John setelah diganti:", scores["John"])

// 	// Menambah kunci baru dengan nilai
// 	scores["David"] = 88
// 	fmt.Println("Nilai David ditambahkan:", scores["David"])

// 	// Menghapus kunci dari map
// 	delete(scores, "Anne")
// 	fmt.Println("Map setelah kunci 'Anne' dihapus:", scores)

// 	// Mengecek apakah kunci ada dalam map
// 	if score, ada := scores["David"]; ada {
// 		fmt.Println("Nilai David ditemukan:", score)
// 	} else {
// 		fmt.Println("Nilai David tidak ditemukan")
// 	}
// }
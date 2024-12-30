package main

import (
	"fmt"
)

const (
	maxPabrikan = 100
	maxMobil    = 100
)

type Mobil struct {
	nama_mobil  string
	model_mobil string
	harga_mobil float64
}

type Pabrikan struct {
	nama_pabrikan string
	banyak_mobil  [maxMobil]Mobil
	hitung_mobil  int
}

var (
	pabrikan      [maxPabrikan]Pabrikan
	hitung_pabrik int
)

// Struktur dan variabel global yang diperlukan untuk mendukung program.
// Tambahkan struktur Pabrik dan Mobil sesuai kebutuhan.

func main() {
	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Edit Pabrikan")
		fmt.Println("3. Hapus Pabrikan")
		fmt.Println("4. Tambah Mobil")
		fmt.Println("5. Edit Mobil")
		fmt.Println("6. Hapus Mobil")
		fmt.Println("7. Lihat Data Pabrikan dan Mobil")
		fmt.Println("8. Cari Mobil Berdasarkan Pabrikan")
		fmt.Println("9. Cari Mobil Tertentu")
		fmt.Println("10. Tampilkan Pabrikan Terurut Berdasarkan Jumlah Mobil")
		fmt.Println("11. Tambah Penjualan")
		fmt.Println("12. Tampilkan Penjualan Tertinggi")
		fmt.Println("13. Lihat Mobil Terurut Berdasarkan Kriteria")
		fmt.Println("14. Keluar")

		var pilihan int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			TambahPabrikan()
		case 2:
			EditPabrikan()
		case 3:
			HapusPabrik()
		case 4:
			TambahMobil()
		case 5:
			EditMobil()
		case 6:
			HapusMobil()
		case 7:
			LihatData()
		case 8:
			var namaPabrik string
			fmt.Print("Masukkan nama pabrikan untuk mencari mobil: ")
			fmt.Scanln(&namaPabrik)
			CariMobilByPabrikan(namaPabrik)
		case 9:
			var namaMobil string
			fmt.Print("Masukkan nama mobil untuk mencari: ")
			fmt.Scanln(&namaMobil)
			CariMobilTertentu(namaMobil)
		case 10:
			SortPabrikanByJumlahMobil()
		case 11:
			TambahPenjualan()
		case 12:
			TampilkanPenjualanTertinggi()
		case 13:
			LihatMobilTerurut()
		case 14:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}

package main

import (
	"fmt"
)

const maxPenjualan = 100

type Penjualan struct {
	nama_pabrikan string
	nama_mobil    string
	jumlah_unit   int
	total_harga   float64
}

var (
	penjualan   [maxPenjualan]Penjualan
	hitung_jual int
)

func TampilkanPenjualanTertinggi() {
	if hitung_jual == 0 {
		fmt.Println("Belum ada data penjualan.")
		return
	}

	// Salin data penjualan untuk pengurutan
	penjualanCopy := penjualan

	// Urutkan dengan Selection Sort berdasarkan total_harga
	for i := 0; i < hitung_jual-1; i++ {
		maxIdx := i
		for j := i + 1; j < hitung_jual; j++ {
			if penjualanCopy[j].total_harga > penjualanCopy[maxIdx].total_harga {
				maxIdx = j
			}
		}
		penjualanCopy[i], penjualanCopy[maxIdx] = penjualanCopy[maxIdx], penjualanCopy[i]
	}

	fmt.Println("\n3 Penjualan Tertinggi (Selection Sort):")
	for i := 0; i < 3 && i < hitung_jual; i++ {
		jual := penjualanCopy[i]
		fmt.Printf("%d. Pabrikan: %s, Mobil: %s, Total Harga: Rp%.2f\n",
			i+1, jual.nama_pabrikan, jual.nama_mobil, jual.total_harga)
	}

	// Urutkan kembali dengan Insertion Sort berdasarkan jumlah_unit
	for i := 1; i < hitung_jual; i++ {
		key := penjualanCopy[i]
		j := i - 1
		for j >= 0 && penjualanCopy[j].jumlah_unit < key.jumlah_unit {
			penjualanCopy[j+1] = penjualanCopy[j]
			j--
		}
		penjualanCopy[j+1] = key
	}

	fmt.Println("\n3 Penjualan Tertinggi Berdasarkan Unit (Insertion Sort):")
	for i := 0; i < 3 && i < hitung_jual; i++ {
		jual := penjualanCopy[i]
		fmt.Printf("%d. Pabrikan: %s, Mobil: %s, Jumlah Unit: %d\n",
			i+1, jual.nama_pabrikan, jual.nama_mobil, jual.jumlah_unit)
	}
}

func TambahPenjualan() {
	if hitung_jual >= maxPenjualan {
		fmt.Println("Kapasitas data penjualan sudah penuh.")
		return
	}

	fmt.Print("Masukkan nama pabrikan: ")
	var namaPabrik string
	fmt.Scanln(&namaPabrik)

	// Cari pabrikan berdasarkan nama
	pabrikanIndex := sequentialSearchManufacturer(namaPabrik)
	if pabrikanIndex == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama mobil: ")
	var namaMobil string
	fmt.Scanln(&namaMobil)

	// Cari mobil berdasarkan nama
	mobilIndex := sequentialSearchCar(pabrikan[pabrikanIndex], namaMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan jumlah unit yang terjual: ")
	var jumlahUnit int
	fmt.Scanln(&jumlahUnit)

	// Hitung total harga
	hargaMobil := pabrikan[pabrikanIndex].banyak_mobil[mobilIndex].harga_mobil
	totalHarga := float64(jumlahUnit) * hargaMobil

	// Simpan data penjualan
	penjualan[hitung_jual] = Penjualan{
		nama_pabrikan: namaPabrik,
		nama_mobil:    namaMobil,
		jumlah_unit:   jumlahUnit,
		total_harga:   totalHarga,
	}
	hitung_jual++

	fmt.Println("Data penjualan berhasil ditambahkan.")
}
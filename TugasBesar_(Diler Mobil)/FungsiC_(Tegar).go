package main

import (
	"fmt"
)

func SortPabrikanByJumlahMobil() {
	if hitung_pabrik == 0 {
		fmt.Println("Belum ada data pabrikan.")
		return
	}

	// Selection Sort berdasarkan jumlah mobil
	for i := 0; i < hitung_pabrik-1; i++ {
		maxIdx := i
		for j := i + 1; j < hitung_pabrik; j++ {
			if pabrikan[j].hitung_mobil > pabrikan[maxIdx].hitung_mobil {
				maxIdx = j
			}
		}
		if maxIdx != i {
			pabrikan[i], pabrikan[maxIdx] = pabrikan[maxIdx], pabrikan[i]
		}
	}

	// Menampilkan Pabrikan Terurut Berdasarkan Jumlah Mobil
	fmt.Println("Daftar Pabrikan Terurut Berdasarkan Jumlah Mobil:")
	for i := 0; i < hitung_pabrik; i++ {
		fmt.Printf("%d. %s - Jumlah Mobil: %d\n", i+1, pabrikan[i].nama_pabrikan, pabrikan[i].hitung_mobil)
	}
}

func LihatMobilTerurut() {
	if hitung_pabrik == 0 {
		fmt.Println("Belum ada data pabrikan.")
		return
	}

	type MobilDetail struct {
		tahun_keluar  int
		nama_pabrikan string
		nama_mobil    string
		model_mobil   string
		harga_mobil   float64
	}

	var semuaMobil [100]MobilDetail // Array statis dengan kapasitas maksimum 100 mobil
	jumlahMobil := 0

	// Mengumpulkan data mobil dari semua pabrikan
	for i := 0; i < hitung_pabrik; i++ {
		for j := 0; j < pabrikan[i].hitung_mobil && jumlahMobil < len(semuaMobil); j++ {
			mobil := pabrikan[i].banyak_mobil[j]
			semuaMobil[jumlahMobil] = MobilDetail{
				nama_pabrikan: pabrikan[i].nama_pabrikan,
				nama_mobil:    mobil.nama_mobil,
				model_mobil:   mobil.model_mobil,
				harga_mobil:   mobil.harga_mobil,
				tahun_keluar:  j + 2000, // Contoh penambahan tahun untuk simulasi
			}
			jumlahMobil++
		}
	}

	// Sorting berdasarkan kriteria: tahun keluar, nama mobil, nama pabrikan
	for i := 0; i < jumlahMobil-1; i++ {
		for j := 0; j < jumlahMobil-i-1; j++ {
			if semuaMobil[j].tahun_keluar > semuaMobil[j+1].tahun_keluar ||
				(semuaMobil[j].tahun_keluar == semuaMobil[j+1].tahun_keluar && semuaMobil[j].nama_mobil > semuaMobil[j+1].nama_mobil) ||
				(semuaMobil[j].tahun_keluar == semuaMobil[j+1].tahun_keluar && semuaMobil[j].nama_mobil == semuaMobil[j+1].nama_mobil && semuaMobil[j].nama_pabrikan > semuaMobil[j+1].nama_pabrikan) {
				semuaMobil[j], semuaMobil[j+1] = semuaMobil[j+1], semuaMobil[j]
			}
		}
	}

	// Menampilkan data yang sudah terurut
	fmt.Println("\nDaftar Mobil Terurut:")
	for i := 0; i < jumlahMobil; i++ {
		mobil := semuaMobil[i]
		fmt.Printf("%d. %s - %s (%s) - Rp%.2f - Tahun: %d\n",
			i+1, mobil.nama_pabrikan, mobil.nama_mobil, mobil.model_mobil, mobil.harga_mobil, mobil.tahun_keluar)
	}
}

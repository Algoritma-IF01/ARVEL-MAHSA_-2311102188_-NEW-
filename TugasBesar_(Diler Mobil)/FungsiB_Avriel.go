package main

import (
	"fmt"
)
func CariMobilByPabrikan(namaPabrik string) {
	index := sequentialSearchManufacturer(namaPabrik)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	fmt.Printf("Daftar mobil dari pabrikan %s:\n", namaPabrik)
	for j := 0; j < pabrikan[index].hitung_mobil; j++ {
		car := pabrikan[index].banyak_mobil[j]
		fmt.Printf("%d. %s (%s) - Rp%.2f\n", j+1, car.nama_mobil, car.model_mobil, car.harga_mobil)
	}
}

func CariMobilTertentu(namaMobil string) {
	found := false
	for i := 0; i < hitung_pabrik; i++ {
		indexMobil := sequentialSearchCar(pabrikan[i], namaMobil)
		if indexMobil != -1 {
			found = true
			car := pabrikan[i].banyak_mobil[indexMobil]
			fmt.Printf("Ditemukan di pabrikan %s:\n", pabrikan[i].nama_pabrikan)
			fmt.Printf("%s (%s) - Rp%.2f\n", car.nama_mobil, car.model_mobil, car.harga_mobil)
		}
	}
	if !found {
		fmt.Println("Mobil tidak ditemukan.")
	}
}
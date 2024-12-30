package main

import (
	"fmt"
)

func sequentialSearchManufacturer(namaPabrik string) int {
	for i := 0; i < hitung_pabrik; i++ {
		if pabrikan[i].nama_pabrikan == namaPabrik {
			return i
		}
	}
	return -1
}

func sequentialSearchCar(manufaktur Pabrikan, namaMobil string) int {
	for i := 0; i < manufaktur.hitung_mobil; i++ {
		if manufaktur.banyak_mobil[i].nama_mobil == namaMobil {
			return i
		}
	}
	return -1
}

func binarySearchCar(manufaktur Pabrikan, namaMobil string) int {
	sortCarsByName(&manufaktur)
	low, high := 0, manufaktur.hitung_mobil-1
	for low <= high {
		mid := (low + high) / 2
		if manufaktur.banyak_mobil[mid].nama_mobil == namaMobil {
			return mid
		} else if manufaktur.banyak_mobil[mid].nama_mobil < namaMobil {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func sortCarsByName(manufaktur *Pabrikan) {
	for i := 0; i < manufaktur.hitung_mobil-1; i++ {
		for j := 0; j < manufaktur.hitung_mobil-i-1; j++ {
			if manufaktur.banyak_mobil[j].nama_mobil > manufaktur.banyak_mobil[j+1].nama_mobil {
				manufaktur.banyak_mobil[j], manufaktur.banyak_mobil[j+1] = manufaktur.banyak_mobil[j+1], manufaktur.banyak_mobil[j]
			}
		}
	}
}

func TambahPabrikan() {
	var namaPabrik string
	if hitung_pabrik >= maxPabrikan {
		fmt.Println("Kapasitas pabrikan sudah penuh.")
		return
	}

	fmt.Print("Masukkan nama pabrikan: ")
	fmt.Scanln(&namaPabrik)

	pabrikan[hitung_pabrik] = Pabrikan{nama_pabrikan: namaPabrik, hitung_mobil: 0}
	hitung_pabrik++
	fmt.Println("Pabrikan berhasil ditambahkan.")
}

func EditPabrikan() {
	fmt.Print("Masukkan nama pabrikan yang ingin diubah: ")
	var namaPabrik string
	fmt.Scanln(&namaPabrik)

	index := sequentialSearchManufacturer(namaPabrik)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama baru untuk pabrikan: ")
	var namaBaru string
	fmt.Scanln(&namaBaru)

	pabrikan[index].nama_pabrikan = namaBaru
	fmt.Println("Pabrikan berhasil diubah.")
}

func HapusPabrik() {
	fmt.Print("Masukkan nama pabrikan yang ingin dihapus: ")
	var namaPabrik string
	fmt.Scanln(&namaPabrik)

	index := sequentialSearchManufacturer(namaPabrik)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	for i := index; i < hitung_pabrik-1; i++ {
		pabrikan[i] = pabrikan[i+1]
	}
	hitung_pabrik--
	fmt.Println("Pabrikan berhasil dihapus.")
}

func TambahMobil() {
	fmt.Print("Masukkan nama pabrikan untuk menambah mobil: ")
	var namaPabrik string
	fmt.Scanln(&namaPabrik)

	index := sequentialSearchManufacturer(namaPabrik)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	if pabrikan[index].hitung_mobil >= maxMobil {
		fmt.Println("Kapasitas mobil untuk pabrikan ini sudah penuh.")
		return
	}

	fmt.Print("Masukkan nama mobil: ")
	var namaMobil string
	fmt.Scanln(&namaMobil)

	fmt.Print("Masukkan model mobil: ")
	var model string
	fmt.Scanln(&model)

	fmt.Print("Masukkan harga mobil: ")
	var harga float64
	fmt.Scanln(&harga)

	pabrikan[index].banyak_mobil[pabrikan[index].hitung_mobil] =
		Mobil{nama_mobil: namaMobil, model_mobil: model, harga_mobil: harga}
	pabrikan[index].hitung_mobil++
	fmt.Println("Mobil berhasil ditambahkan.")
}

func EditMobil() {
	fmt.Print("Masukkan nama pabrikan untuk mengedit mobil: ")
	var namaManufaktur string
	fmt.Scanln(&namaManufaktur)

	IndexManufaktur := sequentialSearchManufacturer(namaManufaktur)
	if IndexManufaktur == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama mobil yang ingin diubah: ")
	var namaMobil string
	fmt.Scanln(&namaMobil)

	carIndex := binarySearchCar(pabrikan[IndexManufaktur], namaMobil)
	if carIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama baru untuk mobil: ")
	var namaBaru string
	fmt.Scanln(&namaBaru)

	fmt.Print("Masukkan model baru untuk mobil: ")
	var newModel string
	fmt.Scanln(&newModel)

	fmt.Print("Masukkan harga baru untuk mobil: ")
	var newHarga float64
	fmt.Scanln(&newHarga)

	pabrikan[IndexManufaktur].banyak_mobil[carIndex] =
		Mobil{nama_mobil: namaBaru, model_mobil: newModel, harga_mobil: newHarga}
	fmt.Println("Mobil berhasil diubah.")
}

func HapusMobil() {
	fmt.Print("Masukkan nama pabrikan untuk menghapus mobil: ")
	var namaManufaktur string
	fmt.Scanln(&namaManufaktur)

	IndexManufaktur := sequentialSearchManufacturer(namaManufaktur)
	if IndexManufaktur == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama mobil yang ingin dihapus: ")
	var namaMobil string
	fmt.Scanln(&namaMobil)

	carIndex := binarySearchCar(pabrikan[IndexManufaktur], namaMobil)
	if carIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		return
	}

	for i := carIndex; i < pabrikan[IndexManufaktur].hitung_mobil-1; i++ {
		pabrikan[IndexManufaktur].banyak_mobil[i] = pabrikan[IndexManufaktur].banyak_mobil[i+1]
	}
	pabrikan[IndexManufaktur].hitung_mobil--
	fmt.Println("Mobil berhasil dihapus.")
}

func LihatData() {
	if hitung_pabrik == 0 {
		fmt.Println("Belum ada data pabrikan.")
		return
	}

	for i := 0; i < hitung_pabrik; i++ {
		fmt.Printf("%d. %s\n", i+1, pabrikan[i].nama_pabrikan)
		for j := 0; j < pabrikan[i].hitung_mobil; j++ {
			car := pabrikan[i].banyak_mobil[j]
			fmt.Printf("   %d.%d. %s (%s) - Rp%.2f\n", i+1, j+1, car.nama_mobil, car.model_mobil, car.harga_mobil)
		}
	}
}
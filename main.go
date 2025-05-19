package main

import (
	"fmt"
	"strings"

func main() {
	var proyek []models.Proyek

	for {
		utils.TampilkanMenu()
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			// Tambah proyek (logika sama seperti sebelumnya)
			// Panggil features.TambahProyek

		case 2:
			// Ubah proyek
			// Panggil features.UbahProyek

		case 3:
			// Hapus proyek
			// Panggil features.HapusProyek

		case 4:
			// Cari proyek
			// Panggil features.CariProyek

		case 5:
			// Urutkan proyek
			// Panggil features.SelectionSort / InsertionSort

		case 6:
			// Statistik kategori
			// Panggil features.StatistikKategori

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

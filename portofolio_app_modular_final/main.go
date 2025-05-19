package main

import (
	"fmt"
	"strings"

	"portofolio_app/data"
	"portofolio_app/crud"
	"portofolio_app/fitur"
)

func main() {
	var daftar []data.Proyek

	for {
		fmt.Println("\n[1] Tambah")
		fmt.Println("[2] Ubah")
		fmt.Println("[3] Hapus")
		fmt.Println("[4] Cari Proyek")
		fmt.Println("[5] Urutkan Proyek")
		fmt.Println("[6] Statistik Kategori")
		fmt.Println("[7] Lihat Semua Proyek")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu: ")
		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			var nama, desk, kat, kes, tgl, tek string
			fmt.Print("Nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Deskripsi: ")
			fmt.Scanln(&desk)
			fmt.Print("Kategori: ")
			fmt.Scanln(&kat)
			fmt.Print("Kesulitan: ")
			fmt.Scanln(&kes)
			fmt.Print("Tanggal: ")
			fmt.Scanln(&tgl)
			fmt.Print("Teknologi (pisah koma): ")
			fmt.Scanln(&tek)
			daftar = crud.Tambah(daftar, data.Proyek{nama, desk, strings.Split(tek, ","), kat, tgl, kes})
			fmt.Println("✅ Proyek ditambahkan.")

		case 2:
			var lama, nama, desk, kat, kes, tgl, tek string
			fmt.Print("Nama Lama: ")
			fmt.Scanln(&lama)
			fmt.Print("Nama Baru: ")
			fmt.Scanln(&nama)
			fmt.Print("Deskripsi Baru: ")
			fmt.Scanln(&desk)
			fmt.Print("Kategori Baru: ")
			fmt.Scanln(&kat)
			fmt.Print("Kesulitan Baru: ")
			fmt.Scanln(&kes)
			fmt.Print("Tanggal Baru: ")
			fmt.Scanln(&tgl)
			fmt.Print("Teknologi (pisah koma): ")
			fmt.Scanln(&tek)
			baru := data.Proyek{nama, desk, strings.Split(tek, ","), kat, tgl, kes}
			daftar = crud.Ubah(daftar, lama, baru)
			fmt.Println("✅ Proyek diubah.")

		case 3:
			var nama string
			fmt.Print("Nama yang ingin dihapus: ")
			fmt.Scanln(&nama)
			daftar = crud.Hapus(daftar, nama)
			fmt.Println("✅ Proyek dihapus.")

		case 4:
			var key string
			var mode int
			fmt.Print("Cari berdasarkan [1] Nama [2] Kategori: ")
			fmt.Scanln(&mode)
			fmt.Print("Masukkan kata kunci: ")
			fmt.Scanln(&key)
			hasil := fitur.BinarySearch(daftar, key, mode == 2)
			fitur.Tampilkan(hasil)

		case 5:
			var mode int
			fmt.Print("Urutkan berdasarkan [1] Kesulitan [2] Tanggal: ")
			fmt.Scanln(&mode)
			if mode == 1 {
				fitur.SortingSelection(daftar)
			} else {
				fitur.SortingInsertion(daftar)
			}
			fmt.Println("✅ Diurutkan.")

		case 6:
			stat := fitur.StatistikKategori(daftar)
			for k, v := range stat {
				fmt.Printf("%s: %d proyek\n", k, v)
			}

		case 7:
			fitur.Tampilkan(daftar)

		case 0:
			fmt.Println("Terima kasih!")
			return
		}
	}
}

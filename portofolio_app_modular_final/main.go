package main

import (
	"fmt"
	"portofolio_app/crud"
	"portofolio_app/data"
	"portofolio_app/fitur"
	"strings"
)

func main() {
	var proyek []data.Proyek

	for {
		// Menampilkan menu utama
		fmt.Println("\n=== APLIKASI PORTOFOLIO DATA SCIENCE ===")
		fmt.Println("[1] Tambah Proyek")
		fmt.Println("[2] Ubah Proyek")
		fmt.Println("[3] Hapus Proyek")
		fmt.Println("[4] Cari Proyek")
		fmt.Println("[5] Urutkan Proyek")
		fmt.Println("[6] Statistik Kategori")
		fmt.Println("[7] Lihat Semua Proyek")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1: // Tambah Proyek
			var nama, desk, kat, kes, tgl, tek string
			fmt.Print("Nama Proyek: ")
			fmt.Scanln(&nama)
			fmt.Print("Deskripsi: ")
			fmt.Scanln(&desk)
			fmt.Print("Kategori: ")
			fmt.Scanln(&kat)
			fmt.Print("Kesulitan: ")
			fmt.Scanln(&kes)
			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&tgl)
			fmt.Print("Teknologi (pisah koma): ")
			fmt.Scanln(&tek)

			// Menambahkan proyek baru ke daftar
			daftar := crud.Tambah(proyek, data.Proyek{
				Nama:      nama,
				Deskripsi: desk,
				Teknologi: strings.Split(tek, ","),
				Kategori:  kat,
				Tanggal:   tgl,
				Kesulitan: kes,
			})
			proyek = daftar
			fmt.Println("✅ Proyek ditambahkan.")

		case 2: // Ubah Proyek
			var lama, nama, desk, kat, kes, tgl, tek string
			fmt.Print("Nama Proyek yang ingin diubah: ")
			fmt.Scanln(&lama)
			fmt.Print("Nama Baru: ")
			fmt.Scanln(&nama)
			fmt.Print("Deskripsi Baru: ")
			fmt.Scanln(&desk)
			fmt.Print("Kategori Baru: ")
			fmt.Scanln(&kat)
			fmt.Print("Kesulitan Baru: ")
			fmt.Scanln(&kes)
			fmt.Print("Tanggal Baru (YYYY-MM-DD): ")
			fmt.Scanln(&tgl)
			fmt.Print("Teknologi Baru (pisah koma): ")
			fmt.Scanln(&tek)

			baru := data.Proyek{
				Nama:      nama,
				Deskripsi: desk,
				Teknologi: strings.Split(tek, ","),
				Kategori:  kat,
				Tanggal:   tgl,
				Kesulitan: kes,
			}

			// Mengubah proyek yang sudah ada
			proyek = crud.Ubah(proyek, lama, baru)
			fmt.Println("✅ Proyek diubah.")

		case 3: // Hapus Proyek
			var nama string
			fmt.Print("Nama proyek yang ingin dihapus: ")
			fmt.Scanln(&nama)

			// Menghapus proyek berdasarkan nama
			proyek = crud.Hapus(proyek, nama)
			fmt.Println("✅ Proyek dihapus.")

		case 4: // Cari Proyek
			var key string
			var mode int
			fmt.Print("Cari berdasarkan [1] Nama [2] Kategori: ")
			fmt.Scanln(&mode)
			fmt.Print("Masukkan kata kunci: ")
			fmt.Scanln(&key)

			// Cari proyek berdasarkan nama atau kategori
			var hasil []data.Proyek
			if mode == 1 {
				hasil = fitur.SequentialSearch(proyek, key)
			} else {
				hasil = fitur.BinarySearch(proyek, key, true)
			}

			// Menampilkan hasil pencarian
			fitur.Tampilkan(hasil)

		case 5: // Urutkan Proyek
			var mode int
			fmt.Print("Urutkan berdasarkan [1] Kesulitan [2] Tanggal: ")
			fmt.Scanln(&mode)
			if mode == 1 {
				// Urutkan berdasarkan kesulitan
				fitur.SortingSelection(proyek)
			} else {
				// Urutkan berdasarkan tanggal
				fitur.SortingInsertion(proyek)
			}
			fmt.Println("✅ Proyek berhasil diurutkan.")

		case 6: // Statistik Kategori
			// Menampilkan statistik kategori proyek
			stat := fitur.StatistikKategori(proyek)
			for k, v := range stat {
				fmt.Printf("%s: %d proyek\n", k, v)
			}

		case 7: // Lihat Semua Proyek
			// Menampilkan daftar semua proyek
			fitur.Tampilkan(proyek)

		case 0: // Keluar
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		}
	}
}


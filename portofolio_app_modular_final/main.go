package main

import (
	"fmt"
	"portfolio-app/internal/proyek"
	"portfolio-app/internal/search"
	"portfolio-app/internal/sort"
	"portfolio-app/internal/stats"
	"strings"
)

func main() {
	var proyek []proyek.Proyek

	for {
		tampilkanMenu()
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("\n--- Tambah Proyek ---")
			var nama, deskripsi, kategori, kesulitan, tanggal, teknologiStr string

			fmt.Print("Nama Proyek: ")
			fmt.Scanln(&nama)

			fmt.Print("Deskripsi: ")
			fmt.Scanln(&deskripsi)

			fmt.Print("Kategori: ")
			fmt.Scanln(&kategori)

			fmt.Print("Kesulitan: ")
			fmt.Scanln(&kesulitan)

			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&tanggal)

			fmt.Print("Teknologi (dipisahkan koma): ")
			fmt.Scanln(&teknologiStr)

			var teknologi []string
			if strings.TrimSpace(teknologiStr) != "" {
				rawTeknologi := strings.Split(teknologiStr, ",")
				for _, t := range rawTeknologi {
					teknologi = append(teknologi, strings.TrimSpace(t))
				}
			}

			proyek = proyek.TambahProyek(proyek, nama, deskripsi, kategori, kesulitan, tanggal, teknologi)
			fmt.Println("Proyek berhasil ditambahkan!")

		case 2:
			fmt.Println("\n--- Ubah Proyek ---")
			if len(proyek) == 0 {
				fmt.Println("Belum ada proyek untuk diubah.")
				continue
			}
			var namaLama, namaBaru, deskripsi, kategori, kesulitan, tanggal, teknologiStr string

			fmt.Print("Nama proyek yang ingin diubah: ")
			fmt.Scanln(&namaLama)

			proyekDitemukan := false
			for _, p := range proyek {
				if p.Nama == namaLama {
					proyekDitemukan = true
					break
				}
			}

			if !proyekDitemukan {
				fmt.Println("Proyek dengan nama tersebut tidak ditemukan.")
				continue
			}

			fmt.Print("Nama Baru: ")
			fmt.Scanln(&namaBaru)
			fmt.Print("Deskripsi Baru: ")
			fmt.Scanln(&deskripsi)
			fmt.Print("Kategori Baru: ")
			fmt.Scanln(&kategori)
			fmt.Print("Kesulitan Baru: ")
			fmt.Scanln(&kesulitan)
			fmt.Print("Tanggal Baru (YYYY-MM-DD): ")
			fmt.Scanln(&tanggal)
			fmt.Print("Teknologi Baru (dipisah koma): ")
			fmt.Scanln(&teknologiStr)

			var teknologi []string
			if strings.TrimSpace(teknologiStr) != "" {
				rawTeknologi := strings.Split(teknologiStr, ",")
				for _, t := range rawTeknologi {
					teknologi = append(teknologi, strings.TrimSpace(t))
				}
			}
			proyek = proyek.UbahProyek(proyek, namaLama, namaBaru, deskripsi, kategori, kesulitan, tanggal, teknologi)
			fmt.Println("Proyek berhasil diubah!")

		case 3:
			fmt.Println("\n--- Hapus Proyek ---")
			if len(proyek) == 0 {
				fmt.Println("Belum ada proyek untuk dihapus.")
				continue
			}
			var nama string
			fmt.Print("Nama proyek yang ingin dihapus: ")
			fmt.Scanln(&nama)

			panjangAwal := len(proyek)
			proyek = proyek.HapusProyek(proyek, nama)
			if len(proyek) < panjangAwal {
				fmt.Println("Proyek berhasil dihapus!")
			} else {
				fmt.Println("Proyek dengan nama tersebut tidak ditemukan.")
			}

		case 4:
			fmt.Println("\n--- Cari Proyek ---")
			if len(proyek) == 0 {
				fmt.Println("Belum ada proyek untuk dicari.")
				continue
			}
			fmt.Println("Cari berdasarkan: \n[1] Nama (Sequential Search)\n[2] Kategori (Binary Search)")

			var mode int
			fmt.Print("Pilih mode: ")
			fmt.Scanln(&mode)
			if mode != 1 && mode != 2 {
				fmt.Println("Pilihan mode tidak valid.")
				continue
			}

			var keyword string
			var hasil []proyek.Proyek

			if mode == 1 {
				fmt.Print("Masukkan nama: ")
				fmt.Scanln(&keyword)
				hasil = search.SequentialSearchByName(proyek, keyword)
			} else {
				fmt.Print("Masukkan kategori: ")
				fmt.Scanln(&keyword)
				proyekUntukCari := make([]proyek.Proyek, len(proyek))
				copy(proyekUntukCari, proyek)
				hasil = search.BinarySearchByKategori(proyekUntukCari, keyword)
			}

			if len(hasil) == 0 {
				fmt.Println("Proyek tidak ditemukan.")
			} else {
				fmt.Println("Hasil Pencarian (" + keyword + "):")
				TampilkanSemuaProyek(hasil)
			}

		case 5:
			fmt.Println("\n--- Urutkan Proyek ---")
			if len(proyek) == 0 {
				fmt.Println("Belum ada proyek untuk diurutkan.")
				continue
			}
			fmt.Println("Urutkan berdasarkan: \n[1] Kesulitan (Selection Sort)\n[2] Tanggal (Insertion Sort)")

			var modeSort int
			fmt.Print("Pilih mode: ")
			fmt.Scanln(&modeSort)
			if modeSort != 1 && modeSort != 2 {
				fmt.Println("Pilihan mode tidak valid.")
				continue
			}

			if modeSort == 1 {
				sort.SelectionSort(proyek)
				fmt.Println("Proyek berhasil diurutkan berdasarkan Kesulitan!")
			} else {
				sort.InsertionSort(proyek)
				fmt.Println("Proyek berhasil diurutkan berdasarkan Tanggal!")
			}
			TampilkanSemuaProyek(proyek)

		case 6:
			fmt.Println("\n--- Statistik Kategori ---")
			if len(proyek) == 0 {
				fmt.Println("Belum ada proyek untuk menampilkan statistik.")
				continue
			}
			stat := stats.StatistikKategori(proyek)
			if len(stat) == 0 {
				fmt.Println("Tidak ada kategori proyek yang terdefinisi.")
			} else {
				for k, v := range stat {
					fmt.Printf("%s: %d proyek\n", k, v)
				}
			}

		case 7:
			TampilkanSemuaProyek(proyek)

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Tampilkan semua proyek
func TampilkanSemuaProyek(proyek []proyek.Proyek) {
	if len(proyek) == 0 {
		fmt.Println("Belum ada proyek yang terdaftar.")
		return
	}
	fmt.Println("\n--- Daftar Semua Proyek ---")
	for _, p := range proyek {
		fmt.Printf("Nama      : %s\n", p.Nama)
		fmt.Printf("Deskripsi : %s\n", p.Deskripsi)
		fmt.Printf("Kategori  : %s\n", p.Kategori)
		fmt.Printf("Kesulitan : %s\n", p.Kesulitan)
		fmt.Printf("Tanggal   : %s\n", p.Tanggal)
		fmt.Printf("Teknologi : %s\n", strings.Join(p.Teknologi, ", "))
		fmt.Println("------------------------------")
	}
}

// Menampilkan menu utama
func tampilkanMenu() {
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
}



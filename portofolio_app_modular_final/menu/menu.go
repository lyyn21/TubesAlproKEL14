package menu

import "fmt"

func TampilkanMenu() {
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
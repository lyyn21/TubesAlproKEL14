package fitur

import (
	"fmt"
	"portofolio_app/data"
	"strings"
)

// Menampilkan daftar proyek
func Tampilkan(proyek []data.Proyek) {
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


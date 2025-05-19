package fitur

import (
	"fmt"
	"strings"
	"portofolio_app/data"
)

func Tampilkan(proyek []data.Proyek) {
	if len(proyek) == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}
	for _, p := range proyek {
		fmt.Printf("Nama: %s\nDeskripsi: %s\nKategori: %s\nKesulitan: %s\nTanggal: %s\nTeknologi: %s\n", p.Nama, p.Deskripsi, p.Kategori, p.Kesulitan, p.Tanggal, strings.Join(p.Teknologi, ", "))
		fmt.Println("-----------------------------------")
	}
}

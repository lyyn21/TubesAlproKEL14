package statistik

import (
	"portfolio-app/internal/proyek"
	"fmt"
)

// StatistikKategori menghitung jumlah proyek berdasarkan kategori
func StatistikKategori(proyekList []proyek.Proyek) map[string]int {
	return proyek.StatistikKategori(proyekList)
}

// TampilkanStatistik menampilkan statistik berdasarkan kategori
func TampilkanStatistik(statistik map[string]int) {
	for kategori, jumlah := range statistik {
		fmt.Printf("%s: %d proyek\n", kategori, jumlah)
	}
}

package fitur

import (
	"portofolio_app/data"
	"fmt"
)

// Menampilkan statistik kategori proyek
func StatistikKategori(proyek []data.Proyek) map[string]int {
	statistik := make(map[string]int)
	for _, p := range proyek {
		statistik[p.Kategori]++
	}
	return statistik
}

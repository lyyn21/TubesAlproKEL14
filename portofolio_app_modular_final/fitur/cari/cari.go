package cari

import (
	"portfolio-app/internal/proyek"
	"strings"
)

// CariProyekNama mencari proyek berdasarkan nama menggunakan pencarian linier
func CariProyekNama(proyekList []proyek.Proyek, nama string) []proyek.Proyek {
	var hasil []proyek.Proyek
	for _, p := range proyekList {
		if p.Nama == nama {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

// CariProyekKategori mencari proyek berdasarkan kategori menggunakan pencarian biner
func CariProyekKategori(proyekList []proyek.Proyek, kategori string) []proyek.Proyek {
	return proyek.CariProyekKategori(proyekList, kategori)
}



// internal/crud/crud.go

package crud

import (
	"portfolio-app/internal/proyek"
	"fmt"
)

// TambahProyek menambahkan proyek baru
func TambahProyek(proyekList []proyek.Proyek, nama, deskripsi, kategori, kesulitan, tanggal string, teknologi []string) []proyek.Proyek {
	proyekBaru := proyek.NewProyek(nama, deskripsi, kategori, kesulitan, tanggal, teknologi)
	return proyek.TambahProyek(proyekList, *proyekBaru)
}

// UbahProyek mengubah proyek berdasarkan nama
func UbahProyek(proyekList []proyek.Proyek, namaLama, namaBaru, deskripsi, kategori, kesulitan, tanggal string, teknologi []string) []proyek.Proyek {
	return proyek.UbahProyek(proyekList, namaLama, namaBaru, deskripsi, kategori, kesulitan, tanggal, teknologi)
}

// HapusProyek menghapus proyek berdasarkan nama
func HapusProyek(proyekList []proyek.Proyek, nama string) []proyek.Proyek {
	return proyek.HapusProyek(proyekList, nama)
}


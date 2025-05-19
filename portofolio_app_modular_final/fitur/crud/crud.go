package crud

import (
	"strings"
)

func Tambah(proyek []data.Proyek, p data.Proyek) []data.Proyek {
	return append(proyek, p)
}

func Ubah(proyek []data.Proyek, namaLama string, baru data.Proyek) []data.Proyek {
	for i, p := range proyek {
		if p.Nama == namaLama {
			proyek[i] = baru
		}
	}
	return proyek
}

func Hapus(proyek []data.Proyek, nama string) []data.Proyek {
	var hasil []data.Proyek
	for _, p := range proyek {
		if p.Nama != nama {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

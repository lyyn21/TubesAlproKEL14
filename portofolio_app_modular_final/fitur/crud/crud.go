package crud

import "portofolio_app/data"

// Menambah proyek baru
func Tambah(proyek []data.Proyek, p data.Proyek) []data.Proyek {
	return append(proyek, p)
}

// Mengubah proyek berdasarkan nama lama
func Ubah(proyek []data.Proyek, namaLama string, pBaru data.Proyek) []data.Proyek {
	for i, p := range proyek {
		if p.Nama == namaLama {
			proyek[i] = pBaru
		}
	}
	return proyek
}

// Menghapus proyek berdasarkan nama
func Hapus(proyek []data.Proyek, nama string) []data.Proyek {
	var hasil []data.Proyek
	for _, p := range proyek {
		if p.Nama != nama {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

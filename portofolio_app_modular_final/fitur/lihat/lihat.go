package lihat

import (
    "fmt"
    "portofolio_app/data"
)

func TampilkanSemua(proyek []data.Proyek) {
    fmt.Println("Daftar Semua Proyek:")
    for _, p := range proyek {
        fmt.Printf("Nama: %s | Kategori: %s | Tanggal: %s\n", p.Nama, p.Kategori, p.Tanggal)
    }
}
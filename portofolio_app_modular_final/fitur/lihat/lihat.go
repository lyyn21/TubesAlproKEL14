package lihat

import (
    "fmt"
)

func TampilkanSemua(proyek []data.Proyek) {
    fmt.Println("Daftar Semua Proyek:")
    for _, p := range proyek {
        fmt.Printf("Nama: %s | Kategori: %s | Tanggal: %s\n", p.Nama, p.Kategori, p.Tanggal)
    }
}

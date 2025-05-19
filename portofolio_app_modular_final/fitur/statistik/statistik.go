package statistik

import (
    "fmt"
    "portofolio_app/data"
)

func Jalankan(proyek []data.Proyek) {
    kategoriCount := make(map[string]int)
    for _, p := range proyek {
        kategoriCount[p.Kategori]++
    }

    fmt.Println("Statistik Kategori:")
    for kategori, jumlah := range kategoriCount {
        fmt.Printf("%s: %d proyek
", kategori, jumlah)
    }
}
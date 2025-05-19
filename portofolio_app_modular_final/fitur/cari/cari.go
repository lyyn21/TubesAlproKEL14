package cari

import (
    "fmt"
    "strings"
)

func Jalankan(proyek []data.Proyek) {
    var keyword string
    var pilih int
    fmt.Println("Cari berdasarkan:")
    fmt.Println("1. Nama")
    fmt.Println("2. Kategori")
    fmt.Print("Pilih: ")
    fmt.Scanln(&pilih)

    if pilih == 1 {
        fmt.Print("Masukkan nama proyek: ")
        fmt.Scanln(&keyword)
        for _, p := range proyek {
            if strings.ToLower(p.Nama) == strings.ToLower(keyword) {
                fmt.Printf("%s - %s
", p.Nama, p.Kategori)
            }
        }
    } else if pilih == 2 {
        fmt.Print("Masukkan kategori: ")
        fmt.Scanln(&keyword)
        for _, p := range proyek {
            if strings.ToLower(p.Kategori) == strings.ToLower(keyword) {
                fmt.Printf("%s - %s
", p.Nama, p.Kategori)
            }
        }
    }
}

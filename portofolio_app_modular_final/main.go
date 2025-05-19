package main

import (
    "fmt"
    "portofolio_app/data"
    "portofolio_app/fitur/crud"
    "portofolio_app/fitur/cari"
    "portofolio_app/fitur/urut"
    "portofolio_app/fitur/statistik"
    "portofolio_app/fitur/lihat"
    "portofolio_app/menu"
)

func main() {
    var proyek []data.Proyek

    for {
        menu.TampilkanMenu()
        var pilihan int
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            proyek = crud.Tambah(proyek)
        case 2:
            proyek = crud.Ubah(proyek)
        case 3:
            proyek = crud.Hapus(proyek)
        case 4:
            cari.Jalankan(proyek)
        case 5:
            urut.Jalankan(proyek)
        case 6:
            statistik.Jalankan(proyek)
        case 7:
            lihat.TampilkanSemua(proyek)
        case 0:
            fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
            return
        default:
            fmt.Println("Pilihan tidak valid, coba lagi.")
        }
    }
}
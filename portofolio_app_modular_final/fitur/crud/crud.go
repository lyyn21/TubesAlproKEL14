package crud

import (
    "fmt"
    "strings"
)

func Tambah(proyek []data.Proyek) []data.Proyek {
    var nama, deskripsi, kategori, kesulitan, tanggal, teknologiStr string

    fmt.Print("Nama Proyek: ")
    fmt.Scanln(&nama)
    fmt.Print("Deskripsi: ")
    fmt.Scanln(&deskripsi)
    fmt.Print("Kategori: ")
    fmt.Scanln(&kategori)
    fmt.Print("Kesulitan: ")
    fmt.Scanln(&kesulitan)
    fmt.Print("Tanggal (YYYY-MM-DD): ")
    fmt.Scanln(&tanggal)
    fmt.Print("Teknologi (pisahkan dengan koma): ")
    fmt.Scanln(&teknologiStr)
    teknologi := strings.Split(teknologiStr, ",")

    proyekBaru := data.Proyek{nama, deskripsi, teknologi, kategori, tanggal, kesulitan}
    return append(proyek, proyekBaru)
}

func Ubah(proyek []data.Proyek) []data.Proyek {
    var namaLama, nama, deskripsi, kategori, kesulitan, tanggal, teknologiStr string
    fmt.Print("Nama proyek yang ingin diubah: ")
    fmt.Scanln(&namaLama)
    fmt.Print("Nama Baru: ")
    fmt.Scanln(&nama)
    fmt.Print("Deskripsi: ")
    fmt.Scanln(&deskripsi)
    fmt.Print("Kategori: ")
    fmt.Scanln(&kategori)
    fmt.Print("Kesulitan: ")
    fmt.Scanln(&kesulitan)
    fmt.Print("Tanggal (YYYY-MM-DD): ")
    fmt.Scanln(&tanggal)
    fmt.Print("Teknologi (pisahkan dengan koma): ")
    fmt.Scanln(&teknologiStr)
    teknologi := strings.Split(teknologiStr, ",")

    for i, p := range proyek {
        if p.Nama == namaLama {
            proyek[i] = data.Proyek{nama, deskripsi, teknologi, kategori, tanggal, kesulitan}
        }
    }
    return proyek
}

func Hapus(proyek []data.Proyek) []data.Proyek {
    var nama string
    fmt.Print("Nama proyek yang ingin dihapus: ")
    fmt.Scanln(&nama)
    var hasil []data.Proyek
    for _, p := range proyek {
        if p.Nama != nama {
            hasil = append(hasil, p)
        }
    }
    return hasil
}

package urut

import (
    "fmt"
)

func Jalankan(proyek []data.Proyek) {
    var pilih int
    fmt.Println("Urutkan berdasarkan:")
    fmt.Println("1. Kesulitan")
    fmt.Println("2. Tanggal")
    fmt.Print("Pilih: ")
    fmt.Scanln(&pilih)

    if pilih == 1 {
        for i := 0; i < len(proyek)-1; i++ {
            minIdx := i
            for j := i + 1; j < len(proyek); j++ {
                if proyek[j].Kesulitan < proyek[minIdx].Kesulitan {
                    minIdx = j
                }
            }
            proyek[i], proyek[minIdx] = proyek[minIdx], proyek[i]
        }
    } else if pilih == 2 {
        for i := 1; i < len(proyek); i++ {
            current := proyek[i]
            j := i - 1
            for j >= 0 && proyek[j].Tanggal > current.Tanggal {
                proyek[j+1] = proyek[j]
                j--
            }
            proyek[j+1] = current
        }
    }

    fmt.Println("Hasil urutan:")
    for _, p := range proyek {
        fmt.Printf("%s - %s - %s
", p.Nama, p.Kategori, p.Tanggal)
    }
}

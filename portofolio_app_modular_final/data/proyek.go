package proyek

// Proyek adalah representasi data proyek dalam aplikasi.
type Proyek struct {
    Nama      string
    Deskripsi string
    Teknologi []string
    Kategori  string
    Tanggal   string
    Kesulitan string
}

// NewProyek adalah constructor untuk membuat proyek baru.
func NewProyek(nama, deskripsi, kategori, kesulitan, tanggal string, teknologi []string) *Proyek {
    return &Proyek{
        Nama:      nama,
        Deskripsi: deskripsi,
        Teknologi: teknologi,
        Kategori:  kategori,
        Tanggal:   tanggal,
        Kesulitan: kesulitan,
    }
}

// TambahProyek adalah fungsi untuk menambah proyek baru ke slice.
func TambahProyek(proyek []Proyek, p Proyek) []Proyek {
    return append(proyek, p)
}

// UbahProyek adalah fungsi untuk mengubah data proyek berdasarkan nama.
func UbahProyek(proyek []Proyek, namaLama string, namaBaru, deskripsi, kategori, kesulitan, tanggal string, teknologi []string) []Proyek {
    for i, p := range proyek {
        if p.Nama == namaLama {
            proyek[i] = Proyek{
                Nama:      namaBaru,
                Deskripsi: deskripsi,
                Teknologi: teknologi,
                Kategori:  kategori,
                Tanggal:   tanggal,
                Kesulitan: kesulitan,
            }
        }
    }
    return proyek
}

// HapusProyek adalah fungsi untuk menghapus proyek berdasarkan nama.
func HapusProyek(proyek []Proyek, nama string) []Proyek {
    var hasil []Proyek
    for _, p := range proyek {
        if p.Nama != nama {
            hasil = append(hasil, p)
        }
    }
    return hasil
}

// CariProyekNama adalah fungsi untuk mencari proyek berdasarkan nama menggunakan pencarian linier.
func CariProyekNama(proyek []Proyek, nama string) []Proyek {
    var hasil []Proyek
    for _, p := range proyek {
        if p.Nama == nama {
            hasil = append(hasil, p)
        }
    }
    return hasil
}

// CariProyekKategori adalah fungsi untuk mencari proyek berdasarkan kategori menggunakan pencarian binari.
func CariProyekKategori(proyek []Proyek, kategori string) []Proyek {
    var hasil []Proyek
    for _, p := range proyek {
        if p.Kategori == kategori {
            hasil = append(hasil, p)
        }
    }
    return hasil
}

// SortProyekKesulitan adalah fungsi untuk mengurutkan proyek berdasarkan tingkat kesulitan.
func SortProyekKesulitan(proyek []Proyek) {
    // Menggunakan Selection Sort untuk mengurutkan proyek berdasarkan kesulitan
    n := len(proyek)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if proyek[j].Kesulitan < proyek[minIdx].Kesulitan {
                minIdx = j
            }
        }
        proyek[i], proyek[minIdx] = proyek[minIdx], proyek[i]
    }
}

// SortProyekTanggal adalah fungsi untuk mengurutkan proyek berdasarkan tanggal.
func SortProyekTanggal(proyek []Proyek) {
    // Menggunakan Insertion Sort untuk mengurutkan proyek berdasarkan tanggal
    n := len(proyek)
    for i := 1; i < n; i++ {
        current := proyek[i]
        j := i - 1
        for j >= 0 && proyek[j].Tanggal > current.Tanggal {
            proyek[j+1] = proyek[j]
            j--
        }
        proyek[j+1] = current
    }
}

// StatistikKategori adalah fungsi untuk menghitung jumlah proyek berdasarkan kategori.
func StatistikKategori(proyek []Proyek) map[string]int {
    statistik := make(map[string]int)
    for _, p := range proyek {
        statistik[p.Kategori]++
    }
    return statistik
}

// TampilkanProyek adalah fungsi untuk menampilkan data proyek ke layar.
func TampilkanProyek(proyek []Proyek) {
    if len(proyek) == 0 {
        fmt.Println("Belum ada proyek yang terdaftar.")
        return
    }
    for _, p := range proyek {
        fmt.Printf("Nama      : %s\n", p.Nama)
        fmt.Printf("Deskripsi : %s\n", p.Deskripsi)
        fmt.Printf("Kategori  : %s\n", p.Kategori)
        fmt.Printf("Kesulitan : %s\n", p.Kesulitan)
        fmt.Printf("Tanggal   : %s\n", p.Tanggal)
        fmt.Printf("Teknologi : %s\n", strings.Join(p.Teknologi, ", "))
        fmt.Println("----------------------------")
    }
}

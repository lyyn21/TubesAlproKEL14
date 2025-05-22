package lihat

import (
	"portfolio-app/internal/proyek"
	"fmt"
)

// TampilkanSemuaProyek menampilkan semua proyek yang ada
func TampilkanSemuaProyek(proyekList []proyek.Proyek) {
	proyek.TampilkanProyek(proyekList)
}

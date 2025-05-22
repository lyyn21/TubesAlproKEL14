package urut

import (
	"portfolio-app/internal/proyek"
	"portfolio-app/internal/sort"
)

// UrutProyekKesulitan mengurutkan proyek berdasarkan kesulitan
func UrutProyekKesulitan(proyekList []proyek.Proyek) {
	sort.SortProyekKesulitan(proyekList)
}

// UrutProyekTanggal mengurutkan proyek berdasarkan tanggal
func UrutProyekTanggal(proyekList []proyek.Proyek) {
	sort.SortProyekTanggal(proyekList)
}


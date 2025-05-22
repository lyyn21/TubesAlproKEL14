package fitur

import (
	"portofolio_app/data"
)

// Mengurutkan proyek berdasarkan kesulitan (Selection Sort)
func SortingSelection(proyek []data.Proyek) {
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

// Mengurutkan proyek berdasarkan tanggal (Insertion Sort)
func SortingInsertion(proyek []data.Proyek) {
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

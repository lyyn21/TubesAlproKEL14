package fitur

import "portofolio_app/data"

func SortingSelection(proyek []data.Proyek) {
	for i := 0; i < len(proyek)-1; i++ {
		min := i
		for j := i + 1; j < len(proyek); j++ {
			if proyek[j].Kesulitan < proyek[min].Kesulitan {
				min = j
			}
		}
		proyek[i], proyek[min] = proyek[min], proyek[i]
	}
}

func SortingInsertion(proyek []data.Proyek) {
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

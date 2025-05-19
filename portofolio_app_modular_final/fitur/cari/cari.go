package fitur

import (
	"strings"
)

func BinarySearch(proyek []data.Proyek, keyword string, byKategori bool) []data.Proyek {
	var hasil []data.Proyek
	low, high := 0, len(proyek)-1
	// Sort dulu
	SortingSelection(proyek)
	for low <= high {
		mid := (low + high) / 2
		var banding string
		if byKategori {
			banding = proyek[mid].Kategori
		} else {
			banding = proyek[mid].Nama
		}
		if strings.ToLower(banding) == strings.ToLower(keyword) {
			hasil = append(hasil, proyek[mid])
			break
		} else if strings.ToLower(banding) < strings.ToLower(keyword) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return hasil
}

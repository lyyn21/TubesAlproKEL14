package fitur

import (
	"portofolio_app/data"
	"strings"
)

// Mencari proyek berdasarkan nama (Sequential Search)
func SequentialSearch(proyek []data.Proyek, keyword string) []data.Proyek {
	var hasil []data.Proyek
	lowerKeyword := strings.ToLower(keyword)
	for _, p := range proyek {
		if strings.ToLower(p.Nama) == lowerKeyword {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

// Mencari proyek berdasarkan kategori (Binary Search)
func BinarySearch(proyek []data.Proyek, keyword string, berdasarkanKategori bool) []data.Proyek {
	if berdasarkanKategori {
		sort.Slice(proyek, func(i, j int) bool {
			return strings.ToLower(proyek[i].Kategori) < strings.ToLower(proyek[j].Kategori)
		})
	}

	var hasil []data.Proyek
	low, high := 0, len(proyek)-1
	target := strings.ToLower(keyword)

	for low <= high {
		mid := (low + high) / 2
		if strings.ToLower(proyek[mid].Kategori) == target {
			hasil = append(hasil, proyek[mid])
			break
		}
		if strings.ToLower(proyek[mid].Kategori) < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return hasil
}


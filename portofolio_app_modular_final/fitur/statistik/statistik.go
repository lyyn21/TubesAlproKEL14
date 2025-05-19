package fitur

import "portofolio_app/data"

func StatistikKategori(proyek []data.Proyek) map[string]int {
	stat := make(map[string]int)
	for _, p := range proyek {
		stat[p.Kategori]++
	}
	return stat
}

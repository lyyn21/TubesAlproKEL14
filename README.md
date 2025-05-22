# Aplikasi Manajemen Proyek Data Science

## Deskripsi

Aplikasi ini dirancang untuk membantu Anda mengelola proyek-proyek dalam bidang data science, seperti menyimpan, mengubah, menghapus, mencari, mengurutkan, dan menampilkan statistik dari proyek-proyek tersebut. Aplikasi ini dibangun menggunakan bahasa pemrograman Go (Golang) yang efisien, serta memanfaatkan struktur data dan algoritma dasar untuk pengoperasian yang sederhana namun efektif.

## Fitur Utama

1. **Manajemen Proyek Data Science**:
   
   Aplikasi ini memungkinkan pengguna untuk mengelola proyek data science mereka. Fitur yang tersedia mencakup:
   
   - **Menambah proyek** baru ke dalam daftar.
   - **Mengubah informasi** proyek yang telah ada.
   - **Menghapus proyek** yang tidak diperlukan.
   
   Setiap proyek dapat mencakup informasi penting seperti:
   - **Nama proyek**
   - **Deskripsi proyek**
   - **Teknologi yang digunakan**
   - **Kategori proyek** (misalnya: Machine Learning, Data Visualization, dll)
   - **Tanggal pembuatan proyek**
   - **Tingkat kesulitan** (misalnya: Mudah, Menengah, Sulit)

2. **Pencatatan Keahlian yang Dipelajari**:

   Aplikasi ini memungkinkan pengguna untuk mencatat **keahlian teknis** yang dipelajari selama pengerjaan proyek. Ini termasuk teknologi dan alat yang digunakan, yang dapat membantu pengguna untuk memantau dan mengevaluasi perkembangan keterampilan mereka di bidang data science.

3. **Pencarian Proyek yang Efisien**:

   Pengguna dapat mencari proyek berdasarkan **nama** atau **kategori** dengan menggunakan dua metode pencarian yang efisien:
   
   - **Sequential Search**: Digunakan untuk pencarian linear berdasarkan nama proyek.
   - **Binary Search**: Digunakan untuk pencarian lebih cepat berdasarkan kategori proyek (setelah data terurut).

4. **Pengurutan Proyek**:

   Aplikasi ini menyediakan fitur pengurutan proyek berdasarkan dua kriteria:
   
   - **Tingkat Kesulitan**: Mengurutkan proyek berdasarkan tingkat kesulitan menggunakan **Selection Sort**.
   - **Tanggal Pembuatan**: Mengurutkan proyek berdasarkan tanggal pembuatan menggunakan **Insertion Sort**.

5. **Statistik Kategori Proyek**:

   Aplikasi ini menampilkan statistik jumlah proyek yang telah diselesaikan berdasarkan kategori tertentu, seperti **Machine Learning**, **Data Visualization**, dan lainnya. Fitur ini memberikan wawasan yang jelas mengenai keahlian yang telah dipelajari dan dikuasai oleh pengguna.

---

Dengan aplikasi ini, pengguna dapat dengan mudah **mengorganisir**, **menilai**, dan **menunjukkan** kemampuan mereka dalam bidang data science melalui proyek-proyek yang telah mereka kerjakan. Aplikasi ini sangat bermanfaat bagi pemula yang ingin membangun portofolio proyek yang komprehensif dan profesional untuk mendukung karier mereka di industri data science.

### Contoh Struktur `Proyek`

```go
type Proyek struct {
    Nama      string
    Deskripsi string
    Teknologi []string
    Kategori  string
    Tanggal   string
    Kesulitan string
}





#### Kesimpulan
Aplikasi Manajemen Proyek Data Science ini adalah solusi yang sederhana namun sangat efektif untuk mengelola berbagai proyek dalam bidang data science. Dengan berbagai fitur seperti menambah, mengubah, menghapus, mencari, dan mengurutkan proyek, aplikasi ini memberikan fleksibilitas dalam mengelola portofolio proyek Anda dengan cara yang terstruktur.

Selain itu, aplikasi ini memungkinkan pengguna untuk dengan mudah melihat statistik berdasarkan kategori proyek yang ada, memberikan wawasan lebih mendalam mengenai distribusi proyek berdasarkan kategori tertentu. Dengan menggunakan algoritma dasar seperti Selection Sort, Insertion Sort, dan Binary Search, aplikasi ini tetap efisien meskipun menangani dataset yang relatif kecil hingga menengah.

Aplikasi ini dibangun dengan menggunakan bahasa pemrograman Go (Golang), yang terkenal dengan kinerja tinggi dan kemudahan penggunaannya. Meskipun aplikasi ini berbasis teks, antarmukanya sangat intuitif dan mudah digunakan, menjadikannya pilihan yang baik bagi mereka yang ingin memulai dengan alat yang sederhana dan efektif untuk mengelola proyek.

Harapan kami, aplikasi ini dapat membantu Anda dalam mengorganisasi dan mengelola proyek-proyek data science dengan lebih mudah dan terstruktur. Kami juga terbuka untuk kontribusi dari komunitas pengembang guna terus meningkatkan aplikasi ini.

Terima kasih telah menggunakan aplikasi ini. Semoga bermanfaat untuk kebutuhan manajemen proyek Anda!

Kelompok 14

 # Aplikasi Manajemen Proyek Data Science

## Deskripsi

Aplikasi ini dirancang untuk membantu Anda mengelola proyek-proyek dalam bidang data science, seperti menyimpan, mengubah, menghapus, mencari, mengurutkan, dan menampilkan statistik dari proyek-proyek tersebut. Aplikasi ini dibangun menggunakan bahasa pemrograman Go (Golang) yang efisien, serta memanfaatkan struktur data dan algoritma dasar untuk pengoperasian yang sederhana namun efektif.

## Fitur Utama

### 1. **Manajemen Proyek Data Science**
   
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

### 2. **Pencatatan Keahlian yang Dipelajari**

   Aplikasi ini memungkinkan pengguna untuk mencatat **keahlian teknis** yang dipelajari selama pengerjaan proyek. Ini termasuk teknologi dan alat yang digunakan, yang dapat membantu pengguna untuk memantau dan mengevaluasi perkembangan keterampilan mereka di bidang data science.

### 3. **Pencarian Proyek yang Efisien**
   
   Pengguna dapat mencari proyek berdasarkan **nama** atau **kategori** dengan menggunakan dua metode pencarian yang efisien:
   
   - **Sequential Search**: Digunakan untuk pencarian linear berdasarkan nama proyek.
   - **Binary Search**: Digunakan untuk pencarian lebih cepat berdasarkan kategori proyek (setelah data terurut).

### 4. **Pengurutan Proyek**
   
   Aplikasi ini menyediakan fitur pengurutan proyek berdasarkan dua kriteria:
   
   - **Tingkat Kesulitan**: Mengurutkan proyek berdasarkan tingkat kesulitan menggunakan **Selection Sort**.
   - **Tanggal Pembuatan**: Mengurutkan proyek berdasarkan tanggal pembuatan menggunakan **Insertion Sort**.

### 5. **Statistik Kategori Proyek**
   
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


Kelompok 14

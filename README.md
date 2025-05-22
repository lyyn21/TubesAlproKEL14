# Aplikasi Manajemen Portofolio Data Science untuk Pemula

## Deskripsi

Aplikasi ini dirancang untuk membantu pengguna dalam mencatat dan memantau proyek-proyek data science yang telah mereka kerjakan. Tujuan utama aplikasi ini adalah untuk membangun portofolio yang kuat dan terstruktur, yang dapat digunakan untuk menunjukkan kemampuan serta pengalaman dalam bidang data science. Aplikasi ini sangat cocok bagi pemula yang ingin memantau perkembangan keterampilan mereka melalui proyek yang telah diselesaikan.

## Fitur Utama
1. **Tambah Proyek**: Menambah proyek baru ke dalam daftar proyek.
2. **Ubah Proyek**: Mengubah informasi proyek yang sudah ada.
3. **Hapus Proyek**: Menghapus proyek berdasarkan nama.
4. **Cari Proyek**:
   - **Pencarian berdasarkan Nama** (Sequential Search)
   - **Pencarian berdasarkan Kategori** (Binary Search)
5. **Urutkan Proyek**:
   - Mengurutkan proyek berdasarkan **Kesulitan** (menggunakan Selection Sort)
   - Mengurutkan proyek berdasarkan **Tanggal** (menggunakan Insertion Sort)
6. **Statistik Kategori**: Menampilkan statistik mengenai jumlah proyek berdasarkan kategori.
7. **Lihat Semua Proyek**: Menampilkan daftar semua proyek yang ada.

## Struktur Data
Aplikasi ini menggunakan struktur data `Proyek`, yang memiliki beberapa atribut penting, seperti:

- **Nama**: Nama proyek.
- **Deskripsi**: Deskripsi singkat proyek.
- **Teknologi**: Teknologi yang digunakan dalam proyek.
- **Kategori**: Kategori proyek (misal: Machine Learning, Data Analysis, dll).
- **Tanggal**: Tanggal pembuatan proyek.
- **Kesulitan**: Tingkat kesulitan proyek (misal: Mudah, Menengah, Sulit).


### Penjelasan Fungsi

1. **Fitur CRUD (Create, Read, Update, Delete)**:
   - Aplikasi ini memungkinkan Anda untuk menambah proyek baru, mengubah proyek yang ada, dan menghapus proyek. Setiap operasi CRUD ini menggunakan slice untuk menambah, mengganti, atau menghapus elemen dalam koleksi proyek.
   
2. **Pencarian dan Pengurutan**:
   - Pencarian dilakukan dengan metode **Sequential Search** (untuk nama proyek) dan **Binary Search** (untuk kategori proyek setelah data diurutkan).
   - Pengurutan menggunakan **Selection Sort** untuk mengurutkan berdasarkan kesulitan dan **Insertion Sort** untuk mengurutkan berdasarkan tanggal. Meskipun algoritma ini efisien untuk dataset kecil hingga sedang, Anda bisa menggantinya dengan algoritma lain jika diperlukan.

3. **Statistik Kategori**:
   - Fungsi statistik akan membantu Anda untuk melihat seberapa banyak proyek yang dimiliki dalam setiap kategori. Ini dapat berguna untuk melacak distribusi proyek Anda.

Dengan penjelasan ini, pengguna bisa lebih memahami bagaimana aplikasi bekerja dan bagaimana cara berinteraksi dengan antarmuka berbasis teks untuk mengelola proyek-proyek data science mereka.


## Kelompok 14


# Program Tagihan Internet (Go)

## Deskripsi Project
Program Tagihan Internet adalah aplikasi berbasis Command Line Interface (CLI) sederhana yang dibangun menggunakan bahasa pemrograman Go (Golang). Project ini dibuat untuk membantu menghitung rincian tagihan layanan internet pelanggan secara otomatis berdasarkan paket yang dipilih, jumlah pemakaian kuota (GB), dan jumlah perangkat terhubung, serta menyediakan fitur simulasi perbandingan jika pelanggan ingin mengganti paket langganannya.

## Fitur-Fitur Program
* **Perhitungan Tagihan Otomatis**: Menghitung biaya dasar (abodemen), kuota progresif berdasarkan pemakaian data (GB), dan denda kelebihan perangkat terhubung.
* **Sistem Diskon & Pajak**: Memberikan diskon sebesar 10% untuk tagihan kotor minimal Rp2.000.000 dengan status pembayaran lancar, serta menambahkan PPN sebesar 11%.
* **Peringatan Pemblokiran Otomatis**: Menampilkan notifikasi pemblokiran jika pemakaian kuota melebihi 50 GB dengan status pembayaran tidak lancar.
* **Simulasi Perubahan Paket Langganan**: Memungkinkan pengguna membandingkan rincian biaya dasar, batas perangkat, dan denda antara paket lama dengan paket baru.

## Daftar Paket Layanan

| Nama Paket | Biaya Dasar (Abodemen) | Batas Perangkat | Denda Kelebihan Perangkat |
| :--- | :--- | :--- | :--- |
| **STARTER** | Rp150.000 | 100 perangkat | Rp75.000 |
| **BUSINESS** | Rp500.000 | 500 perangkat | Rp200.000 |
| **ENTERPRISE** | Rp1.500.000 | 2.000 perangkat | Rp0 (Gratis) |

## Teknologi yang Digunakan
* **Bahasa Pemrograman**: Go (Golang)
* **Standard Library**: `fmt` (input/output terminal) dan `strings` (manipulasi & validasi teks)
* **Version Control**: Git & GitHub

## Cara Menjalankan Program

1. Pastikan **Go** sudah terinstal di komputer Anda (cek dengan perintah `go version`).
2. Clone repositori ini atau buka folder proyek di terminal/command prompt.
3. Jalankan perintah berikut untuk mengeksekusi program:
   ```bash
   go run main.go
   ```
   

## Identitas Mahasiswa
* **Nama**: I Made Dwi Laksmana Putra
* **NIM**: 42630027

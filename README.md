# Program Tagihan Internet (Go)

Aplikasi berbasis CLI sederhana menggunakan bahasa pemrograman Go (Golang) untuk menghitung rincian tagihan layanan internet beserta simulasi perubahan paket langganan.

## Fitur Utama
* **Perhitungan Tagihan Otomatis**: Menghitung biaya dasar (abodemen), kuota progresif berdasarkan pemakaian data (GB), dan denda kelebihan perangkat.
* **Sistem Diskon & Pajak**: Memberikan diskon 10% untuk tagihan kotor minimal Rp2.000.000 dengan status pembayaran lancar, serta menambahkan PPN sebesar 11%.
* **Peringatan Pemblokiran**: Menampilkan notifikasi pemblokiran otomatis jika pemakaian melebihi 50 GB dengan status pembayaran tidak lancar.
* **Pembaruan Paket Langganan**: Memungkinkan pengguna membandingkan detail biaya dan batas perangkat antara paket lama dengan paket baru.

## Daftar Paket Layanan

| Nama Paket | Biaya Dasar (Abodemen) | Batas Perangkat | Denda Kelebihan Perangkat |
| :--- | :--- | :--- | :--- |
| **STARTER** | Rp150.000 | 100 perangkat | Rp75.000 |
| **BUSINESS** | Rp500.000 | 500 perangkat | Rp200.000 |
| **ENTERPRISE** | Rp1.500.000 | 2.000 perangkat | Rp0 (Gratis) |

## Cara Menjalankan Program

1. Pastikan **Go** sudah terinstal di komputer Anda.
2. Buka terminal di folder proyek ini.
3. Jalankan perintah berikut:
   ```bash
   go run main.go
   ```
## Author
* **Dwi Laksana**

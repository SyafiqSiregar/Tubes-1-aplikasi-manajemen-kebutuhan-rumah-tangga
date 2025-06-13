# 🚀 Aplikasi Manajemen Stok Bahan

> **Pantry Keeper**

## 📖 Deskripsi

Aplikasi Manajemen Stok Bahan adalah sistem inventory management yang dikembangkan menggunakan bahasa pemrograman Go. Aplikasi ini dirancang khusus untuk membantu pengelolaan stok bahan makanan dengan interface yang intuitif dan fitur-fitur lengkap untuk monitoring dan tracking inventory.

### ✨ Fitur Utama

🔐 **Sistem Autentikasi**
- Register akun baru dengan validasi username unik
- Login dengan sistem keamanan password
- Manajemen multi-user

📊 **Manajemen Stok**
- Tambah bahan baru ke inventory
- Update stok bahan yang sudah ada
- Kurangi stok dengan tracking otomatis
- Monitoring bahan hampir habis (stok ≤ 5)

🔍 **Pencarian & Sortir**
- Pencarian bahan dengan Sequential Search
- Sorting stok dari terendah ke tertinggi (Selection Sort)
- Filter dan tampilan data yang terorganisir

💰 **Laporan Keuangan**
- Tracking total pengeluaran
- Laporan pengeluaran bulanan
- Riwayat transaksi lengkap dengan tanggal

📈 **Analisis Data**
- Binary Search untuk pencarian transaksi berdasarkan bulan
- Insertion Sort untuk pengurutan transaksi berdasarkan tanggal
- Dashboard monitoring real-time

## 🛠️ Teknologi yang Digunakan

- **Bahasa**: Go (Golang)
- **Algoritma Sorting**: Selection Sort, Insertion Sort
- **Algoritma Searching**: Sequential Search, Binary Search
- **Database**: In-memory arrays (No external database required)
- **UI**: Console-based dengan ASCII art yang menarik

## 🚀 Cara Instalasi

### Prerequisites
- Go 1.19 atau versi lebih baru
- Terminal/Command Prompt
- Git (untuk cloning repository)

### Langkah Instalasi

1. **Clone Repository**
   ```bash
   git clone https://github.com/username/aplikasi-manajemen-stok.git
   cd aplikasi-manajemen-stok
   ```

2. **Compile & Run**
   ```bash
   go run bagian3.go
   ```

   Atau compile terlebih dahulu:
   ```bash
   go build -o inventory-app bagian3.go
   ./inventory-app
   ```

## 📱 Cara Penggunaan

### 1. Registrasi & Login
- Jalankan aplikasi
- Pilih menu **Register** untuk membuat akun baru
- Masukkan username dan password
- Login menggunakan kredensial yang telah dibuat

### 2. Mengelola Stok
```
📋 Menu Utama:
1. Lihat Stok Bahan        - Melihat semua stok (terurut dari terendah)
2. Tambah Bahan           - Menambah bahan baru atau update stok
3. Kurangi Stok           - Mengurangi stok bahan
4. Lihat Bahan Hampir Habis - Monitor stok kritis (≤ 5)
5. Cari Bahan             - Pencarian bahan spesifik
6. Lihat Total Pengeluaran - Dashboard keuangan
7. Laporan Pengeluaran    - Laporan bulanan detail
8. Riwayat Transaksi      - History semua transaksi
```

### 3. Format Input
- **Tanggal**: DD-MM-YYYY (contoh: 15-12-2024)
- **Bulan**: YYYY-MM (contoh: 2024-12)
- **Harga**: Angka desimal (contoh: 15000.50)

## 📊 Struktur Data

### BahanMakanan
```go
type BahanMakanan struct {
    Nama  string   // Nama bahan makanan
    Stok  int      // Jumlah stok tersedia
    Harga float64  // Harga per satuan
}
```

### Transaction
```go
type Transaction struct {
    Nama    string   // Nama bahan
    Jenis   string   // "Tambah" atau "Kurangi"
    Jumlah  int      // Kuantitas transaksi
    Harga   float64  // Harga saat transaksi
    Tanggal string   // Tanggal transaksi (YYYY-MM-DD)
}
```

## 🎯 Algoritma yang Diimplementasikan

| Algoritma | Fungsi | 
|-----------|--------|--------------|
| **Selection Sort** | Mengurutkan stok dari terendah | 
| **Insertion Sort** | Mengurutkan transaksi berdasarkan tanggal | 
| **Sequential Search** | Mencari bahan dalam inventory | 
| **Binary Search** | Mencari transaksi berdasarkan bulan | 

## 🔧 Konfigurasi

Konstanta yang dapat disesuaikan:
```go
const NMAX int = 1024           // Maksimal bahan makanan
const MAX_USERS int = 100       // Maksimal pengguna
const MAX_TRANSACTIONS int = 1000 // Maksimal transaksi
```

## 📸 Screenshots

```
  ╔══════════════════════════════════════════════════════╗
  ║   __        _______ _     ____ ___  __  __ _____     ║
  ║   \ \      / / ____| |   / ___/ _ \|  \/  | ____|    ║
  ║    \ \ /\ / /|  _| | |  | |  | | | | |\/| |  _|      ║
  ║     \ V  V / | |___| |__| |__| |_| | |  | | |___     ║
  ║      \_/\_/  |_____|_____\____\___/|_|  |_|_____|    ║
  ╚══════════════════════════════════════════════════════╝
```

## 👥 Tim Pengembang

- **Syafiq Yusuf Ikhsan** - Lead Developer
- **Dafa Izul Haq** - Co-Developer

*Dikembangkan untuk mata kuliah Algoritma Pemrograman 2025*



## 📞 Kontak & Support

Jika ada pertanyaan atau bug report, silakan:
- Email: syafiqsiregar17@gmail.com
- Email: dafa.i.haq@gmail.com

---

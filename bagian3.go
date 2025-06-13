package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const NMAX int = 1024
const MAX_USERS int = 100
const MAX_TRANSACTIONS int = 1000

type BahanMakanan struct {
	Nama  string
	Stok  int
	Harga float64
}

type User struct {
	Username string
	Password string
}

type Transaction struct {
	Nama    string
	Jenis   string
	Jumlah  int
	Harga   float64
	Tanggal string
}

type arrBahan [NMAX]BahanMakanan

var users [MAX_USERS]User
var userCount int
var transactions [MAX_TRANSACTIONS]Transaction

var transactionCount int
var username, password string
var pilihan int
var dataBahan arrBahan
var dataCount int
var totalPengeluaran float64

func main() {
	start()
}

func start() {
	/*
		IS: Variable pilihan bertipe integer, belum terisi nilai dari user
		FS:  Program menampilkan menu utama dan menerima input pilihan dari user.
	*/
	clearScreen()
	fmt.Println("  ╔══════════════════════════════════════════════════════╗")
	fmt.Println("  ║   __        _______ _     ____ ___  __  __ _____     ║")
	fmt.Println("  ║   \\ \\      / / ____| |   / ___/ _ \\|  \\/  | ____|    ║")
	fmt.Println("  ║    \\ \\ /\\ / /|  _| | |  | |  | | | | |\\/| |  _|      ║")
	fmt.Println("  ║     \\ V  V / | |___| |__| |__| |_| | |  | | |___     ║")
	fmt.Println("  ║      \\_/\\_/  |_____|_____\\____\\___/|_|  |_|_____|    ║")
	fmt.Println("  ╚══════════════════════════════════════════════════════╝")

	fmt.Println("      ╔═══════════════════════════════════════════╗")
	fmt.Println("      ║          🚀 Aplikasi Manajemen            ║")
	fmt.Println("      ║    Created by Syafiq Yusuf Ikhsan &       ║")
	fmt.Println("      ║             Dafa Izul Haq                 ║")
	fmt.Println("      ║      Algoritma Pemrograman 2025           ║")
	fmt.Println("      ╠═══════════════════════════════════════════╣")
	fmt.Println("      ║   🌟 APLIKASI MANAGEMENT STOK BAHAN 🌟    ║")
	fmt.Println("      ╠═══════════════════════════════════════════╣")
	fmt.Println("      ║                1. Register                ║")
	fmt.Println("      ║                2. Login                   ║")
	fmt.Println("      ║                0. Keluar                  ║")
	fmt.Println("      ╚═══════════════════════════════════════════╝")
	fmt.Print("                   👉 Pilihan Anda : ")

	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		register()
	case 2:
		login()
	case 0:
		fmt.Println("   ╔═══════════════════════════════════════════════════╗")
		fmt.Println("   ║ 🙏 Terima kasih telah menggunakan aplikasi ini!!  ║")
		fmt.Println("   ╚═══════════════════════════════════════════════════╝")
	default:
		fmt.Println("\n      ❌ Masukan tidak valid, silahkan coba lagi.")
		start()
	}
}

func daftarLayanan() {
	/*
		IS: Variable pilihan bertipe integer
		FS:  menu tampil, pilihan diterima dan fungsi dipanggil sesuai input

	*/
	var pilihan int

	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║             🚀 Main Menu               ║")
	fmt.Println("        ╠════════════════════════════════════════╣")
	fmt.Println("        ║  1. Lihat Stok Bahan                   ║")
	fmt.Println("        ║  2. Tambah Bahan                       ║")
	fmt.Println("        ║  3. Kurangi Stok                       ║")
	fmt.Println("        ║  4. Lihat Bahan Hampir Habis           ║")
	fmt.Println("        ║  5. Cari Bahan                         ║")
	fmt.Println("        ║  6. Lihat Total Pengeluaran            ║")
	fmt.Println("        ║  7. Laporan Pengeluaran Bulanan        ║")
	fmt.Println("        ║  8. Lihat Riwayat Transaksi            ║")
	fmt.Println("        ║  0. Keluar                             ║")
	fmt.Println("        ╚════════════════════════════════════════╝")
	fmt.Print("                   👉 Pilihan Anda : ")
	fmt.Scan(&pilihan)
	clearScreen()

	switch pilihan {
	case 0:
		fmt.Println("  ╔═══════════════════════════════════════════════════╗")
		fmt.Println("  ║ 🙏 Terima kasih telah menggunakan aplikasi ini!!  ║")
		fmt.Println("  ╚═══════════════════════════════════════════════════╝")
	case 1:
		lihatStok(dataBahan, dataCount)
	case 2:
		tambahBahan(&dataBahan, &dataCount, &totalPengeluaran, &transactions, &transactionCount)
	case 3:
		kurangiStok(&dataBahan, dataCount, &transactions, &transactionCount)
	case 4:
		lihatHampirHabis(dataBahan, dataCount)
	case 5:
		cariBahan(dataBahan, dataCount)
	case 6:
		lihatPengeluaranBulanan(totalPengeluaran)
	case 7:
		laporanPengeluaranBulanan(transactions, transactionCount)
	case 8:
		riwayatTransaksi(transactions, transactionCount)
	default:
		fmt.Println("\nMasukan tidak valid, silahkan coba lagi")
		daftarLayanan()
	}
}

func register() {
	/*
		IS: Daftar pengguna (users) terdefinisi.
		FS: Menambahkan akun baru ke daftar pengguna jika username belum digunakan.
	*/
	var username, password, confirmPassword string
	var i int

	clearScreen()
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║      🚀 Register Akun       ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                   👤 Username : ")
	fmt.Scan(&username)
	fmt.Print("                   🔑 Password : ")
	fmt.Scan(&password)
	fmt.Print("                   🔑 Confirm Pw : ")
	fmt.Scan(&confirmPassword)

	if password != confirmPassword {
		fmt.Println("\n                 Password tidak cocok!")
		start()
		return
	}

	for i = 0; i < userCount; i++ {
		if users[i].Username == username {
			fmt.Println("\n                 Username sudah digunakan!")
			start()
			return
		}
	}

	users[userCount] = User{Username: username, Password: password}
	userCount++
	fmt.Println("\n              Akun berhasil didaftarkan!")
	start()
}

func login() {
	/*
		IS: inputUsername dan inputPassword
		FS: Mengakses menu layanan jika username dan password sesuai, atau menampilkan pesan kesalahan.

	*/
	var inputUsername, inputPassword string
	var i int

	clearScreen()
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║        🚀 Login Akun        ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                   👤 Username : ")
	fmt.Scan(&inputUsername)
	fmt.Print("                   🔑 Password : ")
	fmt.Scan(&inputPassword)
	clearScreen()

	for i = 0; i < userCount; i++ {
		if users[i].Username == inputUsername && users[i].Password == inputPassword {
			fmt.Println("\n              ======= WELCOME ", users[i].Username, " =======")
			daftarLayanan()
			return
		}
	}

	fmt.Println("\n           ❌ Username atau password salah!")
	start()
}

// (select short)
func lihatStok(dataBahan arrBahan, dataCount int) {
	/*
		IS: Data bahan makanan tersedia.
		FS: Daftar bahan makanan ditampilkan dalam urutan stok terendah ke tertinggi (selection sort).
		(select short)
	*/
	var i, j, minIdx int
	var tempBahan BahanMakanan

	for i = 0; i < dataCount-1; i++ {
		minIdx = i
		for j = i + 1; j < dataCount; j++ {
			if dataBahan[j].Stok < dataBahan[minIdx].Stok {
				minIdx = j
			}
		}
		tempBahan = dataBahan[i]
		dataBahan[i] = dataBahan[minIdx]
		dataBahan[minIdx] = tempBahan
	}
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║      Daftar Stok Bahan :    ║")
	fmt.Println("              ╚═════════════════════════════╝")
	for i = 0; i < dataCount; i++ {
		fmt.Printf("                        %d . %s = %d\n", i+1, dataBahan[i].Nama, dataBahan[i].Stok)
	}

	daftarLayanan()
}

func cariBahan(dataBahan arrBahan, dataCount int) {
	/*
		IS: Data bahan makanan tersedia.
		FS: Daftar bahan makanan ditampilkan dalam urutan stok terendah ke tertinggi(squential search)
	*/
	var nama string
	var i int

	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║     Cari Bahan (Linear)     ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                      Nama Bahan : ")
	fmt.Scan(&nama)

	for i = 0; i < dataCount; i++ {
		if strings.EqualFold(dataBahan[i].Nama, nama) {
			fmt.Println()
			fmt.Println("                   ✅ Bahan Ditemukan:")
			fmt.Printf("\n            %s - Stok: %d - Harga: Rp %.2f\n", dataBahan[i].Nama, dataBahan[i].Stok, dataBahan[i].Harga)
			daftarLayanan()
			return
		}
	}
	fmt.Println("\n                   ❌ Bahan tidak ditemukan.")
	daftarLayanan()
}

func tambahBahan(dataBahan *arrBahan, dataCount *int, totalPengeluaran *float64, transactions *[MAX_TRANSACTIONS]Transaction, transactionCount *int) {
	/*
		IS: totalPengeluaran bertipe float64 dan telah terdefinisi sebagai total pengeluaran yang tercatat.
		FS: Menambahkan bahan baru ke stok dan mencatat transaksi.
	*/
	var nama, tanggal string
	var jumlah, i int
	var harga float64
	var formattedDate string

	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║        Tambah Bahan :       ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                     Nama Bahan : ")
	fmt.Scan(&nama)
	fmt.Print("                     Jumlah     : ")
	fmt.Scan(&jumlah)
	fmt.Print("                     Harga per Satuan: Rp ")
	fmt.Scan(&harga)
	fmt.Print("                     Tanggal (DD-MM-YYYY): ")
	fmt.Scan(&tanggal)

	parsedTime, err := time.Parse("02-01-2006", tanggal)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Gunakan format DD-MM-YYYY.")
		daftarLayanan()
		return
	}
	formattedDate = parsedTime.Format("2006-01-02")

	for i = 0; i < *dataCount; i++ {
		if strings.EqualFold((*dataBahan)[i].Nama, nama) {
			(*dataBahan)[i].Stok += jumlah
			(*dataBahan)[i].Harga = harga
			*totalPengeluaran += float64(jumlah) * harga

			(*transactions)[*transactionCount] = Transaction{
				Nama:    nama,
				Jenis:   "Tambah",
				Jumlah:  jumlah,
				Harga:   harga,
				Tanggal: formattedDate,
			}
			*transactionCount++

			fmt.Println("                   ✅ Stok diperbarui.")
			daftarLayanan()
			return
		}
	}

	(*dataBahan)[*dataCount] = BahanMakanan{Nama: nama, Stok: jumlah, Harga: harga}
	*dataCount++
	*totalPengeluaran += float64(jumlah) * harga

	(*transactions)[*transactionCount] = Transaction{
		Nama:    nama,
		Jenis:   "Tambah",
		Jumlah:  jumlah,
		Harga:   harga,
		Tanggal: formattedDate,
	}
	*transactionCount++

	fmt.Println("                   ✅ Bahan ditambahkan.")
	daftarLayanan()
}

func kurangiStok(dataBahan *arrBahan, dataCount int, transactions *[MAX_TRANSACTIONS]Transaction, transactionCount *int) {
	/*
		IS: Bahan ditemukan dan stok cukup untuk dikurangi.
		FS: Mengurangi stok bahan dan mencatat transaksi.
	*/
	var nama, tanggal string
	var jumlah, i int
	var formattedDate string

	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║        Kurangi Bahan :      ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                     Nama Bahan : ")
	fmt.Scan(&nama)
	fmt.Print("                     Jumlah     : ")
	fmt.Scan(&jumlah)
	fmt.Print("                     Tanggal (DD-MM-YYYY): ")
	fmt.Scan(&tanggal)

	parsedTime, err := time.Parse("02-01-2006", tanggal)
	if err != nil {
		fmt.Println("❌ Format tanggal tidak valid. Gunakan format DD-MM-YYYY.")
		daftarLayanan()
		return
	}
	formattedDate = parsedTime.Format("2006-01-02")

	for i = 0; i < dataCount; i++ {
		if strings.EqualFold((*dataBahan)[i].Nama, nama) {
			if (*dataBahan)[i].Stok >= jumlah {
				(*dataBahan)[i].Stok -= jumlah

				(*transactions)[*transactionCount] = Transaction{
					Nama:    nama,
					Jenis:   "Kurangi",
					Jumlah:  jumlah,
					Harga:   (*dataBahan)[i].Harga,
					Tanggal: formattedDate,
				}
				*transactionCount++

				fmt.Println("                   ✅ Stok dikurangi.")
			} else {
				fmt.Println("                   ❌ Stok tidak cukup!")
			}
			daftarLayanan()
			return
		}
	}

	fmt.Println("                   ❌ Bahan tidak ditemukan.")
	daftarLayanan()
}

func laporanPengeluaranBulanan(transactions [MAX_TRANSACTIONS]Transaction, transactionCount int) {
	/*
		IS: transactions berisi data transaksi
		FS: menampilkan laporan transaksi dan total pengeluaran di bulan yang dicari
	*/
	var bulan string
	var start, end int
	var total float64
	var t Transaction
	var i int

	sortTransaksi(transactions, transactionCount)

	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║      📊 Laporan Pengeluaran Bulanan    ║")
	fmt.Println("        ╚════════════════════════════════════════╝")
	fmt.Print("             Masukkan Bulan (YYYY-MM): ")
	fmt.Scan(&bulan)

	start, end = binarySearchBulan(transactions, transactionCount, bulan)

	if start == -1 || end == -1 {
		fmt.Println("\n     ❌ Tidak ada transaksi pada bulan tersebut.")
		daftarLayanan()
		return
	}

	total = 0.0
	fmt.Println("  ╔═════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║ Nama       │ Jenis   │ Jumlah │ Harga      │ Tanggal            ║")
	fmt.Println("  ╠═════════════════════════════════════════════════════════════════╣")

	for i = 0; i < transactionCount; i++ {
		t = transactions[i]
		if strings.HasPrefix(t.Tanggal, bulan) {
			fmt.Printf("  ║ %-10s │ %-7s │ %6d │ Rp %8.2f │ %s        ║\n",
				t.Nama, t.Jenis, t.Jumlah, t.Harga, t.Tanggal)
			total += float64(t.Jumlah) * t.Harga
		}
	}

	fmt.Println("  ╚═════════════════════════════════════════════════════════════════╝")
	fmt.Printf("        💰 Total Pengeluaran Bulan %s: Rp %.2f\n", bulan, total)
	daftarLayanan()
}

// isertion sort
func sortTransaksi(transactions [MAX_TRANSACTIONS]Transaction, transactionCount int) {
	/*
		IS: transactions sudah terdefinisi dengan sejumlah data transaksi.
		FS: transactions terurut ascending berdasarkan tanggal transaksi.
		(isertion sort)
	*/
	var i int
	var j int

	for i = 1; i < transactionCount; i++ {
		key := transactions[i]
		j = i - 1
		for j >= 0 && transactions[j].Tanggal > key.Tanggal {
			transactions[j+1] = transactions[j]
			j--
		}
		transactions[j+1] = key
	}
}

func lihatHampirHabis(dataBahan arrBahan, dataCount int) {
	/*
		IS: dataBahan terdefinisi dengan dataCount sebagai jumlahnya.
		FS: Menampilkan bahan dengan stok <= 5.
	*/
	var i int

	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║      Bahan Hampir Habis (Stok <= 5)    ║")
	fmt.Println("        ╚════════════════════════════════════════╝")

	for i = 0; i < dataCount; i++ {
		if dataBahan[i].Stok <= 5 {
			fmt.Printf("                      - %s      (%d)\n", dataBahan[i].Nama, dataBahan[i].Stok)
		}
	}
	daftarLayanan()
}

func lihatPengeluaranBulanan(totalPengeluaran float64) {
	/*
		IS: totalPengeluaran terdefinisi.
		FS: menampilkan Total pengeluaran bulanan.
	*/

	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║        Total Pengeluaran Bulanan       ║")
	fmt.Println("        ╚════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("                     Total: Rp %.2f\n", totalPengeluaran)
	daftarLayanan()
}

func riwayatTransaksi(transactions [MAX_TRANSACTIONS]Transaction, transactionCount int) {
	/*
		IS: Slice transactions sudah terisi jika ada transaksi.
		FS: Menampilkan daftar seluruh transaksi yang pernah dilakukan.
	*/
	var i int
	var t Transaction

	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║        📊 Riwayat Transaksi            ║")
	fmt.Println("        ╚════════════════════════════════════════╝")
	if transactionCount == 0 {
		fmt.Println("\n                   ❌ Tidak ada transaksi.")
		daftarLayanan()
		return
	}
	fmt.Println()
	fmt.Println("	Nama	Jenis	Jumlah	Harga		Tanggal")
	fmt.Println("	──────────────────────────────────────────────────")
	for i = 0; i < transactionCount; i++ {
		t = transactions[i]
		fmt.Printf("	%s	%s	%d	Rp %.2f	%s\n",
			t.Nama, t.Jenis, t.Jumlah, t.Harga, t.Tanggal)
	}
	fmt.Println()
	daftarLayanan()
}

func binarySearchBulan(transactions [MAX_TRANSACTIONS]Transaction, transactionCount int, bulan string) (int, int) {
	/*
		IS: transactions terdefinisi dengan beberapa transaksi.
		FS: Mengembalikan indeks awal dan akhir transaksi untuk bulan yang diberikan.
	*/
	var end, start int

	low, high := 0, transactionCount-1
	start, end = -1, -1

	for low <= high {
		mid := (low + high) / 2
		if strings.HasPrefix(transactions[mid].Tanggal, bulan) {
			start = mid
			high = mid - 1
		} else if transactions[mid].Tanggal < bulan {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low, high = 0, transactionCount-1
	for low <= high {
		mid := (low + high) / 2
		if strings.HasPrefix(transactions[mid].Tanggal, bulan) {
			end = mid
			low = mid + 1
		} else if transactions[mid].Tanggal < bulan {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return start, end
}

func clearScreen() {
	/* IS: -
	   FS: Mengosongkan layar.
	*/
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Platform tidak didukung.")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

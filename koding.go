package main

import (
	"fmt"
)

type Buku struct {
	NamaBuku   string
	JumlahBuku int
}

var (
	daftarBuku = []Buku{
		{"Pemrograman", 10},
		{"Film", 5},
		{"Printing", 20},
	}
	histori []string
)

// Menu
func menu() {
	fmt.Println("Menu Program")
	fmt.Println("1. Lihat informasi Pengguna Program")
	fmt.Println("2. Lihat Daftar Buku")
	fmt.Println("3. Tambah Daftar Buku")
	fmt.Println("4. Tambah Peminjaman Buku")
	fmt.Println("5. Histori Peminjaman Buku")
	fmt.Println("6. Keluar dari Program")
}

func ListDaftarBuku() {
	fmt.Println("1. Pemrograman")
	fmt.Println("2. Film")
	fmt.Println("3. Printing")
}

// 1) Lihat Informasi Pengguna
func LihatInformasiPenggunaProgram() {
	fmt.Println("=========================================")
	fmt.Println("Informasi Pengguna Program")
	fmt.Println("Username: Kayla")
	fmt.Println("Password: 2406357690")
	fmt.Println("Jenis Kelamin: Perempuan")
	fmt.Println("Makanan Favorit: Fettucine")
	fmt.Println("Minuman Favorit: Mactha Latte")
	fmt.Println("=========================================")
	fmt.Println("")
}

// 2) Daftar Buku
func LihatDaftarBuku() {
	fmt.Println("=========================================")
	fmt.Println("Daftar Buku")
	for _, buku := range daftarBuku {
		fmt.Printf("Nama Buku: %s, Jumlah: %d\n", buku.NamaBuku, buku.JumlahBuku)
	}
	fmt.Println("=========================================")
	fmt.Println("")
}

// 3) Tambah Daftar Buku
func TambahDaftarBuku() {
	var namaBuku string
	var jumlahBuku int

	fmt.Println("=========================================")
	fmt.Println("Tambahkan Daftar Buku")
	ListDaftarBuku()
	fmt.Print("Masukkan nama buku yang ingin ditambahkan: ")
	fmt.Scan(&namaBuku)
	fmt.Print("Masukkan jumlah buku yang ingin ditambahkan: ")
	fmt.Scan(&jumlahBuku)

	if jumlahBuku <= 0 {
		fmt.Println("Jumlah buku harus lebih besar dari 0.")
		return
	} else {
		found := false
		for i, buku := range daftarBuku {
			if buku.NamaBuku == namaBuku {
				daftarBuku[i].JumlahBuku += jumlahBuku
				histori = append(histori, fmt.Sprintf("Ditambahkan: %s, Jumlah: %d", buku.NamaBuku, jumlahBuku))
				fmt.Println("Buku berhasil ditambahkan.")
				found = true
				break
			}
		}
		if !found {
			daftarBuku = append(daftarBuku, Buku{NamaBuku: namaBuku, JumlahBuku: jumlahBuku})
			histori = append(histori, fmt.Sprintf("Ditambahkan: %s, Jumlah: %d", namaBuku, jumlahBuku))
			fmt.Println("Buku baru berhasil ditambahkan.")
		}
	}
	fmt.Println("=========================================")
	fmt.Println("")
}

// 4) Tambah Peminjaman Buku
func TambahPeminjamanBuku() {
	var namaBuku string
	var jumlahPinjam int

	fmt.Println("=========================================")
	fmt.Println("Pinjam Buku")
	ListDaftarBuku()
	fmt.Print("Nama buku yang ingin dipinjam: ")
	fmt.Scan(&namaBuku)
	fmt.Print("Jumlah buku yang ingin dipinjam: ")
	fmt.Scan(&jumlahPinjam)
	for i, buku := range daftarBuku {
		if buku.NamaBuku == namaBuku {
			if jumlahPinjam <= 0 {
				fmt.Println("Jumlah peminjaman harus lebih besar dari 0")
				return
			} else if buku.JumlahBuku < jumlahPinjam {
				fmt.Println("Jumlah buku tidak mencukupi!.")
				return
			} else {
				daftarBuku[i].JumlahBuku -= jumlahPinjam
				histori = append(histori, fmt.Sprintf("Dipinjam: %s, Jumlah %d", buku.NamaBuku, jumlahPinjam))
				fmt.Println("Peminjaman buku berhasil.")
				return
			}
		}
	}
	fmt.Println("=========================================")
	fmt.Println("")
}

// 5) Histori Peminjaman Buku
func HistoriPeminjamanBuku() {
	fmt.Println("=========================================")
	fmt.Println("Histori Peminjaman Buku")
	for _, pinjaman := range histori {
		fmt.Println(pinjaman)
	}
	fmt.Println("=========================================")
	fmt.Println("")
}

// 6) Keluar Dari Program
func KeluarDariProgram() {
	fmt.Println("=========================================")
	fmt.Println("= Terima kasih Sudah Datang =")
	fmt.Println("=========================================")
}

func main() {
	var username, password string

	fmt.Println("=========================================")
	fmt.Println("= Selamat Datang di Perpustakaan Vokasi =")
	fmt.Println("=========================================")

	fmt.Print("Silahkan Input Username: ")
	fmt.Scanln(&username)
	fmt.Print("Silahkan Input Password: ")
	fmt.Scanln(&password)
	fmt.Println("=========================================")

	if username != "Kayla" || password != "2406357690" {
		fmt.Println("Login Tidak Berhasil.")
		return
	}
	fmt.Println("Login Sukses!!!!")
	for {
		menu()
		var pilihan int
		fmt.Print("Input Menu yang Anda Inginkan: ")
		fmt.Scanln(&pilihan)
		fmt.Println("=====================================")
		fmt.Println("")

		switch pilihan {
		case 1:
			LihatInformasiPenggunaProgram()
		case 2:
			LihatDaftarBuku()
		case 3:
			TambahDaftarBuku()
		case 4:
			TambahPeminjamanBuku()
		case 5:
			HistoriPeminjamanBuku()
		case 6:
			KeluarDariProgram()
			return
		default:
			fmt.Println("Menu tidak valid, mohon coba lagi.")
		}
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	for {
		fmt.Println("=== Menu ===")
		fmt.Println("1. Mulai install")
		fmt.Println("2. Exit")
		fmt.Print("Pilih menu (1/2): ")

		var choice int
		fmt.Scanln(&choice) // Meminta pengguna untuk memilih menu dan menyimpan pilihannya di `choice`

		switch choice {
		case 1:
			fmt.Println("=== Mulai install ===")
			fmt.Print("Masukkan URL file yang ingin diunduh: ")
			var url string
			fmt.Scanln(&url) // Meminta pengguna untuk memasukkan URL file yang ingin diunduh
			err := downloadFile(url) // Memanggil fungsi `downloadFile` dengan URL yang dimasukkan pengguna
			if err != nil {
				fmt.Printf("Gagal mengunduh file: %v\n", err) // Menampilkan pesan kesalahan jika terjadi error saat mengunduh
			} else {
				fmt.Println("File berhasil diunduh.") // Menampilkan pesan sukses jika file berhasil diunduh
			}
		case 2:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return // Keluar dari program jika pengguna memilih menu "Exit"
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1 atau 2.") // Menampilkan pesan jika pengguna memilih menu yang tidak valid
		}
		fmt.Println()
	}
}

func downloadFile(url string) error {
	fileName := filepath.Base(url) // Mengambil nama file dari URL yang diberikan

	resp, err := http.Get(url) // Mengirim permintaan GET untuk mengunduh file dari URL
	if err != nil {
		return err // Mengembalikan error jika permintaan GET mengalami masalah
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName) // Membuat file kosong dengan nama file yang telah diambil dari URL
	if err != nil {
		return err // Mengembalikan error jika tidak dapat membuat file
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body) // Menyalin isi dari respons HTTP ke file yang telah dibuat
	if err != nil {
		return err // Mengembalikan error jika terjadi masalah saat menyalin
	}

	return nil // Mengembalikan nil jika pengunduhan file berhasil
}

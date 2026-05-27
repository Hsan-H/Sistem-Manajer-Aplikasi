package main

import "fmt"

// Kamus Global
const NMAX int = 100 // Batas Array
type turnamen struct {
	nama     string
	password string
}
type tabTurnamen [NMAX]turnamen

func main() {
	// Kamus Lokal
	var pilihan string

	// Algoritma
	pilihan = "bebas"
	fmt.Println("Aplikasi Manajer Turnamen")
	fmt.Println("Menu:")
	fmt.Println("1. Registrasi Turnamen")
	fmt.Println("2. Login Turnamen")
	fmt.Println("3. Hapus Turnamen")
	fmt.Println("4. Daftar Turnamen")
	fmt.Println("5. Keluar")
	fmt.Print("---> Pilihan Menu:")
	fmt.Scanln(&pilihan)
}

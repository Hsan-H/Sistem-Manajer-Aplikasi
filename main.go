package main

import (
	"bufio"
	"fmt"
	"os"
)

// Kamus global
// Konstanta untuk jumlah maksimum turnamen
const NMAX = 100

// Inisialisasi data turnamen dan array untuk menyimpan data turnamen
type turnamen struct {
	id         string
	name       string
	password   string
	skorMenang int
	skorKalah  int
	nPlayer    int
	pemenang   string
}
type tabTurnamen [NMAX]turnamen

func main() {
	// Kamus lokal
	var turnamen tabTurnamen
	var pilihan, n int
	var keluar bool

	// Algoritma
	// Atur jumlah turnamen saat ini
	n = 0
	// Menu utama
	for keluar = false; !keluar; {
		fmt.Println("Menu:")
		fmt.Println("1. Registrasi Turnamen")
		fmt.Println("2. Login Turnamen")
		fmt.Println("3. Hapus Turnamen")
		fmt.Println("4. Cari Turnamen")
		fmt.Println("5. List Turnamen")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			RegistrasiTurnamen(&turnamen, &n)
		case 2:
			LoginTurnamen(&turnamen, n)
		case 3:
			HapusTurnamen(&turnamen, &n)
		case 4:
			CariTurnamen(&turnamen, n)
		case 5:
			ListTurnamen(&turnamen, n)
		case 6:
			fmt.Println("Terima kasih.")
			keluar = true
		default:
			fmt.Println("Salah input, mas.")
		}
	}
}

func RegistrasiTurnamen(dataTurnamen *tabTurnamen, n *int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menginput data turnamen baru ke dalam array dataTurnamen dan menambahkan jumlah turnamen n
	*/
	// Kamus lokal
	var name string
	var scanner *bufio.Scanner

	// Algoritma
	// Menu registrasi turnamen
	fmt.Println("Registrasi")
	fmt.Print("Masukkan ID Turnamen: ")
	fmt.Scan(&dataTurnamen[*n].id)
	// Bersihkan input buffer
	fmt.Scanln()
	// Settingan nama turnamen menggunakan bufio.Scanner
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama Turnamen: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	// Simpan nama turnamen ke dalam data turnamen
	dataTurnamen[*n].name = name
	fmt.Print("Masukkan Password Turnamen: ")
	fmt.Scan(&dataTurnamen[*n].password)
	fmt.Println("Setting skor awal...")
	fmt.Print("Masukkan skor menang: ")
	fmt.Scan(&dataTurnamen[*n].skorMenang)
	fmt.Print("Masukkan skor kalah: ")
	fmt.Scan(&dataTurnamen[*n].skorKalah)
	// Inisialisasi jumlah player dan pemenang
	dataTurnamen[*n].nPlayer = 0
	dataTurnamen[*n].pemenang = ""
	// Tambahkan jumlah turnamen
	*n++
	// Konfirmasi registrasi berhasil
	fmt.Println("Registrasi berhasil.")
	// Masuk ke dalam menu turnamen
	MenuTurnamen()
}

func LoginTurnamen(dataTurnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menginput id dan password untuk login ke dalam turnamen
	*/
	// Kamus lokal
	var id string
	var password string
	var pilihan, idx int
	var ketemu bool

	// Algoritma
	// Menu login turnamen
	fmt.Println("Login")
	fmt.Print("Masukkan id Turnamen: ")
	fmt.Scanln(&id)
	// Cari index turnamen berdasarkan ID menggunakan sequential search
	idx = sequentialSearchbyID(*dataTurnamen, id, n)
	if idx != -1 {
		// Inisialisasi variabel
		// Loop untuk verifikasi password
		for ketemu = false; !ketemu; {
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)
			if password == dataTurnamen[idx].password {
				fmt.Println("Login berhasil.")
				ketemu = true
				MenuTurnamen()
			} else {
				// Jika password salah, tawarkan pilihan untuk mengulang atau keluar
				fmt.Println("Salah, mas. Mau ngulang atau keluar?")
				fmt.Println("1. Ulangi")
				fmt.Println("2. Keluar")
				fmt.Print("Pilih: ")
				fmt.Scan(&pilihan)
				if pilihan == 2 {
					fmt.Println("Keluar dari login.")
					ketemu = true
				}
			}
		}
	} else {
		fmt.Println("Turnamen-nya gak ada.")
	}
}

func HapusTurnamen(dataTurnamen *tabTurnamen, n *int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menginput id untuk mencari turnamen yang akan dihapus dan menghapus turnamen tersebut dari array dataTurnamen
	*/
	// Kamus lokal
	var id string
	var idx, i int

	// Algoritma
	// Menu hapus turnamen
	fmt.Println("Hapus Turnamen")
	// Minta input ID turnamen
	fmt.Print("Masukkan id Turnamen: ")
	fmt.Scanln(&id)
	// Cari index turnamen berdasarkan ID menggunakan sequential search
	idx = sequentialSearchbyID(*dataTurnamen, id, *n)
	if idx != -1 {
		// Hapus turnamen dengan menggeser elemen setelah index ke kiri
		for i = idx; i < *n-1; i++ {
			dataTurnamen[i] = dataTurnamen[i+1]
		}
		// Kurangi jumlah turnamen
		*n--
		fmt.Println("Turnamen berhasil dihapus.")
	} else {
		fmt.Println("Turnamen tidak ditemukan.")
	}
}

func CariTurnamen(dataTurnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menginput id untuk mencari turnamen dan menampilkan informasi turnamen yang ditemukan berdasarkan id tersebut
	*/
	// Kamus lokal
	var id string
	var idx int

	// Algoritma
	// Menu cari turnamen
	fmt.Print("Masukkan id Turnamen yang ingin dicari: ")
	fmt.Scanln(&id)
	// Cari index turnamen berdasarkan ID menggunakan sequential search
	idx = sequentialSearchbyID(*dataTurnamen, id, n)
	if idx != -1 {
		// Tampilkan informasi turnamen yang ditemukan
		fmt.Println("\rData Turnamen ", dataTurnamen[idx].name)
		fmt.Printf("ID: %s\n", dataTurnamen[idx].id)
		fmt.Printf("Skor Menang: %d\n", dataTurnamen[idx].skorMenang)
		fmt.Printf("Skor Kalah: %d\n", dataTurnamen[idx].skorKalah)
		fmt.Printf("Jumlah Player: %d Player\n", dataTurnamen[idx].nPlayer)
		fmt.Printf("Pemenang: %s\n", dataTurnamen[idx].pemenang)
	} else {
		fmt.Println("Turnamen tidak ditemukan.")
	}
}
func ListTurnamen(dataTurnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menampilkan daftar turnamen yang terdaftar di dalam array dataTurnamen
	*/
	// Kamus lokal
	var i int

	// Algoritma
	// Menu list turnamen
	if n == 0 {
		fmt.Println("Belum ada turnamen yang terdaftar.")
	} else {
		// Urutkan turnamen berdasarkan id menggunakan insertion sort
		ascInsertionSortTurnamenByID(dataTurnamen, n)
		fmt.Println("Daftar Turnamen:")
		for i = 0; i < n; i++ {
			fmt.Printf("%d. %s \n", i+1, dataTurnamen[i].name)
			fmt.Println("   ID:", dataTurnamen[i].id)
			fmt.Println("   Jumlah Player:", dataTurnamen[i].nPlayer, "Player")
			fmt.Println("   Pemenang:", dataTurnamen[i].pemenang)
			fmt.Println()
		}
	}
}
func MenuTurnamen() {}

// Kumpulan fungsi sorting dan searching
func ascInsertionSortTurnamenByID(Turnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array Turnamen yang berisi n
		F.S mengurutkan array turnamen berdasarkan id yang bersifat string secara ascending
	*/

	// Kamus Lokal
	var i, j int
	var x turnamen

	// Algoritma
	for i = 0; i < n; i++ {
		x = Turnamen[i]
		j = i - 1
		for j >= 0 && Turnamen[j].id > x.id {
			Turnamen[j+1] = Turnamen[j]
			j = j - 1
		}
		Turnamen[j+1] = x
	}
}

func binarySearch(Turnamen *tabTurnamen, ketemu *bool, index *int, n int, nama string) {
	/*
		I.S terdefinisi array Turnamen yang berisi n, nama sebagai variabel yang dicari, index sebagai hasil index apabila ketemu
		F.S mencari nilai turnamen berdasarkan nama sehingga ditemukan nilai index dan ketemu
	*/

	// Kamus Lokal
	var left, mid, right int

	// Algoritma
	left = 0
	right = n - 1
	mid = (left + right) / 2
	for left <= right && Turnamen[mid].nama != nama {
		if nama < Turnamen[mid].nama {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	// Memastikan nama turnamen tersebut ada
	*ketemu = mid >= 0 && Turnamen[mid].nama == nama

	// Menyimpan mid sebagai index
	*index = mid
}

func sequentialSearchbyID(Turnamen tabTurnamen, id string, n int) int {
	/*
		I.S terdefinisi array Turnamen yang berada di index idx yang array Playernya berjumlah n dan variabel id sebagai variabel yang dicari
		F.S mengembalikan index dari id di dalam data nama Player
	*/

	// Kamus Lokal
	var i, idx int
	var ketemu bool

	// Algoritma
	idx = -1
	ketemu = false
	i = 0
	for i < n && !ketemu {
		if id == Turnamen[i].id {
			idx = i
			ketemu = true
		}
		i++
	}
	return idx
}

func selectionSort(Player *tabPlayer, n int) {
	/*
		I.S terdefinisi Player berjumlah n
		F.S mengurutkan array Player berdasarkan skor
	*/

	// Kamus Lokal
	var i, j, nilaimax int
	// Algoritma
	for i = 0; i < n; i++ {
		nilaimax = i
		for j = i + 1; j < n; j++ {
			if Player[j].skor > Player[nilaimax].skor {
				nilaimax = j
			}
		}
		if nilaimax != i {
			Player[i], Player[nilaimax] = Player[nilaimax], Player[i]
		}
	}
}

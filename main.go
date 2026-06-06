package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Kamus global
// Konstanta untuk jumlah maksimum turnamen
const NMAX = 100

// Inisialisasi data turnamen dan array untuk menyimpan data turnamen
type turnamen struct {
	id, name, password, pemenang   string
	skorMenang, skorKalah, nPemain int
	pemain                         tabPemain
}
type tabTurnamen [NMAX]turnamen

// Inisialisasi data pemain dan array untuk menyimpan data pemain
type pemain struct {
	id, name            string
	menang, kalah, skor int
}
type tabPemain [NMAX]pemain

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
			time.Sleep(3 * time.Second)
			keluar = true
		default:
			fmt.Println("Salah input, mas.")
			time.Sleep(3 * time.Second)
		}
	}
}

func RegistrasiTurnamen(dataTurnamen *tabTurnamen, n *int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi nO
		F.S menginput data turnamen baru ke dalam array dataTurnamen dan menambahkan jumlah turnamen n
	*/
	// Kamus lokal
	var name string
	var idx int
	var scanner *bufio.Scanner

	// Algoritma
	// Menu registrasi turnamen
	if *n >= NMAX {
		fmt.Println("Maaf, total yang dapat didaftarkan adalah 100.")
		time.Sleep(3 * time.Second)
	} else {
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
		dataTurnamen[*n].nPemain = 0
		dataTurnamen[*n].pemenang = ""
		// Simpan index turnamen dari n
		idx = *n
		// Tambahkan jumlah turnamen
		*n++
		// Konfirmasi registrasi berhasil
		fmt.Println("Registrasi berhasil.")
		time.Sleep(3 * time.Second)
		// Masuk ke dalam menu turnamen
		RegistrasiPemain(dataTurnamen, idx)
		MenuTurnamen(dataTurnamen, idx)
	}
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
	var ketemu, keluar bool

	// Algoritma
	// Menu login turnamen
	fmt.Println("Login")
	fmt.Print("Masukkan id Turnamen: ")
	fmt.Scan(&id)
	fmt.Scanln()
	// Cari index turnamen berdasarkan ID menggunakan sequential search
	idx = sequentialSearchbyID(*dataTurnamen, id, n)
	if idx != -1 {
		// Inisialisasi variabel
		// Loop untuk verifikasi password
		for ketemu = false; !ketemu; {
			fmt.Print("Masukkan password: ")
			fmt.Scan(&password)
			fmt.Scanln()
			if password == dataTurnamen[idx].password {
				fmt.Println("Login berhasil.")
				ketemu = true
				time.Sleep(3 * time.Second)
				MenuTurnamen(dataTurnamen, idx)
			} else {
				// Jika password salah, tawarkan pilihan untuk mengulang atau keluar
				for keluar = false; !keluar; {
					fmt.Println("Salah, mas. Mau ngulang atau keluar?")
					fmt.Println("1. Ulangi")
					fmt.Println("2. Keluar")
					fmt.Print("Pilih: ")
					fmt.Scan(&pilihan)
					switch pilihan {
					case 1:
						// Ulangi input password
						fmt.Println("Silakan coba lagi.")
						time.Sleep(3 * time.Second)
						keluar = true
					case 2:
						fmt.Println("Keluar dari login.")
						time.Sleep(3 * time.Second)
						ketemu = true
						keluar = true
					default:
						fmt.Println("Salah input, mas.")
						time.Sleep(3 * time.Second)
					}
				}
			}
		}
	} else {
		fmt.Println("Turnamen-nya gak ada.")
		time.Sleep(3 * time.Second)
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
	fmt.Scan(&id)
	fmt.Scanln()
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
		time.Sleep(3 * time.Second)
	} else {
		fmt.Println("Turnamen tidak ditemukan.")
		time.Sleep(3 * time.Second)
	}
}

func CariTurnamen(dataTurnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menginput id untuk mencari turnamen dan menampilkan informasi turnamen yang ditemukan berdasarkan id tersebut
	*/
	// Kamus lokal
	var id, apapun string
	var idx int

	// Algoritma
	// Menu cari turnamen
	fmt.Print("Masukkan id Turnamen yang ingin dicari: ")
	fmt.Scan(&id)
	fmt.Scanln()
	// Cari index turnamen berdasarkan ID menggunakan sequential search
	idx = sequentialSearchbyID(*dataTurnamen, id, n)
	if idx != -1 {
		// Tampilkan informasi turnamen yang ditemukan
		fmt.Println("\rData Turnamen ", dataTurnamen[idx].name)
		fmt.Printf("ID: %s\n", dataTurnamen[idx].id)
		fmt.Printf("Skor Menang: %d\n", dataTurnamen[idx].skorMenang)
		fmt.Printf("Skor Kalah: %d\n", dataTurnamen[idx].skorKalah)
		fmt.Printf("Jumlah Player: %d Player\n", dataTurnamen[idx].nPemain)
		fmt.Printf("Pemenang: %s\n", dataTurnamen[idx].pemenang)
		fmt.Print("Kembali ke menu utama? (tekan apapun) ")
		fmt.Scan(&apapun)

	} else {
		fmt.Println("Turnamen tidak ditemukan.")
		time.Sleep(3 * time.Second)
	}
}
func ListTurnamen(dataTurnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi n
		F.S menampilkan daftar turnamen yang terdaftar di dalam array dataTurnamen
	*/
	// Kamus lokal
	var i int
	var apapun string

	// Algoritma
	// Menu list turnamen
	if n == 0 {
		fmt.Println("Belum ada turnamen yang terdaftar.")
		time.Sleep(3 * time.Second)
	} else {
		// Urutkan turnamen berdasarkan id menggunakan insertion sort
		ascInsertionSortTurnamenByID(dataTurnamen, n)
		fmt.Println("Daftar Turnamen:")
		for i = 0; i < n; i++ {
			fmt.Printf("%d. %s \n", i+1, dataTurnamen[i].name)
			fmt.Println("   ID:", dataTurnamen[i].id)
			fmt.Println("   Jumlah Player:", dataTurnamen[i].nPemain, "Player")
			fmt.Println("   Pemenang:", dataTurnamen[i].pemenang)
			fmt.Println()
		}
		fmt.Print("Kembali ke menu utama? (tekan apapun) ")
		fmt.Scan(&apapun)
	}
}
func RegistrasiPemain(dataTurnamen *tabTurnamen, idx int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi idx dan jumlah player di dalam turnamen tersebut
		F.S menginput data pemain baru ke dalam array pemain di dalam dataTurnamen pada index idx dan menambahkan jumlah player di turnamen tersebut
	*/
	// kamus lokal
	var scanner *bufio.Scanner
	var name string
	var pilihan int

	// algoritma
	if dataTurnamen[idx].nPemain >= NMAX {
		fmt.Println("Maaf, total player yang dapat didaftarkan adalah 100.")
		time.Sleep(3 * time.Second)
	} else {
		for keluar := false; !keluar; {
			fmt.Println("Menu Registrasi Pemain:")
			fmt.Println("1. Tambah Pemain")
			fmt.Println("2. Selesai")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				fmt.Print("Masukkan ID Pemain: ")
				fmt.Scan(&dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].id)
				// Bersihkan input buffer
				fmt.Scanln()
				// Settingan nama Pemain menggunakan bufio.Scanner
				scanner = bufio.NewScanner(os.Stdin)
				fmt.Print("Masukkan Nama Player: ")
				if scanner.Scan() {
					name = scanner.Text()
				}
				// Simpan nama player ke dalam data turnamen
				dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].name = name
				fmt.Print("Masukkan jumlah menang: ")
				fmt.Scan(&dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].menang)
				fmt.Print("Masukkan jumlah kalah: ")
				fmt.Scan(&dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].kalah)
				// Hitung skor player berdasarkan jumlah menang dan kalah
				dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].skor = dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].menang*dataTurnamen[idx].skorMenang + dataTurnamen[idx].pemain[dataTurnamen[idx].nPemain].kalah*dataTurnamen[idx].skorKalah
				// Tambahkan jumlah player di turnamen
				dataTurnamen[idx].nPemain++
				fmt.Println("Player berhasil ditambahkan.")
			case 2:
				descSelectionSortBySkor(&dataTurnamen[idx].pemain, dataTurnamen[idx].nPemain)
				dataTurnamen[idx].pemenang = pemenangTurnamen(dataTurnamen[idx].pemain, dataTurnamen[idx].nPemain)
				fmt.Println("Selesai registrasi player.")
				time.Sleep(3 * time.Second)
				keluar = true
			default:
				fmt.Println("Salah input, mas.")
			}
		}
	}
}
func pemenangTurnamen(pemain tabPemain, n int) string {
	/*
		I.S terdefinisi array pemain yang berjumlah n
		F.S mengembalikan nama sang juara
	*/

	// Kamus Lokal
	var i, j, k, max int
	var menang [NMAX]string

	// Algoritma
	// Mencari Nilai Ekstrim Maksimum untuk Mencari Nilai yang Juara
	if n != 0 {
		max = pemain[0].skor
		j = 1
		for j < n {
			if max < pemain[j].skor {
				max = pemain[j].skor
			}
			j++
		}
		for i = 0; i < n; i++ {
			if max == pemain[i].skor {
				menang[k] = pemain[i].name
				k++
			}
		}
		return rekursif(&menang, k)
	}
	return "-"
}
func rekursif(datapemenang *[NMAX]string, k int) string {
	/*
		I.S terdefinisi array datapemenang yang berjumlah k
		F.S mengembalikan nama dalam sebuah string sehingga apabila lebih dari satu pemenang dapat dikembalikan
	*/

	// Algoritma
	switch k {
	case 0:
		return "-"
	case 1:
		return datapemenang[0]
	default:
		return datapemenang[k-1] + ", " + rekursif(datapemenang, k-1)
	}
}
func MenuTurnamen(dataTurnamen *tabTurnamen, idx int) {
	/*
		I.S terdefinisi array dataTurnamen yang berisi idx dan jumlah player di dalam turnamen tersebut
		F.S menampilkan menu turnamen untuk mengelola data pemain, skor, dan ranking di dalam turnamen tersebut
	*/
	// kamus lokal
	var keluar bool
	var pilihan int

	// Algoritma
	for keluar = false; !keluar; {
		fmt.Println("Pemenang Turnamen: ", dataTurnamen[idx].pemenang)
		fmt.Println("Menu Turnamen:")
		fmt.Println("1. Registrasi Pemain")
		fmt.Println("2. Edit Pemain")
		fmt.Println("3. Edit Skor")
		fmt.Println("4. Ranking Pemain")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			RegistrasiPemain(dataTurnamen, idx)
		case 2:
			EditPemain(dataTurnamen, idx)
		case 3:
			EditSkor(dataTurnamen, idx)
		case 4:
			RankingPemain(dataTurnamen, idx)
		case 5:
			fmt.Println("Keluar dari menu turnamen.")
			time.Sleep(3 * time.Second)
			keluar = true
		default:
			fmt.Println("Salah input, mas.")
			time.Sleep(3 * time.Second)
		}
	}
}
func EditPemain(dataTurnamen *tabTurnamen, idx int) {
	/*
		I.S terdefinisi data turnamen dengan index idx
		F.S mengubah data pemain dan mengembalikan dataTurnamen
	*/
	// Kamus lokal
	var id string
	var index, pilihan int
	var keluar bool

	// Algoritma
	// Menu cari turnamen
	fmt.Print("Masukkan id Pemain yang ingin dicari: ")
	fmt.Scan(&id)
	fmt.Scanln()
	// Cari index turnamen berdasarkan ID menggunakan sequential search
	index = descBinarySearchByID(dataTurnamen[idx].pemain, dataTurnamen[idx].nPemain, id)
	if index != -1 {
		fmt.Println("\r")
		for keluar = false; !keluar; {
			fmt.Printf("Perubahan Data Pemain dengan ID %s\n", dataTurnamen[idx].pemain[index].id)
			fmt.Println("Pilih Data Pemain yang hendak diedit:")
			fmt.Printf("\n1. Nama Pemain: %s\n", dataTurnamen[idx].pemain[index].name)
			fmt.Printf("2. Jumlah Menang: %d\n", dataTurnamen[idx].pemain[index].menang)
			fmt.Printf("3. Jumlah Kalah: %d\n", dataTurnamen[idx].pemain[index].kalah)
			fmt.Println("4. Selesai")
			fmt.Print("Pilih data yang ingin diedit: ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				fmt.Print("Masukkan Nama Pemain baru: ")
				fmt.Scan(&dataTurnamen[idx].pemain[index].name)
				fmt.Scanln()
				fmt.Println("Nama Pemain berhasil diubah.")
			case 2:
				fmt.Print("Masukkan Jumlah Menang baru: ")
				fmt.Scan(&dataTurnamen[idx].pemain[index].menang)
				fmt.Scanln()
				fmt.Println("Jumlah Menang berhasil diubah.")
			case 3:
				fmt.Print("Masukkan Jumlah Kalah baru: ")
				fmt.Scan(&dataTurnamen[idx].pemain[index].kalah)
				fmt.Scanln()
				fmt.Println("Jumlah Kalah berhasil diubah.")
			case 4:
				fmt.Print("Keluar dari editor pemain")
				time.Sleep(3 * time.Second)
				keluar = true
			default:
				fmt.Println("Salah input, mas.")
			}
		}
	} else {
		fmt.Println("Pemain tidak ditemukan.")
		time.Sleep(3 * time.Second)
	}
}
func RankingPemain(dataTurnamen *tabTurnamen, idx int) {
	/*
		I.S terdefinisi Turnamen berindex idx
		F.S menampilkan ranking Player kepada user
	*/
	// Kamus lokal
	var i int
	var apapun string

	// Algoritma
	fmt.Println("Rangking")
	for i = 1; i <= dataTurnamen[idx].nPemain; i++ {
		fmt.Printf("   %d. ID: %s \n", i, dataTurnamen[idx].pemain[i-1].id)
		fmt.Println("      Name: ", dataTurnamen[idx].pemain[i-1].name)
		fmt.Println("      Skor: ", dataTurnamen[idx].pemain[i-1].skor)
	}
	fmt.Print("Kembali ke menu utama? (tekan apapun) ")
	fmt.Scan(&apapun)
}
func EditSkor(dataTurnamen *tabTurnamen, idx int) {
	/*
		I.S terdefinisi data turnamen dengan index idx
		F.S mengubah data skormenag dan skor kalah
	*/
	// Kamus lokal
	var pilihan int
	var keluar bool

	// Algoritma
	for keluar = false; !keluar; {
		fmt.Println("Pilih skor yang hendak diedit")
		fmt.Println("1. Menang =", dataTurnamen[idx].skorMenang)
		fmt.Println("2. Kalah =", dataTurnamen[idx].skorKalah)
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Print("Masukkan skor menang yang baru: ")
			fmt.Scan(&dataTurnamen[idx].skorMenang)
			fmt.Scanln()
			fmt.Println("Skor menang berhasil diubah.")
		case 2:
			fmt.Print("Masukkan skor kalah yang baru: ")
			fmt.Scan(&dataTurnamen[idx].skorKalah)
			fmt.Scanln()
			fmt.Println("Skor kalah berhasil diubah.")
		case 3:
			fmt.Println("Keluar dari editor skor")
			time.Sleep(3 * time.Second)
			keluar = true
		default:
			fmt.Println("Salah input, mas.")
		}
	}
}

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

func descBinarySearchByID(pemain tabPemain, n int, id string) int {
	/*
		I.S terdefinisi array pemain yang berisi n dan variabel id sebagai variabel yang dicari
		F.S mengembalikan index dari id di dalam data nama Player jika ditemukan, jika tidak ditemukan mengembalikan -1
	*/

	// Kamus Lokal
	var left, mid, right int
	var ketemu bool

	// Algoritma
	left = 0
	right = n - 1
	mid = (left + right) / 2
	for left <= right && pemain[mid].id != id {
		if id < pemain[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	// Memastikan nama turnamen tersebut ada
	ketemu = mid >= 0 && pemain[mid].id == id

	if ketemu {
		return mid
	} else {
		return -1
	}
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

func descSelectionSortBySkor(Pemain *tabPemain, n int) {
	/*
		I.S terdefinisi Pemain berjumlah n
		F.S mengurutkan array Pemain berdasarkan skor secara descending
	*/

	// Kamus Lokal
	var i, j, nilaimax int
	// Algoritma
	for i = 0; i < n; i++ {
		nilaimax = i
		for j = i + 1; j < n; j++ {
			if Pemain[j].skor > Pemain[nilaimax].skor {
				nilaimax = j
			}
		}
		if nilaimax != i {
			Pemain[i], Pemain[nilaimax] = Pemain[nilaimax], Pemain[i]
		}
	}
}

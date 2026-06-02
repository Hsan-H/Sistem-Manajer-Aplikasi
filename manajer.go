package main

import "fmt"

// Kamus Global
const NMAX = 100

// type untuk player
type player struct {
	id       string
	nama     string
	password string
	menang   int
	kalah    int
	skor     int
}
type tabPlayer [NMAX]player

// type untuk Turnamen
type turnamen struct {
	nama     string
	password string
	player   tabPlayer
	Pemenang string
	nPlayer  int
}
type tabTurnamen [NMAX]turnamen

// type untuk Menentukan Juara
type kemenangan [NMAX]string

func main() {
	// Kamus Lokal
	var keluar bool
	var pilihan string
	var Turnamen tabTurnamen
	var n int

	// Algoritma
	keluar = false
	for !keluar {
		// Text Interface Utama
		// \033[H  -> Memindahkan kursor ke pojok kiri atas (home)
		// \033[2J -> Menghapus seluruh isi layar
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Println("				Manajer Turnamen")
		fmt.Println("=====================================================================================")
		fmt.Println("   Menu:")
		fmt.Println("   1. Registrasi Turnamen")
		fmt.Println("   2. Login Turnamen")
		fmt.Println("   3. List Turnamen")
		fmt.Println("   4. Keluar")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Print("   Pilihan: ")
		fmt.Scanln(&pilihan)

		// Switch
		if pilihan == "1" {
			registrasiTurnamen(&Turnamen, &n)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "2" {
			Login(&Turnamen, n)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "3" {
			listTurnamen(Turnamen, n)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "4" {
			keluar = true
		} else {
			fmt.Println("Tidak ada pilihan yang anda input.")
		}
	}

}

func registrasiTurnamen(Turnamen *tabTurnamen, n *int) {
	/*
		I.S terdefinisi array Turnamen yang berisi n
		F.S pengisian data array Turnamen
	*/

	// Algoritma
	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================================================================")
	fmt.Println("				Manajer Turnamen")
	fmt.Println("=====================================================================================")
	fmt.Println("   Registrasi")

	// Input pada go, spasi menandakan pemisahan sehingga penggunaan dua kalimat tidak bisa diganti spasi
	fmt.Println("   Perhatikan, Ganti Spasi dengan _ !")

	// Pengisian Data Nama
	fmt.Print("   Nama Turnamen: ")
	fmt.Scanln(&Turnamen[*n].nama)

	// Pengisian Data Password
	fmt.Print("   Password: ")
	fmt.Scanln(&Turnamen[*n].password)

	// Menambah jumlah Turnamen
	*n++
}

func Login(Turnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array Turnamen yang berisi n
		F.S pengecekan data yang di input sebagai verifikasi sebelum memasuki procedure interfaceTurnamen dengan menggunakan binary seaarch yang diurut dengan insertion sort
	*/

	// Kamus Lokal
	var idx int
	var nama, password, pilihan string
	var ketemuNama, ketemuPass, nyerah bool

	// Algoritma
	// Text Interface Login
	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================================================================")
	fmt.Println("				Manajer Turnamen")
	fmt.Println("=====================================================================================")
	fmt.Print("   Login")

	// Pengisian Data Nama untuk dicocokan dengan Data Nama yang ada
	fmt.Print("   Nama Turnamen: ")
	fmt.Scanln(&nama)

	// Mengurutkan Data Turnamen Untuk Binary Search
	insertionSortTurnamen(Turnamen, n)

	// Binary Search untuk mencari dapat turnamen agar dapat dicocokkan dengan password
	binarySearch(Turnamen, &ketemuNama, &idx, n, nama)

	// Berfungsi untuk Autentifikasi Password dan Dapat Dicoba Berkali Kali
	if ketemuNama {
		for !ketemuPass && nyerah == false {
			// Pengisian Data Password untuk dicocokan dengan Data Nama
			fmt.Print("   Password: ")
			fmt.Scanln(&password)

			// Memastikan Password yang diinput benar
			if password == Turnamen[idx].password {
				ketemuPass = true
			} else {
				// Memberikan pilihan untuk menyerah atau mengulang password
				fmt.Print("   Ketik 1 untuk menyerah dan Ketik apapun untuk memasukkan ulang password: ")
				fmt.Scanln(&pilihan)

				// Logika agar perulangan berakhir bila manajer memilih menyerah
				if pilihan == "1" {
					nyerah = true
				}
			}
		}
	} else {
		fmt.Println("   Maaf. Turnamen tersebut belum didaftarkan. Silahkan registrasi.")
	}

	// Setelah Nama dan Password Benar dapat Memasuki ke Sistem Manajer Turnamen
	if ketemuNama == true && ketemuPass == true {
		interfaceTurnamen(Turnamen, idx)
	}
}

func insertionSortTurnamen(Turnamen *tabTurnamen, n int) {
	/*
		I.S terdefinisi array Turnamen yang berisi n
		F.S mengurutkan array turnamen berdasarkan nama yang bersifat string
	*/

	// Kamus Lokal
	var i, j int
	var x turnamen

	// Algoritma
	// https://www.geeksforgeeks.org/dsa/insertion-sort-algorithm/
	for i = 0; i < n; i++ {
		x = Turnamen[i]
		j = i - 1
		for j >= 0 && Turnamen[j].nama > x.nama {
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

func interfaceTurnamen(Turnamen *tabTurnamen, idxT int) {
	/*
		I.S terdefinisi array Turnamen yang berindex idxT
		F.S mengeluarkan output interface untuk registrasi, edit, ranking, dan juara player serta edit skor
	*/

	// Kamus Lokal
	var skorMenang, skorKalah int
	var keluar bool
	var apapun, pilihan string

	// Algoritma
	keluar = false
	for keluar == false {
		// Text Interface Turnamen
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
		fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
		fmt.Println("=====================================================================================")
		fmt.Println("   Juara dan Ranking ditentukan oleh skor.")
		fmt.Println("   JBerikut merupakan skor tiap menang dan kalah saat ini")
		fmt.Println("   Menang =", skorMenang)
		fmt.Println("   Kalah =", skorKalah)
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Println("   Menu")
		fmt.Println("   1. Tambah Data Pemain")
		fmt.Println("   2. Edit Data Pemain")
		fmt.Println("   3. Atur Skor")
		fmt.Println("   4. Rangking Pemain")
		fmt.Println("   5. Juara Turnamen")
		fmt.Println("   6. Keluar")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Print("   Pilihan: ")
		fmt.Scanln(&pilihan)

		// Switch
		if pilihan == "1" {
			interfaceRegistrasiPlayer(Turnamen, idxT, skorMenang, skorKalah)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "2" {
			interfacePengeditanPlayer(Turnamen, idxT, skorMenang, skorKalah)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "3" {
			interfaceSkoringPlayer(Turnamen, idxT, &skorMenang, &skorKalah)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "4" {
			interfaceRankingTurnamen(*Turnamen, idxT)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "5" {
			// Text Interface Ranking
			fmt.Print("\033[H\033[2J")
			fmt.Println("=====================================================================================")
			fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
			fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
			fmt.Println("=====================================================================================")
			fmt.Println("   Juara Turnamen.")
			fmt.Printf("   Selamat kepada %s atas kemenanganya. \n", Turnamen[idxT].Pemenang)
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Print("   Ketik apapun untuk Keluar: ")
			fmt.Scanln(&apapun)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "6" {
			keluar = true
		} else {
			fmt.Println("   Tidak ada pilihan yang anda input.")
		}
	}
}

func interfaceRegistrasiPlayer(Turnamen *tabTurnamen, idx, skorMenang, skorKalah int) {
	/*
		I.S terdefinisi array Turnamen yang berindex idxT dan variavel skorMenang dan skorKalah bersifat integer
		F.S pengisian data array player yang berda di array Turnamen
	*/

	// Kamus Lokal
	var keluar bool
	var apapun string

	// Algoritma
	keluar = false

	for !keluar {
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Printf("   Turnamen %s \n \n", Turnamen[idx].nama)
		fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
		fmt.Println("=====================================================================================")
		fmt.Println("   Registrasi")
		fmt.Println("   Perhatikan, Ganti Spasi dengan _ !")

		// Pengisian Data Nama
		fmt.Print("   Nama: ")
		fmt.Scanln(&Turnamen[idx].player[Turnamen[idx].nPlayer].nama)

		// Pengisian Data ID
		fmt.Print("   ID: ")
		fmt.Scanln(&Turnamen[idx].player[Turnamen[idx].nPlayer].id)

		// Pengisian Data Kemenangan
		fmt.Print("   Jumlah Kemenangan: ")
		fmt.Scanln(&Turnamen[idx].player[Turnamen[idx].nPlayer].menang)

		// Pengisian Data Kekalahan
		fmt.Print("   Jumlah Kekalahan: ")
		fmt.Scanln(&Turnamen[idx].player[Turnamen[idx].nPlayer].kalah)

		// Menghitung skor Player (nilai skor sebelum diedit adalah 0)
		Turnamen[idx].player[Turnamen[idx].nPlayer].skor = Turnamen[idx].player[Turnamen[idx].nPlayer].menang*skorMenang + Turnamen[idx].player[Turnamen[idx].nPlayer].kalah*skorKalah

		// Menambah jumlah Turnamen
		Turnamen[idx].nPlayer++

		// Menawarkan untuk Mengakhiri Mengisi Data Player atau Menambah Player
		fmt.Print("   Ketik 1 untuk menyudahi pengisan atau ketik apapun untuk menambah pemain: ")
		fmt.Scanln(&apapun)

		// Logika untuk Keluar
		if apapun == "1" {
			keluar = true
		}
	}

	// Menghitung Kejuaraan
	Turnamen[idx].Pemenang = Pemenang(Turnamen[idx].player, Turnamen[idx].nPlayer)
}

func Pemenang(Player tabPlayer, n int) string {
	/*
		I.S terdefinisi array Player yang berjumlah n
		F.S mengembalikan nama sang juara
	*/

	// Kamus Lokal
	var i, j, k, max int
	var menang kemenangan

	// Algoritma
	// Mencari Nilai Ekstrim Maksimum untuk Mencari Nilai yang Juara
	max = Player[0].skor
	j = 1
	for j < n {
		if max < Player[j].skor {
			max = Player[j].skor
		}
		j++
	}
	for i = 0; i < n; i++ {
		if max == Player[i].skor {
			menang[k] = Player[i].nama
			k++
		}
	}
	return rekursif(menang, k)
}

func rekursif(datapemenang kemenangan, k int) string {
	/*
		I.S terdefinisi array datapemenang yang berjumlah k
		F.S mengembalikan nama dalam sebuah string sehingga apabila lebih dari satu pemenang dapat dikembalikan
	*/

	// Algoritma
	if k == 0 {
		return ""
	} else if k == 1 {
		return datapemenang[0]
	} else {
		return rekursif(datapemenang, k-1) + ", " + datapemenang[k-1]
	}
}

func interfacePengeditanPlayer(Turnamen *tabTurnamen, idx, skorMenang, skorKalah int) {
	/*
		I.S terdefinisi array yang berada di index idx dan terdefinisi skorMenang dan skorKalah
		F.S menampilkan interface untuk mengedit data array Player yang berada di array Turnamen
	*/

	// Kamus Lokal
	var keluar bool
	var idxedit int
	var nama, pilihan string

	// Algoritma
	keluar = false

	for !keluar {
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Printf("   Turnamen %s \n \n", Turnamen[idx].nama)
		fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
		fmt.Println("=====================================================================================")
		fmt.Println("   Menu Edit Pemain")
		fmt.Println("   1. Edit Pemain")
		fmt.Println("   2. Keluar")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Print("   Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan == "1" {
			// Pengisian Data Nama yang Hendak diedit
			fmt.Print("   Masukkan Nama Pemain: ")
			fmt.Scanln(&nama)

			// Sequential Search untuk mencari index dari nama pemain
			idxedit = sequentialSearch(*Turnamen, nama, idx, Turnamen[idx].nPlayer)

			// Logika untuk back-up bila nama yang di input tidak ada
			if idxedit == -1 {
				fmt.Println("   Silahkan lakukan ragistrasi player / lakukan input nama lagi.")
			} else {
				interfacePengeditanDataPlayer(Turnamen, idx, idxedit, Turnamen[idx].nPlayer, skorMenang, skorKalah)
			}

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "2" {
			keluar = true
		} else {
			fmt.Println("   Tidak ada pilihan yang anda input.")
		}
	}
}

func sequentialSearch(Turnamen tabTurnamen, nama string, idx, n int) int {
	/*
		I.S terdefinisi array Turnamen yang berada di index idx yang array Playernya berjumlah n dan variabel nama sebagai variabel yang dicari
		F.S mengembalikan index dari nama di dalam data nama Player
	*/

	// Kamus Lokal
	var i, idxedit int
	var ketemu bool

	// Algoritma
	idxedit = -1
	ketemu = false
	i = 0
	for i < n && !ketemu {
		if nama == Turnamen[idx].player[i].nama {
			idxedit = i
			ketemu = true
		}
		i++
	}
	return idxedit
}

func interfacePengeditanDataPlayer(Turnamen *tabTurnamen, idx, idxedit, n, skorMenang, skorKalah int) {
	/*
		I.S terdefinisi array Turnamen yang berada di index idx yang array Playernya yang berada di index idxedit dan berjumlah n serta didefinisikan skorMenang dan skorKalah
		F.S menampilkan interface agar user dapat mengedit data player
	*/

	// Kamus lokal
	var keluar bool
	var pilihan string

	// Algoritma
	keluar = false

	for !keluar {
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Printf("   Turnamen %s \n \n", Turnamen[idx].nama)
		fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
		fmt.Println("=====================================================================================")
		fmt.Println("   Silahkan Pilih Bagian yang hendak diedit")
		fmt.Println("   1. Nama: ", Turnamen[idx].player[idxedit].nama)
		fmt.Println("   2. ID :", Turnamen[idx].player[idxedit].id)
		fmt.Println("   3. Jumlah Kemenangan :", Turnamen[idx].player[idxedit].menang)
		fmt.Println("   4. Jumlah Kekalahan :", Turnamen[idx].player[idxedit].kalah)
		fmt.Println("   5. Selesai")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Print("   Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan == "1" {
			fmt.Println("   Perhatikan, Ganti Spasi dengan _ !")

			// Pengeditan nama player
			fmt.Print("   Masukkan Nama Pemain: ")
			fmt.Scanln(&Turnamen[idx].player[idxedit].nama)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "2" {
			// Pengeditan id player
			fmt.Print("   Masukkan ID: ")
			fmt.Scanln(&Turnamen[idx].player[idxedit].id)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "3" {
			// Pengeditan Jumlah Kemenangan player
			fmt.Print("   Masukkan Jumlah Kemenangan: ")
			fmt.Scanln(&Turnamen[idx].player[idxedit].menang)

			// Menghitung skor Player
			Turnamen[idx].player[idxedit].skor = Turnamen[idx].player[idxedit].menang*skorMenang + Turnamen[idx].player[idxedit].kalah*skorKalah

			// Menghitung Kejuaraan
			Turnamen[idx].Pemenang = Pemenang(Turnamen[idx].player, n)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "4" {
			// Menghitung Jumlah Kekalahan
			fmt.Print("   Masukkan Jumlah Kekalahan: ")
			fmt.Scanln(&Turnamen[idx].player[idxedit].kalah)

			// Menghitung skor Player
			Turnamen[idx].player[idxedit].skor = Turnamen[idx].player[idxedit].menang*skorMenang + Turnamen[idx].player[idxedit].kalah*skorKalah

			// Menghitung Kejuaraan
			Turnamen[idx].Pemenang = Pemenang(Turnamen[idx].player, n)

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "5" {
			keluar = true
		} else {
			fmt.Println("   Tidak ada pilihan yang anda input.")
		}
	}
}

func interfaceSkoringPlayer(Turnamen *tabTurnamen, idx int, skorMenang, skorKalah *int) {
	/*
		I.S terdefinisi array Turnamen yang berada di index idx yang array Playernya yang berjumlah n serta didefinisikan skorMenang dan skorKalah
		F.S memberikan interface agar user dapat mengedit data skorMenang dan skorKalah
	*/

	// Kamus Lokal
	var pilihan string
	var keluar bool
	var i int

	// Algoritma
	keluar = false
	for !keluar {
		// Text Interface Turnamen
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Printf("   Turnamen %s \n \n", Turnamen[idx].nama)
		fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
		fmt.Println("=====================================================================================")
		fmt.Println("   Pilih skor yang hendak diedit")
		fmt.Println("   1. Menang =", *skorMenang)
		fmt.Println("   2. Kalah =", *skorKalah)
		fmt.Println("   3. Keluar")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Print("   Pilihan: ")
		fmt.Scanln(&pilihan)

		// Switch
		if pilihan == "1" {
			// Pengeditan Skor Menang
			fmt.Print("   Skor Menang: ")
			fmt.Scanln(skorMenang)

			// Perulangan data untuk Menghitung Seluruh Skor Player
			for i = 0; i < Turnamen[idx].nPlayer; i++ {
				Turnamen[idx].player[i].skor = Turnamen[idx].player[i].menang**skorMenang + Turnamen[idx].player[i].kalah**skorKalah
			}

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "2" {
			// Pengeditan Skor Kalah
			fmt.Print("   Skor Kalah: ")
			fmt.Scanln(skorKalah)

			// Perulangan data untuk Menghitung Seluruh Skor Player
			for i = 0; i < Turnamen[idx].nPlayer; i++ {
				Turnamen[idx].player[i].skor = Turnamen[idx].player[i].menang**skorMenang + Turnamen[idx].player[i].kalah**skorKalah
			}

			// Reset pilihan menjadi nol
			pilihan = "0"
		} else if pilihan == "3" {
			keluar = true
		} else {
			fmt.Println("   Tidak ada pilihan yang anda input.")
		}
	}
}

func interfaceRankingTurnamen(Turnamen tabTurnamen, idx int) {
	/*
		I.S terdefinisi Turnamen berindex idx
		F.S menampilkan ranking Player kepada user
	*/

	// Kamus Lokal
	var apapun string
	var i int

	// Algoritma
	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================================================================")
	fmt.Printf("   Turnamen %s \n \n", Turnamen[idx].nama)
	fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
	fmt.Println("=====================================================================================")
	fmt.Println("   Ranking Pemain Berdasarkan Skor")

	// Selection Sort
	selectionSort(&Turnamen[idx].player, Turnamen[idx].nPlayer)

	// Perulangan Output untuk Menampilkan Ranking
	for i = 1; i <= Turnamen[idx].nPlayer; i++ {
		fmt.Printf("   %d. Nama: %s \n", i, Turnamen[idx].player[i-1].nama)
		fmt.Println("      ID: ", Turnamen[idx].player[i-1].id)
		fmt.Println("      Skor: ", Turnamen[idx].player[i-1].skor)
	}

	// Untuk Balik ke Interface Turnamen
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Print("   Ketik apapun untuk Keluar: ")
	fmt.Scanln(&apapun)
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

func listTurnamen(Turnamen tabTurnamen, n int) {
	/*
		I.S terdefinisi array Turnamen yang berisi n bilangan bulat
		F.S menampilkan data array Turnamen berupa nama turnamen dan pemenangnya
	*/

	// Kamus Lokal
	var apapun string
	var i int

	// Algoritma
	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================================================================")
	fmt.Println("				Manajer Turnamen")
	fmt.Println("=====================================================================================")
	fmt.Println("   List Turnamen")

	// Mengeluarkan Daftar List
	for i = 1; i <= n; i++ {
		fmt.Printf("   %d. Turnamen %s \n", i, Turnamen[i-1].nama)
		fmt.Println("       Pemenang: ", Turnamen[i-1].Pemenang)
	}

	// Untuk Balik ke Interface Utama
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Print("   Ketik apapun untuk Keluar: ")
	fmt.Scanln(&apapun)
}

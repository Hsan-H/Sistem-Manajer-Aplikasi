package main

import "fmt"

// Kamus Global
const NMAX = 1000

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
	Masuk    int
	nama     string
	password string
	player   tabPlayer
	Pemenang string
}
type tabTurnamen [NMAX]turnamen

// type untuk Menentukan Juara
type kemenangan [NMAX]string

func main() {
	// Kamus Lokal
	var keluar bool = false
	var pilihan, n int
	var Turnamen tabTurnamen
	var apapun string

	// Algoritma
	for keluar == false {
		// Text Interface Utama
		// \033[H  -> Memindahkan kursor ke pojok kiri atas (home)
		// \033[2J -> Menghapus seluruh isi layar
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Println("				Manajer Turnamen")
		fmt.Println("=====================================================================================")
		fmt.Println("   Menu")
		fmt.Println("   1. Registrasi Turnamen")
		fmt.Println("   2. Login Turnamen")
		fmt.Println("   3. List Turnamen")
		fmt.Println("   4. Keluar")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Print("   Pilihan: ")
		fmt.Scanln(&pilihan)

		// Switch
		if pilihan == 1 {
			// Text Interface Registrasi
			fmt.Print("\033[H\033[2J")
			fmt.Println("=====================================================================================")
			fmt.Println("				Manajer Turnamen")
			fmt.Println("=====================================================================================")
			fmt.Println("   Registrasi")
			fmt.Println("   Perhatikan, Ganti Spasi dengan _ !")
			fmt.Print("   Nama Turnamen: ")
			fmt.Scanln(&Turnamen[n].nama)
			fmt.Print("   Password: ")
			fmt.Scanln(&Turnamen[n].password)

			// Menambah jumlah Turnamen
			n++
			Turnamen[n].Masuk = n
		} else if pilihan == 2 {
			Login(&Turnamen, n)
		} else if pilihan == 3 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("=====================================================================================")
			fmt.Println("				Manajer Turnamen")
			fmt.Println("=====================================================================================")
			fmt.Println("   List Turnamen")
			list(Turnamen, n)
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Print("   Ketik apapun untuk Keluar: ")
			fmt.Scanln(&apapun)
		} else if pilihan == 4 {
			keluar = true
		} else {
			fmt.Println("Tidak ada pilihan yang anda input.")
		}
	}

}

func Login(Turnamen *tabTurnamen, n int) {
	// Kamus Lokal
	var left, right, mid, pilihan int
	var nama, password string
	var ketemuNama, ketemuPass, nyerah bool

	// Algoritma
	// Text Interface Login
	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================================================================")
	fmt.Println("				Manajer Turnamen")
	fmt.Println("=====================================================================================")
	fmt.Print("   Nama Turnamen: ")
	fmt.Scanln(&nama)

	// Binary Search untuk mencari dapat turnamen agar dapat dicocokkan dengan password
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
	ketemuNama = mid >= 0 && Turnamen[mid].nama == nama

	if ketemuNama {
		for !ketemuPass && nyerah == false {
			fmt.Print("   Password: ")
			fmt.Scanln(&password)
			// Memastikan Password yang diinput benar
			if password == Turnamen[mid].password {
				ketemuPass = true
			} else {
				// Memberikan pilihan untuk menyerah atau mengulang password
				fmt.Print("   Ketik 1 untuk menyerah dan Ketik apapun untuk memasukkan ulang password: ")
				fmt.Scanln(&pilihan)
				if pilihan == 1 {
					nyerah = true
				}
			}
		}
	} else {
		fmt.Println("   Maaf. Turnamen tersebut belum didaftarkan. Silahkan registrasi.")
	}
	if ketemuNama == true && ketemuPass == true {
		interfaceTurnamen(Turnamen, mid)
	}
}

func interfaceTurnamen(Turnamen *tabTurnamen, idxT int) {
	// Kamus Lokal
	var pilihan, n, skorMenang, skorKalah, pilihedit, pilihfield, pilihskor, idxedit int
	var keluar, keluaredit, keluarfield, ketemuedit, keluarskor bool
	var nama, apapun string
	keluar = false
	keluaredit = false
	keluarfield = false
	keluarskor = false

	// Algoritma
	for keluar == false {
		// Text Interface Turnamen
		fmt.Print("\033[H\033[2J")
		fmt.Println("=====================================================================================")
		fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
		fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
		fmt.Println("=====================================================================================")
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
		if pilihan == 1 {
			// Text Interface Registrasi
			fmt.Print("\033[H\033[2J")
			fmt.Println("=====================================================================================")
			fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
			fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
			fmt.Println("=====================================================================================")
			fmt.Println("   Registrasi")
			fmt.Println("   Perhatikan, Ganti Spasi dengan _ !")
			fmt.Print("   Nama: ")
			fmt.Scanln(&Turnamen[idxT].player[n].nama)
			fmt.Print("   ID: ")
			fmt.Scanln(&Turnamen[idxT].player[n].id)
			fmt.Print("   Jumlah Kemenangan: ")
			fmt.Scanln(&Turnamen[idxT].player[n].menang)
			fmt.Print("   Jumlah Kekalahan: ")
			fmt.Scanln(&Turnamen[idxT].player[n].kalah)

			// Menghitung skor Player
			Turnamen[idxT].player[n].skor = Turnamen[idxT].player[n].menang*skorMenang + Turnamen[idxT].player[n].kalah*skorKalah

			// Menambah jumlah Turnamen
			n++

			// Menghitung Kejuaraan
			Turnamen[idxT].Pemenang = Pemenang(Turnamen[idxT].player, n)
		} else if pilihan == 2 {
			// Text Interface Edit Data
			for !keluaredit {
				fmt.Print("\033[H\033[2J")
				fmt.Println("=====================================================================================")
				fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
				fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
				fmt.Println("=====================================================================================")
				fmt.Println("   Menu Edit Pemain")
				fmt.Println("   1. Edit Pemain")
				fmt.Println("   2. Keluar")
				fmt.Println("-------------------------------------------------------------------------------------")
				fmt.Print("   Pilihan: ")
				fmt.Scanln(&pilihedit)
				if pilihedit == 1 {
					fmt.Print("   Masukkan Nama Pemain: ")
					fmt.Scanln(&nama)

					//Sequential Search untuk mencari nama pemain yang hendak di edit
					ketemuedit = false
					i := 0
					for i < n && !ketemuedit {
						if nama == Turnamen[idxT].player[i].nama {
							idxedit = i
							ketemuedit = true
						}
						i++
					}

					for !keluarfield {
						// Text Interface Bagian Edit
						fmt.Print("\033[H\033[2J")
						fmt.Println("=====================================================================================")
						fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
						fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
						fmt.Println("=====================================================================================")
						fmt.Println("   Silahkan Pilih Bagian yang hendak diedit")
						fmt.Println("   1. Nama: ", Turnamen[idxT].player[idxedit].nama)
						fmt.Println("   2. ID :", Turnamen[idxT].player[idxedit].id)
						fmt.Println("   3. Jumlah Kemenangan :", Turnamen[idxT].player[idxedit].menang)
						fmt.Println("   4. Jumlah Kekalahan :", Turnamen[idxT].player[idxedit].kalah)
						fmt.Println("   5. Selesai")
						fmt.Println("-------------------------------------------------------------------------------------")
						fmt.Print("   Pilihan: ")
						fmt.Scanln(&pilihfield)
						if pilihfield == 1 {
							fmt.Println("   Perhatikan, Ganti Spasi dengan _ !")
							fmt.Print("   Masukkan Nama Pemain: ")
							fmt.Scanln(&Turnamen[idxT].player[idxedit].nama)
						} else if pilihfield == 2 {
							fmt.Print("   Masukkan ID: ")
							fmt.Scanln(&Turnamen[idxT].player[idxedit].id)
						} else if pilihfield == 3 {
							fmt.Print("   Masukkan Jumlah Kemenangan: ")
							fmt.Scanln(&Turnamen[idxT].player[idxedit].menang)

							// Menghitung skor Player
							Turnamen[idxT].player[idxedit].skor = Turnamen[idxT].player[idxedit].menang*skorMenang + Turnamen[idxT].player[idxedit].kalah*skorKalah

							// Menghitung Kejuaraan
							Turnamen[idxT].Pemenang = Pemenang(Turnamen[idxT].player, n)
						} else if pilihfield == 4 {
							fmt.Print("   Masukkan Jumlah Kekalahan: ")
							fmt.Scanln(&Turnamen[idxT].player[idxedit].kalah)

							// Menghitung skor Player
							Turnamen[idxT].player[idxedit].skor = Turnamen[idxT].player[idxedit].menang*skorMenang + Turnamen[idxT].player[idxedit].kalah*skorKalah

							// Menghitung Kejuaraan
							Turnamen[idxT].Pemenang = Pemenang(Turnamen[idxT].player, n)
						} else if pilihfield == 5 {
							keluarfield = true
						} else {
							fmt.Println("   Tidak ada pilihan yang anda input.")
						}
					}
				} else if pilihedit == 2 {
					keluaredit = true
				} else {
					fmt.Println("   Tidak ada pilihan yang anda input.")
				}
			}
		} else if pilihan == 3 {
			for !keluarskor {
				// Text Interface Turnamen
				fmt.Print("\033[H\033[2J")
				fmt.Println("=====================================================================================")
				fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
				fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
				fmt.Println("=====================================================================================")
				fmt.Println("   Pilih skor yang hendak diedit")
				fmt.Println("   1. Menang =", skorMenang)
				fmt.Println("   2. Kalah =", skorKalah)
				fmt.Println("   3. Keluar")
				fmt.Println("-------------------------------------------------------------------------------------")
				fmt.Print("   Pilihan: ")
				fmt.Scanln(&pilihskor)

				// Switch
				if pilihskor == 1 {
					fmt.Print("   Skor Menang: ")
					fmt.Scanln(&skorMenang)
					// Perulangan data yang telah diurut
					for i := 0; i < n; i++ {
						Turnamen[idxT].player[i].skor = Turnamen[idxT].player[i].menang*skorMenang + Turnamen[idxT].player[i].kalah*skorKalah
					}
				} else if pilihskor == 2 {
					fmt.Print("   Skor Kalah: ")
					fmt.Scanln(&skorKalah)
					for i := 0; i < n; i++ {
						Turnamen[idxT].player[i].skor = Turnamen[idxT].player[i].menang*skorMenang + Turnamen[idxT].player[i].kalah*skorKalah
					}
				} else if pilihskor == 3 {
					keluarskor = true
				} else {
					fmt.Println("   Tidak ada pilihan yang anda input.")
				}
			}
		} else if pilihan == 4 {
			// Text Interface Ranking
			fmt.Print("\033[H\033[2J")
			fmt.Println("=====================================================================================")
			fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
			fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
			fmt.Println("=====================================================================================")
			fmt.Println("   Ranking Pemain Berdasarkan Skor")

			rank(Turnamen[idxT].player, n)

			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Print("   Ketik apapun untuk Keluar: ")
			fmt.Scanln(&apapun)
		} else if pilihan == 5 {
			// Text Interface Ranking
			fmt.Print("\033[H\033[2J")
			fmt.Println("=====================================================================================")
			fmt.Printf("   Turnamen %s \n \n", Turnamen[idxT].nama)
			fmt.Println("   Aplikasi Manajer Turnamen dari Hasan dan Rahmat")
			fmt.Println("=====================================================================================")
			fmt.Println("   Juara Turnamen.")
			fmt.Printf("   Selamat kepada %s atas kemenanganya. \n", Pemenang(Turnamen[idxT].player, n))
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Print("   Ketik apapun untuk Keluar: ")
			fmt.Scanln(&apapun)
		} else if pilihan == 6 {
			keluar = true
		} else {
			fmt.Println("   Tidak ada pilihan yang anda input.")
		}
	}
}

func rank(Player tabPlayer, n int) {
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
	for i := 1; i <= n; i++ {
		fmt.Printf("   %d. Nama: %s \n", i, Player[i-1].nama)
		fmt.Println("      ID: ", Player[i-1].id)
		fmt.Println("      Skor: ", Player[i-1].skor)
	}
}

func Pemenang(Player tabPlayer, n int) string {
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
	// Algoritma
	if k == 0 {
		return ""
	} else if k == 1 {
		return datapemenang[0]
	} else {
		return rekursif(datapemenang, k-1) + ", " + datapemenang[k-1]
	}
}

func list(Turnamen tabTurnamen, n int) {
	// Kamus Lokal
	var i, j int
	var x turnamen

	// Insetion Sort untuk Mengurutkan List Turnamen agar Turnamen Terbaru Berada Bagian Atas
	// https://www.geeksforgeeks.org/dsa/insertion-sort-algorithm/
	for i := 0; i < n; i++ {
		x = Turnamen[i]
		j = i - 1
		for j >= 0 && Turnamen[j].Masuk < x.Masuk {
			Turnamen[j+1], Turnamen[j] = Turnamen[j], Turnamen[j+1]
			j = j - 1
		}
		Turnamen[j+1] = x
	}

	// Mengeluarkan daftar list yang telah diurut
	for i = 1; i <= n; i++ {
		fmt.Printf("   %d. Turnamen %s \n", i, Turnamen[i-1].nama)
		fmt.Println("       Pemenang: ", Turnamen[i-1].Pemenang)
	}
}

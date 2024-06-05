package main

import (
	"fmt"
)

const NMAX = 31

type Urutan struct {
	tanggal int
	noUrut  int
	noKamar int
}

var urutanList [NMAX][9]int
var namaList [NMAX][9]string

func penambahanUrutan(tanggal int) int {
	for i := 0; i < 9; i++ {
		if urutanList[tanggal-1][i] == 0 {
			urutanList[tanggal-1][i] = i + 1
			return i + 1
		}
	}
	return 0
}

func DapatkanNoKamar(tanggal int, noUrut int) string {
	noKamar := fmt.Sprintf("%d", tanggal)
	if noUrut < 10 {
		if tanggal < 10 {
			noKamar += "0"
		}
	}
	noKamar += fmt.Sprintf("%d", noUrut)
	return noKamar
}

func DapatkanWaktuCheckIn(noUrut int) string {
	if noUrut >= 1 && noUrut <= 6 {
		return fmt.Sprintf("%d Siang", noUrut)
	} else if noUrut >= 7 && noUrut <= 9 {
		return fmt.Sprintf("%d Pagi", noUrut)
	}
	return ""
}

func daftarAsrama() {
	var nama string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	var tanggal int
	fmt.Print("Input tanggal (1-31, 0 untuk berhenti): ")
	fmt.Scanln(&tanggal)
	if tanggal == 0 {
		return
	}
	if tanggal < 1 || tanggal > NMAX {
		fmt.Println("Tanggal gak bener. Masukin angka antara 1 dan", NMAX)
		return
	}

	noUrut := penambahanUrutan(tanggal)
	if noUrut != 0 {
		namaList[tanggal-1][noUrut-1] = nama
		fmt.Printf("Nama: %s, Tanggal: %d, Nomor Urut: %d\n", nama, tanggal, noUrut)
	} else {
		fmt.Println("Full.")
	}
}

func perolehNomorKamar() {
	var tanggal int
	var noUrut int
	fmt.Print("Masukkan tanggal (1-31): ")
	fmt.Scanln(&tanggal)
	if tanggal < 1 || tanggal > NMAX {
		fmt.Println("Tanggal gak bisa. Masukkin angka antara 1 dan", NMAX)
		return
	}

	fmt.Print("Masukkan nomor urut (1-9): ")
	fmt.Scanln(&noUrut)
	if noUrut < 1 || noUrut > 9 {
		fmt.Println("Nomor urut salah. input angka antara 1 dan 9.")
		return
	}

	if urutanList[tanggal-1][noUrut-1] == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	nama := namaList[tanggal-1][noUrut-1]
	nomorKamar := DapatkanNoKamar(tanggal, noUrut)
	waktuCheckIn := DapatkanWaktuCheckIn(noUrut)
	fmt.Printf("Nama: %s, Tanggal: %d, Nomor Urut: %d, Nomor Kamar: %s, Jam: %s\n", nama, tanggal, noUrut, nomorKamar, waktuCheckIn)
}

func editData() {
	var tanggal int
	var noUrut int
	var namaBaru string
	fmt.Print("Masukkan tanggal (1-31): ")
	fmt.Scanln(&tanggal)
	if tanggal < 1 || tanggal > NMAX {
		fmt.Println("Tanggal gak bener. Masukin angka antara 1 dan", NMAX)
		return
	}

	fmt.Print("Masukkan nomor urut (1-9): ")
	fmt.Scanln(&noUrut)
	if noUrut < 1 || noUrut > 9 {
		fmt.Println("Nomor urut salah. input angka antara 1 dan 9.")
		return
	}

	if urutanList[tanggal-1][noUrut-1] == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	nama := namaList[tanggal-1][noUrut-1]
	nomorKamar := DapatkanNoKamar(tanggal, noUrut)
	waktuCheckIn := DapatkanWaktuCheckIn(noUrut)
	fmt.Printf("Data saat ini - Nama: %s, Tanggal: %d, Nomor Urut: %d, Nomor Kamar: %s, Jam: %s\n", nama, tanggal, noUrut, nomorKamar, waktuCheckIn)

	fmt.Print("Masukkan nama baru: ")
	fmt.Scanln(&namaBaru)

	namaList[tanggal-1][noUrut-1] = namaBaru
	fmt.Printf("Data berhasil diubah: Tanggal %d, Nomor Urut %d, Nama Baru: %s\n", tanggal, noUrut, namaBaru)
}

func selectionSort(arr *[NMAX * 9]Urutan, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].tanggal < arr[minIdx].tanggal {
				minIdx = j
			} else if arr[j].tanggal == arr[minIdx].tanggal && arr[j].noUrut < arr[minIdx].noUrut {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func jadwalCheckIn() {
	var urutans [NMAX * 9]Urutan
	index := 0

	for t := 0; t < NMAX; t++ {
		for u := 0; u < 9; u++ {
			if urutanList[t][u] != 0 {
				urutans[index] = Urutan{
					tanggal: t + 1,
					noUrut:  urutanList[t][u],
					noKamar: u + 1,
				}
				index++
			}
		}
	}

	selectionSort(&urutans, index)

	fmt.Println("Jadwal Check-in Mahasiswa:")
	fmt.Println("Check-in bulan depan. Check-in Sesuai Tanggal Pendaftaran.")
	for i := 0; i < index; i++ {
		nama := namaList[urutans[i].tanggal-1][urutans[i].noUrut-1]
		noKamar := DapatkanNoKamar(urutans[i].tanggal, urutans[i].noKamar)
		waktuCheckIn := DapatkanWaktuCheckIn(urutans[i].noUrut)
		fmt.Printf("Nama: %s, Tanggal: %d, Nomor Urut: %d, Nomor Kamar: %s, Jam: %s\n", nama, urutans[i].tanggal, urutans[i].noUrut, noKamar, waktuCheckIn)
	}
}

func binarySearch(arr [NMAX * 9]Urutan, tanggal int, n int) int {
	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid].tanggal == tanggal {
			for mid > 0 && arr[mid-1].tanggal == tanggal {
				mid--
			}
			return mid
		}
		if arr[mid].tanggal < tanggal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func cariCheckInBerdasarkanTanggal() {
	var urutans [NMAX * 9]Urutan
	index := 0

	for t := 0; t < NMAX; t++ {
		for u := 0; u < 9; u++ {
			if urutanList[t][u] != 0 {
				urutans[index] = Urutan{
					tanggal: t + 1,
					noUrut:  urutanList[t][u],
					noKamar: u + 1,
				}
				index++
			}
		}
	}

	selectionSort(&urutans, index)

	var tanggal int
	fmt.Print("Masukkan tanggal yang dicari (1-31): ")
	fmt.Scanln(&tanggal)

	startIndex := binarySearch(urutans, tanggal, index)
	if startIndex == -1 {
		fmt.Println("Tidak ada data yang ditemukan untuk tanggal tersebut.")
		return
	}

	fmt.Println("Data ditemukan:")
	for i := startIndex; i < index && urutans[i].tanggal == tanggal; i++ {
		nama := namaList[urutans[i].tanggal-1][urutans[i].noUrut-1]
		noKamar := DapatkanNoKamar(urutans[i].tanggal, urutans[i].noKamar)
		waktuCheckIn := DapatkanWaktuCheckIn(urutans[i].noUrut)
		fmt.Printf("Nama: %s, Tanggal: %d, Nomor Urut: %d, Nomor Kamar: %s, Jam: %s\n", nama, urutans[i].tanggal, urutans[i].noUrut, noKamar, waktuCheckIn)
	}
}

func sequentialSearchByName(nama string) [NMAX * 9]Urutan {
	var hasil [NMAX * 9]Urutan
	index := 0

	for t := 0; t < NMAX; t++ {
		for u := 0; u < 9; u++ {
			if namaList[t][u] == nama {
				hasil[index] = Urutan{
					tanggal: t + 1,
					noUrut:  urutanList[t][u],
					noKamar: u + 1,
				}
				index++
			}
		}
	}
	return hasil
}

func cariCheckInBerdasarkanNama() {
	var nama string
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scanln(&nama)

	hasil := sequentialSearchByName(nama)
	index := 0
	for hasil[index].tanggal != 0 && index < NMAX*9 {
		index++
	}
	if index > 0 {
		fmt.Println("Data ditemukan:")
		for i := 0; i < index; i++ {
			noKamar := DapatkanNoKamar(hasil[i].tanggal, hasil[i].noKamar)
			waktuCheckIn := DapatkanWaktuCheckIn(hasil[i].noUrut)
			fmt.Printf("Nama: %s, Tanggal: %d, Nomor Urut: %d, Nomor Kamar: %s, Jam: %s\n", nama, hasil[i].tanggal, hasil[i].noUrut, noKamar, waktuCheckIn)
		}
	} else {
		fmt.Println("Tidak ada data yang ditemukan untuk nama tersebut.")
	}
}

func hapusData() {
	var tanggal int
	var noUrut int
	fmt.Print("Masukkan tanggal (1-31): ")
	fmt.Scanln(&tanggal)
	if tanggal < 1 || tanggal > NMAX {
		fmt.Println("Tanggal gak bener. Masukin angka antara 1 dan", NMAX)
		return
	}

	fmt.Print("Masukkan nomor urut (1-9): ")
	fmt.Scanln(&noUrut)
	if noUrut < 1 || noUrut > 9 {
		fmt.Println("Nomor urut salah. input angka antara 1 dan 9.")
		return
	}

	if urutanList[tanggal-1][noUrut-1] == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	namaList[tanggal-1][noUrut-1] = ""
	urutanList[tanggal-1][noUrut-1] = 0
	fmt.Println("Data berhasil dihapus.")
}

func main() {
	var running bool = true
	for running {
		fmt.Println("----------------------------------------")
		fmt.Println("                  MENU                 ")
		fmt.Println("----------------------------------------")
		fmt.Println("A. Daftar Asrama")
		fmt.Println("B. Peroleh Nomor Kamar")
		fmt.Println("C. Edit Data Peserta")
		fmt.Println("D. Jadwal Check-in Mahasiswa")
		fmt.Println("E. Cari Waktu Check-in Berdasarkan Tanggal")
		fmt.Println("F. Cari Waktu Check-in Peserta Berdasarkan Nama")
		fmt.Println("G. Hapus Data")
		fmt.Println("----------------------------------------")

		fmt.Print("Pilih menu (A-G atau X untuk exit): ")
		var pilihan string
		fmt.Scanln(&pilihan)

		if pilihan == "A" || pilihan == "a" {
			daftarAsrama()
		} else if pilihan == "B" || pilihan == "b" {
			perolehNomorKamar()
		} else if pilihan == "C" || pilihan == "c" {
			editData()
		} else if pilihan == "D" || pilihan == "d" {
			jadwalCheckIn()
		} else if pilihan == "E" || pilihan == "e" {
			cariCheckInBerdasarkanTanggal()
		} else if pilihan == "F" || pilihan == "f" {
			cariCheckInBerdasarkanNama()
		} else if pilihan == "G" || pilihan == "g" {
			hapusData()
		} else if pilihan == "X" || pilihan == "x" {
			running = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih antara A-G atau X untuk exit.")
		}
	}
}

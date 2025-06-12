package main

import (
	"fmt"
)

type Tanggal struct {
	hari, bulan, tahun int
}

type Penerbit struct {
	nama, alamat, telepon string
}

type Buku struct {
	kodeBuku     string
	judul        string
	pengarang    string
	tahunTerbit  int
	tanggalMasuk Tanggal
	penerbit     Penerbit
}

const maxBuku = 100

var daftarBuku1 [maxBuku]Buku
var daftarBuku2 [maxBuku]Buku
var daftarBuku3 [maxBuku]Buku
var jumlahBuku1, jumlahBuku2, jumlahBuku3 int

func tambahBuku(data *[maxBuku]Buku, jumlah *int) {
	if *jumlah < maxBuku {
		var b Buku
		fmt.Print("Masukkan kode buku: ")
		fmt.Scan(&b.kodeBuku)
		fmt.Print("Masukkan judul: ")
		fmt.Scan(&b.judul)
		fmt.Print("Masukkan pengarang: ")
		fmt.Scan(&b.pengarang)
		fmt.Print("Masukkan tahun terbit: ")
		fmt.Scan(&b.tahunTerbit)
		fmt.Print("Masukkan tanggal masuk (dd mm yyyy): ")
		fmt.Scan(&b.tanggalMasuk.hari, &b.tanggalMasuk.bulan, &b.tanggalMasuk.tahun)

		fmt.Print("Masukkan nama penerbit: ")
		fmt.Scan(&b.penerbit.nama)
		fmt.Print("Masukkan alamat penerbit: ")
		fmt.Scan(&b.penerbit.alamat)
		fmt.Print("Masukkan telepon penerbit: ")
		fmt.Scan(&b.penerbit.telepon)

		data[*jumlah] = b
		*jumlah++
		fmt.Println("Buku berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

func tampilkanBuku(data [maxBuku]Buku, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data buku.")
		return
	}
	for i := 0; i < jumlah; i++ {
		fmt.Printf("\n%d.\n", i+1)
		fmt.Println("Kode Buku:", data[i].kodeBuku)
		fmt.Println("Judul Buku:", data[i].judul)
		fmt.Println("Pengarang:", data[i].pengarang)
		fmt.Println("Tahun Terbit:", data[i].tahunTerbit)
		fmt.Printf("Tanggal Masuk: %02d-%02d-%d\n", data[i].tanggalMasuk.hari, data[i].tanggalMasuk.bulan, data[i].tanggalMasuk.tahun)
		fmt.Println("Penerbit:", data[i].penerbit.nama)
		fmt.Println("Alamat Penerbit:", data[i].penerbit.alamat)
		fmt.Println("Telepon Penerbit:", data[i].penerbit.telepon)
	}
}

func cariBukuSequential(data [maxBuku]Buku, jumlah int, judul string) int {
	for i := 0; i < jumlah; i++ {
		if data[i].judul == judul {
			return i
		}
	}
	return -1
}

func cariBukuBinary(data *[maxBuku]Buku, jumlah int, judul string) int {
	sortJudul(data, jumlah, true)
	kiri, kanan := 0, jumlah-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah].judul == judul {
			return tengah
		} else if data[tengah].judul < judul {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func editBuku(data *[maxBuku]Buku, jumlah int) {
	var judul string
	fmt.Print("Masukkan judul buku yang ingin diedit: ")
	fmt.Scan(&judul)
	idx := cariBukuSequential(*data, jumlah, judul)
	if idx == -1 {
		fmt.Println("Buku tidak ditemukan.")
		return
	}
	fmt.Println("Masukkan data baru untuk buku:")
	tambahBuku(data, &idx)
	fmt.Println("Data buku berhasil diubah.")
}

func hapusBuku(data *[maxBuku]Buku, jumlah *int) {
	var judul string
	fmt.Print("Masukkan judul buku yang ingin dihapus: ")
	fmt.Scan(&judul)
	idx := cariBukuSequential(*data, *jumlah, judul)
	if idx == -1 {
		fmt.Println("Buku tidak ditemukan.")
		return
	}
	for i := idx; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}
	*jumlah--
	fmt.Println("Buku berhasil dihapus.")
}

func sortJudul(data *[maxBuku]Buku, jumlah int, ascending bool) {
	for i := 0; i < jumlah-1; i++ {
		idx := i
		for j := i + 1; j < jumlah; j++ {
			if (ascending && data[j].judul < data[idx].judul) ||
				(!ascending && data[j].judul > data[idx].judul) {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func sortTahunTerbit(data *[maxBuku]Buku, jumlah int, ascending bool) {
	for i := 1; i < jumlah; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && ((ascending && data[j].tahunTerbit > temp.tahunTerbit) ||
			(!ascending && data[j].tahunTerbit < temp.tahunTerbit)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func menu() {
	for {
		fmt.Println("\nPilih Array Buku:")
		fmt.Println("1. Buku Koleksi 1")
		fmt.Println("2. Buku Koleksi 2")
		fmt.Println("3. Buku Koleksi 3")
		fmt.Println("4. Keluar")
		fmt.Print("Pilihan: ")
		var pilihArray int
		fmt.Scan(&pilihArray)

		var data *[maxBuku]Buku
		var jumlah *int
		switch pilihArray {
		case 1:
			data, jumlah = &daftarBuku1, &jumlahBuku1
		case 2:
			data, jumlah = &daftarBuku2, &jumlahBuku2
		case 3:
			data, jumlah = &daftarBuku3, &jumlahBuku3
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		for {
			fmt.Println("\nMenu Perpustakaan:")
			fmt.Println("1. Tambah Buku")
			fmt.Println("2. Tampilkan Buku")
			fmt.Println("3. Urutkan Judul (ASC)")
			fmt.Println("4. Urutkan Judul (DESC)")
			fmt.Println("5. Urutkan Tahun Terbit (ASC)")
			fmt.Println("6. Urutkan Tahun Terbit (DESC)")
			fmt.Println("7. Cari Buku (Sequential)")
			fmt.Println("8. Cari Buku (Binary)")
			fmt.Println("9. Edit Buku")
			fmt.Println("10. Hapus Buku")
			fmt.Println("11. Kembali")
			fmt.Print("Pilih menu: ")
			var pilihan int
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				tambahBuku(data, jumlah)
			case 2:
				tampilkanBuku(*data, *jumlah)
			case 3:
				sortJudul(data, *jumlah, true)
			case 4:
				sortJudul(data, *jumlah, false)
			case 5:
				sortTahunTerbit(data, *jumlah, true)
			case 6:
				sortTahunTerbit(data, *jumlah, false)
			case 7:
				var judul string
				fmt.Print("Masukkan judul: ")
				fmt.Scan(&judul)
				idx := cariBukuSequential(*data, *jumlah, judul)
				if idx != -1 {
					fmt.Println("Buku ditemukan di indeks:", idx)
				} else {
					fmt.Println("Buku tidak ditemukan.")
				}
			case 8:
				var judul string
				fmt.Print("Masukkan judul: ")
				fmt.Scan(&judul)
				idx := cariBukuBinary(data, *jumlah, judul)
				if idx != -1 {
					fmt.Println("Buku ditemukan di indeks:", idx)
				} else {
					fmt.Println("Buku tidak ditemukan.")
				}
			case 9:
				editBuku(data, *jumlah)
			case 10:
				hapusBuku(data, jumlah)
			case 11:
				return
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		}
	}
}

func main() {
	menu()
}

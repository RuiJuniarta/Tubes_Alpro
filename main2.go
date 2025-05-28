package main

import (
	"fmt"
	"sort"
	"strings"
)

type Mahasiswa struct {
	NIM   string
	Nama  string
	Prodi string
}

type MataKuliah struct {
	Kode string
	Nama string
	SKS  int
}

type Nilai struct {
	NIM        string
	KodeMK     string
	NilaiHuruf string
}

var daftarMahasiswa []Mahasiswa
var daftarMatkul []MataKuliah
var daftarNilai []Nilai

func main() {
	for {
		fmt.Println("\n=== Aplikasi Dosen Wali ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tambah Mata Kuliah")
		fmt.Println("3. Tambah Nilai")
		fmt.Println("4. Lihat Daftar Mahasiswa")
		fmt.Println("5. Cari Mahasiswa (NIM/Nama)")
		fmt.Println("6. Urutkan Mahasiswa (IPK/Nama)")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahMahasiswa()
		case 2:
			tambahMatkul()
		case 3:
			tambahNilai()
		case 4:
			tampilkanMahasiswa()
		case 5:
			cariMahasiswa()
		case 6:
			urutkanMahasiswa()
		case 7:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahMahasiswa() {
	var nim, nama, prodi string

	fmt.Print("NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Prodi: ")
	fmt.Scanln(&prodi)

	daftarMahasiswa = append(daftarMahasiswa, Mahasiswa{NIM: nim, Nama: nama, Prodi: prodi})
	fmt.Println("Mahasiswa ditambahkan.")
}

func tambahMatkul() {
	var kode, nama string
	var sks int

	fmt.Print("Kode MK: ")
	fmt.Scanln(&kode)

	fmt.Print("Nama MK: ")
	fmt.Scanln(&nama)

	fmt.Print("SKS: ")
	fmt.Scanln(&sks)

	daftarMatkul = append(daftarMatkul, MataKuliah{Kode: kode, Nama: nama, SKS: sks})
	fmt.Println("Mata kuliah ditambahkan.")
}

func tambahNilai() {
	var nim, kode, nilai string

	fmt.Print("NIM Mahasiswa: ")
	fmt.Scanln(&nim)

	fmt.Print("Kode MK: ")
	fmt.Scanln(&kode)

	fmt.Print("Nilai Huruf (A/B/C/D/E): ")
	fmt.Scanln(&nilai)
	nilai = strings.ToUpper(nilai)

	daftarNilai = append(daftarNilai, Nilai{NIM: nim, KodeMK: kode, NilaiHuruf: nilai})
	fmt.Println("Nilai ditambahkan.")
}

func konversiNilai(huruf string) float64 {
	switch huruf {
	case "A":
		return 4.0
	case "B":
		return 3.0
	case "C":
		return 2.0
	case "D":
		return 1.0
	default:
		return 0.0
	}
}

func hitungIPK(nim string) float64 {
	totalSKS := 0
	totalNilai := 0.0

	for _, n := range daftarNilai {
		if n.NIM == nim {
			for _, mk := range daftarMatkul {
				if mk.Kode == n.KodeMK {
					bobot := konversiNilai(n.NilaiHuruf)
					totalNilai += bobot * float64(mk.SKS)
					totalSKS += mk.SKS
				}
			}
		}
	}

	if totalSKS == 0 {
		return 0.0
	}
	return totalNilai / float64(totalSKS)
}

func tampilkanMahasiswa() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("\nDaftar Mahasiswa:")
	for _, m := range daftarMahasiswa {
		fmt.Printf("%s - %s - %s - IPK: %.2f\n", m.NIM, m.Nama, m.Prodi, hitungIPK(m.NIM))
	}
}

func cariMahasiswa() {
	var keyword string
	fmt.Print("Masukkan NIM atau Nama: ")
	fmt.Scanln(&keyword)
	keyword = strings.ToLower(keyword)

	found := false
	for _, m := range daftarMahasiswa {
		if strings.Contains(strings.ToLower(m.NIM), keyword) || strings.Contains(strings.ToLower(m.Nama), keyword) {
			fmt.Printf("%s - %s - %s - IPK: %.2f\n", m.NIM, m.Nama, m.Prodi, hitungIPK(m.NIM))
			found = true
		}
	}

	if !found {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}
}

func urutkanMahasiswa() {
	var pilihan int
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. IPK (Tinggi ke Rendah)")
	fmt.Println("2. Nama (A-Z)")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		sort.Slice(daftarMahasiswa, func(i, j int) bool {
			return hitungIPK(daftarMahasiswa[i].NIM) > hitungIPK(daftarMahasiswa[j].NIM)
		})
		fmt.Println("Data diurutkan berdasarkan IPK.")
	case 2:
		sort.Slice(daftarMahasiswa, func(i, j int) bool {
			return daftarMahasiswa[i].Nama < daftarMahasiswa[j].Nama
		})
		fmt.Println("Data diurutkan berdasarkan Nama.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}

	tampilkanMahasiswa()
}

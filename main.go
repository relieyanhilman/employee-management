package main

import (
	"fmt"
)

// format struktur karyawan
type Karyawan struct {
	ID      int
	Nama    string
	Usia    int
	Jabatan string
	Gaji    float64
}

//fungsi CRUD karyawan

// Ini fungsi untuk menambahkan karyawan ke slice utama
func TambahKaryawan(karyawanBaru Karyawan, karyawan *[]Karyawan) {
	*karyawan = append(*karyawan, karyawanBaru)
}

func DisplayKaryawan(karyawan *[]Karyawan) {
	fmt.Println(*karyawan)
}

func HapusKaryawan(id int, karyawan *[]Karyawan) []Karyawan {

	if len(*karyawan) == 0 {
		fmt.Println("data karyawan yang dimaksud tidak ditemukan")
		return *karyawan
	}

	for i, individu := range *karyawan {
		if individu.ID == id {
			(*karyawan)[i] = (*karyawan)[len(*karyawan)-1]
			*karyawan = (*karyawan)[:len(*karyawan)-1]
			return *karyawan
		}
	}

	fmt.Println("data karyawan yang dimaksud tidak ditemukan")
	return *karyawan

}

func PerbaruiKaryawan(id int, nama *string, usia *int, jabatan *string, karyawan *[]Karyawan) *Karyawan {
	//cari karyawan berdasarkan id
	for i := range *karyawan {
		if (*karyawan)[i].ID == id {
			if nama != nil {
				(*karyawan)[i].Nama = *nama
			}
			if usia != nil {
				(*karyawan)[i].Usia = *usia
			}
			if jabatan != nil {
				(*karyawan)[i].Jabatan = *jabatan
			}
			return &(*karyawan)[i]
		}
	}
	return nil
}

func CariKaryawan(id *int, nama *string, usia *int, jabatan *string, karyawan *[]Karyawan) []Karyawan {
	listHasil := []Karyawan{}
	for _, individu := range *karyawan {
		matched := true
		if id != nil && *id != individu.ID {
			matched = false
		}
		if nama != nil && *nama != individu.Nama {
			matched = false
		}
		if usia != nil && *usia != individu.Usia {
			matched = false
		}
		if jabatan != nil && *jabatan != individu.Jabatan {
			matched = false
		}
		if matched {
			listHasil = append(listHasil, individu)
		}
	}
	if len(listHasil) == 0 {
		fmt.Println("Data karyawan tidak ditemukan")
		return listHasil
	}

	return listHasil
}

func main() {
	dataKaryawan := []Karyawan{}
	TambahKaryawan(Karyawan{ID: 1, Nama: "rieco hilman", Usia: 18, Jabatan: "junior", Gaji: 2000000}, &dataKaryawan)
	DisplayKaryawan(&dataKaryawan)

	TambahKaryawan(Karyawan{ID: 2, Nama: "relieyan hilman", Usia: 23, Jabatan: "Senior", Gaji: 5000000}, &dataKaryawan)
	DisplayKaryawan(&dataKaryawan)
	namaLengkap := "Relieyan Ramadhan Hilman"
	PerbaruiKaryawan(2, &namaLengkap, nil, nil, &dataKaryawan)

	DisplayKaryawan(&dataKaryawan)
}

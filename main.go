package main

import (
	"errors"
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

func HapusKaryawan(id int, karyawan *[]Karyawan) error {

	if len(*karyawan) == 0 {
		return errors.New("data karyawan tidak ditemukan")
	}

	for i, individu := range *karyawan {
		if individu.ID == id {
			(*karyawan)[i] = (*karyawan)[len(*karyawan)-1]
			*karyawan = (*karyawan)[:len(*karyawan)-1]
			return nil
		}
	}

	return errors.New("data karyawan tidak ditemukan")

}

func PerbaruiKaryawan(id int, nama *string, usia *int, jabatan *string, karyawan *[]Karyawan) error {
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
			return nil
		}
	}
	return errors.New("data karyawan tidak ditemukan")
}

func CariKaryawan(id *int, nama *string, usia *int, jabatan *string, karyawan *[]Karyawan) ([]Karyawan, error) {
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
		return nil, errors.New("data karyawan tidak ditemukan")
	}

	return listHasil, nil
}

func main() {
	dataKaryawan := []Karyawan{}
	TambahKaryawan(Karyawan{ID: 1, Nama: "rieco hilman", Usia: 18, Jabatan: "junior", Gaji: 2000000}, &dataKaryawan)
	DisplayKaryawan(&dataKaryawan)

	// TambahKaryawan(Karyawan{ID: 2, Nama: "relieyan hilman", Usia: 23, Jabatan: "Senior", Gaji: 5000000}, &dataKaryawan)
	// DisplayKaryawan(&dataKaryawan)
	namaLengkap := "Relieyan Ramadhan Hilman"

	err := PerbaruiKaryawan(2, &namaLengkap, nil, nil, &dataKaryawan)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// DisplayKaryawan(&dataKaryawan)

	err = HapusKaryawan(2, &dataKaryawan)
	if err != nil {
		fmt.Println("Error:", err)
	}

	usiaYangDIcari := 15
	cariKaryawanUsia23, err := CariKaryawan(nil, nil, &usiaYangDIcari, nil, &dataKaryawan)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Hasil Pencarian:", cariKaryawanUsia23)

	DisplayKaryawan(&dataKaryawan)

}

package employee

import (
	"errors"
)

// format struktur karyawan
type Karyawan struct {
	ID      int
	Nama    string
	Usia    int
	Jabatan string
	Gaji    float64
}

var DataPusatKaryawan = []Karyawan{}
var PointerDataPusatKaryawan = &DataPusatKaryawan

//fungsi CRUD karyawan

// Ini fungsi untuk menambahkan karyawan ke slice utama
func TambahKaryawan(karyawanBaru Karyawan, karyawan *[]Karyawan) {
	if len(*karyawan) == 0 {
		karyawanBaru.ID = 1
	} else {
		karyawanBaru.ID = (*karyawan)[(len(*karyawan)-1)].ID + 1
	}
	*karyawan = append(*karyawan, karyawanBaru)
}

func DisplayKaryawan() []Karyawan {
	return *PointerDataPusatKaryawan
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

func PerbaruiKaryawan(id int, nama *string, usia *int, jabatan *string, gaji *float64, karyawan *[]Karyawan) error {
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
			if gaji != nil {
				(*karyawan)[i].Gaji = *gaji
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

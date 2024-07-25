package main

import "fmt"

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
func TambahKaryawan(karyawanBaru Karyawan, karyawan []Karyawan) []Karyawan {
	karyawan = append(karyawan, karyawanBaru)
	return karyawan
}

func DisplayKaryawan(karyawan []Karyawan) {
	fmt.Println(karyawan)
}

// func HapusKaryawan(id int) int {
// 	return 1
// }

// func PerbaruiKaryawan(id int) {}

func CariKaryawan(nama *string, usia *int, jabatan *string, karyawan []Karyawan) any {
	listHasil := []Karyawan{}
	for _, individu := range karyawan {
		matched := true
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
		return nil
	}

	return listHasil
}

func main() {
	dataKaryawan := []Karyawan{}
	dataKaryawan = TambahKaryawan(Karyawan{ID: 1, Nama: "rieco hilman", Usia: 23, Jabatan: "junior", Gaji: 2000000}, dataKaryawan)
	DisplayKaryawan(dataKaryawan)

	dataKaryawan = TambahKaryawan(Karyawan{ID: 2, Nama: "relieyan hilman", Usia: 23, Jabatan: "Senior", Gaji: 5000000}, dataKaryawan)

	// namaDicari := "relieyan hilman"
	usiaDicari := 23
	// jabatanDicari := "Senior"

	pencarianKaryawan := CariKaryawan(nil, &usiaDicari, nil, dataKaryawan)
	fmt.Println("hasil pencarian: ", pencarianKaryawan)
}

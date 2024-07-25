package main

import (
	"errors"
	"testing"
)

func TestTambahKaryawan(t *testing.T) {
	kumpulanKaryawan := []Karyawan{}
	karyawanBaru := Karyawan{1, "afu arifan", 25, "mobile app developer", 6000000}
	TambahKaryawan(karyawanBaru, &kumpulanKaryawan)
	if len(kumpulanKaryawan) != 1 {
		t.Fatalf("jumlah karyawan sekarang seharusnya 1, tapi kita mendapatkan %d", len(kumpulanKaryawan))
	}

	if kumpulanKaryawan[0].ID != karyawanBaru.ID {
		t.Errorf("data ID karyawan seharusnya %d, tapi tertulis %d", karyawanBaru.ID, kumpulanKaryawan[0].ID)
	}

	if kumpulanKaryawan[0].Nama != karyawanBaru.Nama {
		t.Errorf("data Nama karyawan seharusnya %s, tapi tertulis %s", karyawanBaru.Nama, kumpulanKaryawan[0].Nama)
	}

	if kumpulanKaryawan[0].Usia != karyawanBaru.Usia {
		t.Errorf("data usia karyawan seharusnya %d, tapi tertulis %d", karyawanBaru.Usia, kumpulanKaryawan[0].Usia)
	}

	if kumpulanKaryawan[0].Jabatan != karyawanBaru.Jabatan {
		t.Errorf("data jabatan karyawan seharusnya %s, tapi tertulis %s", karyawanBaru.Jabatan, kumpulanKaryawan[0].Jabatan)
	}

	if kumpulanKaryawan[0].Gaji != karyawanBaru.Gaji {
		t.Errorf("data gaji karyawan seharusnya %f, tapi tertulis %f", karyawanBaru.Gaji, kumpulanKaryawan[0].Gaji)
	}
}

func TestHapusKaryawan(t *testing.T) {
	tests := []struct {
		name          string
		initialData   []Karyawan
		idToDelete    int
		expectedData  []Karyawan
		expectedError error
	}{
		{
			name:          "Hapus dari daftar kosong",
			initialData:   []Karyawan{},
			idToDelete:    1,
			expectedData:  []Karyawan{},
			expectedError: errors.New("data karyawan tidak ditemukan"),
		},
		{
			name:          "Hapus karyawan yang ada",
			initialData:   []Karyawan{{ID: 1, Nama: "Afu Arifan", Usia: 25, Jabatan: "Developer", Gaji: 6000000}},
			idToDelete:    1,
			expectedData:  []Karyawan{},
			expectedError: nil,
		},
		{
			name:          "Hapus karyawan yang tidak ada",
			initialData:   []Karyawan{{ID: 1, Nama: "Afu Arifan", Usia: 25, Jabatan: "Developer", Gaji: 6000000}},
			idToDelete:    2,
			expectedData:  []Karyawan{{ID: 1, Nama: "Afu Arifan", Usia: 25, Jabatan: "Developer", Gaji: 6000000}},
			expectedError: errors.New("data karyawan tidak ditemukan"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataKaryawan := tt.initialData
			err := HapusKaryawan(tt.idToDelete, &dataKaryawan)

			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}

			if len(dataKaryawan) != len(tt.expectedData) {
				t.Fatalf("Expected %d karyawan, got %d", len(tt.expectedData), len(dataKaryawan))
			}

			for i, k := range dataKaryawan {
				if k != tt.expectedData[i] {
					t.Errorf("Expected %v, got %v", tt.expectedData[i], k)
				}
			}
		})
	}
}

func TestPerbaruiKaryawan(t *testing.T) {
	tests := []struct {
		name          string
		initialData   []Karyawan
		idToUpdate    int
		nama          *string
		usia          *int
		jabatan       *string
		expectedData  []Karyawan
		expectedError error
	}{
		{
			name:          "Perbarui karyawan yang ada - semua data",
			initialData:   []Karyawan{{ID: 1, Nama: "Afu Arifan", Usia: 25, Jabatan: "Developer", Gaji: 6000000}},
			idToUpdate:    1,
			nama:          stringPointer("Afu Arifan Updated"),
			usia:          intPointer(26),
			jabatan:       stringPointer("Senior Developer"),
			expectedData:  []Karyawan{{ID: 1, Nama: "Afu Arifan Updated", Usia: 26, Jabatan: "Senior Developer", Gaji: 6000000}},
			expectedError: nil,
		},
		{
			name:          "Perbarui karyawan yang ada - sebagian data",
			initialData:   []Karyawan{{ID: 2, Nama: "Budi", Usia: 30, Jabatan: "Manager", Gaji: 8000000}},
			idToUpdate:    2,
			nama:          nil,
			usia:          intPointer(31),
			jabatan:       nil,
			expectedData:  []Karyawan{{ID: 2, Nama: "Budi", Usia: 31, Jabatan: "Manager", Gaji: 8000000}},
			expectedError: nil,
		},
		{
			name:          "Perbarui karyawan yang tidak ada",
			initialData:   []Karyawan{{ID: 3, Nama: "Cici", Usia: 22, Jabatan: "Intern", Gaji: 4000000}},
			idToUpdate:    4,
			nama:          stringPointer("Dewi"),
			usia:          intPointer(23),
			jabatan:       stringPointer("Junior Developer"),
			expectedData:  []Karyawan{{ID: 3, Nama: "Cici", Usia: 22, Jabatan: "Intern", Gaji: 4000000}},
			expectedError: errors.New("data karyawan tidak ditemukan"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataKaryawan := tt.initialData
			err := PerbaruiKaryawan(tt.idToUpdate, tt.nama, tt.usia, tt.jabatan, &dataKaryawan)

			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Fatalf("Expected error %v, got %v", tt.expectedError, err)
			}

			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}

			if len(dataKaryawan) != len(tt.expectedData) {
				t.Fatalf("Expected %d karyawan, got %d", len(tt.expectedData), len(dataKaryawan))
			}

			for i, k := range dataKaryawan {
				if k != tt.expectedData[i] {
					t.Errorf("Expected %v, got %v", tt.expectedData[i], k)
				}
			}
		})
	}
}

func TestCariKaryawan(t *testing.T) {
	tests := []struct {
		name          string
		initialData   []Karyawan
		id            *int
		nama          *string
		usia          *int
		jabatan       *string
		expectedData  []Karyawan
		expectedError error
	}{
		{
			name:          "Cari karyawan berdasarkan ID",
			initialData:   []Karyawan{{ID: 1, Nama: "Afu Arifan", Usia: 25, Jabatan: "Developer", Gaji: 6000000}},
			id:            intPointer(1),
			nama:          nil,
			usia:          nil,
			jabatan:       nil,
			expectedData:  []Karyawan{{ID: 1, Nama: "Afu Arifan", Usia: 25, Jabatan: "Developer", Gaji: 6000000}},
			expectedError: nil,
		},
		{
			name:          "Cari karyawan berdasarkan nama",
			initialData:   []Karyawan{{ID: 2, Nama: "Budi", Usia: 30, Jabatan: "Manager", Gaji: 8000000}},
			id:            nil,
			nama:          stringPointer("Budi"),
			usia:          nil,
			jabatan:       nil,
			expectedData:  []Karyawan{{ID: 2, Nama: "Budi", Usia: 30, Jabatan: "Manager", Gaji: 8000000}},
			expectedError: nil,
		},
		{
			name:          "Cari karyawan berdasarkan jabatan",
			initialData:   []Karyawan{{ID: 3, Nama: "Cici", Usia: 22, Jabatan: "Intern", Gaji: 4000000}},
			id:            nil,
			nama:          nil,
			usia:          nil,
			jabatan:       stringPointer("Intern"),
			expectedData:  []Karyawan{{ID: 3, Nama: "Cici", Usia: 22, Jabatan: "Intern", Gaji: 4000000}},
			expectedError: nil,
		},
		{
			name:          "Cari karyawan berdasarkan kombinasi parameter",
			initialData:   []Karyawan{{ID: 4, Nama: "Dewi", Usia: 23, Jabatan: "Junior Developer", Gaji: 5000000}},
			id:            intPointer(4),
			nama:          stringPointer("Dewi"),
			usia:          intPointer(23),
			jabatan:       stringPointer("Junior Developer"),
			expectedData:  []Karyawan{{ID: 4, Nama: "Dewi", Usia: 23, Jabatan: "Junior Developer", Gaji: 5000000}},
			expectedError: nil,
		},
		{
			name:          "Karyawan tidak ditemukan",
			initialData:   []Karyawan{{ID: 5, Nama: "Eko", Usia: 28, Jabatan: "Senior Developer", Gaji: 9000000}},
			id:            intPointer(6),
			nama:          nil,
			usia:          nil,
			jabatan:       nil,
			expectedData:  nil,
			expectedError: errors.New("data karyawan tidak ditemukan"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataKaryawan := tt.initialData
			result, err := CariKaryawan(tt.id, tt.nama, tt.usia, tt.jabatan, &dataKaryawan)

			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Fatalf("Expected error %v, got %v", tt.expectedError, err)
			}

			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}

			if len(result) != len(tt.expectedData) {
				t.Fatalf("Expected %d karyawan, got %d", len(tt.expectedData), len(result))
			}

			for i, k := range result {
				if k != tt.expectedData[i] {
					t.Errorf("Expected %v, got %v", tt.expectedData[i], k)
				}
			}
		})
	}
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}

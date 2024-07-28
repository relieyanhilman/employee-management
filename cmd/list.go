package cmd

import (
	"fmt"

	"github.com/relieyanhilman/latihan1_pengelolaanDataKaryawan/employee"
	"github.com/urfave/cli/v2"
)

func DisplayKaryawanCmd(c *cli.Context) error {
	karyawans := employee.DisplayKaryawan()
	if len(karyawans) == 0 {
		fmt.Println("Tidak ada karyawan.")
	} else {
		for _, k := range karyawans {
			fmt.Printf("ID: %d, Nama: %s, Usia: %d, Jabatan: %s, Gaji: %0.1f\n", k.ID, k.Nama, k.Usia, k.Jabatan, k.Gaji)
		}
	}
	return nil
}

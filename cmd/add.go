package cmd

import (
	"fmt"

	"github.com/relieyanhilman/latihan1_pengelolaanDataKaryawan/employee"
	"github.com/urfave/cli/v2"
)

func TambahKaryawanCmd(c *cli.Context) error {
	nama := c.String("nama")
	usia := c.Int("usia")
	jabatan := c.String("jabatan")
	gaji := c.Float64("gaji")

	karyawanBaru := employee.Karyawan{
		Nama:    nama,
		Usia:    usia,
		Jabatan: jabatan,
		Gaji:    gaji,
	}

	employee.TambahKaryawan(karyawanBaru, employee.PointerDataPusatKaryawan)
	fmt.Printf("Karyawan %s berhasil ditambahkan\n", nama)
	return nil
}

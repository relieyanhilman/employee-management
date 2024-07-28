package cmd

import (
	"fmt"

	"github.com/relieyanhilman/latihan1_pengelolaanDataKaryawan/employee"
	"github.com/urfave/cli/v2"
)

func PerbaruiKaryawanCmd(c *cli.Context) error {
	id := c.Int("id")
	var nama *string
	var usia *int
	var jabatan *string
	var gaji *float64

	if c.IsSet("nama") {
		n := c.String("nama")
		nama = &n
	}
	if c.IsSet("usia") {
		u := c.Int("usia")
		usia = &u
	}
	if c.IsSet("jabatan") {
		j := c.String("jabatan")
		jabatan = &j
	}
	if c.IsSet("gaji") {
		g := c.Float64("gaji")
		gaji = &g
	}

	err := employee.PerbaruiKaryawan(id, nama, usia, jabatan, gaji, employee.PointerDataPusatKaryawan)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Karyawan dengan ID %d berhasil diperbarui\n", id)
	}
	return nil
}

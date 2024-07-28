package cmd

import (
	"fmt"

	"github.com/relieyanhilman/latihan1_pengelolaanDataKaryawan/employee"
	"github.com/urfave/cli/v2"
)

func HapusKaryawanCmd(c *cli.Context) error {
	id := c.Int("id")
	err := employee.HapusKaryawan(id, employee.PointerDataPusatKaryawan)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Karyawan dengan ID %d berhasil dihapus\n", id)
	}
	return nil
}

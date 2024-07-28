package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/relieyanhilman/latihan1_pengelolaanDataKaryawan/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "employee-management",
		Usage: "Manajemen data karyawan",
		Commands: []*cli.Command{
			{
				Name:   "tambah",
				Usage:  "Menambah karyawan baru",
				Action: cmd.TambahKaryawanCmd,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "nama", Usage: "Nama karyawan", Required: true},
					&cli.IntFlag{Name: "usia", Usage: "Usia karyawan", Required: true},
					&cli.StringFlag{Name: "jabatan", Usage: "Jabatan karyawan", Required: true},
					&cli.Float64Flag{Name: "gaji", Usage: "Gajikaryawan", Required: true},
				},
			},
			{
				Name:   "list",
				Usage:  "Menampilkan daftar karyawan",
				Action: cmd.DisplayKaryawanCmd,
			},
			{
				Name:   "hapus",
				Usage:  "Menghapus karyawan berdasarkan ID",
				Action: cmd.HapusKaryawanCmd,
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "id", Usage: "ID karyawan yang akan dihapus", Required: true},
				},
			},
			{
				Name:   "update",
				Usage:  "Memperbarui data karyawan",
				Action: cmd.PerbaruiKaryawanCmd,
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "id", Usage: "ID karyawan", Required: true},
					&cli.StringFlag{Name: "nama", Usage: "Nama karyawan"},
					&cli.IntFlag{Name: "usia", Usage: "Usia karyawan"},
					&cli.StringFlag{Name: "jabatan", Usage: "Jabatan karyawan"},
					&cli.Float64Flag{Name: "gaji", Usage: "Gaji karyawan"},
				},
			},
		},
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Remove the newline character
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}
		// kode di bawah ini yang perlu dimodifikasi untuk bisa mencegah
		// bug string pake spasi yang dianggap command yang berbeda
		args := parseArgs(input)
		err = app.Run(args)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

func parseArgs(input string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false

	for _, char := range input {
		switch char {
		case ' ':
			if inQuotes {
				currentArg.WriteRune(char)
			} else {
				if currentArg.Len() > 0 {
					args = append(args, currentArg.String())
					currentArg.Reset()
				}
			}
		case '"':
			inQuotes = !inQuotes
		default:
			currentArg.WriteRune(char)
		}
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	// Adding a fake "cli-app" as the first argument
	return append([]string{"cli-app"}, args...)
}

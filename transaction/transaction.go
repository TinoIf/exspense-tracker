package transaction

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	Deskripsi string
	Harga     float64
	Tipe      string
}

var Laporan = []Transaction{}

func AddTransaction(Lapor []Transaction) []Transaction {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan Deskripsi Pengeluaran atau Pemasukkan : ")
	deskripsi, _ := reader.ReadString('\n')
	deskripsi = strings.TrimSpace(deskripsi)

	fmt.Println("Masukkan Harga : ")
	inputHarga, _ := reader.ReadString('\n')
	harga, err := strconv.ParseFloat(inputHarga, 64)

	if err != nil {
		fmt.Println("Konversi Harga Gagal Dilakukan, Masukkan Harga tanpa huruf")
	}

	fmt.Println("Masukkan Jenis nya Ketik Pengeluaran / Pemasukkan")
	tipe, _ := reader.ReadString('\n')

	olahLaporan := Transaction{
		Deskripsi: deskripsi,
		Harga:     harga,
		Tipe:      tipe,
	}

	Lapor = append(Lapor, olahLaporan)

	return Lapor
}

func DisplayTransaction(Lapor []Transaction) {

	fmt.Println("======================== Berikut Laporan Transaksi Kamu ==============================")
	fmt.Println("No\tDeskripsi\t\tHarga\t\tTipe")
	for index, value := range Lapor {
		fmt.Print(index+1, "\t", value.Deskripsi, "\t\t", value.Harga, "\t", value.Tipe)
	}
}

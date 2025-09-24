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
	fmt.Print("Masukkan Deskripsi Transaksi: ")
	deskripsi, _ := reader.ReadString('\n')
	deskripsi = strings.TrimSpace(deskripsi)

	fmt.Print("Masukkan Harga : ")
	inputHarga, _ := reader.ReadString('\n')
	inputHarga = strings.TrimSpace(inputHarga)
	harga, err := strconv.ParseFloat(inputHarga, 64)

	if err != nil {
		fmt.Println("Konversi Harga Gagal Dilakukan, Masukkan Angka tanpa huruf")
	}

	fmt.Print("Masukkan Jenis nya Ketik Pengeluaran / Pemasukkan : ")
	tipe, _ := reader.ReadString('\n')
	tipe = strings.TrimSpace(tipe)

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
		fmt.Print(index+1, "\t", value.Deskripsi, "\t\t", value.Harga, "\t\t", value.Tipe)
	}
}

func DeleteTransaction(Lapor []Transaction) []Transaction {

	var deletedId int
	var indexDeleted int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan Index Yang Ingin Di Hapus : ")
	deletedIdS, _ := reader.ReadString('\n')
	deletedIdS = strings.TrimSpace(deletedIdS)
	deletedId, _ = strconv.Atoi(deletedIdS)
	fmt.Printf("Index nya adalah %d", deletedId)

	if deletedId < 0 || deletedId > len(Lapor) {
		fmt.Println("Maaf Index Yang Kamu Masukan Tidak ada dan Tidak Valid")
		return Lapor
	}

	for index := range Lapor {
		if index == deletedId {
			indexDeleted = index - 1
			break
		}
	}

	Lapor = append(Lapor[:indexDeleted], Lapor[indexDeleted+1:]...)
	return Lapor

}

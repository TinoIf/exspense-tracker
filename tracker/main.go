package main

import (
	"bufio"
	"exspense-tracker/transaction"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var pilihan string
	for {
		fmt.Println()
		fmt.Println("================================= Exspense Tracker =================================")
		fmt.Println("================================= Pilihlah Menu    =================================")
		fmt.Println("1. Masukkan Transaksi")
		fmt.Println("2. Tampilkan Semua Catatan Transaksi")
		fmt.Println("3. Hapus Transaksi")
		fmt.Println("4. Keluar Program")
		fmt.Print("Masukkan Nilai yang diPilih : ")
		reader := bufio.NewReader(os.Stdin)
		pilihan, _ = reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)
		pilihanMenu, _ := strconv.Atoi(pilihan)

		if pilihanMenu == 1 {
			transaction.Laporan = transaction.AddTransaction(transaction.Laporan)
			continue
		}
		if pilihanMenu == 2 {
			transaction.DisplayTransaction(transaction.Laporan)
			continue
		}
		if pilihanMenu == 3 {
			transaction.Laporan = transaction.DeleteTransaction(transaction.Laporan)
			continue
		}
		if pilihanMenu == 4 {
			fmt.Println("=========== Terimakasih Telah Menggunakan Program =============")
			break
		}
	}

}

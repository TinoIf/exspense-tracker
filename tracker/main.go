package main

import (
	"exspense-tracker/transaction"
	"fmt"
)

func main() {
	var pilihan int

	fmt.Println("================================= Exspense Tracker =================================")
	fmt.Println("================================= Pilihlah Menu    =================================")
	fmt.Println("1. Masukkan Transaksi")
	fmt.Println("2. Tampilkan Semua Catatan Transaksi")
	fmt.Println("3. Hapus Transaksi")
	fmt.Println("4. Keluar Program")
	fmt.Print("Masukkan Nilai yang diPilih : ")
	fmt.Scan(&pilihan)

	for {
		if pilihan == 1 {
			transaction.Laporan = transaction.AddTransaction(transaction.Laporan)
		}
		if pilihan == 2 {
			transaction.DisplayTransaction(transaction.Laporan)
		}
		if pilihan == 3 {
			//Ini Pilihan 3 Untuk Menghapus Salah Satu Transaksi
		}
		if pilihan == 4 {
			fmt.Println("=========== Terimakasih Telah Menggunakan Program =============")
			break
		}
	}

}

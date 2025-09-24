# Aplikasi Expense Tracker Sederhana 📊

Ini adalah aplikasi CLI (*Command-Line Interface*) sederhana untuk melacak pemasukan dan pengeluaran harian. Aplikasi ini dibuat menggunakan bahasa Go dan dirancang untuk membantu pengguna mencatat transaksi keuangan mereka dengan mudah melalui terminal.

Proyek ini memanfaatkan `struct` untuk merepresentasikan data transaksi dan `slice` untuk menyimpan daftar semua transaksi yang telah dicatat.

---

## ## Fitur

* 💰 **Menambah Pemasukan**: Mencatat transaksi baru sebagai pemasukan.
* 💸 **Menambah Pengeluaran**: Mencatat transaksi baru sebagai pengeluaran.
* 🧾 **Melihat Semua Transaksi**: Menampilkan daftar lengkap semua pemasukan dan pengeluaran yang telah dicatat.

---

## ## Struktur Proyek

Proyek ini diorganisir menjadi dua file utama dalam dua direktori untuk memisahkan logika utama dengan definisi tipe data transaksi.

```
tracker/
├── go.mod
├── trackker
|   └── main.go                // File utama, berisi logika menu dan interaksi pengguna
└── transaction/
    └── transaction.go      // Mendefinisikan struct Transaksi dan fungsi terkait
```

---

## ## Struktur Data

Inti dari aplikasi ini adalah `struct` `Transaksi` yang menyimpan detail dari setiap catatan keuangan.

```go
// Didefinisikan di dalam file transaction/transaction.go
package transaction

// Transaksi merepresentasikan satu catatan keuangan, baik itu pemasukan atau pengeluaran.
type Transaksi struct {
    Deskripsi string
    Jumlah    float64 // Menggunakan float64 untuk menampung nilai desimal
    Tipe      string  // "Pemasukan" atau "Pengeluaran"
}
```

---

## ## Cara Menjalankan

Untuk menjalankan aplikasi ini, ikuti langkah-langkah berikut:

1.  **Clone Repositori (Jika sudah di GitHub)**
    ```bash
    git clone [URL_REPOSITORI_ANDA]
    cd ke [URL_REPOSITORI_ANDA]
    ```

2.  **Masuk ke Direktori Proyek**
    ```bash
    cd tracker
    ```

3.  **Jalankan Aplikasi**
    Perintah `go run` akan secara otomatis mengompilasi dan menjalankan file `main.go` beserta paket `transaction` yang dibutuhkan.
    ```bash
    go run main.go
    ```
    Setelah itu, menu aplikasi akan muncul di terminal Anda.

---

## ## Konsep Go yang Dipelajari

Proyek ini merupakan latihan yang baik untuk memahami konsep-konsep dasar dan fundamental dalam bahasa Go, antara lain:
* **Struct**: Untuk membuat tipe data kustom.
* **Slice**: Untuk mengelola koleksi data dinamis (`[]Transaksi`).
* **Functions**: Untuk memecah logika menjadi bagian-bagian yang dapat digunakan kembali.
* **Packages**: Mengorganisir kode ke dalam paket `main` dan `transaction`.
* **Control Flow**: Penggunaan `if` dan `for` untuk alur aplikasi.
* **User Input**: Menerima dan memproses input dari pengguna melalui terminal.
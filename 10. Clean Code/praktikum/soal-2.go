package main

import "fmt"

/*
Menghapus struct Mobil pada kode sebelumnya. Dikarenakan struct mobil hanya mengembalikan nilai dari struct kendaraan
tanpa menambah fungsional lagi, Mungkin akan berbeda jika mobil memiliki jenis-jenis lain. Maka dari itu menghapus struct
mobil dan mengganti struct kendaraan menjadi mobil, serta menghapus variabel `totalroda` dikarenakan variabel ini tidak
digunakan pada program maka lebih baik dihapus supaya tidak membingungkan
*/
type Mobil struct {
	KecepatanPerJam int
}

// Mengganti variabel "m" menjadi lebih deskriptif dengan menggantinya dengan mobil agar mudah dipahami
func (mobil *Mobil) Melaju() {
	mobil.TambahKecepatan(10)
}

/* 
Menyederhanakan bentuk penjumlahan yang sebelumnya mobil.kecepatanperjam = mobil.kecepatanperjam + kecepatanbaru
menjadi mobil.kecepatanperjam += kecepatanbaru karena sama saja. kode menjadi lebih sederhana
*/
func (mobil *Mobil) TambahKecepatan(kecepatanbaru int) {
	mobil.KecepatanPerJam += kecepatanbaru
}


/* 
Juga menerapkan kejelasan func dengan menggunakan uppercase, jika ingin bisa di ekspor berarti huruf depannya uppercase
jika tidak huruf depannya di lowercase namun sambungan kata dibelakangnya tetap uppercase, contoh :

tambahKecepatan() agar mudah dibaca

dari pada -> tambahkecepatan()
serta pada variabel di fungsi main, menggunakan uppercase pada kata kedua untuk memperjelas dan mudah dibaca
*/

func main() {

	mobilCepat := Mobil{}
	mobilCepat.Melaju()
	mobilCepat.Melaju()
	mobilCepat.Melaju()

	mobilLamban := Mobil{}
	mobilLamban.Melaju()
	fmt.Printf("mobilCepat kecepatan: %d\n", mobilCepat.KecepatanPerJam)
	fmt.Printf("mobilLamban kecepatan: %d\n", mobilLamban.KecepatanPerJam)

}
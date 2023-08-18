package Soal

import "fmt"

// Soal ke-1 Mencetak Segitiga Asterik
func Asterik() {
	Clear()
	fmt.Println("Mencetak Segitiga Asterik")
	Line()
	A := 5
	for i := 0; i < A; i++ {
		// Menambahkan spasi di awal setiap baris
		for space := 0; space < A-i-1; space++ {
			fmt.Print(" ")
		}

		// Mencetak bintang
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println() 
	}
	Footer()
	Prioritas2()
}

// Soal ke-2 Mencetak Faktor Bilangan
func Faktor() {
	var angka int
	Clear()
	Line()
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	fmt.Println("")
	fmt.Printf("Faktor dari %d adalah: \n", angka)
	for i := 1; i <= angka; i++ {
		if angka %i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	Footer()
	Prioritas2()
}


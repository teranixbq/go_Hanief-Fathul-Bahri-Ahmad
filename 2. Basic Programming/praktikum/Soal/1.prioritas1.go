package Soal

import 
    "fmt"

// Soal ke-1 Menghitung Luas Trapesium
func Trapesium() {
    var a,b,t int

    Clear()
    Line()
    fmt.Print("Panjang Sisi Atas  (a) : ")
    fmt.Scan(&a)
    fmt.Print("Panjang Sisi Bawah (b) : ")
    fmt.Scan(&b)
    fmt.Print("Tinggi             (t) : ")
    fmt.Scan(&t)
    Line()
    luas := (a+b)/2 *t

    fmt.Println("Luas Trapesium = ", luas)
    Footer()
    Prioritas1()
}

// Soal ke-2 Menentukan Bilangan Ganjil Genap
func Genapganjil () {
    var angka int

    Clear()
    Line()
    fmt.Print("Masukkan Angka : ")
    fmt.Scan(&angka)
    fmt.Println("")
    if angka % 2 == 0 {
        fmt.Println(angka," Adalah bilangan Genap")
    } else {
        fmt.Println(angka," Adalah bilangan Ganjil")
    }
    Footer()
    Prioritas1()
}

// Soal ke-3 Menentukan Grade
func Grade() {
    var grade int

    Clear()
    Line()
    fmt.Print("Masukkan Nilai : ")
    fmt.Scan(&grade)
    fmt.Println("")

    switch {
	case grade >= 80 && grade <= 100:
		fmt.Println("Nilai Anda A")
	case grade >= 65 && grade <= 79:
		fmt.Println("Nilai Anda B")
	case grade >= 50 && grade <= 64:
		fmt.Println("Nilai Anda C")
	case grade >= 35 && grade <= 49:
		fmt.Println("Nilai Anda D")
	case grade >= 0 && grade <= 34:
		fmt.Println("Nilai Anda E")
	default:
		fmt.Println("Nilai Invalid")
	}
    Footer()
    Prioritas1()
}

// Soal ke-4 Buzz Fizz
func Buzzfizz () {
    Clear()
    Line()
    for i := 1; i <= 100; i++{	

        if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0{
			fmt.Println("Fizz")
		}else {
			fmt.Println(i)
		}

	}

	fmt.Println("selesai")
    Footer()
    Prioritas1()

}


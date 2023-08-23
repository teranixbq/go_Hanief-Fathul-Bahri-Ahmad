package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var pilih int
	clear()	
	fmt.Println("	Praktikum	")
	line()	
	fmt.Println(" 1. Mendeteksi Bilangan Prima ")
	fmt.Println(" 2. Perhitungan Perpangkatan ")
	line()
	fmt.Print("Input Pilihan : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1 :
		var number int
		clear()
		fmt.Print("Input Angka : ")
		fmt.Scan(&number)
		fmt.Println(primeNumber(number))
		back()
		
    case 2 :
		var x,n int
		clear()
		fmt.Print("Input x : ")
		fmt.Scan(&x)
		fmt.Print("Input n : ")
		fmt.Scan(&n)
		fmt.Println(pow(x,n))
		back()
    default :
		os.Exit(0)
	}
}


// Soal 1 Explorasi
func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		
		if number%i == 0 {
			return false
		} 
	}
	return true

}

// Soal 2 Explorasi
func pow(x, n int) int {
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x 
		n /= 2 

	}
	return result
}

func line () {
	Line := "━━━━━━━━━━━━━━━━━━━━━━━━━━"
	fmt.Println(Line)
}

// Clear Linux
func clear () {
    c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func back() {
	line()
	fmt.Print("Kembali Tekan [Enter]")
	fmt.Scanln()
	main()
}
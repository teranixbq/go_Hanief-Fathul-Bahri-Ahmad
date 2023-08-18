package Soal

import (
    "fmt"
	"os/exec"
	"os"
)


var pilih int	
func Menu() {
	Clear()
	Line()
	fmt.Println("	Praktikum	")
	Line()	
	fmt.Println(" 1. Soal Prioritas 1 ")
	fmt.Println(" 2. Soal Prioritas 2 ")
	fmt.Println(" 3. Soal Explorasi ")
	fmt.Println(" 4. Keluar ")
	inputmenu()
	fmt.Scan(&pilih)

	switch pilih {
	case 1 :
		Prioritas1()
    case 2 :
        Prioritas2()
	case 3 :
		Explorasi()
	case 4 :
		os.Exit(0)
    default :
		Menu()
	}
}


func Prioritas1 () {
	Clear()
	Line()
	fmt.Println("	Prioritas 1	")
	Line()	
	fmt.Println(" 1. Luas Trapesium ")
	fmt.Println(" 2. Menentukan Ganjil Genap ")
	fmt.Println(" 3. Menentukan Grade ")
	fmt.Println(" 4. Fizz Buzz ")
	fmt.Println(" 5. Kembali ")
	fmt.Println(" 6. Keluar ")
	inputmenu()
	fmt.Scan(&pilih)

	switch pilih {
	case 1 :
		Trapesium()
    case 2 :
        Genapganjil()
	case 3 :
		Grade()
	case 4 :
		Buzzfizz()
	case 5 :
		Menu()
	case 6 :
		os.Exit(0)
    default :
        Prioritas1()
	}

}

func Prioritas2 (){

	Clear()
	Line()
	fmt.Println("	Prioritas 2	")
	Line()	
	fmt.Println(" 1. Mencetak segitiga asterik ")
	fmt.Println(" 2. Mencetak Faktor Bilangan ")
	fmt.Println(" 3. Kembali ")
	fmt.Println(" 4. Keluar ")
	inputmenu()
	fmt.Scan(&pilih)

	switch pilih {
	case 1 :
		Asterik()
    case 2 :
        Faktor()
	case 3 :
		Menu()
	case 4 :
		os.Exit(0)
    default :
        Prioritas2()
	}
}


func Line () {
	Line := "━━━━━━━━━━━━━━━━━━━━━━━━━━"
	fmt.Println(Line)
}

func Line2(){
	Line := "--------------------------"
	fmt.Println(Line)
}

// Clear Linux
func Clear () {
    c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Footer() {
	Line()
    fmt.Print("Kembali Tekan [Enter]")
    fmt.Scanln()
}

func inputmenu() {
	Line()
	fmt.Print("Input Pilihan : ")
}
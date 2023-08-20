package Soal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Explorasi() {

	reverse := ""
	alltext := bufio.NewScanner(os.Stdin)

	Clear()
	fmt.Println("Apakah Palindrome ? ")
	Line()
	fmt.Print("Masukkan Kata\t : ")

	// Membaca baris input dan langsung menggabungkannya ke variabel 'kata'
	alltext.Scan()
	kata := alltext.Text()

	for i := len(kata) - 1; i >= 0; i-- {
		reverse += string(kata[i])
	}

	fmt.Println("Captured\t :", reverse)
	Line()

	switch {
		case strings.EqualFold(reverse, kata):
			fmt.Println("\"" + kata + "\" Adalah Kata Palindrome")
        default:
			fmt.Println("\"" + kata + "\" Bukan Kata Palindrome")
	}

	Footer()
	Menu()
}

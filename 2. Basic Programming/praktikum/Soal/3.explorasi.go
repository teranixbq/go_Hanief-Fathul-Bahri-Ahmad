package Soal

import (
	"fmt"
	"strings"
)

func Explorasi() {
	kata := ""
	reverse := ""

	Clear()
	fmt.Println("Apakah Palindrome ? ")
	Line()
	fmt.Print("Masukkan Kata\t : ")
	fmt.Scanln(&kata)

	for i := len(kata) - 1; i >= 0; i-- {
		reverse += string(kata[i])
	}

	fmt.Println("Captured\t :", reverse)
	Line()

	switch {
	case strings.EqualFold(reverse, kata):
		fmt.Println("\"" + kata + "\" adalah Kata Palindrome")
	default:
		fmt.Println("\"" + kata + "\" Bukan Kata Palindrome")
	}
	Footer()
	Menu()
}
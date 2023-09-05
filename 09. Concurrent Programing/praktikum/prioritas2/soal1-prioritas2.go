package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func countChar(data string, resultCh chan map[string]int) {
	defer close(resultCh)
	charlist := make(map[string]int)

	// Menghitung frekuensi huruf disimpan pada slice
	for _, value := range data {
		if unicode.IsLetter(value) {
			lowerchar := strings.ToLower(string(value)) // Merubah semua text menjadi lowercase
			charlist[lowerchar]++
		}
	}



	// Menyimpan key map charlist pada slice untuk di sorting
	keys := make([]string, 0, len(charlist))
	for key := range charlist {
		keys = append(keys, key)
	}

	// Menghitung total huruf
	var values int
	for _, value := range charlist {
		values += value
	}
	
	sort.Strings(keys) // Sorting

	// Menampilkan charlist berdasarkan key yang sudah diurutkan pada slice keys
	fmt.Println("---------------")
	for _, key := range keys {
		fmt.Printf("Huruf %s = %d\n", key, charlist[key])
	}
	fmt.Println("---------------")
	fmt.Println("Jumlah Alfabet	= ", len(charlist))
	fmt.Println("Total Huruf	= ", values)

	//  mengirim hasil perhitungan karakter huruf yang ada di dalam variabel charlist ke channel resultCh
	resultCh <- charlist

}

func main() {
	data := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	resultCh := make(chan map[string]int)
	fmt.Println("Data : \n",data)
	go countChar(data, resultCh)
	<-resultCh // Menerima hasil perhitungan dari goroutine

}

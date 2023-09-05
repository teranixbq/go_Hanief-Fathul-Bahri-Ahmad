package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"unicode"
)

func countChar(data string, resultCh chan map[string]int, mutex *sync.Mutex) {
	defer close(resultCh)
	charlist := make(map[string]int)

	// Menghitung frekuensi huruf disimpan pada map charlist
	for _, value := range data {
		if unicode.IsLetter(value) {
			lowerchar := strings.ToLower(string(value)) // Merubah semua teks menjadi lowercase
			mutex.Lock() // Mengunci mutex sebelum mengakses charlist
			charlist[lowerchar]++
			mutex.Unlock() // Melepaskan kunci mutex setelah selesai mengakses charlist
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
	
	mutex.Lock() // Mengunci mutex sebelum sorting dan menampilkan hasil
	sort.Strings(keys) // Sorting
	fmt.Println("---------------")
	for _, key := range keys {
		fmt.Printf("Huruf %s = %d\n", key, charlist[key])
	}
	fmt.Println("---------------")
	fmt.Println("Jumlah Alfabet	= ", len(charlist))
	fmt.Println("Total Huruf	= ", values)

	mutex.Unlock() // Melepaskan kunci mutex setelah selesai menampilkan hasil

	// Mengirim hasil perhitungan karakter huruf yang ada di dalam variabel charlist ke channel resultCh
	resultCh <- charlist
}

func main() {
	var mutex sync.Mutex
	alltext := bufio.NewScanner(os.Stdin)
	fmt.Print("Input Text : ")
	alltext.Scan()
	text := alltext.Text()

	resultCh := make(chan map[string]int)
	
	go countChar(text, resultCh, &mutex)
	<-resultCh // Menerima hasil perhitungan dari goroutine
}

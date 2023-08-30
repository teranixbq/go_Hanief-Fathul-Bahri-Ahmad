package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	// Periksa apakah b merupakan substring dari a dan sebaliknya
	if strings.Contains(a, b) {
		return b
	} else if strings.Contains(b, a) {
		return a
	} else {
		return "Tidak ada data yang sama"
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // "AKA"
	fmt.Println(Compare("KANGOORO", "KANG"))  // "KANG"
	fmt.Println(Compare("KI", "KIJANG"))      // "KI"
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // "KUPU"
	fmt.Println(Compare("ILALANG", "ILA"))    // "ILA"
}

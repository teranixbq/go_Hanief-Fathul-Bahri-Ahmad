package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	nilaijump := make([]int, n)

	// Inisialisasi nilai batu pertama
	nilaijump[0] = 0

	// Inisialisasi nilai untuk batukedua
	nilaijump[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < n; i++ {
		
		opsi1 := nilaijump[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		opsi2 := nilaijump[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))

		// Memilih nilai minimum dari kedua opsi
		nilaijump[i] = int(math.Min(float64(opsi1), float64(opsi2)))
	}

	return nilaijump[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
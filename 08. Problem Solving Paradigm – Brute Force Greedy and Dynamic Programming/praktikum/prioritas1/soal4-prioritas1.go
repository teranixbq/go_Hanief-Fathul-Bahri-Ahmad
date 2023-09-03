package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	// Loop melalui semua kemungkinan nilai x dari 1 hingga a
	for x := 1; x <= a; x++ {
		// Loop melalui semua kemungkinan nilai y dari 1 hingga a
		for y := 1; y <= a; y++ {
			// Hitung nilai z berdasarkan persamaan x + y + z = a
			z := a - x - y

			// Memeriksa apakah x, y, dan z adalah bilangan bulat positif
			if x > 0 && y > 0 && z > 0 {
				// Memeriksa apakah x*y*z = b dan x^2 + y^2 + z^2 = c
				if x*y*z == b && (x*x)+(y*y)+(z*z) == c {
					// Menemukan solusi
					fmt.Printf("%d %d %d\n", x, y, z)
					return
				}
			}
		}
	}

	// Jika tidak ada solusi yang ditemukan
	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}

package main

import (
	"fmt"
	"math"
)

func difference(arr [][]int, n int) int {
	var d1,d2 int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				d1 += arr[i][j]
			}
			
			if i == n-j-1 {
				d2 += arr[i][j]
			}
		}
	}

	return int(math.Abs(float64(d1 - d2)))
}

func main() {
	jumlaharray := 3

	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	fmt.Println(difference(data, jumlaharray))
}


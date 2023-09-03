package main

import "fmt"

func fibonacci(number int) int {
	var memotop = map [int]int{}
	nilai, isCount := memotop[number]
		if isCount {
			return nilai
		}
		if number <= 1 {
			memotop[number]=number
		}else {
			memotop[number]=fibonacci(number-1) + fibonacci(number-2)
		}
    return memotop[number]
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
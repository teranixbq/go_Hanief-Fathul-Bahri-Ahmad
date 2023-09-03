package main

import "fmt"

func primeX(number int) int {
	Count := 0
	num := 2
	if number <= 0 {
		return -1 
	}

	isPrime := func(num int) bool {
		if num <= 1 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	for {
		if isPrime(num) {
			Count++
			if Count == number {
				return num
			}
		}
		num++
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
package main

import "fmt" 

func MaximumBuyProduct(money int, productPrice []int) int {
    var result int

    for _, product := range productPrice {
        if money >= product && money%product == 0 {
			// Setiap produk yang bisa dibeli akan menambah nilai result
            result++
        }
    }
    return result
}

func main() {
    fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))      // 3
    fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) // 4
    fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))   // 4
    fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}))           // 1
    fmt.Println(MaximumBuyProduct(0, []int{10000, 30000}))                        // 0
}

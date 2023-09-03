package main

import "fmt"

func MaxSequence(arr []int) int {
    if len(arr) == 0 {
        return 0
    }

    maxtmp := arr[0]
    maxakhir := arr[0]

    for i := 1; i < len(arr); i++ {
        if arr[i] > maxtmp+arr[i] {
            maxtmp = arr[i]
        } else {
            maxtmp += arr[i]
        }

        if maxtmp > maxakhir {
            maxakhir = maxtmp
        }
    }

    return maxakhir
}

func main() {
    fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
    fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
    fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
    fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
    fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}

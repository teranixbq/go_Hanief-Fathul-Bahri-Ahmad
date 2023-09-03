package main

import (
    "fmt"
    "strconv"
)

func binary(n int) []string {
    ans := make([]string, n+1)
    for i := 0; i <= n; i++ {
        ans[i] = strconv.FormatInt(int64(i), 2)
    }

    return ans
}

func main() {
    fmt.Println(binary(2))
    fmt.Println(binary(3))
}

package main

import (
    "fmt"
)

func pascal(numRows int) [][]int {
    data := make([][]int, numRows)

    for i := 0; i < numRows; i++ {
  
        row := make([]int, i+1)
        row[0], row[i] = 1, 1
        
        for j := 1; j < i; j++ {
            row[j] = data[i-1][j-1] + data[i-1][j]
        }

        data[i] = row
        fmt.Print(row)
		
    }

    return data
}

func main() {
   
    pascal(5)
	fmt.Println()
}

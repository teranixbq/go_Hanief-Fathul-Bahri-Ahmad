package main

import (
    "fmt"
    "time"
)

func Kelipatan(x,n int) {
    for i := 1; i<= n ; i++ {
        if i%x == 0 {
            fmt.Printf("Kelipatan %d: %d\n", x, i)
			time.Sleep(3 * time.Second) // menerapkan interval 3 detik
        }
    }
    
}

func printwait(){
	for i := 1;  ; i++ {
        fmt.Println("Loading...")
		time.Sleep(time.Second)
    }
}
func main() {
    x := 5
    jangkauan := 20
	
    go Kelipatan(x, jangkauan)
	go printwait()

	time.Sleep(10*time.Second)
	fmt.Println("selesai")
  
    
}

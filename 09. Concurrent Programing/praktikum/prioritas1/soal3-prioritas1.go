package main

import (
	"fmt"
	"time"
)

func cetakKelipatan3(ch chan int, jumlah int) {
	defer close(ch) // Menutup channel setelah selesai mengirim data
	for i := 1; i <= jumlah; i++ {
		if i%3 == 0 {
			time.Sleep(time.Second)
			ch <- i
		}
	}
}

func main() {
	bufferSize := 10 // Ukuran buffer channel yang gunakan
	ch := make(chan int, bufferSize)

	// Memanggil goroutine untuk mencetak kelipatan 3
	go cetakKelipatan3(ch, 50)

	// Menerima data dari channel
	for kelipatan := range ch {
		fmt.Printf("Kelipatan 3: %d\n", kelipatan)
	}
	fmt.Println("selesai")
}

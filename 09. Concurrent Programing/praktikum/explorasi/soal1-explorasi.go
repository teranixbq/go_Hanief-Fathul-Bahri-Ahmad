package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	// URL yang ingin Anda ambil data JSON-nya
	url := "https://fakestoreapi.com/products"

	// Membuat channel untuk mengirim hasil penguraian
	productChannel := make(chan Product)

	// Goroutine untuk menguraikan respons JSON
	go fetchProducts(url, productChannel)

	// Goroutine untuk mencetak produk dari channel
	go printProducts(productChannel)

    time.Sleep(5*time.Second)
}

func fetchProducts(url string, ch chan<- Product) {
	// Melakukan permintaan GET ke URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Gagal melakukan GET:", err)
		return
	}
	defer response.Body.Close()

	// Membaca body respons ke dalam variabel data
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Gagal membaca respons:", err)
		return
	}

	// respons JSON
	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("Gagal mendapatkan respons JSON:", err)
		return
	}

	// Mengirim produk ke channel 
	for _, product := range products {
		ch <- product
	}
	close(ch)
}

func printProducts(ch <-chan Product) {
	fmt.Println("         Produk Data         ")
	for product := range ch {
		fmt.Println("=============================")
		fmt.Println("Title      : ", product.Title)
		fmt.Println("Price      : ", product.Price)
		fmt.Println("Category   : ", product.Category)
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Soal-1 : ")
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
	fmt.Println("")

	fmt.Println("Soal-2 : ")
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))
	fmt.Println("")

	fmt.Println("Soal-3 : ")
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]

}

// Soal-1 Prioritas 1
func ArrayMerge(a, b []string) []string {
	merge := append([]string{}, a...) 

	for _, arrayB := range b {
		duplicate := false
		for _, dataMerge := range merge {
			if dataMerge == arrayB {
				duplicate = true
				break
			}
		}
		if !duplicate {
			merge = append(merge, arrayB)
		}
	}

	return merge
}



// Soal-2 Prioritas 1
func Mapping(a []string) map[string]int {
    data := make(map[string]int)

    for _, value := range a {
        data[value]++
    }

    return data
}

// Soal-3 Prioritas 1
func munculSekali(angka string) []int {
	datamap := make(map[string]int)

	// Menghitung frekuensi setiap karakter
	for _, dataangka := range angka {
		datamap[string(dataangka)]++
	}

	result := []int{}
	for _, dataangka := range angka {
		key := string(dataangka)
		count := datamap[key]
		if count == 1 {
			num, _ := strconv.Atoi(key)
			result = append(result, num)
		}
	}

	return result
}
	


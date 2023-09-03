package main

import (
	"fmt"
)

func convromawi(number int) string {

	if number > 3999 {
		return "melebihi batas"
	}

	romawilist:= []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	convert := ""
	for _, romwainum := range romawilist {
		for number >= romwainum.value{
			convert += romwainum.digit
			number -= romwainum.value
		}
	}

	return convert
}

func main() {
	fmt.Println("1984	: ",convromawi(1984)) // Output: MCMLXXXIV
	fmt.Println("3999	: ",convromawi(3999)) // Output: MMMCMXCIX
	fmt.Println("512	: ",convromawi(512))   // Output: XLIX
}

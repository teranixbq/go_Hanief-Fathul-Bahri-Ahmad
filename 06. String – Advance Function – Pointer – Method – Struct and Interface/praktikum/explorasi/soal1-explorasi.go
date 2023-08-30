package main

import "fmt"

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string

	for _, char := range s.name {
		
		nameEncode += string(122 - (char - 97))
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode string

	for _, char := range s.name {
		nameDecode += string(122 - (char - 97))
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Println("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Println("\nDecode of student's name " + a.name + " is : " + c.Decode())
	}
}

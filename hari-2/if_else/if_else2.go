package main

import "fmt"

func main() {

	//operator Logical
	var umur uint = 32
	var umurA uint = 24
	var umurB uint = 24

	if umur == 10 && umurB == umurA {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}

	if umur == 10 || umurB == umurA {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}
}

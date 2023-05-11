package main

import "fmt"

func main() {
	//Condition statemant

	if true {
		fmt.Println("Kondisi", true)
	} else {
		fmt.Println("Kondisi", false)
	}

	if false {
		fmt.Println("Kondisi", true)
	} else {
		fmt.Println("Kondisi", false)
	}

	var umur uint = 32

	if umur == 10 {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}

	if umur > 10 {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}

	var umurLuthfi uint = 23
	if umurLuthfi < umur {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}

	var umurA uint = 24
	var umurB uint = 24

	if umurLuthfi >= umurB {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}

	if umurA >= umurB {
		fmt.Println("Kondisi Benar")
	} else {
		fmt.Println("Kondisi Salah")
	}

}

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

	var uangSaku = 2000000

	if uangSaku >= 13000 {
		fmt.Println("Uang Sakunya Banyak")
	} else if uangSaku >= 10000 {
		fmt.Println("Uang Saku Cukup")
	} else if uangSaku >= 5000 {
		fmt.Println("Kondisi Salah")
	} else {
		fmt.Println("uang Saku sedikit")
	}

	yen := uangSaku / 1000
	credCard := true
	payPal := true

	if yen >= 100 && credCard {
		fmt.Println("Bisa Belanja Online")
	} else if yen >= 100 && payPal {
		fmt.Println("Bisa Belanja Online")
	} else {
		fmt.Println("Tidak bisa belanja online")
	}

	if yen > 700 {
		fmt.Println("Uang Yen Anda", yen)
	}

	if usd := uangSaku / 15000; usd >= 100 {
		fmt.Println("Bisa beli ps dengan USD : ", usd)
	} else if usd < 100 {
		fmt.Println("Cukup untuk membuat membeli")
	} else {
		fmt.Println("Tidak bisa beli ps dengan USD : ", usd)
	}

}

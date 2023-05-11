package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Helo World")

	//string
	var address string = "Jalan Durian"
	var boolean bool = true
	var suhuChelcius int8 = -2
	var umur uint8 = 85

	fmt.Println("Alamat : ", address)
	fmt.Println("Benar : ", boolean)
	fmt.Println("Suhu : ", suhuChelcius)
	fmt.Println("Unur : ", umur)

	var uint32 = 8999

	cetak := fmt.Sprintln(uint32)

	fmt.Println(cetak)

	var angkaPertama int8 = 45
	var angkaKedua int8 = 25

	var jumlah int8 = angkaPertama + angkaKedua
	fmt.Println("Jumlah", jumlah)
	var kurang int8 = angkaPertama - angkaKedua
	fmt.Println("Kurang", kurang)
	var bagi int8 = angkaPertama / angkaKedua
	fmt.Println("Bagi", bagi)
	var kali int8 = angkaPertama * angkaKedua
	fmt.Println("Kali", kali)

	var value = []string{}

	var isinya = reflect.TypeOf(value)

	fmt.Println(isinya)

}

package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type cetak interface {
}

func test(test int) {
	fmt.Println("test : ", test)
}

func CekKebenaran(benargak bool) (string, error) {
	var Cek = "Apliakasi Aman"
	if benargak {
		return Cek, nil
	} else {
		return "", errors.New("aplikasi error")
	}
}

func main() {
	test(22222)

	aman, err := CekKebenaran(true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aman)
	}

	fmt.Println()
	fmt.Println(jalanJalan("leleleel"))

	hasil := jumlah(7, 7)

	fmt.Println(hasil)

	nilai := nilaiSiswa("Ahmad", 90)

	fmt.Println(nilai)

	var angkaHari uint32 = 29000

	intToString := strconv.Itoa(int(angkaHari))

	fmt.Println(angkaHari)

	fmt.Println(reflect.TypeOf(intToString))

	hKurang, hBagi := kurangDanBagi(100, 10)

	fmt.Println(hKurang)
	fmt.Println(hBagi)

}

func kurangDanBagi(angka1, angka2 uint32) (uint32, uint32) {
	kr := angka1 - angka2
	bg := angka1 / angka2
	return kr, bg
}

func nilaiSiswa(nama string, nilai uint32) (resp string) {
	stringNilai := strconv.FormatUint(uint64(nilai), 10)
	resp = "Nama : " + nama + " " + stringNilai

	return
}

func jumlah(angka1 int, angka2 int) int {
	var hasil = angka1 + angka2
	return hasil
}

func jalanJalan(test string) (resp string) {
	resp = "jalan jalan" + "di Malaka 4"

	return
}

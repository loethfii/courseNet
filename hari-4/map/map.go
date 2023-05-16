package main

import (
	"fmt"
)

func main() {
	var barang map[int]string

	fmt.Println(barang) // nil

	var nilaiSiswa = map[string]uint32{}

	nilaiSiswa["Akbar"] = 30
	nilaiSiswa["Aji"] = 60
	nilaiSiswa["Masaid"] = 77

	fmt.Println(nilaiSiswa)

	alamat := map[string]string{
		"Juan":   "Cilincing",
		"Deka":   "Jakarta",
		"Firman": "Semarang",
	}

	fmt.Println(alamat)

	for i, v := range alamat {
		fmt.Printf("Nama : %s Alamat saya : %s\n", i, v)
	}

	var def = []map[int][]map[string]string{{1: {{"2": "3", "344": "33121"}, {"233": "33"}}}}

	fmt.Println(def)
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")

	var alamat1 = []map[string]string{
		map[string]string{"Name": "Luthfi", "Alamat": "Jakarta"},
		map[string]string{"Name": "Jaka", "Alamat": "Jogja", "pekerjaan": "Youtuber"},
		map[string]string{"Name": "Jaka", "Alamat": "Jogja", "pekerjaan": "Youtuber", "hobi": "Music"},
	}

	for _, v := range alamat1 {
		fmt.Println("Name : ", v["Name"])
		fmt.Println("Alamat : ", v["Alamat"])
		fmt.Println("pekerjaan : ", v["pekerjaan"])
		fmt.Println("hobi : ", v["hobi"])
		fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||")
	}

	//Map
	//
	x := map[string][]string{
		"Buah":  {"Jambu", "Jeruk", "Mangga"},
		"Mobil": {"Bmw", "Jaguar", "Toyota"},
	}

	for key, val := range x {
		fmt.Printf("janis2 %s2 : ", key)
		for _, v := range val {
			fmt.Printf("%s, ", v)
		}
		fmt.Println()
	}

}

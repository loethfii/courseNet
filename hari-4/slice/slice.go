package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println("Slice")
	//var a = [900]string{"a", "v", "fff", "ppp", "", "", ""}
	//fmt.Println("a : ", a)
	//fmt.Println("reflect.TypeOf(a)", reflect.TypeOf(a))
	//fmt.Println("a[0]", a[0])
	//fmt.Println("a[0]", a[1])
	//
	//fmt.Println("len(a)", len(a))
	//fmt.Println("len(a)", cap(a))
	//
	var foods = []string{"egg", "food", "ZZa"}
	fmt.Println("len(foods)", len(foods))
	fmt.Println("len(foods)", cap(foods))

	var newFoods = foods[0:2]
	fmt.Println("len(newFoods)", len(newFoods))
	fmt.Println("cap(newFoods)", cap(newFoods))

	fmt.Println()

	newFoods = append(newFoods, "Martabak")
	newFoods = append(newFoods, "Martabak")
	newFoods = append(newFoods, "Martabak")

	fmt.Println("len(newFoods)", len(newFoods))
	fmt.Println("cap(newFoods)", cap(newFoods))

	fmt.Println()
	newFoods = append(newFoods, "Martabak")
	fmt.Println("len(newFoods)", len(newFoods))
	fmt.Println("cap(newFoods)", cap(newFoods))

	fmt.Println()
	newFoods = append(newFoods, "Martabak")
	fmt.Println("len(newFoods)", len(newFoods))
	fmt.Println("cap(newFoods)", cap(newFoods))

	fmt.Println()

	var siswa1 = []string{"rahman", "rahman", "rahman", "rahman", "rahman", "rahman", "rahman"}

	fmt.Println("panjang : ", len(siswa1))
	fmt.Println("kapasistas : ", cap(siswa1))

	var siswa2 = siswa1[0:4]

	fmt.Println("panjang : ", len(siswa2))
	fmt.Println("kapasistas : ", cap(siswa2))
	fmt.Println()

	siswa2 = append(siswa2, "Budi")

	fmt.Println("panjang : ", len(siswa2))
	fmt.Println("kapasistas : ", cap(siswa2))
	fmt.Println()

	siswa2 = append(siswa2, "Soekarno")

	fmt.Println("panjang : ", len(siswa2))
	fmt.Println("kapasistas : ", cap(siswa2))
	fmt.Println()

	siswa2 = append(siswa2, "Soehokgie")

	fmt.Println("panjang : ", len(siswa2))
	fmt.Println("kapasistas : ", cap(siswa2))

	fmt.Println()

	siswa2 = append(siswa2, "Soehokgie")

	fmt.Println("panjang : ", len(siswa2))
	fmt.Println("kapasistas : ", cap(siswa2))

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var mahasiswa = []string{"3", "3", "3", "3", "3", "3", "3"}

	var guru []string

	guru = append(guru, mahasiswa...)

	fmt.Println(mahasiswa)
	fmt.Println(guru)

	var angka = []int{3, 4, 5, 5, 4, 3, 3, 3, 3, 3}

	angka1 := angka[0:3]

	angka1 = append(angka1, 2, 4, 5, 6, 7, 8, 5, 5)

	if cap(angka1) > cap(angka) {
		err := errors.New("Kapasitas melebihi batas")
		fmt.Println(err)
	} else {
		fmt.Println("Panjang : ", len(angka1))
		fmt.Println("cap : ", cap(angka1))
	}

}

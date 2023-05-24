package main

import "fmt"

type Makanan struct {
	Buah  Buah
	Sayur Sayur
	Kue   string
}

type Jus struct {
	Buah   []Buah
	Toping map[int]string
}

type Buah struct {
	Nama  string
	Rasa  string
	Harga string
}

type Sayur struct {
	nama  string
	warna string
	harga string
}

func main() {
	fmt.Println("struct")

	var menu1 Makanan

	menu1.Buah = Buah{
		Nama:  "Melon",
		Rasa:  "Manis",
		Harga: "20000",
	}

	menu1.Sayur = Sayur{
		nama:  "Wortel",
		warna: "Orenge",
		harga: "90.000",
	}
	menu1.Kue = "Nastar"

	fmt.Println(menu1)

	var menu2 = Jus{
		Buah: []Buah{
			{
				Nama:  "Strawbery",
				Rasa:  "Asem",
				Harga: "90.000",
			},
		},
		Toping: map[int]string{1: "BOba", 2: "Ceres"},
	}

	fmt.Println(menu2)
	//
	//var menu2 = Makanan{
	//	Buah: "Berry", Kue: "Bolu",
	//}
	//
	//fmt.Println("Menu 1", menu1)
	//fmt.Println("Menu 2 Buah", menu2.Buah)
	//fmt.Println("Menu 2 Buah", menu2.Kue)

}

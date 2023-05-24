package main

import "fmt"

type Buah struct {
	nama  string
	harga string
}

type OlahanBuah struct {
	pembuat  string
	nama     string
	namaBuah string
}

func (b Buah) hargaBuah() {
	fmt.Println("Harga Buah : ", b.harga)
}

func (b Buah) beliBuah(pembeli string) (resp string) {
	resp = "Buah " + b.nama + " Dibeli Oleh " + pembeli + " dengan Harga " + b.harga
	return resp
}

func membuatJuice(pembuat string, buah Buah) (resp string) {
	resp = "Buah " + buah.nama + " Dibuat Oleh " + pembuat
	return resp
}

func mengelolaBuah(buah Buah, pembuat, namaOlahan string) (resp OlahanBuah) {
	resp.nama = namaOlahan
	resp.pembuat = pembuat
	resp.namaBuah = buah.nama
	return resp
}

func (b Buah) mengelolaBuah1(pembuat, namaOlahan string) (resp OlahanBuah) {
	resp.nama = namaOlahan
	resp.pembuat = pembuat
	resp.namaBuah = b.nama
	return resp
}

func (b Buah) test(buah Buah) (resp string) {
	fmt.Println("b ", b.nama)
	fmt.Println("buah ", buah.nama)
	return "test"
}

func main() {
	fmt.Println("Method 1")

	melon := Buah{
		nama:  "Melon",
		harga: "150000",
	}

	fmt.Println(melon)

	melon.hargaBuah()

	pembeli := melon.beliBuah("Luthfi")

	fmt.Println(pembeli)

	res := mengelolaBuah(melon, "Luthfi", "rujak")

	fmt.Println(res)

	asinan := melon.mengelolaBuah1("Jawa", "Asinan")

	fmt.Println(asinan)

	test := melon.test(melon)

	fmt.Println(test)

}

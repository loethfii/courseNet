package main

import "fmt"

type Beli struct {
	Barang      Barang
	TanggalBeli string
	AlamatKirim string
}
type Barang struct {
	Name   string
	Jumlah string
}

func (b Beli) Detail() (detail string) {
	detail = fmt.Sprintf("Kamu Beli Barang %s Jumlahnya %s Tanggal Beli %s Alamat kirim %s", b.Barang, b.Barang.Name, b.TanggalBeli, b.AlamatKirim)
	return detail
}

func main() {
	PS4 := Barang{Name: "PS4", Jumlah: "90"}
	Orang1 := Beli{
		Barang: Barang{Name: PS4.Name},
	}

	res := Orang1.Detail()

	fmt.Println(res)
}

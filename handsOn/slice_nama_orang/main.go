package main

import (
	"cari_nama/model"
	"cari_nama/repository"
	"fmt"
)

type newModelName struct {
	model.NameMuhammad
}

func (nm newModelName) cetakName() {
	fmt.Println(nm.Name)
}

func terimaData(nameChannel []model.NameMuhammad, name chan<- string) {
	for _, v := range nameChannel {
		fmt.Println(v)
	}
}

func main() {

	var nameChannel = make(chan string)

	cekData := repository.CariMuhammad(model.DataName())

	terimaData(cekData, nameChannel)

	fmt.Println(cekData)

	data := newModelName{model.NameMuhammad{"Bagus"}}

	data.cetakName()
}

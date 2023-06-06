package main

import (
	"fmt"
	"runtime"
	"sync"
)

// restoran - AGung gsa
// ikan -> cuci -> simpan frozen
// daging -> potong -> simpan frozen
// sayur ->kupas -> simpan dibox->simpan di chiller
//buah -> cuci -> simpan chiller
//bawang -> kupas-> simpan di box -> simpan chiller
//selain itu buang ke tong sampah

// metode
var cuciChannel = make(chan string)
var potongCHannel = make(chan string)
var kupasChannel = make(chan string)

// simpan
var simpanFrozenChannel = make(chan string)
var SimpanChillerChannel = make(chan string)
var SimpanBoxChannel = make(chan string)

// waitgrup
var wg sync.WaitGroup

func cekBarang(belanja []string) {
	defer wg.Done()
	for _, barang := range belanja {
		switch barang {
		case "ikan":
			fmt.Println("Menemukan", barang)
			cuciChannel <- barang
			close(cuciChannel)
		case "daging":
			fmt.Println("Menemukan", barang)
			potongCHannel <- barang
			close(potongCHannel)
		case "sayur":
			fmt.Println("Menemukan ", barang)
			kupasChannel <- barang
			close(kupasChannel)
		case "buah":
			fmt.Println("Menemukan ", barang)
			cuciChannel <- barang
			close(cuciChannel)
		case "bawang":
			fmt.Println("Menemukan ", barang)
			kupasChannel <- barang
			close(kupasChannel)
		default:
			fmt.Println("Dibuang aja", barang, "nya")
		}
	}

}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Go routine")
	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bambu", "buah", "sterofoam", "bawang",
		"ikan", "ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang",
		"buah", "sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang",
		"buah", "sterofoam", "bawang",
	}

	//method

	wg.Add(7)
	go cekBarang(belanja)
	go cuci()
	go simpanFrozen()
	go potong()
	//simpan
	go kupas()
	go simpanBox()
	go simpanChiller()

	wg.Wait()

	fmt.Println("Selesai")
	fmt.Println("========================================================")
}

// methode
func cuci() {
	defer wg.Done()
	for itemCuci := range cuciChannel {
		fmt.Println("Dicuci dulu")
		if itemCuci == "ikan" {
			simpanFrozenChannel <- itemCuci
		} else if itemCuci == "buah" {
			SimpanChillerChannel <- itemCuci
		}
	}
	close(simpanFrozenChannel)
	close(SimpanChillerChannel)
}

func potong() {
	defer wg.Done()
	for itemPotong := range potongCHannel {
		fmt.Println(itemPotong, " dipotong dulu")
		if itemPotong == "daging" {
			simpanFrozenChannel <- itemPotong
		}
	}
	close(simpanFrozenChannel)
}

func kupas() {
	defer wg.Done()
	for itemKupas := range kupasChannel {
		fmt.Println(itemKupas, "dikupas dulu")
		if itemKupas == "sayur" {
			SimpanBoxChannel <- itemKupas
		}
	}
	close(SimpanBoxChannel)
}

// simpan
func simpanFrozen() {
	defer wg.Done()
	for itemSimpanFrozen := range simpanFrozenChannel {
		fmt.Println(itemSimpanFrozen, "Simpan Frozen")
	}
}

func simpanBox() {
	defer wg.Done()
	for itemSimpanBox := range SimpanBoxChannel {
		fmt.Println(itemSimpanBox, "Simpan box")
		if itemSimpanBox == "sayur" {
			SimpanChillerChannel <- itemSimpanBox
		}
	}

	close(SimpanChillerChannel)
}

func simpanChiller() {
	defer wg.Done()
	for itemSimpanChiller := range SimpanChillerChannel {
		fmt.Println(itemSimpanChiller, "Simpan chiller")
	}
}

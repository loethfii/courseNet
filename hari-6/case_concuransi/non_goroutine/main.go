package main

import (
	"fmt"
)

func main() {
	fmt.Println("Non Go routine")

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bambu", "buah", "sterofoam", "bawang", "ikan",
		"ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang", "buah",
		"sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang", "buah", "sterofoam", "bawang",
	}

	for _, barang := range belanja {
		switch barang {
		case "ikan":
			fmt.Println("Menemukan", barang)
			cuci(barang)
			//ikan -> cuci ->simpan di frozen
		case "daging":
			fmt.Println("menemukan", barang)
			potong(barang)
			//daging -> potong -> simpan di frozen
		case "buah":
			fmt.Println("Menemukan", barang)
			cuci(barang)
			//ikan -> cuci -> simpan di chiler
		case "sayur":
			fmt.Println("Menemukan", barang)
			kupas(barang)
			//sayur ->kupas ->simpan di box ->simpan di chiler
		case "bawang":
			fmt.Println("Menemukan", barang)
			kupas(barang)
			//bawang -> kupas -> simpan di box -> simpan di chiller
		default:
			fmt.Println("Dibuang aja")
		}
	}

}

// metode
func cuci(barang string) {
	fmt.Println("Mencuci", barang)
	if barang == "ikan" {
		simpanDiFrozen(barang)
	} else if barang == "buah" {
		simpanDiChiller(barang)
	}
}

func potong(barang string) {
	fmt.Println("Memotong", barang)
	if barang == "daging" {
		simpanDiFrozen(barang)
	}
}

func kupas(barang string) {
	fmt.Println("Mengupas", barang)
	if barang == "sayur" {
		simpanDiBox(barang)
		simpanDiChiller(barang)
	} else if barang == "bawang" {
		simpanDiBox(barang)
		simpanDiChiller(barang)
	}
}

// simpan
func simpanDiFrozen(barang string) {
	fmt.Println(barang, "Simpan di frozen")
}

func simpanDiChiller(barang string) {
	fmt.Println(barang, "Simpan di chiller")
}

func simpanDiBox(barang string) {
	fmt.Println(barang, "simpan di box")
}

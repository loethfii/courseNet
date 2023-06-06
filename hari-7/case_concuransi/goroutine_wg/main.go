package main

import (
	"fmt"
	"sync"
)

func menmukanBarang(belanja []string, wg *sync.WaitGroup) {
	for _, barang := range belanja {
		switch barang {
		case "ikan":
			fmt.Println("Menemukan", barang)
			wg.Add(1)
			go cuci(barang, wg)
			
			//ikan -> cuci ->simpan di frozen
		case "daging":
			fmt.Println("menemukan", barang)
			wg.Add(1)
			go potong(barang, wg)
			
			//daging -> potong -> simpan di frozen
		case "buah":
			fmt.Println("Menemukan", barang)
			wg.Add(1)
			go cuci(barang, wg)
			
			//ikan -> cuci -> simpan di chiler
		case "sayur":
			fmt.Println("Menemukan", barang)
			wg.Add(1)
			go kupas(barang, wg)
			
			//sayur ->kupas ->simpan di box ->simpan di chiler
		case "bawang":
			fmt.Println("Menemukan", barang)
			wg.Add(1)
			go kupas(barang, wg)
			
			//bawang -> kupas -> simpan di box -> simpan di chiller
		default:
			fmt.Println("Dibuang aja")
		}
	}
	wg.Wait()
}

func main() {
	fmt.Println("Goroutine Waitgrub")
	
	var wg sync.WaitGroup
	
	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bambu", "buah", "sterofoam", "bawang", "ikan",
		"ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang", "buah",
		"sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang", "buah", "sterofoam", "bawang",
	}
	
	menmukanBarang(belanja, &wg)
	
}

// metode
func cuci(barang string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mencuci", barang)
	if barang == "ikan" {
		simpanDiFrozen(barang)
	} else if barang == "buah" {
		simpanDiChiller(barang)
	}
}

func potong(barang string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memotong", barang)
	if barang == "daging" {
		simpanDiFrozen(barang)
		
	}
	
}

func kupas(barang string, wg *sync.WaitGroup) {
	defer wg.Done()
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

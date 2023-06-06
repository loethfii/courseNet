package main

import (
	"fmt"
	"sync"
	"test_ch/repository"
)

func main() {
	var wg sync.WaitGroup

	var daftarBelanja = make(chan string)

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bambu", "buah", "sterofoam", "bawang",
		"ikan", "ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang",
		"buah", "sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang",
		"buah", "sterofoam", "bawang",
	}

	wg.Add(8)
	go repository.TerimaBarang(belanja, daftarBelanja, &wg)
	go repository.KirimKeFungsiLain(daftarBelanja)
	go repository.Cuci()
	go repository.SimpanFrozen()
	go repository.Potong()

	go repository.Kupas()
	go repository.SimpanBox()
	go repository.SimpanChiller()

	wg.Wait()

	fmt.Println("selesai")
}

package repository

import (
	"fmt"
	"sync"
)

func TerimaBarang(belanja []string, daftarBelanja chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, barang := range belanja {
		daftarBelanja <- barang
	}
	close(daftarBelanja)
}

func KirimKeFungsiLain(daftarBelanja <-chan string) {
	defer wg.Done()
	for barang := range daftarBelanja {
		switch barang {
		case "ikan":
			fmt.Println("menemukan ikan")
			cuciChannel <- barang
		case "daging":
			potongCHannel <- barang
		case "sayur":
			kupasChannel <- barang
		case "buah":
			cuciChannel <- barang
		case "bawang":
			kupasChannel <- barang
		default:
			fmt.Println("error")
		}
	}
	close(cuciChannel)
	close(potongCHannel)
	close(potongCHannel)
	close(kupasChannel)
	close(cuciChannel)
	close(kupasChannel)
}

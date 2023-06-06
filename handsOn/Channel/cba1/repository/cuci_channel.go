package repository

import "fmt"

func Cuci() {
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

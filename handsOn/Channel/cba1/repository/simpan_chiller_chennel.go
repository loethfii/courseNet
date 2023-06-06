package repository

import "fmt"

func SimpanChiller() {
	defer wg.Done()
	for itemSimpanChiller := range SimpanChillerChannel {
		fmt.Println(itemSimpanChiller, "Simpan chiller")
	}
}

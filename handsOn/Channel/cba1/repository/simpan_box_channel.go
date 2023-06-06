package repository

import "fmt"

func SimpanBox() {
	defer wg.Done()
	for itemSimpanBox := range SimpanBoxChannel {
		fmt.Println(itemSimpanBox, "Simpan box")
		if itemSimpanBox == "sayur" {
			SimpanChillerChannel <- itemSimpanBox
		}
	}

	close(SimpanChillerChannel)
}

package repository

import "fmt"

func Kupas() {
	defer wg.Done()
	for itemKupas := range kupasChannel {
		fmt.Println(itemKupas, "dikupas dulu")
		if itemKupas == "sayur" {
			SimpanBoxChannel <- itemKupas
		}
	}
	close(SimpanBoxChannel)
}

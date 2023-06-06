package repository

import "fmt"

func Potong() {
	defer wg.Done()
	for itemPotong := range potongCHannel {
		fmt.Println(itemPotong, " dipotong dulu")
		if itemPotong == "daging" {
			simpanFrozenChannel <- itemPotong
		}
	}
	close(simpanFrozenChannel)
}

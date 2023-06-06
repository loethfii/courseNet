package repository

import "fmt"

func SimpanFrozen() {
	defer wg.Done()
	for itemSimpanFrozen := range simpanFrozenChannel {
		fmt.Println(itemSimpanFrozen, "Simpan Frozen")
	}
}

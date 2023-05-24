package handler

import (
	"fmt"
	"handsOn/repository"
)

//OutPut to main

func DetailBarang(d repository.Item) {
	value := d.GetDetail()
	fmt.Println(value)
}

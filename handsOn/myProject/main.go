package main

import (
	"handsOn/handler"
	"handsOn/repository"
)

func main() {
	stock := repository.NewStockBarang{}
	handler.DetailBarang(&stock)
}

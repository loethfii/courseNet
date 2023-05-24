package repository

import (
	"handsOn/model"
)

type Item interface {
	GetDetail() []model.StockBarang
}

type NewStockBarang struct {
	model.StockBarang
}

func (s *NewStockBarang) GetDetail() []model.StockBarang {
	valueBarang := model.DataBarang

	return valueBarang
}

func (s *NewStockBarang) InputBarang() {

}


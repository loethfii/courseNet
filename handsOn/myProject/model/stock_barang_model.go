package model

type StockBarang struct {
	Merk        string
	JenisBarang string
	Harga       float64
	Qty         uint32
}

var DataBarang = []StockBarang{
	{
		Merk:        "Nuvo",
		JenisBarang: "Sabun",
		Harga:       90.0000,
		Qty:         90,
	},
}

package model

type Barang struct {
	ID           int
	KodeProduct  string
	NamaBarang   string
	JumlahBarang uint32
	Kategori     Kategori
}

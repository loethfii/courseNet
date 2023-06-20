package models

import "time"

type Mobil struct {
	Id           int        `json:"id"`
	NamaMobil    string     `json:"nama_mobil"`
	TipeMobil    string     `json:"tipe_mobil"`
	JenisMobil   string     `json:"jenis_mobil"`
	BahanBakar   string     `json:"bahan_bakar"`
	Isi_Silinder string     `json:"isi_Silinder"`
	Warna        string     `json:"warna"`
	Transmisi    string     `json:"transmisi"`
	Harga        uint32     `json:"harga"`
	Qty          uint32     `json:"qty"`
	TimeStamp    *time.Time `json:"time_stamp"`
}

type ArrMobil = []Mobil

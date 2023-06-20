package models

import "time"

type FormPembelian struct {
	Id                 int        `json:"id"`
	NamaLengkapPembeli string     `json:"nama_lengkap_pembeli"`
	NomerKTP           string     `json:"nomer_ktp"`
	AlamatRumah        string     `json:"alamat_rumah"`
	NomerDebit         string     `json:"nomer_debit"`
	Mobil              Mobil      `json:"id_mobil"`
	HarusInden         bool       `json:"harus_inden"`
	LamanInden         string     `json:"laman_inden"`
	CustomPlat         string     `json:"custom_plat"`
	TambahanKit        string     `json:"tambahan_kit"`
	Sales              Sales      `json:"id_sales"`
	TimeStamp          *time.Time `json:"time_stamp"`
}

type ArrFormPembelian = []FormPembelian

type NewFormPembelian struct {
	NamaLengkapPembeli string `json:"nama_lengkap_pembeli"`
	NomerKTP           string `json:"nomer_ktp"`
	AlamatRumah        string `json:"alamat_rumah"`
	NomerDebit         string `json:"nomer_debit"`
	IdMobil            int    `json:"id_mobil"`
	HarusInden         bool   `json:"harus_inden"`
	LamanInden         string `json:"lama_inden"`
	CustomPlat         bool   `json:"custom_plat"`
	TambahanKit        string `json:"tambahan_kit"`
	IdSales            int    `json:"id_sales"`
}

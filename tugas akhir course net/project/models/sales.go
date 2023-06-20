package models

import "time"

type Sales struct {
	Id          int        `json:"id"`
	NamaSales   string     `json:"nama_sales"`
	Nip         string     `json:"nip"`
	NomerTelpon string     `json:"nomer_telpon"`
	Bagian      string     `json:"bagian"`
	TimeStamp   *time.Time `json:"time_stamp"`
}

type ArrSales []Sales

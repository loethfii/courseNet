package models

import "time"

type Pembayaran struct {
	Id            int           `json:"id"`
	FormPembelian FormPembelian `json:"form_pembelian"`
	BuktiTransfer string        `json:"bukti_transfer"`
	IsConfirm     bool          `json:"is_confirm"`
	CreatedAt     *time.Time    `json:"created_at"`
}

type ArrPembayaran = []Pembayaran

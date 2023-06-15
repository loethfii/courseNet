package models

type Employee struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Alamat string `json:"alamat"`
	Posisi string `json:"posisi"`
}

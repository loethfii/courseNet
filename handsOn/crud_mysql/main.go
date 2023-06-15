package main

import (
	"crud_mysql/repository"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var pilihYangMana int

	fmt.Println("Daftar fungsi : ")
	fmt.Println("1 . Tambah Barang ")
	fmt.Println("2 . Tambah Kategori ")
	fmt.Println("3 . Lihat Data Barang ")
	fmt.Println("4 . Lihat Data Kategori ")
	fmt.Println("5 . Update Barang ")
	fmt.Println("6 . Hapus Barang ")
	fmt.Printf("Pilih fungsi yang ingin digunakan : ")
	fmt.Scanln(&pilihYangMana)

	switch pilihYangMana {
	case 1:
		repository.AddBarang()
	case 2:
		repository.AddKategori()
	case 3:
		repository.ShowAllBarang()
	case 4:
		repository.ShowAllCategories()
	case 5:
		repository.UpdateBarang()
	case 6:
		repository.DeleteOneData()

	}
}

package repository

import (
	"bufio"
	"crud_mysql/db_connect"
	"crud_mysql/model"
	"fmt"
	"log"
	"os"
)

func AddBarang() {
	db, err := db_connect.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Masukkan Kode Produk : ")
	kodeProduk, _ := reader.ReadString('\n')

	fmt.Printf("Masukkan Nama Barang : ")
	nama_barang, _ := reader.ReadString('\n')

	fmt.Printf("Masukkan Jumlah Stock : ")
	jumlahStock, _ := reader.ReadString('\n')

	//kategori yang tersedia
	rows, err := db.Query("select * from kategori")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var response []model.Kategori

	for rows.Next() {
		var data = model.Kategori{}

		var err = rows.Scan(&data.ID, &data.NamaKategori)
		if err != nil {
			log.Println(err.Error())
			return
		}

		response = append(response, data)
	}
	for _, each := range response {
		fmt.Println(each.ID, each.NamaKategori)
	}

	fmt.Printf("Masukkan Ketegori : ")
	ketegoriId, _ := reader.ReadString('\n')

	_, err = db.Exec("insert into barang (kode_produk, nama_barang, jumlah_stok,kategori_id) values (?, ?, ?, ?)", kodeProduk, nama_barang, jumlahStock, ketegoriId)

	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println()

	fmt.Println("Berhasil input barang")
}

func ShowAllBarang() {
	db, err := db_connect.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("select barang.kode_produk, barang.nama_barang, barang.jumlah_stok, kategori.nama from barang inner join kategori on kategori.id = barang.kategori_id order by barang.id desc")

	if err != nil {
		log.Printf(err.Error())
		return
	}

	defer rows.Close()

	var response []model.Barang

	for rows.Next() {
		var data = model.Barang{}
		err := rows.Scan(&data.KodeProduct, &data.NamaBarang, &data.JumlahBarang, &data.Kategori.NamaKategori)
		if err != nil {
			log.Println(err.Error())
			return
		}

		response = append(response, data)
	}

	for _, v := range response {
		fmt.Println("=========================================================================")
		fmt.Println("Kode Produk : ", v.KodeProduct)
		fmt.Println("Nama Barang : ", v.NamaBarang)
		fmt.Println("Jumlah Stok : ", v.JumlahBarang)
		fmt.Println("Kategori Produk : ", v.Kategori.NamaKategori)
		fmt.Println()
	}

	log.Println("Sukses menampilkan data")
}

func UpdateBarang() {
	db, err := db_connect.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from barang")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var barangMap map[int]model.Barang
	barangMap := make(map[int]model.Barang)
	increment := 0

	for rows.Next() {
		data := model.Barang{}
		increment++
		err := rows.Scan(&data.ID, &data.KodeProduct, &data.NamaBarang, &data.JumlahBarang, &data.Kategori.NamaKategori)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		barangMap[increment] = data
	}

	for _, v := range barangMap {
		fmt.Println(v)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Mau id mana yang di ubah : ")
	id, _ := reader.ReadString('\n')

	fmt.Printf("Masukkan Nama Barang :  ")
	namaBarang, _ := reader.ReadString('\n')

	fmt.Printf("Masukkan Jumlah Stock : ")
	jumlahStok, _ := reader.ReadString('\n')

	fmt.Printf("Masukkan kategori id : ")
	kategoriId, _ := reader.ReadString('\n')

	_, err = db.Exec("UPDATE barang set nama_barang= ?,jumlah_stok = ?, kategori_id = ?  WHERE id = ?", namaBarang, jumlahStok, kategoriId, id)

	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("Data Berhasil di update")

}

func DeleteOneData() {
	db, err := db_connect.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Mau id mana yang di hapus : ")
	id, _ := reader.ReadString('\n')

	_, err = db.Exec("delete from barang where id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data berhasil dihapus")
}

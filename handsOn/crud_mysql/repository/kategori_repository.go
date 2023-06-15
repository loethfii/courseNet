package repository

import (
	"bufio"
	"crud_mysql/db_connect"
	"crud_mysql/model"
	"fmt"
	"log"
	"os"
)

func AddKategori() {
	db, err := db_connect.ConnectDB()
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Masukkan Nama Kategori : ")
	nama, _ := reader.ReadString('\n')

	_, err = db.Exec("insert into kategori (nama) values (?)", nama)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Berhasil input ketegori")
}

func ShowAllCategories() {
	db, err := db_connect.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM kategori")

	if err != nil {
		log.Printf(err.Error())
		return
	}

	defer rows.Close()

	var response []model.Kategori

	for rows.Next() {
		var data = model.Kategori{}
		err := rows.Scan(&data.ID, &data.NamaKategori)
		if err != nil {
			log.Println(err.Error())
			return
		}

		response = append(response, data)
	}

	for _, v := range response {
		fmt.Println("=========================================================================")
		fmt.Println("Kategori ID : ", v.ID)
		fmt.Println("Nama Kategori : ", v.NamaKategori)
		fmt.Println()
	}

	log.Println("Sukses menampilkan data")
}

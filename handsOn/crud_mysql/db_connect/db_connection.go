package db_connect

import (
	"database/sql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/kelola_warung")
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	log.Println("Koneksi berhasil")
	return db, err
}

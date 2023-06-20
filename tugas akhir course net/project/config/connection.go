package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dealer_mobil?parseTime=true")
	
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	
	log.Printf("connection success")
	
	return db, nil
}

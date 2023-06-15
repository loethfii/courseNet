package config

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_basic_sql_2")
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	log.Printf("config success")
	
	return db, err
}

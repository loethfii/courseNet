package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type article struct {
	id     uint
	title  string
	writer string
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_basic_sql_2")
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	log.Printf("connection success")
	
	return db, err
}

func main() {
	fmt.Println("Tugas Article")
	
	execInsert()
	query()
	queryRow()
	execDelete()
}

func execDelete() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	id := 1
	
	_, err = db.Exec("DELETE from article where id=?", id)
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println("Delete success")
}

func queryRow() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	defer db.Close()
	
	article := article{}
	id := 2
	err = db.QueryRow("SELECT * FROM article WHERE id = ?", id).Scan(&article.id, &article.title, &article.writer)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	fmt.Println(article.id, article.title, article.writer)
	
}

func query() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	rows, err := db.Query("SELECT * FROM article")
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	defer rows.Close()
	
	var resp []article
	
	for rows.Next() {
		var a = article{}
		var err = rows.Scan(&a.id, &a.title, &a.writer)
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		resp = append(resp, a)
	}
	
	for _, v := range resp {
		fmt.Println("id : ", v.id)
		fmt.Println("title : ", v.title)
		fmt.Println("writer : ", v.writer)
	}
}

func execInsert() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	title := "5 game rpg terbaik"
	writer := "gilang"
	
	_, err = db.Exec("insert into article (title, writer) values (?,?)", title, writer)
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println("Insert sukses")
}

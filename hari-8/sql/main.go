package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Employee struct {
	id     uint   `json:"id"`
	name   string `json:"name"`
	alamat string `json:"alamat"`
	posisi string `json:"posisi"`
}

func connectDB() (*sql.DB, error) {
	//template
	//db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3307)/databasename")
	//php my admin
	//db, err := sql.Open("mysql", "user:@tcp(127.0.0.1:3307)/golang_basic_sql_1")
	
	//workbanch
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_basic_sql_2")
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	log.Printf("connection success")
	
	return db, err
}
func main() {
	fmt.Println("SQL Server")
	//execInsert()
	
	//query()
	//queryRow()
	//prepare()
	//execUpdate()
	//execDelete()
	//testPerbedaaan()
}

func testPerbedaaan() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	var empl = Employee{}
	
	//query row hanya bisa menampilan 1 data
	err = db.QueryRow("SELECT * FROM employees").Scan(&empl.id, &empl.name, &empl.alamat, &empl.posisi)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Println(empl)
	
	//res , err := db.Exec("SELECT * FROM employees")
}

func execDelete() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	id := 1
	
	_, err = db.Exec("DELETE from employees where id=?", id)
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println("Delete success")
}

func execUpdate() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	id := 1
	name := "Muhammmad iqbal"
	alamat := "jakarta"
	posisi := "staff"
	
	_, err = db.Exec("UPDATE employees set name=?, alamat=?, posisi=? where id=?", name, alamat, posisi, id)
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println("Update sukses")
}

func prepare() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	defer db.Close()
	
	q, err := db.Prepare("SELECT * FROM employees WHERE name = ?")
	
	if err != nil {
		log.Println(err.Error())
		return
	}
	
	emplyee1 := Employee{}
	//emplyee2 := Employee{}
	
	q.QueryRow(1).Scan(&emplyee1.id, &emplyee1.name, &emplyee1.alamat, &emplyee1.posisi)
	
	fmt.Println(emplyee1.id, emplyee1.name, emplyee1.alamat, emplyee1.posisi)
}

func queryRow() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	defer db.Close()
	
	emp := Employee{}
	name := "luthfi"
	err = db.QueryRow("SELECT * FROM employees WHERE name = ?", name).Scan(&emp.id, &emp.name, &emp.alamat, &emp.posisi)
	if err != nil {
		fmt.Println("err")
		return
	}
	
	fmt.Println(emp)
	
}

func query() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	rows, err := db.Query("SELECT * FROM employees")
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	defer rows.Close()
	
	var resp []Employee
	
	for rows.Next() {
		var a = Employee{}
		var err = rows.Scan(&a.id, &a.name, &a.alamat, &a.posisi)
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		resp = append(resp, a)
	}
	
	fmt.Println(resp)
	
	for _, v := range resp {
		fmt.Println(v)
	}
}

func execInsert() {
	db, err := connectDB()
	if err != nil {
		log.Printf(err.Error())
	}
	
	defer db.Close()
	
	name := "luthfi"
	alamat := "jakarta"
	posisi := "staff"
	
	_, err = db.Exec("insert into employees (name, alamat, posisi) values (?,?,?)", name, alamat, posisi)
	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println("Insert sukses")
}

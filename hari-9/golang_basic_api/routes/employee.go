package routes

import (
	"encoding/json"
	"fmt"
	"golang_basic_api/config"
	"golang_basic_api/models"
	"log"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get /employee")
	
	fmt.Fprintln(w, "Halaman Home")
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDB()
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
	
	var resp []models.Employee
	
	for rows.Next() {
		var a = models.Employee{}
		var err = rows.Scan(&a.Id, &a.Name, &a.Alamat, &a.Posisi)
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		resp = append(resp, a)
	}
	
	json.NewEncoder(w).Encode(resp)
	
}

func PostEmployees(w http.ResponseWriter, r *http.Request) {
	log.Printf("Post /employees/post")
	
	if r.Method == "POST" {
		
		db, err := config.ConnectDB()
		if err != nil {
			log.Printf(err.Error())
		}
		
		defer db.Close()
		
		var employee = models.Employee{}
		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		_, err = db.Exec("insert into employees (name, alamat, posisi) values (?,?,?)", employee.Name, employee.Alamat, employee.Posisi)
		
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		log.Printf("Insert Sukses")
		
		json.NewEncoder(w).Encode(employee)
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya post", http.StatusMethodNotAllowed)
		return
	}
}

func PutEmployees(w http.ResponseWriter, r *http.Request) {
	log.Printf("Post /employees/post")
	
	if r.Method == "PUT" {
		
		db, err := config.ConnectDB()
		if err != nil {
			log.Printf(err.Error())
		}
		
		defer db.Close()
		
		var employee = models.Employee{}
		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		_, err = db.Exec("UPDATE employees set name=?, alamat=?, posisi=? where id=?", employee.Name, employee.Alamat, employee.Posisi, employee.Id)
		
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		log.Printf("Update Sukses")
		
		json.NewEncoder(w).Encode(employee)
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya post", http.StatusMethodNotAllowed)
		return
	}
}

func DeleteEmployees(w http.ResponseWriter, r *http.Request) {
	log.Printf("Post /employees/post")
	
	if r.Method == "DELETE" {
		
		db, err := config.ConnectDB()
		if err != nil {
			log.Printf(err.Error())
		}
		
		defer db.Close()
		
		var employee = models.Employee{}
		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		_, err = db.Exec("DELETE from employees where id=?", employee.Id)
		
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		log.Printf("Delete Sukses")
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya DELETE", http.StatusMethodNotAllowed)
		return
	}
}

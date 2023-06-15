package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang_basic_api/routes"
	"log"
	"net/http"
)

func main() {
	log.Printf("Golang Basic API")
	//
	http.HandleFunc("/", routes.GetHome)
	http.HandleFunc("/employees", routes.GetEmployees)
	http.HandleFunc("/employees/submit", routes.PostEmployees)
	http.HandleFunc("/employees/update", routes.PutEmployees)
	http.HandleFunc("/employees/delete", routes.DeleteEmployees)
	
	http.HandleFunc("/article", routes.GetArticle)
	http.HandleFunc("/article/submit", routes.PostArticle)
	http.HandleFunc("/article/put", routes.PutArticle)
	http.HandleFunc("/article/delete", routes.DeleteArticle)
	
	log.Printf("localhost 3000")
	http.ListenAndServe(":3000", nil)
}

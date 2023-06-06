package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
}

type Articles []Article

type Employee struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Position string `json:"position"`
}

type Emplyees []Employee

func main() {
	fmt.Println("Web server HTTP")
	
	http.HandleFunc("/", getHome)
	http.HandleFunc("/test", getTest)
	http.HandleFunc("/article", getArticle)
	http.HandleFunc("/article/submit", postArticle)
	
	http.HandleFunc("/employees", getEmployee)
	http.HandleFunc("/employees/submit", postEmployee)
	
	log.Printf("Service runing : ")
	
	http.ListenAndServe(":3000", nil)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get /employee")
	employee := Emplyees{
		Employee{Name: "Test Judul", Address: "Test Penulis", Email: "luthfi@yahoo.com", Position: "staff"},
	}
	
	json.NewEncoder(w).Encode(employee)
}

func postEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Post /employees/submit")
	if r.Method == "POST" {
		var employee Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		
		if err != nil {
			log.Printf("Bad Request", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		
		if employee.Name == "" || employee.Position == "" || employee.Email == "" || employee.Address == "" {
			log.Printf("Employee tidak boleh kosong")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		
		json.NewEncoder(w).Encode(employee)
		return
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya post", http.StatusMethodNotAllowed)
		return
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	path := r.Header["Referer"]
	fmt.Println(path)
	w.Write([]byte("Halaman Home"))
}

func getTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman Home"))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get /articles")
	articles := Articles{
		Article{Judul: "Test Judul", Penulis: "Test Penulis"},
	}
	
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Post /article/submit")
	if r.Method == "POST" {
		var article Article
		err := json.NewDecoder(r.Body).Decode(&article)
		
		if err != nil {
			log.Printf("Bad Request", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		
		if article.Judul == "" || article.Penulis == "" {
			log.Printf("Judul atau penulis tidak boleh kosong")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		
		json.NewEncoder(w).Encode(article)
		return
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya post", http.StatusMethodNotAllowed)
		return
	}
}

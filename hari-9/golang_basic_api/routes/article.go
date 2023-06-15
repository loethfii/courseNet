package routes

import (
	"encoding/json"
	"golang_basic_api/config"
	"golang_basic_api/models"
	"log"
	"net/http"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDB()
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
	
	var resp models.Articles
	
	for rows.Next() {
		var a = models.Article{}
		var err = rows.Scan(&a.Id, &a.Title, &a.Writer)
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		resp = append(resp, a)
	}
	
	json.NewEncoder(w).Encode(resp)
	
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Post /article/submit")
	
	if r.Method == "POST" {
		
		db, err := config.ConnectDB()
		if err != nil {
			log.Printf(err.Error())
		}
		
		defer db.Close()
		
		var article = models.Article{}
		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		_, err = db.Exec("insert into article (title, writer) values (?,?)", article.Title, article.Writer)
		
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		log.Printf("Insert Sukses")
		
		json.NewEncoder(w).Encode(article)
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya post", http.StatusMethodNotAllowed)
		return
	}
}

func PutArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("put /article/put")
	
	if r.Method == "PUT" {
		
		db, err := config.ConnectDB()
		if err != nil {
			log.Printf(err.Error())
		}
		
		defer db.Close()
		
		var article = models.Article{}
		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		_, err = db.Exec("UPDATE article set title=?, writer=? where id=?", article.Title, article.Writer, article.Id)
		
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		log.Printf("Update Sukses")
		
		json.NewEncoder(w).Encode(article)
	} else {
		log.Printf("Invalid method")
		http.Error(w, "Invalid method / method harusnya PUT", http.StatusMethodNotAllowed)
		return
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete /article/delete")
	
	if r.Method == "DELETE" {
		
		db, err := config.ConnectDB()
		if err != nil {
			log.Printf(err.Error())
		}
		
		defer db.Close()
		
		var article = models.Article{}
		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		
		_, err = db.Exec("DELETE from article where id=?", article.Id)
		
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

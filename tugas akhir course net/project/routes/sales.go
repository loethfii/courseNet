package routes

import (
	"encoding/json"
	"fmt"
	"github.com/loethfii/courseNet/config"
	"github.com/loethfii/courseNet/models"
	"log"
	"net/http"
	"strings"
)

func GetSales(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
		}

		defer db.Close()

		rows, err := db.Query("select * from sales order by id desc")

		if err != nil {
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		defer rows.Close()

		var response models.ArrSales

		for rows.Next() {
			var data = models.Sales{}
			var err = rows.Scan(&data.Id, &data.NamaSales, &data.Nip, &data.NomerTelpon, &data.Bagian, &data.TimeStamp)
			if err != nil {
				log.Printf(err.Error())
				return
			}

			response = append(response, data)
		}

		json.NewEncoder(w).Encode(response)
	} else {
		log.Println("Method harus GET", http.StatusMethodNotAllowed)
	}
}

func PostSales(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		var sales = models.Sales{}

		err = json.NewDecoder(r.Body).Decode(&sales)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		_, err = db.Exec("insert into sales (nama_sales, nip, nomer_telpon, bagian) values (?,?,?,?)", sales.NamaSales, sales.Nip, sales.NomerTelpon, sales.Bagian)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		log.Printf("Insert Sukses")

		//json.NewEncoder(w).Encode(sales)

		//jalanin Get Sales

		var response models.Sales

		err = db.QueryRow("select * from sales ORDER BY id DESC LIMIT 1").Scan(&response.Id, &response.NamaSales, &response.Nip, &response.NomerTelpon, &response.Bagian, &response.TimeStamp)

		if err != nil {
			fmt.Println(err.Error())
		}

		json.NewEncoder(w).Encode(response)

	} else {
		log.Println("Method harus POST", http.StatusMethodNotAllowed)
	}
}

func UpdateSales(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		id := strings.TrimPrefix(r.URL.Path, "/sales/")

		sales := models.Sales{}

		err = json.NewDecoder(r.Body).Decode(&sales)

		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = db.Exec("UPDATE sales set nama_sales = ?, nip = ?, nomer_telpon = ?, bagian =? WHERE id = ? ",
			sales.NamaSales,
			sales.Nip,
			sales.NomerTelpon,
			sales.Bagian,
			id,
		)

		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Printf("Berhasil UPDATE")
		w.Write([]byte("Berhasil Di Upadete"))

	} else {
		log.Println("Method harus PUT", http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method harus PUT", http.StatusMethodNotAllowed)
	}
}

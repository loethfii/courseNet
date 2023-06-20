package routes

import (
	"encoding/json"
	"fmt"
	"github.com/loethfii/courseNet/config"
	"github.com/loethfii/courseNet/models"
	"log"
	"net/http"
	"reflect"
)

func GetMobil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
		}

		defer db.Close()

		rows, err := db.Query("select * from mobil order by id desc")

		if err != nil {
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		defer rows.Close()

		var response models.ArrMobil

		for rows.Next() {
			var data = models.Mobil{}
			var err = rows.Scan(&data.Id, &data.NamaMobil, &data.TipeMobil, &data.JenisMobil, &data.BahanBakar, &data.Isi_Silinder, &data.Transmisi, &data.Warna, &data.Harga, &data.Qty, &data.TimeStamp)
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

func PostMobil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		var mobil = models.Mobil{}

		err = json.NewDecoder(r.Body).Decode(&mobil)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		_, err = db.Exec("insert into mobil (nama_mobil, tipe_mobil,jenis_mobil, bahan_bakar, isi_silinder, transmisi, warna, harga, qty) values (?,?,?,?,?,?,?,?,?)",
			mobil.NamaMobil,
			mobil.TipeMobil,
			mobil.JenisMobil,
			mobil.BahanBakar,
			mobil.Isi_Silinder,
			mobil.Transmisi,
			mobil.Warna,
			mobil.Harga,
			mobil.Qty,
		)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		log.Printf("Insert Sukses")

		var response models.Mobil

		err = db.QueryRow("select * from mobil ORDER BY id DESC LIMIT 1").Scan(
			&response.Id,
			&response.NamaMobil,
			&response.TipeMobil,
			&response.JenisMobil,
			&response.BahanBakar,
			&response.Isi_Silinder,
			&response.Transmisi,
			&response.Warna,
			&response.Harga,
			&response.Qty,
			&response.TimeStamp,
		)

		if err != nil {
			fmt.Println(err.Error())
		}

		json.NewEncoder(w).Encode(response)

	} else {
		log.Println("Method harus POST", http.StatusMethodNotAllowed)
	}
}

func CariNamaMobil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		namaMobil := r.URL.Query().Get("nama")

		fmt.Println(namaMobil)
		fmt.Println(reflect.TypeOf(namaMobil))

		rows, err := db.Query("SELECT * FROM mobil where nama_mobil like ?", "%"+namaMobil+"%")

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var response models.ArrMobil

		for rows.Next() {

			data := models.Mobil{}

			rows.Scan(
				&data.Id,
				&data.NamaMobil,
				&data.TipeMobil,
				&data.BahanBakar,
				&data.Isi_Silinder,
				&data.Transmisi,
				&data.Warna,
				&data.Harga,
				&data.Qty,
			)

			response = append(response, data)
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(response)

	} else {
		fmt.Println("Method haru GET", http.StatusMethodNotAllowed)
	}
}

func GetMobilById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		mobil := models.Mobil{}

		err = json.NewDecoder(r.Body).Decode(&mobil)
		if err != nil {
			panic(err.Error())
		}

		//namaMobil := r.URL.Query().Get("nama")
		//id := strings.TrimPrefix(r.URL.Path, "/mobil/")
		//
		//fmt.Println(reflect.TypeOf(id))
		//
		//idInt, _ := strconv.Atoi(id)

		//fmt.Println(idInt)
		//fmt.Println(reflect.TypeOf(idInt))

		err = db.QueryRow("SELECT * FROM mobil where id = ?", mobil.Id).Scan(
			&mobil.Id,
			&mobil.NamaMobil,
			&mobil.TipeMobil,
			&mobil.JenisMobil,
			&mobil.BahanBakar,
			&mobil.Isi_Silinder,
			&mobil.Transmisi,
			&mobil.Warna,
			&mobil.Harga,
			&mobil.Qty,
			&mobil.TimeStamp,
		)

		err = json.NewEncoder(w).Encode(mobil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {
		fmt.Println("Method harus GET", http.StatusMethodNotAllowed)
	}
}

package routes

import (
	"encoding/json"
	"fmt"
	"github.com/loethfii/courseNet/config"
	"github.com/loethfii/courseNet/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetPembayaran(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
		}

		defer db.Close()

		rows, err := db.Query("select * from pembayaran order by id desc")

		defer rows.Close()

		var pembayaran models.ArrPembayaran

		for rows.Next() {
			data := models.Pembayaran{}

			err = rows.Scan(&data.Id, &data.FormPembelian.Id, &data.BuktiTransfer, &data.IsConfirm, &data.CreatedAt)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			pembayaran = append(pembayaran, data)
		}

		err = json.NewEncoder(w).Encode(pembayaran)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Berhasil GET Pembayaran")
	} else {
		fmt.Fprintln(w, "Method Harus GET", http.StatusMethodNotAllowed)
	}
}

func PostPembayaran(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		if err != nil {
			log.Printf(err.Error())
			return
		}

		idFormulirPembelian := r.FormValue("id_formulir_pembelian")

		file, handler, err := r.FormFile("bukti_transfer")
		if err != nil {
			//http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err.Error())
			return
		}
		defer file.Close()

		// Membuat file baru di server untuk menyimpan file yang diunggah
		newFile, err := os.Create("./assets/img/" + handler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// Menyalin isi file yang diunggah ke file baru di server
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("insert into pembayaran (id_formulir_pembelian,bukti_transfer) values (?,?)",
			idFormulirPembelian,
			handler.Filename,
		)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		log.Printf("Insert Sukses")

		if err == nil {
			fmt.Fprintln(w, "Berhasil Input")
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	} else {
		log.Println("Method harus POST", http.StatusMethodNotAllowed)
	}
}

func ConfirmPembayaran(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		id := strings.TrimPrefix(r.URL.Path, "/pembayaran/confirm/")

		var pembayaran = models.Pembayaran{}

		err = db.QueryRow("SELECT pembayaran.is_confirm,  formulir_pembelian.id_mobil, formulir_pembelian.harus_inden FROM pembayaran INNER JOIN formulir_pembelian ON pembayaran.id_formulir_pembelian = formulir_pembelian.id where pembayaran.id = ?", id).Scan(
			&pembayaran.IsConfirm, &pembayaran.FormPembelian.Mobil.Id, &pembayaran.FormPembelian.HarusInden,
		)

		if err != nil {
			log.Println(err.Error())
		}

		//err = db.QueryRow("select formulir_pembelian.harus_inden from pembayaran inner join formulir_pembelian on pembayaran.id_formulir_pembelian = formulir_pembelian.id where pembayaran.id = ? order by pembayaran.id desc;").Scan(&formPembelianTampil.Mobil.Qty)

		if pembayaran.IsConfirm == false {
			_, err = db.Exec("UPDATE pembayaran set is_confirm = ? WHERE id = ? ", true, id)

			if err != nil {
				log.Println(err.Error())
				return
			}

			if pembayaran.FormPembelian.HarusInden == false {
				_, err = db.Exec("UPDATE mobil set qty = qty - 1 WHERE id = ? ", pembayaran.FormPembelian.Mobil.Id)
				if err != nil {
					log.Println(err.Error())
					return
				}
			}

			log.Printf("Berhasil Dikonfirmasi")
			fmt.Fprintln(w, "Berhasil Dikonfirmasi")
		} else {
			fmt.Fprintln(w, "Pembayaran yang sudah terkonfirmasi tidak dapat dikonfirmasi lagi!")
		}

	} else {
		log.Println("Method harus PUT", http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method harus PUT", http.StatusMethodNotAllowed)
	}
}

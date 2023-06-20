package routes

import (
	"encoding/json"
	"fmt"
	"github.com/loethfii/courseNet/config"
	"github.com/loethfii/courseNet/models"
	"log"
	"net/http"
)

func GetFormulirPembelian(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
		}

		defer db.Close()

		rows, err := db.Query(`select formulir_pembelian.id, formulir_pembelian.nama_lengkap_pembeli, formulir_pembelian.nomer_ktp,
			formulir_pembelian.alamat_rumah,formulir_pembelian.nomer_debit, mobil.id, mobil.nama_mobil, mobil.tipe_mobil, mobil.bahan_bakar,
            mobil.transmisi, mobil.warna, mobil.isi_silinder, mobil.harga,
            formulir_pembelian.harus_inden, formulir_pembelian.lama_inden,
            formulir_pembelian.custom_plat, formulir_pembelian.tambahan_kit,
			sales.nama_sales, formulir_pembelian.time_stamp from formulir_pembelian
			inner join mobil on  formulir_pembelian.id_mobil = mobil.id
			inner join sales on formulir_pembelian.id_sales = sales.id
			order by formulir_pembelian.id desc`)

		if err != nil {
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		defer rows.Close()

		var response models.ArrFormPembelian

		for rows.Next() {
			var data = models.FormPembelian{}
			var err = rows.Scan(
				&data.Id, &data.NamaLengkapPembeli, &data.NomerKTP, &data.AlamatRumah, //sudah
				&data.NomerDebit, &data.Mobil.Id, &data.Mobil.NamaMobil, &data.Mobil.TipeMobil, &data.Mobil.BahanBakar, //sudah
				&data.Mobil.Transmisi, &data.Mobil.Warna, &data.Mobil.Isi_Silinder, &data.Mobil.Harga, //sudah
				&data.HarusInden, &data.LamanInden, &data.CustomPlat, &data.TambahanKit, //sudah
				&data.Sales.NamaSales, &data.TimeStamp, //sudah
			)
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

func PostFormulirPembelian(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		defer db.Close()

		var formPembelian = models.NewFormPembelian{}

		var formPembelianTampil = models.FormPembelian{}

		err = json.NewDecoder(r.Body).Decode(&formPembelian)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		err = db.QueryRow("select qty from mobil where id = ?", formPembelian.IdMobil).Scan(&formPembelianTampil.Mobil.Qty)

		if formPembelianTampil.Mobil.Qty < 1 && formPembelian.HarusInden == true {
			var _, err2 = db.Exec(`insert into formulir_pembelian(
                               nama_lengkap_pembeli,
                               nomer_ktp,
                               alamat_rumah,
                               nomer_debit,
                               id_mobil,
                               harus_inden,
                               lama_inden,
                               custom_plat,
                               tambahan_kit,
                               id_sales) values (?,?,?,?,?,?,?,?,?,?)`,
				formPembelian.NamaLengkapPembeli,
				formPembelian.NomerKTP,
				formPembelian.AlamatRumah,
				formPembelian.NomerDebit,
				formPembelian.IdMobil,
				formPembelian.HarusInden,
				formPembelian.LamanInden,
				formPembelian.CustomPlat,
				formPembelian.TambahanKit,
				formPembelian.IdSales,
			)
			if err2 != nil {
				log.Printf(err.Error())
				return
			}

			log.Printf("Insert Dengan Inden Sukses")

			w.Write([]byte("Insert Dengan Inden Sukses"))
		} else if formPembelianTampil.Mobil.Qty < 1 {
			fmt.Fprintln(w, "Mobil Sedang Tidak Tersedia Harus Inden")
		}

		if formPembelianTampil.Mobil.Qty >= 1 && formPembelian.HarusInden == false {
			var _, err2 = db.Exec(`insert into formulir_pembelian(
                               nama_lengkap_pembeli,
                               nomer_ktp,
                               alamat_rumah,
                               nomer_debit,
                               id_mobil,
                               harus_inden,
                               lama_inden,
                               custom_plat,
                               tambahan_kit,
                               id_sales) values (?,?,?,?,?,?,?,?,?,?)`,
				formPembelian.NamaLengkapPembeli,
				formPembelian.NomerKTP,
				formPembelian.AlamatRumah,
				formPembelian.NomerDebit,
				formPembelian.IdMobil,
				formPembelian.HarusInden,
				formPembelian.LamanInden,
				formPembelian.CustomPlat,
				formPembelian.TambahanKit,
				formPembelian.IdSales,
			)
			if err2 != nil {
				log.Printf(err.Error())
				return
			}

			log.Printf("Insert Tanpa Inden Sukses")

			w.Write([]byte("Insert Tanpa Inden Sukses"))
		} else if formPembelianTampil.Mobil.Qty >= 1 && formPembelian.HarusInden == true {
			fmt.Fprintln(w, "Ngapain Inden Mobil Tersedia Ko!")
		}

	} else {
		log.Println("Method harus POST", http.StatusMethodNotAllowed)
	}
}

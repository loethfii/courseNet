package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/loethfii/courseNet/routes"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.EndPoint)
	http.HandleFunc("/sales", routes.GetSales)
	http.HandleFunc("/sales/post", routes.PostSales)
	http.HandleFunc("/sales/", routes.UpdateSales)

	http.HandleFunc("/mobil", routes.GetMobil)
	http.HandleFunc("/mobil/carinama", routes.CariNamaMobil) //gagal
	http.HandleFunc("/mobil/post", routes.PostMobil)
	http.HandleFunc("/mobil/cari", routes.GetMobilById)

	http.HandleFunc("/formpembelian", routes.GetFormulirPembelian)
	http.HandleFunc("/formpembelian/post", routes.PostFormulirPembelian)

	http.HandleFunc("/pembayaran", routes.GetPembayaran)
	http.HandleFunc("/pembayaran/post", routes.PostPembayaran)
	http.HandleFunc("/pembayaran/confirm/", routes.ConfirmPembayaran)

	fmt.Printf("Terhubung server http://localhost:3000/")

	http.ListenAndServe(":3000", nil)
}

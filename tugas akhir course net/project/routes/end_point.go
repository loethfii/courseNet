package routes

import (
	"fmt"
	"net/http"
)

func EndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	fmt.Fprintln(w, "GET Sales 			: <a href='/sales'>localhost:3000/sales</a> <br>")
	fmt.Fprintln(w, "GET Mobil			: <a href='/mobil'>localhost:3000/mobil</a><br>")
	fmt.Fprintln(w, "GET Form Pembelian 	: <a href='/formpembelian'>localhost:3000/formpembelian</a><br>")
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Warga struct {
	Name      string `json:"name"`
	JmlahAnak string `json:"jmlah_anak"`
}

var dataWarga = []Warga{{Name: "Rudi", JmlahAnak: "3"}}

func listenItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(dataWarga)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Write(result)
		return
	}
}

func main() {
	http.HandleFunc("/", listenItem)
	
	http.ListenAndServe(":3000", nil)
}

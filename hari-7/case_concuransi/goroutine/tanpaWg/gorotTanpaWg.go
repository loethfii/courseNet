package main

import (
	"fmt"
)

// restoran - AGung gsa
// ikan -> cuci -> simpan frozen
// daging -> potong -> simpan frozen
// sayur ->kupas -> simpan dibox->simpan di chiller
//buah -> cuci -> simpan chiller
//bawang -> kupas-> simpan di box -> simpan chiller
//selain itu buang ke tong sampah

// metode
var cuciChannel = make(chan string)
var potongCHannel = make(chan string)
var kupasChannel = make(chan string)

// simpan
var simpanFrozenChannel = make(chan string)
var SimpanChillerChannel = make(chan string)
var SimpanBoxChannel = make(chan string)


func cekBarang(belanja []string) {
    for _, barang := range belanja {
        switch barang {
        case "ikan":
            fmt.Println("Menemukan", barang)
            cuciChannel <- barang
        case "daging":
            fmt.Println("Menemukan", barang)
            potongCHannel <- barang
        case "sayur":
            fmt.Println("Menemukan ", barang)
            kupasChannel <- barang
        case "buah":
            fmt.Println("Menemukan ", barang)
            cuciChannel <- barang
        case "bawang":
            fmt.Println("Menemukan ", barang)
            kupasChannel <- barang
        default:
            fmt.Println("Dibuang aja", barang, "nya")
        }
    }
}

func main() {
    fmt.Println("Go routine")
    belanja := []string{
        "ikan", "daging", "plastik", "sayur", "plastik", "bambu", "buah", "sterofoam", "bawang",
        "ikan", "ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang",
        "buah", "sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang",
        "buah", "sterofoam", "bawang",
    }


    //method
    go cekBarang(belanja)
    go cuci()
    go simpanFrozen()
    go potong()
    //simpan
    go kupas()
    go simpanBox()
    go simpanChiller()

    fmt.Scanln()

    fmt.Println("Selesai")
    fmt.Println("========================================================")
}

// methode
func cuci() {
    for itemCuci := range cuciChannel {
        fmt.Println("Dicuci dulu")
        if itemCuci == "ikan" {
            simpanFrozenChannel <- itemCuci
        } else if itemCuci == "buah" {
            SimpanChillerChannel <- itemCuci
        }
    }
}

func potong() {
  
    for itemPotong := range potongCHannel {
        fmt.Println(itemPotong, " dipotong dulu")
        if itemPotong == "daging" {
            simpanFrozenChannel <- itemPotong
        }
    }
}

func kupas() {
 
    for itemKupas := range kupasChannel {
        fmt.Println(itemKupas, "dikupas dulu")
        if itemKupas == "sayur" {
            SimpanBoxChannel <- itemKupas
        }
    }
}

// simpan
func simpanFrozen() {
    for itemSimpanFrozen := range simpanFrozenChannel {
        fmt.Println(itemSimpanFrozen, "Simpan Frozen")
    }
}

func simpanBox() {
    for itemSimpanBox := range SimpanBoxChannel {
        fmt.Println(itemSimpanBox, "Simpan box")
        if itemSimpanBox == "sayur" {
            SimpanChillerChannel <- itemSimpanBox
        }
    }
}

func simpanChiller() {
       for itemSimpanChiller := range SimpanChillerChannel {
        fmt.Println(itemSimpanChiller, "Simpan chiller")
    }
}
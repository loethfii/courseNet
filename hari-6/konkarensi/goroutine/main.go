package main

import (
	"fmt"
	"runtime"
)

func print(angka int, message string) {
	for i := 0; i <= angka; i++ {
		fmt.Println(angka, message, i)
	}
}

func mobil(angka int, message string) {
	for i := 0; i <= angka; i++ {
		fmt.Println(angka, message, i)
	}
}

func main() {
	fmt.Println("Gorotine")

	fmt.Println("Angka 1")
	fmt.Println("Angka 2")
	fmt.Println("Angka 4")
	fmt.Println("Angka 5")
	fmt.Println("Angka 6")

	print(4, "Halli cetak")

	runtime.GOMAXPROCS(15)
\
}

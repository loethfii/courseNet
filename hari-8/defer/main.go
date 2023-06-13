package main

import "fmt"

func buka() {
	fmt.Println("buka")
}

func tutup() {
	fmt.Println("Tutup")
}

func main() {
	fmt.Println("defer")
	
	defer tutup()
	
	buka()
	fmt.Println("AMbil makanan")
	fmt.Println("AMbil minuman")
	fmt.Println("AMbil wedang")
}

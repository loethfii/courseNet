package main

import (
	"fmt"
)

func dataName() []string {
	name := []string{
		"Jaka Maulana",
		"Rahmat Darmawan",
		"Piki Muhammad Diki",
		"Bima Lampung",
		"Firman Hutabrat",
	}

	return name
}

func main() {

	data := dataName()

	fmt.Println(data)

	var sayHay = make(chan string)
	sapaDulu := func(name string) {
		var sapaOrang = fmt.Sprintf("Hallo %v, Selamat Datang ", name)
		sayHay <- sapaOrang
	}

	for _, val := range data {
		go sapaDulu(val)
	}

	for i := 1; i <= len(data); i++ {
		fmt.Println(i, <-sayHay)
	}

}

package main

import "fmt"

func panggilSapa(s chan string) {
	fmt.Println(<-s)
}

func main() {
	var sapa = make(chan string)

	var cetaSapa = func(who string) {
		var str = "Hallo Nama kamu " + who + " yah!"
		sapa <- str
	}

	go cetaSapa("Rahman")

	panggilSapa(sapa)
}

package main

import (
	"fmt"
	"runtime"
)

var Message = make(chan string)

func hello(msg string) {
	var data = "Hello " + msg
	Message <- data
}

func main() {

	runtime.GOMAXPROCS(2)
	go hello("Nurul")
	go hello("Siska")

	var isi1 = <-Message

	fmt.Println(isi1)
}

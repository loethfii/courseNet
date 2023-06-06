package main

import (
	"fmt"
	"sync"
)

func hello(angka int, message string, wg *sync.WaitGroup) {
	for i := 0; i <= angka; i++ {
		fmt.Println(angka, message, i)
	}

	wg.Done()
}

func main() {
	fmt.Println("Wait grup")

	var wg sync.WaitGroup

	//wg. add(3)
	wg.Add(1)

	go hello(1, "Hello Angga", &wg)

	//wg. add(3)
	wg.Add(1)

	go hello(1, "Hello Rahman", &wg)

	wg.Add(1)

	go hello(1, "Hello Bima", &wg)

	wg.Wait()

}

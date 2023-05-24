package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello1() {
	fmt.Println("Hello 1")
}

func printHello2() {
	fmt.Println("Hello 2")
}

func printHello3() {
	fmt.Println("Hello 3")
}

type data interface {
	printHello1()
	printHello2()
	printHello3()
}

func main() {

	go printHello1()
	time.Sleep(1 * time.Second)
	go printHello2()
	time.Sleep(1 * time.Second)
	go printHello3()
	time.Sleep(1 * time.Second)

	time.Sleep(3 * time.Second)

	var wg sync.WaitGroup

	var start = time.Now()

	fmt.Println("Start")
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("go rotine ke", i)
		}(i)
	}

	elapsed := time.Since(start)

	fmt.Println("Waktu eksekusi:", elapsed.Seconds())
	fmt.Println("Finish")

}

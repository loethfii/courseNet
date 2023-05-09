package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type n struct {
	name string
}

func main() {
	fmt.Println("Hello")

	fmt.Printf("Score %09d, %.2f, %t\n", 900, 0.7877, true)

	var name string
	name = "Luthfi"

	var age int
	age = 90

	var weight float64
	weight = 8.8

	var smoking bool
	smoking = true

	fmt.Println(name, age, weight, smoking)

	salary := 900000

	salary = 2999222

	salary = 222

	fmt.Println(salary)

	var x = 900

	uang := "RP" + strconv.Itoa(x)

	fmt.Println(uang)

	fmt.Print("Masukkan Nama Anda : ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Nama ku " + input)

	var nameLaptop string

	_, err := fmt.Scanln(&nameLaptop)

	fmt.Println(nameLaptop, err.Error())

}

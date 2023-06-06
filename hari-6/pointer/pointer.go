package main

import "fmt"

func main() {
	fmt.Println("Pointer")

	var a = 90
	fmt.Println("value ", a)
	fmt.Println("value ", &a)

	var b = 90
	fmt.Println("value ", b)
	fmt.Println("value ", &b)

	var c = a
	fmt.Println("value ", c)
	fmt.Println("value ", &c)

	var d = &c

	fmt.Println("pointer d : ", d)

	fmt.Println("==================================================")
	var new = 900
	var new1 = &new
	fmt.Println("Pointer new", &new)
	fmt.Println("Pointer new1", &new1)
	fmt.Println("Pointer new1 ambil dari pointer new", new1)

	fmt.Println("==================================================")

	var nn = 800
	var new2 = &nn

	fmt.Println("new3 :", &new2)

	var new3 = &new2

	fmt.Println("new3 :", new3)

	fmt.Println("==================================================")

	fmt.Println("==================================================")

	//copy value
	var newNme = "Rizki"
	fmt.Println(newNme) //rizki

	var newNme3 = newNme
	newNme3 = "saepul"
	fmt.Println(newNme)  //juan //berganti dari Rizki ke juan karna ubah pointer
	fmt.Println(newNme3) //value dari newNme3

	//pointer copy

	var newNme2 = &newNme

	*newNme2 = "Juan"

	fmt.Println(newNme)   //juan //berganti dari Rizki ke juan karna ubah pointer
	fmt.Println(newNme2)  //memory newNme2
	fmt.Println(*newNme2) //value dari newNme2

	fmt.Println("==================================================")

	num1 := 99
	num2 := &num1
	a = 70
	*num2 = a

	fmt.Println(*num2)
}

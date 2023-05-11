package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = [5]string{}

	a[0] = "Hello"
	a[1] = "Hello"
	a[2] = "Hello"
	a[3] = "Hello"
	a[4] = "Hello"

	fmt.Println("a", a)
	fmt.Println("reflect.TypeOf(a)", reflect.TypeOf(a))

	var b = [7]int{2, 31, 23, 4, 3}

	fmt.Println(b)
	fmt.Println(len(b))
	//
	for i := 0; i < len(b); i++ {
		fmt.Println("Angka ke-", i, b[i])
	}

	var kue = [3]string{"nastar", "Bolu", "Brownies"}

	for _, v := range kue {
		fmt.Println("Kue : ", v)
	}
}

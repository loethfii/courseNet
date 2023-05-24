package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Interface")

	var x interface{}
	fmt.Println(x)

	fmt.Println("Nil : ", reflect.TypeOf(x))

	x = "interface in string"
	fmt.Println(x)

	fmt.Println("string : ", reflect.TypeOf(x))

	x = []string{"asdasd", "adasdas"}
	fmt.Println(x)
	fmt.Println("slice : ", reflect.TypeOf(x))

	xyz := map[string]any{
		"nama":  "Luthfi",
		"nilai": 90,
	}

	fmt.Println("Map : ", reflect.TypeOf(xyz))

	fmt.Println(xyz)
}

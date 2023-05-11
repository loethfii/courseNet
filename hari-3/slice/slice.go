package main

import "fmt"

func main() {
	//fmt.Println("Slice")
	//var a = [900]string{"a", "v", "fff", "ppp", "", "", ""}
	//fmt.Println("a : ", a)
	//fmt.Println("reflect.TypeOf(a)", reflect.TypeOf(a))
	//fmt.Println("a[0]", a[0])
	//fmt.Println("a[0]", a[1])
	//
	//fmt.Println("len(a)", len(a))
	//fmt.Println("len(a)", cap(a))

	var foods = []string{"egg", "food", "ZZa"}
	fmt.Println("len(foods)", len(foods))
	fmt.Println("len(foods)", cap(foods))

	var newFoods = foods[0:2]
	fmt.Println("len(newFoods)", len(newFoods))
	fmt.Println("len(newFoods)", cap(newFoods))
}

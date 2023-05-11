package main

import "fmt"

func main() {
	//var a = [5]string{}
	//
	//a[0] = "Hello"
	//a[1] = "Hello"
	//a[2] = "Hello"
	//a[3] = "Hello"
	//a[4] = "Hello"
	//
	//fmt.Println("a", a)
	//fmt.Println("reflect.TypeOf(a)", reflect.TypeOf(a))
	//
	//var b = [7]int{2, 31, 23, 4, 3}
	//
	//fmt.Println(b)
	//fmt.Println(len(b))
	////
	//for i := 0; i < len(b); i++ {
	//	fmt.Println("Angka ke-", i, b[i])
	//}
	//
	//var kue = [3]string{"nastar", "Bolu", "Brownies"}
	//
	//for _, v := range kue {
	//	fmt.Println("Kue : ", v)
	//}

	//c := [...]uint{1, 3, 4, 5, 6, 2, 3}
	//
	//fmt.Println(c)
	//
	//a1 := [2][3]int{
	//	{13, 4, 5},
	//	{41, 4, 5},
	//}
	//
	//fmt.Println("a1", a1)

	var angka [5]int

	angka[0] = 2
	angka[1] = 34

	var huruf [5]string

	huruf[0] = "GXX"
	huruf[1] = "FFGH"
	huruf[2] = "JHHJ"

	fmt.Println(angka)
	fmt.Println(huruf)
}

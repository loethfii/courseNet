package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var x int8 = 10
	var y int16 = 10

	jumlah := int16(x) + y

	fmt.Println("JUmlah", jumlah)

	var a uint16 = 21
	jumlahUint16 := a + uint16(y)
	fmt.Println(jumlahUint16)

	var b string = "123"

	tint, err := strconv.Atoi(b)
	fmt.Println(tint)
	fmt.Println("Error : ", err)
	fmt.Println(reflect.ValueOf(tint))
	fmt.Println(reflect.ValueOf(tint).Type())

	//str con itoa
	var c uint16 = 123
	cint := strconv.Itoa(int(c))
	fmt.Println(reflect.ValueOf(cint))
	fmt.Println(reflect.ValueOf(cint).Type())

	//uint32 to string
	var cuint32 uint32 = 8000
	konvuint32 := strconv.FormatUint(uint64(cuint32), 10)
	fmt.Println(konvuint32)

	//string to float

	d := "3.19"
	dfloat, err2 := strconv.ParseFloat(d, 64)
	fmt.Println(err2)
	fmt.Println(reflect.ValueOf(dfloat))
	fmt.Println(reflect.ValueOf(dfloat).Type())

}

package main

import "fmt"

type ContohMethod struct {
	Name    string
	Address string
}

func (c ContohMethod) cetak() {
	fmt.Println(c)
}

func main() {
	ct1 := ContohMethod{Name: "Angga", Address: "Jakarta"}
	ct1.cetak()
}

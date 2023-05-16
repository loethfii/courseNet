package main

import (
	"errors"
	"fmt"
)

func test(test int) {
	fmt.Println("test : ", test)
}

func CekKebenaran(benargak bool) (string, error) {
	var Cek = "Apliakasi Aman"
	if benargak {
		return Cek, nil
	} else {
		return "", errors.New("aplikasi error")
	}
}

func main() {
	test(22222)

	aman, err := CekKebenaran(true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aman)
	}
}

package main

import "fmt"

type Animal interface {
	walk() int
	makesound() string
	breath() string
}

type Anjing struct {
	Legs        int
	Sound       string
	Respiretory string
}

func (an Anjing) breath() string {
	var a = an.Respiretory
	return a
}

func (an Anjing) makesound() string {
	var a = an.Sound
	return a
}

func (an Anjing) walk() int {
	var a = an.Legs
	return a
}

func main() {
	var anjing1 Animal

	anjing1 = Anjing{4, "wof wof", "lungs"}

	fmt.Println(anjing1.walk(), anjing1.makesound(), anjing1.breath())

}

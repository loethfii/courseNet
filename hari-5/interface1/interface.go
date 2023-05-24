package main

import (
	"fmt"
	"math"
)

type Bilangan struct {
	a int64
	b int64
}

func (b Bilangan) tambah() (resp int64) {
	resp = b.a + b.b

	return
}

// struct P L
type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	diameter float64
}

//end struct P L

// method Luas
func (p Persegi) Luas() float64 {
	return p.sisi * p.sisi
}

func (l Lingkaran) Luas() float64 {
	jari2 := l.diameter / 2
	return math.Pi * math.Pow(jari2, 2)
}

//end method Luas

// method keliling
func (p Persegi) Keliling() float64 {
	return 4 * p.sisi
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.diameter
}

//end method keliling

type Calculate interface {
	Luas() float64
	Keliling() float64
}

func main() {

	//struct
	fmt.Println("Interface 1")

	var x = Bilangan{10, 29}

	fmt.Println(x.tambah())

	//interface
	//interface persegi
	var persegi Calculate
	persegi = Persegi{30}
	fmt.Println("Luas persegi : ", persegi.Luas())
	fmt.Println("Keliling persegi : ", persegi.Keliling())

	//interface lingkaran
	var lingkaran Calculate
	lingkaran = Lingkaran{30}
	fmt.Println("Luas Lingkaran : ", lingkaran.Luas())
	fmt.Println("Keliling Lingkaran : ", lingkaran.Keliling())

}

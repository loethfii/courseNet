package main

import "fmt"

type Product interface {
	Showdetails() (string, string, float64)
	CalculatePrice(jumlahBarang int) float64
}

func CetakDetail(p Product) {
	title, author, price := p.Showdetails()

	fmt.Printf("Judul Buku : %s\n", title)
	fmt.Printf("Penulis : %s\n", author)
	fmt.Printf("Harga : %.3f\n", price)
}

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (books Book) Showdetails() (string, string, float64) {
	var title = books.Title
	var author = books.Author
	var price = books.Price

	return title, author, price
}

func (books Book) CalculatePrice(jumlahBarang int) float64 {
	var hasil = float64(jumlahBarang) * books.Price

	return hasil

}

type Clothing struct {
	Brand string
	Size  string
	Price float64
}

func (cloth Clothing) Showdetails() (string, string, float64) {
	var brand = cloth.Brand
	var size = cloth.Size
	var price = cloth.Price

	return brand, size, price
}

func (cloth Clothing) CalculatePrice(jumlahBarang int) float64 {
	var hasil = float64(jumlahBarang) * cloth.Price

	return hasil

}

func main() {
	var bukuLaskarPelangi Product

	bukuLaskarPelangi = Book{Title: "Buku Laska Pelangi", Author: "Ganesha Hadi Cahyo", Price: 70.000}

	title, author, price := bukuLaskarPelangi.Showdetails()

	//versi variable lansung
	fmt.Println(title, author, price)

	fmt.Println("================================================================")

	//versi function
	CetakDetail(bukuLaskarPelangi)
	fmt.Println("================================================================")

	harga := bukuLaskarPelangi.CalculatePrice(3)

	fmt.Printf("Harga : %.3f\n", harga)

	fmt.Println("========================================================")

	var jersiBola Product

	jersiBola = Clothing{Brand: "Adidas", Size: "XL", Price: 900.000}

	brand, size, pr := jersiBola.Showdetails()

	fmt.Println("=======================================================")
	fmt.Printf("Brand : %s\n", brand)
	fmt.Printf("Brand : %s\n", size)
	fmt.Printf("Brand : %.3f\n", pr)
	fmt.Println("=======================================================")
	asilHarga := jersiBola.CalculatePrice(2)
	fmt.Printf("Harga : %.3f \n", asilHarga)

}

package main

import "fmt"

func main() {
	fmt.Println("Switch Case")

	swVar := "Biru"
	switch swVar {
	case "Hijau":
		fmt.Println("Case Pertama")
	case "Kuning":
		fmt.Println("case Kedua")
	case "Biru":
		fmt.Println("case Ke tiga")
	default:
		fmt.Println("default")
	}

	switchNum := 100
	switch {
	case switchNum >= 100:
		fmt.Println("Case Pertama")
	case switchNum >= 30:
		fmt.Println("case Kedua")
	case switchNum >= 10:
		fmt.Println("case Ke tiga")
	default:
		fmt.Println("default")
	}

	fmt.Println("")
	auth := "Guest"
	switch auth {
	case "Admin":
		fmt.Println("Halaman Admin dan reporting")
		fallthrough
	case "Supervisor":
		fmt.Println("Mengakses Halaman pengunjung")
		fallthrough
	case "Guest":
		fmt.Println("Mengakses Halaman menu")
	default:
		fmt.Println("Diarhkan ke resitration")
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	speed := 65
	fast := speed >= 80
	slow := speed < 80

	fmt.Printf("Going fast %t\n", fast)
	fmt.Printf("Going fast %t\n", slow)

	speed = 100
	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	novel := "Laskar pelangi"

	fmt.Println(`Hallo namaku Aji kerjaan ku membaca buku
	novel yang ku suka ialah`, novel, `Hal ini bisa membuatku senang
	Setiap Hari`)

	color := "blue"
	fmt.Println("blue colors?", color == "blue" || color == "dark blue")

	n, err := strconv.Atoi("Purnomo")
	if err != nil {
		fmt.Println("Terdapat Error : ", err)
	} else {
		fmt.Println(n)
	}

	angka := 1
	switch angka {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}

	city := "Indonesia"
	switch city {
	case "Indonesia":
		fmt.Println("Jakarta")
	}

	city = "Tangerang"
	switch city {
	case "Tangerang", "Serang":
		fmt.Println("Jawa")
	case "Kalimantan":
		fmt.Println("Kalimantan")
	}

	city = "Ciamik soro"
	switch city {
	case "Tangerang", "Serang":
		fmt.Println("Jawa")
	case "Kalimantan":
		fmt.Println("Kalimantan")
	default:
		fmt.Println("Tidak dikenal")
	}

	i := 111
	switch {
	case i > 100:
		fmt.Print(">100")
		break
	case i > 0:
		fmt.Print(">0")
		break
	default:
		fmt.Println("Ini default")
	}
}

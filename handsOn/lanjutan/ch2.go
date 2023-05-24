package main

import "fmt"

func cetakMessage(inp chan string) {
	var increment = 0
	fmt.Println(increment+1, <-inp)
}

func siswaName() []string {
	name := []string{
		"Jaka Maulana",
		"Rahmat Darmawan",
		"Piki Muhammad Diki",
		"Bima Lampung",
		"Firman Hutabrat",
	}

	return name
}

func main() {
	var message = make(chan string)

	sName := siswaName()

	for _, each := range sName {
		go func(who string) {
			res := fmt.Sprintf("Hallo, %s Selamat Datang Di Kotaku", who)

			message <- res
		}(each)

	}

	for i := 1; i <= len(sName); i++ {
		cetakMessage(message)
	}

}

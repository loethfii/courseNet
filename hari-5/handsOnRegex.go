package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//contoh regex

	var contohText = "Banana Warna kuining"

	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regex.FindAllString(strings.ToLower(contohText), 1)

	fmt.Println(res1)

	// cari slice index

	fmt.Println("====================================")

	names := []string{
		"Daniel Saipul",
		"Kona Yuki",
		"Riwata",
		"John Doe",
		"Daniel Smith",
		"Michael Johnson",
		"Sarah Davis",
		"Daniel Brown",
		"Emily Wilson",
		"Daniel Taylor",
	}

	cari := "Daniel"

	regex2 := regexp.MustCompile(cari)

	tampungTr := []string{}
	for _, v := range names {
		if regex2.MatchString(v) {
			fmt.Println(v)
		}

		tr := regex2.MatchString(v)

		if tr == true {
			con := strconv.FormatBool(tr)
			tampungTr = append(tampungTr, con)

		}

	}

	if ketemu := len(tampungTr); ketemu <= 0 {
		fmt.Println("Nama tidak ditemukan")
	} else {
		fmt.Println("Ketemu ", len(tampungTr), " Data")
	}

	fmt.Println("=================================")
	fmt.Println("=================================")

	var angka1 = []int{2, 4, 5, 5, 4, 44, 4, 4, 23, 3, 3, 3, 45, 5656, 5}

	cariAngka := 5

	regex3 := regexp.MustCompile(strconv.Itoa(cariAngka))

	for _, v := range angka1 {
		if regex3.MatchString(strconv.Itoa(v)) {
			fmt.Println(v)
		}

	}
}

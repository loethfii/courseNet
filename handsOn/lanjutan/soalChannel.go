package main

import "fmt"

func countEventOdd(val []int32, evenCh chan int32, oddCh chan int32) {
	// evenCh itu angka genap
	// OddCh itu angka ganjil
	filterGenap := []int32{}
	filterGanjil := []int32{}

	for _, v := range val {
		if v%2 == 0 {
			filterGenap = append(filterGenap, v)
		} else {
			filterGanjil = append(filterGanjil, v)
		}
	}

	var genap = []int32{}
	for i := 1; i <= len(filterGenap); i++ {
		genap = append(genap, <-evenCh)
	}

	var ganjil = []int32{}
	for i := 1; i <= len(filterGanjil); i++ {
		ganjil = append(ganjil, <-oddCh)
	}

	fmt.Println("genap : ", genap)
	fmt.Println("Ganjil : ", ganjil)
}

func main() {
	var num = []int32{4, 1, 2, 3, 4, 56, 7, 5, 4, 5, 7, 91, 91, 91}

	evenCh := make(chan int32)
	oddCh := make(chan int32)

	for _, v := range num {
		go func(v int32) {
			if v%2 == 0 {
				evenCh <- v
			} else {
				oddCh <- v
			}
		}(v)
	}

	countEventOdd(num, evenCh, oddCh)

}

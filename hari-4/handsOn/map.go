package main

import "fmt"

type Siswa struct {
	Name       string
	NameFather string
	Class      string
}

func main() {
	var dataSiswa = map[int]Siswa{
		1: {Name: "Dono Putra", NameFather: "Irawan", Class: "4A"},
		2: {Name: "Silaban Coki", NameFather: "Paradede", Class: "4A"},
		3: {Name: "Silaban Coki", NameFather: "Paradede", Class: "4A"},
	}
	fmt.Println("No	Name		Father Name		Class")

	for i, v := range dataSiswa {
		fmt.Println("============================================")
		fmt.Printf("%d	%s		%s		%s\n", i, v.Name, v.NameFather, v.Class)
	}
}

package main

import (
	"fmt"
)

type Fungsi interface {
	AddStudent() string
	ShowStudent() []Student
}

type Student struct {
	Name   string
	NISN   string
	Alamat string
	NoHp   string
}

var newStudent = []Student{}

func (add Student) AddStudent() string {

	nStudent := Student{
		Name:   add.Name,
		NISN:   add.NISN,
		Alamat: add.Alamat,
		NoHp:   add.NoHp,
	}
	newStudent = append(newStudent, nStudent)
	return "Berhasil Di tambahkan"
}

func (show Student) ShowStudent() []Student {

	return newStudent
}

func cetakSemua() {
	for i, v2 := range newStudent {
		fmt.Printf("%d | %s %s %s %s\n", i+1, v2.Name, v2.NISN, v2.Alamat, v2.NoHp)

	}
}

func main() {
	var st1 Fungsi
	st1 = Student{
		Name:   "Ridwan",
		NISN:   "121212",
		Alamat: "Jakarta",
		NoHp:   "0895396173140",
	}

	success := st1.AddStudent()

	fmt.Println(success)
	fmt.Println(st1.ShowStudent())

	st2 := Student{
		Name:   "Gozali",
		NISN:   "19920002",
		Alamat: "asdasdas",
		NoHp:   "1212312312312",
	}

	success2 := st2.AddStudent()
	fmt.Println(success2)
	cetak := st2.ShowStudent()
	fmt.Println(cetak)

	var st3 Fungsi
	st3 = Student{
		Name:   "Maman",
		NISN:   "Firman",
		Alamat: "Magelang",
		NoHp:   "9009900900",
	}

	success2 = st3.AddStudent()

	cetakSemua()

}

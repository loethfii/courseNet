package main

import "fmt"

type student struct {
	Id     string
	Name   string `json:"student_name"`
	Adress string `json:"address"`
}

func main() {
	var studentData = student{"1", "Luthfi", "Malaka 4"}

	fmt.Println(studentData)
}

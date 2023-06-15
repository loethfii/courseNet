package models

type Article struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Writer string `json:"writer"`
}

type Articles []Article

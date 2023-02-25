package models

type Filter struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Genres    []int  `json:"genres"`
}

type NewBook struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Genres    []int  `json:"genres"`
}

type Book struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Author    string `db:"author" json:"author"`
	Publisher string `db:"publisher" json:"publisher"`
}

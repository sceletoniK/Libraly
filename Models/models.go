package models

type Filter struct {
	Name      string         `json:"name"`
	Author    string         `json:"author"`
	Publisher string         `json:"publisher"`
	Genres    map[string]int `json:"genres"`
}

type NewBook struct {
	Name      string         `json:"name"`
	Author    string         `json:"author"`
	Publisher string         `json:"publisher"`
	Genres    map[string]int `json:"genres"`
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Admin    int    `db:"admin" json:"admin"`
}

type Book struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Author    string `db:"author" json:"author"`
	Publisher string `db:"publisher" json:"publisher"`
}

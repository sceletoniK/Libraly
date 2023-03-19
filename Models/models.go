package models

import (
	"time"
)

type Filter struct {
	Name      string         `json:"name"`
	Author    string         `json:"author"`
	Publisher string         `json:"publisher"`
	Genres    map[string]int `json:"genres"`
}

type NewBook struct {
	Id        int            `db:"id" json:"id"`
	Name      string         `json:"name"`
	Author    string         `json:"author"`
	Publisher string         `json:"publisher"`
	Genres    map[string]int `json:"genres"`
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Admin    bool   `db:"admin" json:"admin"`
}

type Book struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Author    string `db:"author" json:"author"`
	Publisher string `db:"publisher" json:"publisher"`
}

type Session struct {
	RefreshToken string    `db:"refreshtoken" json:"refreshToken"`
	ClientId     int       `db:"clientid" json:"clientId"`
	ExpiresAt    time.Time `db:"expiresat" json:"expiresAt"`
}

type Cart struct {
	ClientId int `db:"clientid" json:"clientId"`
	BookId   int `db:"bookid" json:"bookId"`
}

type BookInstance struct {
	InstanceId int `db:"id" json:"instanceId"`
	BookId     int `db:"originalid" json:"bookId"`
}

type Rent struct {
	ClientId      int       `db:"clientid" json:"clientId"`
	InstanceId    int       `db:"instancebookid" json:"instanceId"`
	RequestDate   time.Time `db:"requestdate" json:"requestDate"`
	StartRentDate time.Time `db:"startrentdate" json:"startRentDate"`
	Deadline      time.Time `db:"deadline" json:"deadline"`
}

type FilterRent struct {
	ClientId int  `json:"clientId"`
	Started  bool `json:"Started"`
	Expired  bool `json:"Expired"`
}

type RentHistory struct {
	ClientId      int       `db:"clientid" json:"clientId"`
	BookId        int       `db:"bookid" json:"bookId"`
	StartRentDate time.Time `db:"startrentdate" json:"startRentDate"`
}

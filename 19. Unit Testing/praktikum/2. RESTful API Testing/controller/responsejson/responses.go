package responsejson

import (
	"time"
	"praktikum/rest/model"
)

type ResponseUser struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseBook struct {
	ID        uint      `json:"id"`
	Judul     string    `json:"judul,omitempty"`
	Penulis   string    `json:"penulis"`
	Penerbit  string    `json:"penerbit"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseLogin struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name,omitempty"`
	Token	  string	`json:"token"`
}

type ResponCreate struct {
	Status  string
	Data    model.User
}

type ResponSuccesSlice struct {
	Status  string
	Data    []model.User
}

type ResponCreateBooks struct {
	Status  string
	Data    model.Books
}
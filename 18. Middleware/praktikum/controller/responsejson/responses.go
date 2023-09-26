package responsejson

import "time"

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
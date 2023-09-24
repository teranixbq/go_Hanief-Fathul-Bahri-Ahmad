package responsejson

import "time"

type ResponseBlogs struct {
	ID        uint      `json:"id"`
	// UserID    uint     `json:"user_id,omitempty"`
	Judul     string    `json:"judul"`
	Konten    string    `json:"konten"`
	CreatedAt time.Time `json:"created_at"`
	User      ResponseUser `json:"user"`
}

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

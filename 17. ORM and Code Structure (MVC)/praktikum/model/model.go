package model

import "gorm.io/gorm"

type Config struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

type User struct {
    gorm.Model
    Name     string `json:"name"      form:"name"`
    Email    string `json:"email"     form:"email"`
    Password string `json:"password"  form:"password"`
}

type Books struct {
	gorm.Model
	Judul    string `json:"judul"    form:"judul"`
	Penulis  string `json:"penulis"  form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}


type Blogs struct {
    gorm.Model
    UserID   uint `json:"user_id" form:"user_id"`
    Judul    string `json:"judul"  form:"judul"`
    Konten   string `json:"konten" form:"konten"`
	User	User
}



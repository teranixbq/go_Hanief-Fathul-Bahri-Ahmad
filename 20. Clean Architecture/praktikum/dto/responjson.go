package dto

import (
	"clean/model"
	"time"
)

type ResponseLogin struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResponseGetAll struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FromAll(data model.User) ResponseGetAll {
	return ResponseGetAll{
		ID:        data.ID,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}
}

func FromAllData(data []model.User) []ResponseGetAll {
	result := []ResponseGetAll{}
	for key := range data {
		result = append(result, FromAll(data[key]))
	}
	return result
}

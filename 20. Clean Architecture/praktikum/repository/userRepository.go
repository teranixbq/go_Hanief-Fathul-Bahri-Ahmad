package repository

import (
	"clean/middleware"
	"clean/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(email, password string) (model.User, string, error)
	Create(data model.User) error
	SelectAll() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(data model.User) error {
	return ur.db.Create(&data).Error
}

func (ur *userRepository) SelectAll() ([]model.User, error) {
	var users []model.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) Login(email, password string) (model.User, string, error) {
	var data model.User

	conf := ur.db.Where("email = ?", email).First(&data)
	if conf.Error != nil {
		return model.User{}, "", conf.Error
	}

	if conf.RowsAffected == 0 {
		return model.User{}, "", errors.New("ser not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password)); err != nil {
		return model.User{}, "", errors.New("ncorrect password")
	}

	token := ""
	var errorToken error
	token, errorToken = middleware.CreateToken(int(data.ID))
	if errorToken != nil {
		return model.User{}, "", errorToken
	}

	return data, token, nil
}

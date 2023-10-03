package usercase

import (
	"clean/dto"
	"clean/model"
	"clean/repository"
	"errors"

	// "errors"

	// "github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Login(email, password string) (model.User, string, error)
	Create(payloads dto.CreateUserRequest) error
	SelectAll() ([]model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: ur}
}

func (us *userUsecase) Create(payload dto.CreateUserRequest) error {
	if payload.Email == "" || payload.Password == "" {
		return errors.New("error. email dan password harus diisi")
	}
	securePassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Email:    payload.Email,
		Password: string(securePassword),
	}

	if err := us.userRepository.Create(user); err != nil {
		return err
	}

	return nil
}

func (us *userUsecase) SelectAll() ([]model.User, error) {
	users, err := us.userRepository.SelectAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userUsecase) Login(email, password string) (model.User, string, error) {

	user := model.User{
		Email:    email,
		Password: password,
	}

	data, token, err := us.userRepository.Login(user.Email, user.Password)
	if err != nil {
		return model.User{}, "", err
	}
	return data, token, nil
}

// func (us *userUsecase) Login(email, password string) (model.User, string, error) {
//     userData, token, err := us.userRepository.Login(email, "")
//     if err != nil {
//         return model.User{}, "", err
//     }

//     err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
//     if err != nil {
//         return model.User{}, "", errors.New("password salah")
//     }

//     return userData, token, nil
// }

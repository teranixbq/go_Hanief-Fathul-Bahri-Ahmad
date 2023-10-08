package usercase

import (
	"clean/dto"
	"clean/mocks"
	"clean/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_Create(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	userUsecase := NewUserUseCase(userRepositoryMock)

	createUserRequest := dto.CreateUserRequest{
		Email:    "john@example.com",
		Password: "password",
	}

	userRepositoryMock.On("Create", mock.Anything).Return(nil)
	err := userUsecase.Create(createUserRequest)
	userRepositoryMock.AssertCalled(t, "Create", mock.Anything)
	assert.NoError(t, err)
}

func TestUserUsecase_CreateWithError(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	userUsecase := NewUserUseCase(userRepositoryMock)

	createUserRequest := dto.CreateUserRequest{
		Email:    "john@example.com",
		Password: "password",
	}

	expectedError := errors.New("some error occurred")
	userRepositoryMock.On("Create", mock.Anything).Return(expectedError)
	err := userUsecase.Create(createUserRequest)
	userRepositoryMock.AssertCalled(t, "Create", mock.Anything)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestUserUsecase_SelectAll(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	userUsecase := NewUserUseCase(userRepositoryMock)

	expectedUsers := []model.User{
		{Email: "user1@example.com", Password: "password1"},
		{Email: "user2@example.com", Password: "password2"},
	}

	userRepositoryMock.On("SelectAll").Return(expectedUsers, nil)
	users, err := userUsecase.SelectAll()
	userRepositoryMock.AssertCalled(t, "SelectAll")
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestUserUsecase_SelectAllWithError(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	userUsecase := NewUserUseCase(userRepositoryMock)

	expectedError := errors.New("some error occurred")

	userRepositoryMock.On("SelectAll").Return(nil, expectedError)
	users, err := userUsecase.SelectAll()
	userRepositoryMock.AssertCalled(t, "SelectAll")
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, users)
}

func TestUserUsecase_Login(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	userUsecase := NewUserUseCase(userRepositoryMock)

	email := "john@example.com"
	password := "password"

	expectedUser := model.User{
		Email:    email,
		Password: password,
	}
	expectedToken := "token123"

	userRepositoryMock.On("Login", email, password).Return(expectedUser, expectedToken, nil)
	user, token, err := userUsecase.Login(email, password)

	userRepositoryMock.AssertCalled(t, "Login", email, password)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.Equal(t, expectedToken, token)
}

func TestUserUsecase_LoginWithError(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	userUsecase := NewUserUseCase(userRepositoryMock)
	email := "john@example.com"
	password := "password"

	expectedError := errors.New("authentication failed")
	userRepositoryMock.On("Login", email, password).Return(model.User{}, "", expectedError)

	user, token, err := userUsecase.Login(email, password)

	userRepositoryMock.AssertCalled(t, "Login", email, password)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, model.User{}, user)
	assert.Empty(t, token)
}

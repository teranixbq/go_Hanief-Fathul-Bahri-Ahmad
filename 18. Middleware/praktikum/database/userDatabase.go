package database

import (
	"errors"
	"praktikum/rest/model"
	"praktikum/rest/config"
    "praktikum/rest/middleware"
)

func UserLoginAuth(email string, password string) (model.User, string, error) {
    var data model.User

    conf := config.DB.Where("email = ? AND password = ?", email, password).First(&data)
    if conf.Error != nil {
        return model.User{}, "", conf.Error
    }

    token := ""
    var errorToken error
    if conf.RowsAffected > 0 {
        token, errorToken = middleware.CreateToken(int(data.ID), data.Name)
        if errorToken != nil {
            return model.User{}, "", errorToken
        }
    } 
    return data, token, nil
}


// insert data user baru
func InsertUsers(user model.User) error {
	err := config.DB.Save(&user)
	if err.Error != nil {
		return errors.New("failed insert data")
	}
	return nil
}

// select semua data
func SelectAllUsers() ([]model.User, error) {
    var datausers []model.User

    if err := config.DB.Find(&datausers).Error; err != nil {
        return nil, errors.New("gagal mengambil data")
    }

    return datausers, nil
}

// select data berdasarkan ID
func SelectById(id int) (model.User, error) {
    var databyid model.User
    if err := config.DB.Where("id = ?", id).First(&databyid).Error; err != nil {
        return model.User{}, errors.New("gagal mengambil data")
    }
	
    return databyid, nil
}

// update data berdasrkan ID
func UpdateUser(id int,updatedUser model.User) error {
    user := model.User{}
    if err := config.DB.First(&user, id).Error; err != nil {
        return err
    }

    user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

    if err := config.DB.Save(&user).Error; err != nil {
        return err
    }

    return nil
}

// delete data berdasarkan ID
func DeleteUser(id int) error {
	user := model.User{}
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return errors.New("gagal menghapus data pengguna")
	}

	return nil
}
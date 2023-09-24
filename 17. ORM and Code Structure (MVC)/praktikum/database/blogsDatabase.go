package database

import (
	"errors"
	"praktikum/rest/model"
	"praktikum/rest/config"
)

//insert Data Blogs
func InsertBlogs(blogs model.Blogs) error {
	var user model.User
    if err := config.DB.First(&user, blogs.UserID).Error; err != nil {
        return errors.New("user dengan UserID tidak ditemukan")
    }

	if err := config.DB.Save(&blogs).Error; err != nil {
		return errors.New("gagal menambahkan data")
	}
	return nil
}

//select semua data blogs
func SelectAllBlogs() ([]model.Blogs, error) {
    var dataBlogs []model.Blogs

    if err := config.DB.Preload("User").Find(&dataBlogs).Error; err != nil {
        return nil, errors.New("gagal mengambil data")
    }

    return dataBlogs, nil
}

//select data berdasarkan ID
func SelectBlogById(id int) ([]model.Blogs, error) {
    var databyid []model.Blogs
    if err := config.DB.Preload("User").Where("id = ?", id).First(&databyid).Error; err != nil {
        return nil, errors.New("gagal mengambil data")
    }
    return databyid, nil
}

// update data berdasarkan ID
func UpdateBlogs(id int, updatedBlogs model.Blogs) error {
    blogs := model.Blogs{}
    if err := config.DB.First(&blogs, id).Error; err != nil {
        return err
    }

    var user model.User
    if err := config.DB.First(&user, updatedBlogs.UserID).Error; err != nil {
        return errors.New("user dengan UserID tidak ditemukan")
    }

    blogs.UserID = updatedBlogs.UserID
    blogs.Judul = updatedBlogs.Judul
    blogs.Konten = updatedBlogs.Konten

    if err := config.DB.Save(&blogs).Error; err != nil {
        return err
    }

    return nil
}

//delete data berdasarkan ID
func DeleteBlogs(id int) error {
	Blogs := model.Blogs{}
	if err := config.DB.Delete(&Blogs, id).Error; err != nil {
		return errors.New("gagal menghapus data ")
	}

	return nil
}


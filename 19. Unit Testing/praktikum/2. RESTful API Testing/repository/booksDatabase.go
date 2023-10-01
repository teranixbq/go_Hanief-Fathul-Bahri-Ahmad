package repository

import (
	"errors"
	"praktikum/rest/model"
	"praktikum/rest/config"
)

// insert data buku baru
func InsertBooks(books model.Books) error {
	err := config.DB.Save(&books)
	if err.Error != nil {
		return errors.New("failed")
	}
	return nil
}

// select semua data
func SelectAllBooks() ([]model.Books, error) {
    var dataBooks []model.Books

    if err := config.DB.Find(&dataBooks).Error; err != nil {
        return nil, errors.New("gagal mengambil data")
    }

    return dataBooks, nil
}

// select data berdasarkan ID
func SelectBooksById(id int) (model.Books, error) {
    var databyid model.Books
    if err := config.DB.Where("id = ?", id).First(&databyid).Error; err != nil {
        return model.Books{}, errors.New("gagal mengambil data")
    }
    return databyid, nil
}

// update data berdasrkan ID
func UpdateBooks(id int,updatedBooks model.Books) error {
    books := model.Books{}
    if err := config.DB.First(&books, id).Error; err != nil {
        return err
    }

    books.Judul = updatedBooks.Judul
	books.Penulis = updatedBooks.Penulis
	books.Penerbit = updatedBooks.Penerbit

    if err := config.DB.Save(&books).Error; err != nil {
        return err
    }

    return nil
}

// delete data berdasarkan ID
func DeleteBooks(id int) error {
	Books := model.Books{}
	if err := config.DB.Delete(&Books, id).Error; err != nil {
		return errors.New("gagal menghapus data pengguna")
	}

	return nil
}

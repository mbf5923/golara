package models

import (
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID    int
	Name  string
	Email string
}

type BookModel struct {
	database *gorm.DB `container:"name"`
}

func NewBookModel() *BookModel {
	myApp := BookModel{}
	err := container.Fill(&myApp)
	if err != nil {
		panic(err)
	}
	return &myApp
}

func (bookModel *BookModel) CreateBook(Book *Book) (err error) {
	err = bookModel.database.Create(Book).Error
	if err != nil {
		return err
	}
	return nil
}

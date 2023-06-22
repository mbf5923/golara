package models

import (
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID          uint   `gorm:"primaryKey;"`
	UserId      uint   `gorm:"type:int;not null" json:"user_id"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(500);not null"`
	Price       uint64 `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

// Books return all books by user
func (bookModel *BookModel) Books(userId int, book *[]Book) (err error) {
	err = bookModel.database.Where("user_id = ?", userId).Find(book).Error
	if err != nil {
		return err
	}
	return nil

}

package models

import (
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Name     string
	Email    string
	Password string
}

type UserModel struct {
	database *gorm.DB `container:"name"`
}

func NewUserModel() *UserModel {
	myApp := UserModel{}
	err := container.Fill(&myApp)
	if err != nil {
		panic(err)
	}
	return &myApp
}

// CreateUser create a user
func (userModel *UserModel) CreateUser(User *User) (err error) {
	err = userModel.database.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUsers get users
func (userModel *UserModel) GetUsers(User *[]User) (err error) {
	err = userModel.database.Order("id desc").Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser get user by id
func (userModel *UserModel) GetUser(User *User, id int) (err error) {
	err = userModel.database.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser update user
func (userModel *UserModel) UpdateUser(User *User) (err error) {
	userModel.database.Save(User)
	return nil
}

// DeleteUser delete user
func (userModel *UserModel) DeleteUser(User *User, id int) (err error) {
	userModel.database.Where("id = ?", id).Delete(User)
	return nil
}

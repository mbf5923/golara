package models

import (
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint    `gorm:"column:id;type:INT NOT NULL;primaryKey;autoIncrement:true"`
	Name      string  `gorm:"column:name;type:varchar(255);unique;not null"`
	Email     string  `gorm:"column:email;type:varchar(255);unique;not null"`
	Password  string  `gorm:"column:password;type:varchar(255);not null"`
	Active    bool    `gorm:"column:active;type:bool;default:false"`
	ApiKey    *string `gorm:"column:api_key;type:varchar(255);unique;null" json:"omitempty" sql:"DEFAULT:NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
func (userModel *UserModel) GetUser(User *User, id uint) (err error) {
	err = userModel.database.Where("id = ?", id).Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByApiKey get user by api key
func (userModel *UserModel) GetUserByApiKey(User *User, apiKey string) (err error) {
	err = userModel.database.Where("api_key = ?", apiKey).Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// CheckExistUserByEmail check exist user by email
func (userModel *UserModel) CheckExistUserByEmail(email string) (exists bool, err error) {
	err = userModel.database.Table("users").Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error
	return
}

// UpdateUser update user
func (userModel *UserModel) UpdateUser(User *User) (err error) {
	userModel.database.Save(User)
	return nil
}

// DeleteUser delete user
func (userModel *UserModel) DeleteUser(User *User, id uint) (err error) {
	userModel.database.Where("id = ?", id).Delete(User)
	return nil
}

// LoginUser login user
func (userModel *UserModel) LoginUser(user *User, email string, password string) (err error) {
	err = userModel.database.Where("email = ? AND password = ?", email, password).Find(&user).Error
	return
}

// UpdateApiKey update api key
func (userModel *UserModel) UpdateApiKey(User *User) (err error) {
	userModel.database.Save(User)
	return nil
}

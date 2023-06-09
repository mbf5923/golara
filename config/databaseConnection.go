package config

import (
	"github.com/sirupsen/logrus"
	"gorm-test/app/models"
	"gorm-test/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	databaseURI := make(chan string, 1)
	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- utils.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(mysql.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Book{},
		//&model.EntityStudent{},
	)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
		return db.Debug()
	}
	return db
}

package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/models"
)

func SignUp(db *gorm.DB, user *models.Users) error {
	//Save update value in database, if the value doesn't have primary key, will insert it
	//1. check whether email id exists or not
	//2. if email is unique then save user
	err := db.Save(user).Error //book here is without id
	if err != nil {
		return err
	}
	return nil
}

func SignIn(db *gorm.DB, book *models.Book) error {
	//Save update value in database, if the value doesn't have primary key, will insert it
	getUser := models.Users{}
	err := db.Select("users.*").Group("users.email_id").Where("users.email_id= ?", user.Email_Id).First(&getUser).Error //book here is with id
	if err != nil {
		return err
	}
	log.Println(getUser)

	if user.Password != getUser.Password {
		return errors.New("Password Incorrect")
	}

	log.Println("existing signin")
	return nil
}

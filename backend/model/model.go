package model

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) Create() error {
	user := User{}
	if err := db.Where(&User{Username: u.Username}).First(&user).Error; err != gorm.ErrRecordNotFound {
		if err == nil {
			return errors.New("username exist")
		} else {
			return err
		}
	}

	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) Find() error {
	user := User{}
	if err := db.Where(&User{Username: u.Username, Password: u.Password}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("username or password error")
		}
		return err
	}
	return nil
}

func Init() {
	dsn := "root:111111@tcp(127.0.0.1:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&User{}, &Topic{}, &TopicOption{}, &Vote{}); err != nil {
		panic("migrate failed")
	}
}

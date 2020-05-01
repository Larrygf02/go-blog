package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(70);unique;not null"`
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(60);not null"`
	Password string `gorm:"type: varchar(80); not null"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) Login(db *gorm.DB) (*User, bool) {
	var userFind User
	count := 0
	db.Where(&u).First(&userFind).Count(&count)
	if count != 0 {
		return &userFind, true
	}
	return &User{}, false
}

func (u *User) GetStories(db *gorm.DB) (*[]Storie, error) {
	var stories []Storie
	var err error
	err = db.Model(&u).Related(&stories).Error
	if err != nil {
		return &[]Storie{}, err
	}
	return &stories, nil
}

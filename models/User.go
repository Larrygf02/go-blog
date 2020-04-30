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

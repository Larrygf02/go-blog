package models

import "github.com/jinzhu/gorm"

type Storie struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string
	User    User `gorm:"foreignkey:UserId"`
	UserId  uint
}

func (s *Storie) SaveStorie(db *gorm.DB) (*Storie, error) {
	var err error
	err = db.Create(&s).Error
	if err != nil {
		return &Storie{}, err
	}
	return s, nil
}

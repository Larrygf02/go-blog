package models

import "github.com/jinzhu/gorm"

type Storie struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string
	User    User `gorm:"foreignkey:UserId"`
	UserId  uint
}

type StorieVisit struct {
	gorm.Model
	User     User `gorm:"foreignkey:UserId;not null"`
	UserId   uint
	Storie   Storie `gorm:"foreignkey:StorieId;not null"`
	StorieId uint
}

type StorieApplause struct {
	gorm.Model
	User     User `gorm:"foreignkey:UserId;not null"`
	UserId   uint
	Storie   Storie `gorm:"foreignkey:StorieId;not null"`
	StorieId uint
	Count    int
}

/* Storie */
func (s *Storie) SaveStorie(db *gorm.DB) (*Storie, error) {
	var err error
	err = db.Create(&s).Error
	if err != nil {
		return &Storie{}, err
	}
	return s, nil
}

/* Storie Visit*/
func (sv *StorieVisit) SaveStorieVisit(db *gorm.DB) (*StorieVisit, error) {
	var err error
	err = db.Create(&sv).Error
	if err != nil {
		return &StorieVisit{}, err
	}
	return sv, nil
}

func (sv *StorieVisit) GetAll(db *gorm.DB) (*[]StorieVisit, int) {
	count := 0
	var stories_visit []StorieVisit
	db.Find(&stories_visit).Count(&count)
	return &stories_visit, count
}

/* Storie Applause */
func (sa *StorieApplause) Save(db *gorm.DB) (*StorieApplause, error) {
	var err error
	err = db.Create(&sa).Error
	if err != nil {
		return &StorieApplause{}, err
	}
	return sa, nil
}

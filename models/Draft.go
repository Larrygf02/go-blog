package models

import "github.com/jinzhu/gorm"

type Draft struct {
	gorm.Model
	ID      int    `gorm:"primary_key;auto_increment" json:"id"`
	Title   string `gorm:"type:varchar(100);not null"`
	Content string
	User    User `gorm:"foreignkey:UserId"`
	UserId  uint
}

/* DRAFTS */
func (d *Draft) Save(db *gorm.DB) (*Draft, error) {
	var err error
	err = db.Create(&d).Error
	if err != nil {
		return &Draft{}, err
	}
	return d, nil
}

func (d *Draft) Update(db *gorm.DB) (*Draft, error) {
	var err error
	var updated Draft
	err = db.Model(&updated).Where("id = ?", d.ID).Updates(Draft{Content: d.Content, Title: d.Title}).Error
	if err != nil {
		return &Draft{}, err
	}
	db.First(&updated, d.ID)
	return &updated, nil
}

func (d *Draft) GetByID(db *gorm.DB) (*Draft, error) {
	var err error
	var found Draft
	err = db.First(&found, d.ID).Error
	if err != nil {
		return &Draft{}, err
	}
	return &found, nil
}

/* func (d *Draft) Filter(db *gorm.DB) */

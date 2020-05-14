package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type User struct {
	gorm.Model
	ID        int           `gorm:"primary_key;auto_increment"`
	Nickname  string        `gorm:"type:varchar(70);unique;not null"`
	Name      string        `gorm:"type:varchar(50)"`
	Email     string        `gorm:"type:varchar(60);not null"`
	Password  string        `gorm:"type: varchar(80); not null"`
	Favorites pq.Int64Array `gorm:"type:integer[]"`
	Archiveds pq.Int64Array `gorm:"type:integer[]"`
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
	//count := 0
	err := db.Where(&User{Nickname: u.Nickname, Password: u.Password}).First(&userFind).Error
	//fmt.Println(count)
	if err != nil {
		return &User{}, false
	}
	return &userFind, true
}

func (u *User) GetByID(db *gorm.DB) (*User, bool) {
	err := db.First(&u, u.ID).Error
	if err != nil {
		return &User{}, false
	}
	return u, true
}

/* Stories */
func (u *User) GetStories(db *gorm.DB) (*[]Storie, error) {
	var stories []Storie
	var err error
	err = db.Model(&u).Related(&stories).Error
	if err != nil {
		return &[]Storie{}, err
	}
	return &stories, nil
}

/* Drafts */
func (u *User) GetDrafts(db *gorm.DB) (*[]Draft, error) {
	var drafts []Draft
	var err error
	err = db.Model(&u).Related(&drafts).Error
	if err != nil {
		return &[]Draft{}, err
	}
	return &drafts, nil
}

/* Favorites */
func (u *User) SaveFavorites(db *gorm.DB, data interface{}) (*User, error) {
	var updated User
	err := db.Model(&updated).Where("id = ?", u.ID).Updates(User{Favorites: u.Favorites}).Error
	if err != nil {
		return &User{}, err
	}
	db.First(&updated, u.ID)
	return &updated, nil
}

func (u *User) GetFavorites(db *gorm.DB) ([]Storie, error) {
	var stories []Storie
	favorites := []int64(u.Favorites)
	err := db.Where(favorites).Find(&stories).Error
	if err != nil {
		return []Storie{}, err
	}
	return stories, nil
}

/* Archiveds */
func (u *User) SaveArchiveds(db *gorm.DB, data interface{}) (*User, error) {
	var updated User
	err := db.Model(&updated).Where("id = ?", u.ID).Updates(User{Archiveds: u.Archiveds}).Error
	if err != nil {
		return &User{}, err
	}
	db.First(&updated, u.ID)
	return &updated, nil
}

func (u *User) GetArchiveds(db *gorm.DB) ([]Storie, error) {
	var stories []Storie
	Archiveds := []int64(u.Archiveds)
	err := db.Where(Archiveds).Find(&stories).Error
	if err != nil {
		return []Storie{}, err
	}
	return stories, nil
}

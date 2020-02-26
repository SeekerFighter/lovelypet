package model

import (
	"github.com/jinzhu/gorm"
	"lovelypet/src/com/lovelypet/cache"
)

type User struct {
	Name string `json:"name" gorm:"size:255"`
	Tel  string `json:"tel"`
	Pwd  string `json:"pwd"`
	gorm.Model
}

func NewUser(name, tel, pwd string) *User {
	return &User{Name: name, Tel: tel, Pwd: pwd}
}

func (u *User) IsUserExist() bool {
	return cache.IsExist(u,"tel=?", u.Tel)
}

func (u *User) Insert() bool {
	return cache.Insert(u)
}
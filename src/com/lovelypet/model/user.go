package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lovelypet/src/com/lovelypet/cache"
)

type User struct {
	gorm.Model
	Name string  `json:"name" gorm:"size:255"`
	Tel string `json:"tel"`
	Pwd string `json:"pwd"`
}

func NewUser(name,tel,pwd string) *User {
	return &User{Name:name,Tel:tel,Pwd:pwd}
}

func (u *User)IsUserExist() bool {
	return !cache.DBClient.Where("tel=?",u.Tel).First(u).RecordNotFound()
}

func (u *User)Insert() bool {
	fmt.Println("Insert() called with:",u)
	err :=  cache.DBClient.Create(u).Error
	if err != nil {
		fmt.Println("Insert() called err:",err)
		return false
	}
	return true
}
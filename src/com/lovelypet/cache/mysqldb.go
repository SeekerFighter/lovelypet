package cache

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DBClient *gorm.DB

func init() {
	fmt.Println("mysql init() called...")
	var err error
	DBClient, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lovelypet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	DBClient.SingularTable(true)
	DBClient.DB().SetMaxIdleConns(10)
	DBClient.DB().SetMaxOpenConns(100)
}

func CreateTable(models ...interface{})  {
	for _,model := range models{
		if !DBClient.HasTable(model) {
			DBClient.CreateTable(model)
		}
	}
}
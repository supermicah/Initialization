package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/micah?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

func Gorm() {
	GormQuery()
}

func GormInsert() {
	user := User{
		Id:       1,
		UserName: "11",
		Password: "11",
		Email:    "11",
	}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func GormQuery() {
	user := User{}
	db.First(&user, 1)
	fmt.Println(fmt.Sprintf("%+v", user))

	tu := User{}
	db.Table("users").First(&tu, 1)
	fmt.Println(fmt.Sprintf("%+v", tu))
}

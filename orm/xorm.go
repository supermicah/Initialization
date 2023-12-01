package orm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/micah?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	}
}

func Xorm() {
	XormQuery()
}

func XormQuery() {
	user := User{}
	exist, err := engine.Table("users").Where("id = ?", 1).Get(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exist)
	fmt.Println(user)
}

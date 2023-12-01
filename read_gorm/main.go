package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"micah.wiki/demo/readgorm/dal/query"
)

func main() {
	dsn := "micah:shijinping@tcp(192.168.80.5:3306)/pandora?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	q := query.Use(db)
	s := q.SysMenu
	a, err := s.WithContext(context.Background()).Where(s.MenuID.Eq(1)).First()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(a)

}

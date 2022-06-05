package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func TestMysql() {
	db, err := gorm.Open("mysql", "root:220108@tcp(192.168.31.219:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	defer db.Close()
}

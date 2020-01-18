package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	d, err := gorm.Open("mysql", "goactions:goactions@(172.21.0.2:3306)/goactions?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Errorf("error connector database %s", err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

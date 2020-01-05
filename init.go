package main

import (
	"demoProject/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	PROTOCOL = "tcp"
	HOST     = "localhost"
	PORT     = "3306"
	DATABASE = "pro_db"
)

//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
func Init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USERNAME,
		PASSWORD,
		PROTOCOL,
		HOST,
		PORT,
		DATABASE,
	)
	db, err := gorm.Open("mysql", dsn)
	checkErr(err)
	defer func() {
		err = db.Close()
		checkErr(err)
	}()
	db.LogMode(true)

	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)

	db.AutoMigrate(&models.User{})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

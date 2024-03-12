/*
 * @Author: Aina
 * @Date: 2024-03-12 21:26:47
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 22:18:19
 * @FilePath: /ginEssential/common/database.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package common

import (
	"fmt"
	"ginEssential/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "rootroot"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	fmt.Println("database connection successfully setablished")

	db.AutoMigrate(&model.User{})
	DB = db
	return db

}
func GetDB() *gorm.DB {
	return DB
}

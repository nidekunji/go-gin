/*
 * @Author: Aina
 * @Date: 2024-03-12 16:16:17
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 21:05:26
 * @FilePath: /go/ginEssential/main.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm: "type:varchar(20); not null"`
	Telephone string `gorm: "varchar(110);not null; unique"`
	Password  string `gorm: "size:255; not null"`
}

func main() {
	db := InitDB()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current working directory:", dir)
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		fmt.Println(c.Request.PostForm)
		// 数据验证
		println(len(telephone))
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": len(telephone)})
			return
		}
		if len(password) <= 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		// 判断手机号是否存在
		log.Println(telephone, name, password)
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}
		// 创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		// 如果名字没有，随机生成10个字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
func RandomString(n int) string {
	var letters = []byte("adbckdiglirohgkABCDEFGHIKJOPUYQEJBF")
	result := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
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

	db.AutoMigrate(&User{})
	return db

}

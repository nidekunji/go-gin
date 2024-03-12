/*
 * @Author: Aina
 * @Date: 2024-03-12 21:25:44
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 22:06:02
 * @FilePath: /ginEssential/controller/UserController.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	// 获取参数
	DB := common.GetDB()
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
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// 如果名字没有，随机生成10个字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

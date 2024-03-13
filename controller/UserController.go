/*
 * @Author: Aina
 * @Date: 2024-03-12 21:25:44
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-13 22:30:22
 * @FilePath: /ginEssential/controller/UserController.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package controller

import (
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	// 获取参数
	DB := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号长度不对")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": len(telephone)})
		return
	}
	if len(password) <= 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	// 如果名字没有，随机生成10个字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Login(c *gin.Context) {
	//获取参数

	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号长度不对")
		return
	}
	if len(password) <= 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	DB := common.GetDB()
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在！")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回信息
	response.Success(c, gin.H{"token": token}, "登录成功")

}
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取用户信息成功")
}

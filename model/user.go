/*
 * @Author: Aina
 * @Date: 2024-03-12 21:27:14
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 21:27:27
 * @FilePath: /go/ginEssential/model/user.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm: "type:varchar(20); not null"`
	Telephone string `gorm: "varchar(110);not null; unique"`
	Password  string `gorm: "size:255; not null"`
}

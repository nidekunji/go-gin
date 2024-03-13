/*
 * @Author: Aina
 * @Date: 2024-03-13 22:04:43
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-13 22:09:23
 * @FilePath: /ginEssential/dto/user_dto.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package dto

import "ginEssential/model"

type UserDto struct {
	Name      string `json: "name"`
	Telephone string `json: "telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

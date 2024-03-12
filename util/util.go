/*
 * @Author: Aina
 * @Date: 2024-03-12 21:33:15
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 21:33:23
 * @FilePath: /go/ginEssential/util/util.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("adbckdiglirohgkABCDEFGHIKJOPUYQEJBF")
	result := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

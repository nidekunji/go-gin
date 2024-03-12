/*
 * @Author: Aina
 * @Date: 2024-03-12 16:16:17
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 22:47:09
 * @FilePath: /ginEssential/main.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"ginEssential/common"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get database")
	} else {
		println(sqlDB)
	}

	// // 做一些数据库操作...
	r := gin.Default()
	r = CollectRoute(r)

	// // 当不再需要连接时关闭它
	// err = sqlDB.Close()
	// if err != nil {
	// 	log.Fatal("failed to close database")
	// }

	r.Run() // listen and serve on 0.0.0.0:8080

}

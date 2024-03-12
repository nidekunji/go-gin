/*
 * @Author: Aina
 * @Date: 2024-03-12 22:09:55
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-12 22:43:21
 * @FilePath: /ginEssential/routes.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"ginEssential/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// db := common.GetDB()
	// // 获取底层sql.DB对象以进行进一步操作
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal("failed to get database")
	// }

	// // 做一些数据库操作...

	// // 当不再需要连接时关闭它
	// err = sqlDB.Close()
	// if err != nil {
	// 	log.Fatal("failed to close database")
	// }
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}

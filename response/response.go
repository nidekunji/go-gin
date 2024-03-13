/*
 * @Author: Aina
 * @Date: 2024-03-13 22:17:20
 * @LastEditors: Aina
 * @LastEditTime: 2024-03-13 22:20:51
 * @FilePath: /ginEssential/response/response.go
 * @Description:

 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

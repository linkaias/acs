package middleware

import (
	"APICallerStats/app/serve/controller"
	"APICallerStats/libs/token_lib"
	"APICallerStats/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RecordCount 访问调用记录
func RecordCount(openRecord bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if openRecord == false {
			return
		}
		// 获取path
		path := ctx.Request.URL.Path
		// 获取method
		method := ctx.Request.Method
		// 获取ip
		ip := ctx.ClientIP()
		controller.AddACSRecord(
			&model.BaseRecordModel{
				GroupId: 2,
				Name:    "Admin 访问:" + path,
				Url:     path,
				Method:  method,
				IP:      ip,
				AddTime: time.Now(),
			},
		)
		return
	}
}

func AdminAuthCheck(secret string, expire int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// check ip

		// check token
		err := token_lib.CheckAuthToken("admin", ctx.GetHeader("Authorization"), secret, expire)
		if err != nil {
			ctx.JSON(
				http.StatusOK, gin.H{
					"code": 6,
					"msg":  "身份验证失败！",
				},
			)
			ctx.Abort()
			return
		}
		return
	}
}

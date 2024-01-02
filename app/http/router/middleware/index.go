package middleware

import (
	"APICallerStats/libs/token_lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FrontAuthCheck 校验访问权限
func FrontAuthCheck(secret string, expire int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// check ip

		// check token
		err := token_lib.CheckAuthToken("http", ctx.GetHeader("Authorization"), secret, expire)
		if err != nil {
			ctx.JSON(
				http.StatusOK, gin.H{
					"code": 100,
					"msg":  "身份验证失败！",
				},
			)
			ctx.Abort()
			return
		}
		return
	}
}

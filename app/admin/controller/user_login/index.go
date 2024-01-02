package user_login

import (
	"APICallerStats/app/http/model"
	"APICallerStats/libs/token_lib"
	model2 "APICallerStats/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminUserCont struct {
	Env *model2.Env
}

func (c *AdminUserCont) Auth(ctx *gin.Context) {
	rt := map[string]interface{}{}
	info := new(model.AuthTokenModel)
	_ = ctx.ShouldBind(info)

	if info.User != c.Env.AdminUser || c.Env.AdminPassword != info.Password {
		rt["msg"] = "验证失败！"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}

	// 生成token
	token, exp := token_lib.CreateAccessToken(info.User, c.Env.SecretKey, c.Env.JwtAdminExpireTime, "admin")
	rt["msg"] = "Success！"
	rt["code"] = 200
	rt["data"] = map[string]interface{}{
		"token":    token,
		"username": info.User,
		"exp":      exp,
	}
	ctx.JSON(http.StatusOK, rt)
}

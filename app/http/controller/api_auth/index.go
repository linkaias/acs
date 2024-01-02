package api_auth

import (
	"APICallerStats/app/http/model"
	"APICallerStats/libs/token_lib"
	model2 "APICallerStats/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpAuthCont struct {
	Env *model2.Env
}

func (c *HttpAuthCont) Auth(ctx *gin.Context) {
	rt := map[string]interface{}{}
	info := new(model.AuthTokenModel)
	_ = ctx.ShouldBind(info)

	if info.User != c.Env.HTTPUser || c.Env.HTTPPassword != info.Password {
		rt["msg"] = "验证失败！"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}

	// 生成token
	token, exp := token_lib.CreateAccessToken(info.User, c.Env.SecretKey, c.Env.JWTExpireTime, "http")
	rt["msg"] = "Success！"
	rt["code"] = 200
	rt["data"] = map[string]interface{}{
		"token": token,
		"exp":   exp,
	}
	ctx.JSON(http.StatusOK, rt)
}

package router

import (
	"APICallerStats/app/http/controller/api_auth"
	"APICallerStats/app/http/controller/api_record"
	"APICallerStats/app/http/controller/api_record_group"
	"APICallerStats/app/http/router/middleware"
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/model"
	"github.com/gin-gonic/gin"
)

func InitFrontRouter(r *gin.Engine, recordGroup record_group_dao_interface.ACSRecordGroupDao, env *model.Env) {
	frontGroup := r.Group("front/api")

	// 获取token
	authCont := &api_auth.HttpAuthCont{Env: env}
	frontGroup.POST("auth_token", authCont.Auth)

	frontGroup.Use(middleware.FrontAuthCheck(env.SecretKey, env.JWTExpireTime))
	{
		cont := api_record.NewCurlApiRecordCont()
		frontGroup.POST("record", cont.Add)
	}
	{
		cont := api_record_group.NewApiRecordGroupCont(recordGroup)
		frontGroup.GET("record_group", cont.List)
		frontGroup.POST("record_group", cont.Add)
		frontGroup.PUT("record_group", cont.Update)
		frontGroup.DELETE("record_group", cont.Delete)
	}

}

package router

import (
	"APICallerStats/app/admin/controller/chart_data"
	"APICallerStats/app/admin/controller/record_list"
	"APICallerStats/app/admin/controller/system"
	"APICallerStats/app/admin/controller/user_login"
	"APICallerStats/app/admin/router/middleware"
	"APICallerStats/app/http/controller/api_record_group"
	"APICallerStats/common/chart_data_dao_interface"
	"APICallerStats/common/config_dao_interface"
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/model"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(
	r *gin.Engine, chartDao chart_data_dao_interface.ChartDataDao,
	recordGroup record_group_dao_interface.ACSRecordGroupDao, env *model.Env,
	recordDao record_dao_interface.ACSRecordDao,
	configDao config_dao_interface.ConfigDao,
) {
	adminGroup := r.Group("admin/api")

	// 统计后台访问
	adminGroup.Use(middleware.RecordCount(env.OpenAdminRecord))

	// 获取token
	userCont := user_login.AdminUserCont{Env: env}
	adminGroup.POST("login", userCont.Auth)

	adminGroup.Use(middleware.AdminAuthCheck(env.SecretKey, env.JwtAdminExpireTime))

	// chart
	{
		group := adminGroup.Group("chart")
		cont := &chart_data.AdminChartDataCont{Dao: chartDao, GroupDao: recordGroup, RecordDao: recordDao}

		// 总调用记录统计
		group.GET("all_chart_data", cont.AllData)

		// 按分组统计
		group.GET("get_chart_data_by_group", cont.GetChartDataByGroup)
		// 获取全部分组统计
		group.GET("get_chart_data_by_all_group", cont.GetChartDataByAllGroup)
		// 按请求方法统计
		group.GET("get_chart_data_by_method", cont.GetChartDataByMethod)
		// 按全部请求方法统计
		group.GET("get_chart_data_by_all_method", cont.GetChartDataByAllMethod)

	}
	// record
	{
		group := adminGroup.Group("record")
		cont := &record_list.AdminRecordCont{Dao: recordDao}
		group.GET("list", cont.List)
	}
	// group
	{
		group := adminGroup.Group("group")
		cont := api_record_group.NewApiRecordGroupCont(recordGroup)
		group.GET("record_group", cont.List)
		group.POST("record_group", cont.Add)
		group.PUT("record_group", cont.Update)
		group.DELETE("record_group", cont.Delete)
	}
	// system info
	{
		group := adminGroup.Group("system")
		cont := system.AdminSystemCont{Env: env, ConfigDao: configDao}
		group.GET("info", cont.GetInfo)

		group.GET("get_config", cont.GetConfig)
		group.POST("save_config", cont.SaveConfig)
		group.POST("ping_db", cont.TestDBConn)

	}

}

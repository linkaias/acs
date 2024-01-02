package chart_data

import (
	"APICallerStats/app/admin/controller/help_lib"
	"APICallerStats/app/admin/model"
	"APICallerStats/common/chart_data_dao_interface"
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/common/record_group_dao_interface"
	model2 "APICallerStats/model"
	"github.com/gin-gonic/gin"
)

type AdminChartDataCont struct {
	Dao       chart_data_dao_interface.ChartDataDao
	RecordDao record_dao_interface.ACSRecordDao
	GroupDao  record_group_dao_interface.ACSRecordGroupDao
}

func (c AdminChartDataCont) AllData(ctx *gin.Context) {
	rt := make(map[string]interface{})
	query := new(model.ChartDataRequestModel)
	_ = ctx.ShouldBind(query)

	if err := query.AutoVerify(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}
	info := c.Dao.GetCountByDateRange(query)
	for _, val := range info {
		val.FormatTimeStr(query.TimeDimensions)
	}
	values, timeRange := convertChartData(info, query)
	rt["code"] = 200
	rt["msg"] = "Success！"
	rt["data"] = map[string]interface{}{
		"x_data":           timeRange,
		"y_data_all_count": values,
	}
	ctx.JSON(200, rt)
}

// GetChartDataByGroup 按分组获取数据
func (c AdminChartDataCont) GetChartDataByGroup(ctx *gin.Context) {
	rt := make(map[string]interface{})
	query := new(model.ChartDataRequestModel)
	_ = ctx.ShouldBind(query)

	if err := query.AutoVerify(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}
	if err := query.VerifyGroupId(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}

	info := c.Dao.GetCountByDateRange(query)
	for _, val := range info {
		val.FormatTimeStr(query.TimeDimensions)
	}
	values, timeRange := convertChartData(info, query)
	rt["code"] = 200
	rt["msg"] = "Success！"
	rt["data"] = map[string]interface{}{
		"x_data":           timeRange,
		"y_data_all_count": values,
	}
	ctx.JSON(200, rt)
}

// GetChartDataByMethod 按请求方法获取数据
func (c AdminChartDataCont) GetChartDataByMethod(ctx *gin.Context) {
	rt := make(map[string]interface{})
	query := new(model.ChartDataRequestModel)
	_ = ctx.ShouldBind(query)
	if err := query.AutoVerify(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}
	if err := query.VerifyMethod(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}

	info := c.Dao.GetCountByDateRange(query)
	for _, val := range info {
		val.FormatTimeStr(query.TimeDimensions)
	}
	values, timeRange := convertChartData(info, query)
	rt["code"] = 200
	rt["msg"] = "Success！"
	rt["data"] = map[string]interface{}{
		"x_data":           timeRange,
		"y_data_all_count": values,
	}
	ctx.JSON(200, rt)
}

// GetChartDataByAllGroup 获取全部分组
func (c AdminChartDataCont) GetChartDataByAllGroup(ctx *gin.Context) {
	rt := make(map[string]interface{})
	query := new(model.ChartDataRequestModel)
	_ = ctx.ShouldBind(query)
	if err := query.AutoVerify(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}

	groups, err := c.GroupDao.List()
	if err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}

	timeRanges, allValues := make([]string, 0), make(map[string][]int, 0)
	for _, group := range groups {
		query.GroupId = group.Id
		info := c.Dao.GetCountByDateRange(query)
		for _, val := range info {
			val.FormatTimeStr(query.TimeDimensions)
		}
		values, timeRange := convertChartData(info, query)
		if len(timeRanges) <= 0 {
			timeRanges = timeRange
		}
		allValues[group.Name] = values
	}
	rt["code"] = 200
	rt["msg"] = "Success！"
	respData := map[string]interface{}{}
	respData["x_data"] = timeRanges
	respData["y_data"] = allValues
	rt["data"] = respData
	ctx.JSON(200, rt)
}

// GetChartDataByAllMethod 获取全部method
func (c AdminChartDataCont) GetChartDataByAllMethod(ctx *gin.Context) {
	rt := make(map[string]interface{})
	query := new(model.ChartDataRequestModel)
	_ = ctx.ShouldBind(query)
	if err := query.AutoVerify(); err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}

	methods, err := c.RecordDao.GetMethodsByRecord()
	if err != nil {
		rt["code"] = 101
		rt["msg"] = err.Error()
		ctx.JSON(200, rt)
		return
	}

	timeRanges, allValues := make([]string, 0), make(map[string][]int, 0)
	for _, method := range methods {
		query.Method = method.Method
		info := c.Dao.GetCountByDateRange(query)
		for _, val := range info {
			val.FormatTimeStr(query.TimeDimensions)
		}
		values, timeRange := convertChartData(info, query)
		if len(timeRanges) <= 0 {
			timeRanges = timeRange
		}
		allValues[method.Method] = values
	}
	rt["code"] = 200
	rt["msg"] = "Success！"
	respData := map[string]interface{}{}
	respData["x_data"] = timeRanges
	respData["y_data"] = allValues
	rt["data"] = respData
	ctx.JSON(200, rt)
}

func convertChartData(wait []*model2.ChartDataMysqlCountModel, query *model.ChartDataRequestModel) ([]int, []string) {
	values := make([]int, 0)
	help := help_lib.AdminChartHelp{}
	timeRange := help.GetTimeRangeStr(query)
	for _, timeStr := range timeRange {
		value := findTimeByData(wait, timeStr)
		values = append(values, value)
	}
	return values, timeRange
}
func findTimeByData(
	wait []*model2.ChartDataMysqlCountModel, nowTimeStr string,
) int {
	for _, val := range wait {
		if val.TimeStr == nowTimeStr {
			return int(val.RecordCount)
		}
	}
	return 0
}

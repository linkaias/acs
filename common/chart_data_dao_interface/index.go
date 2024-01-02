package chart_data_dao_interface

import (
	model2 "APICallerStats/app/admin/model"
	"APICallerStats/model"
)

type ChartDataDao interface {
	GetCountByDateRange(query *model2.ChartDataRequestModel) []*model.ChartDataMysqlCountModel
}

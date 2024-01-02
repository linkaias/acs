package mysql_dao

import (
	model2 "APICallerStats/app/admin/model"
	"APICallerStats/common/chart_data_dao_interface"
	"APICallerStats/model"
	"gorm.io/gorm"
)

type chartDataDao struct {
	db *gorm.DB
}

func NewMysqlChartDataDao(db *gorm.DB) chart_data_dao_interface.ChartDataDao {
	return &chartDataDao{
		db: db,
	}
}

func (d chartDataDao) GetCountByDateRange(query *model2.ChartDataRequestModel) []*model.ChartDataMysqlCountModel {

	formatStr := "%Y-%m-%d %H:%i"
	switch query.TimeDimensions {
	case "day":
		formatStr = "%Y-%m-%d"
		break
	case "hour":
		formatStr = "%Y-%m-%d %H"
		break
	case "min":
		formatStr = "%Y-%m-%d %H:%i"
		break
	}
	info := make([]*model.ChartDataMysqlCountModel, 0)
	db := d.db.Model(&model.RecordMysqlModel{})
	db.Where("add_time >= ?", query.BeginTime)
	db.Where("add_time <= ?", query.EndTime)
	if query.GroupId > 0 {
		db.Where("group_id = ?", query.GroupId)
	}
	if query.Method != "" {
		db.Where("method = ?", query.Method)
	}
	db.Group("timeStr")
	db.Select("DATE_FORMAT(add_time, '" + formatStr + "') as timeStr, count(id) as record_count")
	db.Scan(&info)
	return info
}

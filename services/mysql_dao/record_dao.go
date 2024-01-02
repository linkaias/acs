package mysql_dao

import (
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/model"
	"gorm.io/gorm"
	"strings"
)

type recordDao struct {
	db *gorm.DB
}

func NewMysqlRecordDao(db *gorm.DB) record_dao_interface.ACSRecordDao {
	return &recordDao{
		db: db,
	}
}

func (d recordDao) Add(value []*model.BaseRecordModel) error {
	return d.db.Create(convertInsertData(value)).Error
}

func convertInsertData(wait []*model.BaseRecordModel) []*model.RecordMysqlModel {
	info := make([]*model.RecordMysqlModel, 0)
	for _, val := range wait {
		info = append(
			info, &model.RecordMysqlModel{
				GroupId: val.GroupId,
				Name:    val.Name,
				Url:     val.Url,
				Method:  val.Method,
				Ip:      val.IP,
				AddTime: val.AddTime,
			},
		)
	}
	return info
}

func (d recordDao) ClearRecord(endTime int64) error {
	return d.db.Where("add_time < ?", endTime).Delete(&model.RecordMysqlModel{}).Error
}

func (d recordDao) GetMethodsByRecord() ([]*model.BaseRecordModel, error) {
	wait := make([]*model.BaseRecordModel, 0)
	d.db.Model(&model.RecordMysqlModel{}).Group("method").Select("method").Scan(&wait)
	return wait, nil
}

func (d recordDao) List(query *model.BaseRecordModel, filter *model.BaseListRequestModel) (
	int64, []*model.BaseRecordModel, error,
) {
	wait, result := make([]*model.RecordMysqlModel, 0), make([]*model.BaseRecordModel, 0)
	db := d.db.Model(&model.RecordMysqlModel{})
	if query.GroupId != 0 {
		db = db.Where("group_id = ?", query.GroupId)
	}
	if query.Method != "" {
		db = db.Where("method = ?", strings.ToUpper(query.Method))
	}
	if query.Url != "" {
		db = db.Where("url like ?", "%"+query.Url+"%")
	}
	if query.IP != "" {
		db = db.Where("ip = ?", query.IP)
	}
	if filter.StartTime != "" {
		db = db.Where("add_time between ? and ?", filter.StartTime, filter.EndTime)
	}
	var count int64
	if filter.Page == 1 {
		db.Count(&count)
	}

	db.Scopes(
		func(db *gorm.DB) *gorm.DB {
			offset := (filter.Page - 1) * filter.Size
			return db.Offset(offset).Limit(filter.Size)
		},
	)
	db.Order("id desc")
	res := db.Find(&wait)

	for _, val := range wait {
		result = append(
			result, &model.BaseRecordModel{
				Id:      val.Id,
				GroupId: val.GroupId,
				Name:    val.Name,
				Url:     val.Url,
				Method:  val.Method,
				IP:      val.Ip,
				AddTime: val.AddTime,
			},
		)
	}
	return count, result, res.Error
}

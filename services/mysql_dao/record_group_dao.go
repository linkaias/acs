package mysql_dao

import (
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/model"
	"gorm.io/gorm"
)

type recordGroupDao struct {
	db *gorm.DB
}

func NewMysqlRecordGroupDao(db *gorm.DB) record_group_dao_interface.ACSRecordGroupDao {
	return recordGroupDao{
		db: db,
	}
}

func (a recordGroupDao) List() ([]*model.BaseRecordGroupModel, error) {
	info := make([]*model.RecordGroupMysqlModel, 0)
	err := a.db.Model(&model.RecordGroupMysqlModel{}).Find(&info).Error
	if err != nil {
		return nil, err
	}
	res := make([]*model.BaseRecordGroupModel, 0)
	for _, val := range info {
		res = append(
			res, &model.BaseRecordGroupModel{
				Id:        val.Id,
				Name:      val.Name,
				Desc:      val.Desc,
				CreatedAt: val.CreatedAt,
				UpdatedAt: val.UpdatedAt,
			},
		)
	}
	return res, err
}

func (a recordGroupDao) Add(value *model.BaseRecordGroupModel) error {
	return a.db.Create(convertRecordGroupMysqlModel(value)).Error
}
func (a recordGroupDao) GetByGroupId(id int) (*model.BaseRecordGroupModel, error) {
	val := new(model.RecordGroupMysqlModel)
	res := a.db.Model(&val).Where("id = ?", id).First(val)
	return &model.BaseRecordGroupModel{
		Id:        val.Id,
		Name:      val.Name,
		Desc:      val.Desc,
		CreatedAt: val.CreatedAt,
		UpdatedAt: val.UpdatedAt,
	}, res.Error
}
func (a recordGroupDao) UpByGroupId(value *model.BaseRecordGroupModel) error {
	return a.db.Model(&model.RecordGroupMysqlModel{}).Where(
		"id = ?", value.Id,
	).Select("name", "desc").Updates(convertRecordGroupMysqlModel(value)).Error
}
func (a recordGroupDao) DelByGroupId(id int) error {
	return a.db.Where("id = ?", id).Delete(&model.RecordGroupMysqlModel{}).Error
}
func convertRecordGroupMysqlModel(wait *model.BaseRecordGroupModel) *model.RecordGroupMysqlModel {
	val := new(model.RecordGroupMysqlModel)
	val.Id = wait.Id
	val.Name = wait.Name
	val.Desc = wait.Desc
	val.CreatedAt = wait.CreatedAt
	val.UpdatedAt = wait.UpdatedAt
	return val
}

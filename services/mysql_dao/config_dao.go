package mysql_dao

import (
	"APICallerStats/common/config_dao_interface"
	"APICallerStats/model"
	"gorm.io/gorm"
)

type cfgDao struct {
	db *gorm.DB
}

func NewMysqlConfigDao(db *gorm.DB) config_dao_interface.ConfigDao {
	return &cfgDao{
		db: db,
	}
}

func (d cfgDao) Get(key string) string {
	val := new(model.ACSConfigMysqlModel)
	d.db.Model(val).Where("`key` = ?", key).First(&val)
	return val.Value
}
func (d cfgDao) Set(key, value string) error {
	old := new(model.ACSConfigMysqlModel)
	d.db.Model(old).Where("`key` = ?", key).First(&old)
	if old.Id <= 0 {
		return d.db.Create(
			&model.ACSConfigMysqlModel{
				Key:   key,
				Value: value,
			},
		).Error
	} else {
		return d.db.Model(old).Where("`key` = ?", key).Update("value", value).Error
	}
}

func (d cfgDao) Del(key string) error {
	return d.db.Where("`key` = ?", key).Delete(&model.ACSConfigMysqlModel{}).Error
}

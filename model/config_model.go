package model

type ACSConfigMysqlModel struct {
	Id    int64  `json:"id" gorm:"primaryKey;column:id;autoIncrement" form:"id"`
	Key   string `json:"key" gorm:"index;column:key;size:255" form:"key"`
	Value string `json:"value" gorm:"column:value;type:text" form:"value"`
}

func (*ACSConfigMysqlModel) TableName() string {
	return "acs_config"
}

package model

import "time"

type BaseRecordGroupModel struct {
	Id        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Desc      string    `json:"desc" form:"desc"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

// RecordGroupMysqlModel 调用记录分组
type RecordGroupMysqlModel struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id;autoIncrement;type:int(11)" form:"id"`
	Name      string    `json:"name" gorm:"unique;column:name;comment:分组名称;type:varchar(100)" form:"name"`
	Desc      string    `json:"desc" gorm:"column:desc;comment:分组描述;type:varchar(255)" form:"desc"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" form:"updated_at"`
}

func (*RecordGroupMysqlModel) TableName() string {
	return "acs_call_record_group"
}

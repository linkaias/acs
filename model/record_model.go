package model

import "time"

type BaseRecordModel struct {
	Id      int64     `json:"id" form:"id"`
	GroupId int       `json:"group_id" form:"group_id"`
	Name    string    `json:"name" form:"name"`
	Url     string    `json:"url" form:"url"`
	Method  string    `json:"method" form:"method"`
	IP      string    `json:"ip" form:"ip"`
	AddTime time.Time `json:"add_time" form:"add_time"`
}

type RecordMysqlModel struct {
	Id      int64     `json:"id" gorm:"primaryKey;column:id;autoIncrement;type:bigint unsigned" form:"id"`
	GroupId int       `json:"group_id" gorm:"index;column:group_id;comment:分组;type:int(11)" form:"group_id"`
	Name    string    `json:"name" gorm:"column:name;comment:别名;type:varchar(100)" form:"name"`
	Url     string    `json:"url" gorm:"column:url;comment:请求url;type:varchar(255)" form:"url"`
	Method  string    `json:"method" gorm:"index;column:method;comment:请求方法;type:varchar(10)" form:"method"`
	Ip      string    `json:"ip" gorm:"column:ip;comment:请求ip;type:varchar(20)" form:"ip"`
	AddTime time.Time `json:"add_time" gorm:"index;column:add_time;comment:添加时间;type:timestamp" form:"add_time"`
}

func (*RecordMysqlModel) TableName() string {
	return "acs_call_record"
}

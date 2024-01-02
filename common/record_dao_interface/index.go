package record_dao_interface

import "APICallerStats/model"

type ACSRecordDao interface {
	Add(value []*model.BaseRecordModel) error
	ClearRecord(endTime int64) error
	GetMethodsByRecord() ([]*model.BaseRecordModel, error)
	List(query *model.BaseRecordModel, filter *model.BaseListRequestModel) (int64, []*model.BaseRecordModel, error)
}

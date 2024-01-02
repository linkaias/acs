package record_group_dao_interface

import "APICallerStats/model"

type ACSRecordGroupDao interface {
	List() ([]*model.BaseRecordGroupModel, error)
	Add(value *model.BaseRecordGroupModel) error
	GetByGroupId(id int) (*model.BaseRecordGroupModel, error)
	UpByGroupId(value *model.BaseRecordGroupModel) error
	DelByGroupId(id int) error
}

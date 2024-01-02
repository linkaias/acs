package auto_clear_record

import (
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/model"
	"time"
)

type AutoClearRecordCont struct {
	Env *model.Env
}

func (a AutoClearRecordCont) ClearMysqlRecord(dao record_dao_interface.ACSRecordDao) {
	delTime := time.Now().Unix() - int64(86400*a.Env.RecordSaveDays)
	_ = dao.ClearRecord(delTime)
}

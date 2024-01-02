package crons

import (
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/crons/auto_clear_record"
	"APICallerStats/model"
	"github.com/robfig/cron/v3"
)

func RunCron(env *model.Env, recordDao record_dao_interface.ACSRecordDao) {
	c := cron.New(cron.WithSeconds()) //支持秒级

	//每天定时清理record旧数据
	{
		task := auto_clear_record.AutoClearRecordCont{Env: env}
		_, _ = c.AddFunc(
			"0 0 0 */1 * ?", func() {
				task.ClearMysqlRecord(recordDao)
			},
		)
	}

	c.Run()
}

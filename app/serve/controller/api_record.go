package controller

import (
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/model"
	"strings"
	"sync"
	"time"
)

// WaitWriteRecord 等待写入的原始调用记录
var WaitWriteRecord = make(chan *model.BaseRecordModel, 10000)

// 等待批量插入的数据
var waitBatchInsert = make(chan []*model.BaseRecordModel, 1000)

var batchData = new(batchInsert)

type batchInsert struct {
	batch []*model.BaseRecordModel
	lock  sync.Mutex
}

func (b *batchInsert) Append(val *model.BaseRecordModel) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.batch = append(b.batch, val)
}

func (b *batchInsert) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.batch = make([]*model.BaseRecordModel, 0)
}
func (b *batchInsert) Len() int {
	b.lock.Lock()
	defer b.lock.Unlock()
	return len(b.batch)
}

func RunRecordTask(env *model.Env, dao record_dao_interface.ACSRecordDao) {
	// 整理数据
	go func() {
		for wait := range WaitWriteRecord {
			// 转大写
			wait.Method = strings.ToUpper(wait.Method)
			batchData.Append(wait)
			if batchData.Len() >= env.DBBatchInsertNum {
				waitBatchInsert <- batchData.batch
				batchData.Reset()
			}
		}
	}()
	go func() {
		for {
			// 每隔一段时间进行批量插入
			time.Sleep(time.Duration(env.DBBatchInsertGap) * time.Second)
			if batchData.Len() > 0 {
				waitBatchInsert <- batchData.batch
				batchData.Reset()
			}
		}
	}()
	// 入库
	go func() {
		log := &LogClient{}
		for batch := range waitBatchInsert {
			err := dao.Add(batch)
			if err != nil {
				log.Error("batch insert error", err.Error(), nil)
			}
		}
	}()
}

// AddACSRecord 新增调用记录
func AddACSRecord(val *model.BaseRecordModel) {
	WaitWriteRecord <- val
}

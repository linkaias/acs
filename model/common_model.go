package model

import (
	"time"
)

type BaseListRequestModel struct {
	Page      int    `json:"page" form:"page"`
	Size      int    `json:"size" form:"size"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}

func (b *BaseListRequestModel) Verify() {
	if b.Page == 0 {
		b.Page = 1
	}
	if b.Size == 0 || b.Size > 500 {
		b.Size = 10
	}

	if b.StartTime == "" || b.EndTime == "" {
		b.StartTime = ""
		b.EndTime = ""
		return
	}

	begin, _ := time.ParseInLocation("2006-01-02 15:04", b.StartTime, time.Local)
	end, _ := time.ParseInLocation("2006-01-02 15:04", b.EndTime, time.Local)
	// 判断begin end都不为空，且end大于begin
	if !begin.IsZero() && !end.IsZero() && end.Before(begin) {
		b.StartTime, b.EndTime = b.EndTime, b.StartTime
		begin, end = end, begin
	}
	b.StartTime = begin.Format("2006-01-02 15:04:05")
	b.EndTime = end.Format("2006-01-02 15:04:05")
}

type BaseResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

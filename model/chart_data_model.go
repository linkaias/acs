package model

type ChartDataMysqlCountModel struct {
	TimeStr     string `json:"timeStr" gorm:"column:timeStr" form:"timeStr"`
	RecordCount int64  `json:"record_count" gorm:"column:record_count" form:"record_count"`
}

func (c *ChartDataMysqlCountModel) FormatTimeStr(dimensions string) {
	switch dimensions {
	case "day":
		c.TimeStr = c.TimeStr + " 00:00:00"
		break
	case "hour":
		c.TimeStr = c.TimeStr + ":00:00"
		break
	case "min":
		c.TimeStr = c.TimeStr + ":00"
		break
	}
}

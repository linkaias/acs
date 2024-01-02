package help_lib

import (
	"APICallerStats/app/admin/model"
	"time"
)

type AdminChartHelp struct{}

func (h AdminChartHelp) GetTimeRangeStr(query *model.ChartDataRequestModel) []string {
	wait := make([]string, 0)
	if query.BeginTimeDate.IsZero() || query.EndTimeDate.IsZero() {
		return wait
	}
	// 按照维度切割时间
	switch query.TimeDimensions {
	case model.TimeDimensionsMin:
		for i := query.BeginTimeDate; i.Before(query.EndTimeDate); i = i.Add(time.Minute) {
			wait = append(wait, i.Format("2006-01-02 15:04")+":00")
		}
		break
	case model.TimeDimensionsHour:
		for i := query.BeginTimeDate; i.Before(query.EndTimeDate); i = i.Add(time.Hour) {
			wait = append(wait, i.Format("2006-01-02 15")+":00:00")
		}
		break
	case model.TimeDimensionsDay:
		for i := query.BeginTimeDate; i.Before(query.EndTimeDate); i = i.AddDate(0, 0, 1) {
			wait = append(wait, i.Format("2006-01-02")+" 00:00:00")
		}
		break
	case model.TimeDimensionsMonth:
		for i := query.BeginTimeDate; i.Before(query.EndTimeDate); i = i.AddDate(0, 1, 0) {
			wait = append(wait, i.Format("2006-01")+"-01 00:00:00")
		}
	}
	return wait
}

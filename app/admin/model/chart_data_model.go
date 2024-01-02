package model

import (
	"fmt"
	"time"
)

const (
	TimeDimensionsMin   = "min"
	TimeDimensionsHour  = "hour"
	TimeDimensionsDay   = "day"
	TimeDimensionsMonth = "month"
)

type ChartDataRequestModel struct {
	GroupId        int       `json:"group_id" form:"group_id"`               // 分组ID
	Method         string    `json:"method" form:"method"`                   // 请求类别
	TimeDimensions string    `json:"time_dimensions" form:"time_dimensions"` //时间维度
	BeginTime      string    `json:"begin_time" form:"begin_time" `          //开始时间
	BeginTimeDate  time.Time `json:"-" form:"-" `                            //开始时间
	EndTime        string    `json:"end_time" form:"end_time" `              //结束时间
	EndTimeDate    time.Time `json:"-" form:"-" `                            //结束时间
}

func (c *ChartDataRequestModel) AutoVerify() error {
	// 时间范围筛选
	err := c.VerifyTimeRange()
	if err != nil {
		return err
	}

	// 时间粒度
	err = c.VerifyTimeDimensions()
	if err != nil {
		return err
	}

	return c.VerifyTimeRangeLegal()
}

func (c *ChartDataRequestModel) VerifyTimeDimensions() error {
	//判断时间维度是否合法
	if c.TimeDimensions != TimeDimensionsHour &&
		c.TimeDimensions != TimeDimensionsDay &&
		c.TimeDimensions != TimeDimensionsMin &&
		c.TimeDimensions != TimeDimensionsMonth {
		return fmt.Errorf("time dimensions is not valid. allows: min/hour/day/month")
	}
	return nil
}

func (c *ChartDataRequestModel) VerifyTimeRange() error {
	if len(c.BeginTime) > 16 {
		c.BeginTime = c.BeginTime[:len(c.BeginTime)-3]
	}
	if len(c.EndTime) > 16 {
		c.EndTime = c.EndTime[:len(c.EndTime)-3]
	}

	begin, err := time.ParseInLocation("2006-01-02 15:04", c.BeginTime, time.Local)
	if err != nil {
		return err
	}
	end, err := time.ParseInLocation("2006-01-02 15:04", c.EndTime, time.Local)
	if err != nil {
		return err
	}
	// end 添加一秒钟
	end = end.Add(time.Second)

	c.BeginTimeDate = begin
	c.EndTimeDate = end
	if begin.After(end) {
		c.BeginTime, c.EndTime = c.EndTime, c.BeginTime
		c.BeginTimeDate, c.EndTimeDate = c.EndTimeDate, c.BeginTimeDate
	}
	return nil
}

func (c *ChartDataRequestModel) VerifyTimeRangeLegal() error {
	// 如果TimeDimensions是min, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1分钟，并且间隔最大不能超过1天
	if c.TimeDimensions == TimeDimensionsMin {
		timeDiff := c.EndTimeDate.Sub(c.BeginTimeDate)
		if timeDiff <= time.Minute {
			return fmt.Errorf("当前维度是 分钟：时间间隔必须大于1分钟")
		}
		if timeDiff > 24*time.Hour {
			return fmt.Errorf("当前维度是 分钟：间隔最大不能超过1天")
		}
	}
	// 如果TimeDimensions是hour, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1小时并且间隔最大不能超过31天
	if c.TimeDimensions == TimeDimensionsHour {
		timeDiff := c.EndTimeDate.Sub(c.BeginTimeDate)
		if timeDiff <= time.Hour {
			return fmt.Errorf("当前维度是 小时：时间间隔必须大于1小时")
		}
		if timeDiff > 31*24*time.Hour {
			return fmt.Errorf("当前维度是 小时：间隔最大不能超过31天")
		}
	}

	// 如果TimeDimensions是day, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1天，并且间隔最大不能超过1年
	if c.TimeDimensions == TimeDimensionsDay {
		timeDiff := c.EndTimeDate.Sub(c.BeginTimeDate)
		if timeDiff <= 24*time.Hour {
			return fmt.Errorf("当前维度是 小时：时间间隔必须大于1天")
		}
		if timeDiff > 365*24*time.Hour {
			return fmt.Errorf("当前维度是 小时：间隔最大不能超过1年")
		}
	}
	// 如果TimeDimensions是month, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1月，并且间隔最大不能超过10年
	if c.TimeDimensions == TimeDimensionsMonth {
		timeDiff := c.EndTimeDate.Sub(c.BeginTimeDate)
		if timeDiff <= 30*24*time.Hour {
			return fmt.Errorf("当前维度是 月：时间间隔必须大于1月")
		}
		if timeDiff > 10*365*24*time.Hour {
			return fmt.Errorf("当前维度是 月：间隔最大不能超过10年")
		}
	}

	return nil
}

func (c *ChartDataRequestModel) VerifyGroupId() error {
	if c.GroupId <= 0 {
		return fmt.Errorf("group id is not valid")
	}
	return nil
}

func (c *ChartDataRequestModel) VerifyMethod() error {
	if c.Method == "" {
		return fmt.Errorf("method is not valid")
	}
	return nil

	// 检查请求方法method全部转为大写后，是否符合规范。
	//validMethods := []string{"POST", "GET", "DELETE", "PUT", "PATCH", "OPTIONS", "HEAD", "CONNECT", "TRACE"}
	//method := strings.ToUpper(c.Method)
	//for _, m := range validMethods {
	//	if m == method {
	//		return nil
	//	}
	//}
	//return fmt.Errorf("Invalid method: %s ", c.Method)
}


const VerifyTimeRange = (TimeDimension, BeginTimeDate, EndTimeDate)=>{
    // 先判断BeginTimeDate和EndTimeDate是否为date格式，如果不是先转换为日期格式
    if (!(BeginTimeDate instanceof Date)) {
        BeginTimeDate = new Date(BeginTimeDate);
    }
    if (!(EndTimeDate instanceof Date)) {
        EndTimeDate = new Date(EndTimeDate);
    }

    // 如果TimeDimension是min, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1分钟，并且间隔最大不能超过1天
    if (TimeDimension === "min") {
      const timeDiff = EndTimeDate.getTime() - BeginTimeDate.getTime();
      const oneMinute = 60 * 1000;
      const oneDay = 24 * 60 * 60 * 1000;
      if (timeDiff <= oneMinute) {
          // 时间间隔必须大于1分钟
          return "当前维度是 分钟：时间间隔必须大于1分钟"
      }
      if (timeDiff > oneDay) {
        // 间隔最大不能超过1天
        return "当前维度是 分钟：间隔最大不能超过1天"
      }
    }

    // 如果TimeDimensions是hour, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1小时并且间隔最大不能超过31天
    if (TimeDimension === "hour") {
      const timeDiff = EndTimeDate.getTime() - BeginTimeDate.getTime();
      const oneHour = 60 * 60 * 1000;
      const thirtyOneDays = 31 * 24 * 60 * 60 * 1000;
      if (timeDiff <= oneHour) {
        // 时间间隔必须大于1小时
        return "当前维度是 小时：时间间隔必须大于1小时";
      }
      if (timeDiff > thirtyOneDays) {
        // 间隔最大不能超过31天
        return "当前维度是 小时：间隔最大不能超过31天";
      }
    }

    // 如果TimeDimensions是day, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1天，并且间隔最大不能超过1年
    if (TimeDimension === "day") {
      const timeDiff = EndTimeDate.getTime() - BeginTimeDate.getTime();
      const oneDay = 24 * 60 * 60 * 1000;
      const oneYear = 365 * oneDay;
      if (timeDiff <= oneDay) {
        // 时间间隔必须大于1天
        return "当前维度是 天：时间间隔必须大于1天";
      }
      if (timeDiff > oneYear) {
        // 间隔最大不能超过1年
        return "当前维度是 天：间隔最大不能超过1年";
      }
    }
    // 如果TimeDimensions是month, 那么BeginTimeDate和EndTimeDate的时间间隔必须大于1月，并且间隔最大不能超过10年
    if (TimeDimension === "month") {
      const timeDiff = EndTimeDate.getTime() - BeginTimeDate.getTime();
      const oneMonth = 30 * 24 * 60 * 60 * 1000;
      const tenYears = 10 * 365 * 24 * 60 * 60 * 1000;
      if (timeDiff <= oneMonth) {
        // 时间间隔必须大于1月
        return "当前维度是 月：时间间隔必须大于1月";
      }
      if (timeDiff > tenYears) {
        // 间隔最大不能超过10年
        return "当前维度是 月：间隔最大不能超过10年";
      }
    }
    return  ''
}

const TimeRangeShortcuts = () =>{
    return [
        {
            text: '最近15分钟',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 60 * 1000 * 15)
                return [start, end]
            }
        },
        {
            text: '最近1小时',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000)
                return [start, end]
            },
        },
        {
            text: '最近3小时',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 3)
                return [start, end]
            },
        },
        {
            text: '最近6小时',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 6)
                return [start, end]
            },
        },
        {
            text: '最近12小时',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 12)
                return [start, end]
            },
        },
        {
            text: '最近1天',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 24)
                return [start, end]
            }
        },
        {
            text: '最近7天',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
                return [start, end]
            }
        },
        {
            text: '最近15天',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 15)
                return [start, end]
            }
        },
        {
            text: '最近1个月',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
                return [start, end]
            }
        },
        {
            text: '最近3个月',
            value: () => {
                const end = new Date()
                const start = new Date()
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
                return [start, end]
            }
        },
    ]

}

export {
    VerifyTimeRange,
    TimeRangeShortcuts
}
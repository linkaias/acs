import request from "../request";

export const reqGetAllCount = (time_dimensions,begin_time,end_time) => request({
    url: `/chart/all_chart_data?time_dimensions=${time_dimensions}&begin_time=${begin_time}&end_time=${end_time}`,
    method: "get",
})
export const reqGetChartByGroupId = (time_dimensions,begin_time,end_time, group_id) => request({
    url: `/chart/get_chart_data_by_group?time_dimensions=${time_dimensions}&begin_time=${begin_time}&end_time=${end_time}&group_id=${group_id}`,
    method: "get",
})

export const reqGetChartByAllGroupId = (time_dimensions,begin_time,end_time) => request({
    url: `/chart/get_chart_data_by_all_group?time_dimensions=${time_dimensions}&begin_time=${begin_time}&end_time=${end_time}`,
    method: "get",
})

export const reqGetChartByMethod = (time_dimensions,begin_time,end_time, method) => request({
    url: `/chart/get_chart_data_by_method?time_dimensions=${time_dimensions}&begin_time=${begin_time}&end_time=${end_time}&method=${method}`,
    method: "get",
})

export const reqGetChartByAllMethod = (time_dimensions,begin_time,end_time) => request({
    url: `/chart/get_chart_data_by_all_method?time_dimensions=${time_dimensions}&begin_time=${begin_time}&end_time=${end_time}`,
    method: "get",
})
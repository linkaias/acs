
import request from "../request";

//获取全部分组列表
//params={
//       page:PageInfo.value.Page,
//       size: PageInfo.value.Limit,
//       start_time:'',
//       end_time:'',
//       group_id:0,
//       url:"",
//       method:"",
//       ip:""
//     }
export const reqGetRecordList = (params) => request({
    url: `/record/list?page=${params.page}&size=${params.size}&start_time=${params.start_time}&end_time=${params.end_time}&group_id=${params.group_id}&url=${params.url}&method=${params.method}&ip=${params.ip}`,
    method: "get",
    data:params,
})

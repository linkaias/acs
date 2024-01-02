import request from "../request";

//获取全部分组列表
export const reqGetAllGroup = () => request({
    url: `/group/record_group`,
    method: "get",
})

export const reqAddGroup = (data) => request({
    url: `/group/record_group`,
    method: "post",
    data
})

export const reqUpdateGroup = (data) => request({
    url: `/group/record_group`,
    method: "put",
    data
})
export const reqDelGroup = (id) => request({
    url: `/group/record_group?id=${id}`,
    method: "delete",
})
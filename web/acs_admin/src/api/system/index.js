
import request from "../request";

//获取系统信息
export const reqGetSystemInfo = () => request({
    url: `/system/info`,
    method: "get",
})

export const reqGetConfig = () => request({
    url: `/system/get_config`,
    method: "get",
})

export const reqSaveConfig = (data) => request({
    url: `/system/save_config`,
    method: "post",
    data:data
})

export const reqPingDb = (data) => request({
    url: `/system/ping_db`,
    method: "post",
    data:data
})
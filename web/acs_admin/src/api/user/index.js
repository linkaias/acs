import request from "../request";

export const reqGetToken = (info) => request({
    url: `/login`,
    data:info,
    method: "post",
})
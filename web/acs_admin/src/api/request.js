import axios from "axios";
import { GetToken, RemoveToken } from '@/utils/token_utils'
import { ElMessage } from "element-plus"
import {usePublicStore} from "@/stores/public"
const storePublic = usePublicStore();
//底下的代码也是创建axios实例
let requests = axios.create({
    //基础路径
    baseURL: "/admin/api",
    timeout: 6000,
});

//请求拦截器
requests.interceptors.request.use((config) => {
    config.headers['Authorization'] = 'Bearer ' + GetToken()
    if (config.method === 'post' || config.method === 'put') {
        config.headers['Content-Type'] = 'multipart/form-data';
    }
    return config
});

//响应拦截器----当服务器手动请求之后，做出响应（相应成功）会执行的
requests.interceptors.response.use(
    response => {
        let res = response.data
        if (res.code !== 200) {
            // 登录过期
            if (res.code === 6){
                RemoveToken()
                // 最多只弹出一次登陆过期
                if(!storePublic.requestError){
                    storePublic.SetRequestError(true);
                    ElMessage.error("登录过期，请重新登录!")
                    setTimeout(function () {
                        window.location.href = '/login'; // 或者使用路由导航实现跳转
                    },1500)
                }
            } else {
                ElMessage.error("错误："+res.msg);
            }
            return Promise.reject(new Error(res.msg || 'Error'))
        } else {
            return res
        }
    },
    error => {
        return Promise.reject(error)
    }
);

export default requests;
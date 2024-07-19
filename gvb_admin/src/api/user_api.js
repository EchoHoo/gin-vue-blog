import { Service } from "@/services/service"

// 邮箱登录接口
export  function emailLoginApi(data) {
    return  Service.post("/api/email_login", data)
}

// 用户列表
export function userListApi(params){
    return Service.get("/api/users", {params})
}

// 用户创建
export function userCreateApi(data){
    return Service.post("/api/users", data)
}
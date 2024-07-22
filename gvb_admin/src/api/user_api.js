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

// 删除用户 参数是用户的ID列表
export function useRemoveBatchApi(id_List){
    return Service.delete("/api/users", { data: { id_List } });
}

// 修改用户权限和昵称
export function updateUserNickNameApi(data){
    return Service.put(`/api/user_role`, data)
}
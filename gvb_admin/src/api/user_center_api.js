import { Service } from "@/services/service"


export function getUserInfoApi(){
    return Service.get("/api/user_info")
}

export function updateUserInfoApi(data){
    return Service.put("/api/user_info",data)
}

// 发送验证码
export function sendEmailCodeApi(email){
    return Service.post("/api/user_bind_email",{email})   
}
// 绑定邮箱
export function bindEmailApi(data){
    return Service.post("/api/user_bind_email",data)   
}
// 登录人与用户的聊天记录
export function getMessageRecordApi(user_id){
    return Service.post("/api/messages_record",{user_id})   
}

// 发送消息 
export function sendMessageApi(data){
    return Service.post("/api/messages",data)   
}
import axios from "axios";

export const Service = axios.create({
    timeout:7000,
    baseURL:"",
    headers: {
        "Content-Type":"application/json"
    }
});


Service.interceptors.request.use(request => {
        // 一般用于添加用户的token
        return request
})






Service.interceptors.response.use(response => {
    return response.data
})
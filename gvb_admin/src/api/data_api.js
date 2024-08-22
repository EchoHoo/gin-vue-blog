import { Service } from "@/services/service";


export function getWeekDataApi(){
    return Service.get('/api/data_login')
}

export function getDataSumApi(){
    return Service.get('/api/data_sum')
}
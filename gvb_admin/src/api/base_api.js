import { Service } from "@/services/service";


export function baseListApi(url,params){
    return Service.get(url,{params})
}

export function baseDeleteApi(url,id_list){
    return Service.delete(url,{data:{id_list}})
}